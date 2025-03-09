package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Config holds all the settings loaded from settings.json
type Config struct {
	EncryptionKey       string `json:"encryption_key"`
	UseXOREncryption    bool   `json:"use_xor_encryption"`
	UseAESEncryption    bool   `json:"use_aes_encryption"`
	EnableMemoryEvasion bool   `json:"enable_memory_evasion"`
	EnableAntiDebugging bool   `json:"enable_anti_debugging"`
	SyscallMode         string `json:"syscall_mode"`

	Shellcode struct {
		File            string `json:"file"`
		ReverseShellcode bool  `json:"reverse_shellcode"`
		XOREncrypt      bool   `json:"xor_encrypt"`
	} `json:"shellcode"`

	Build struct {
		OutputName   string `json:"output_name"`
		GOOS         string `json:"goos"`
		GOARCH       string `json:"goarch"`
		StripBinaries bool  `json:"strip_binaries"`
	} `json:"build"`

	Logging struct {
		EnableDebug bool   `json:"enable_debug"`
		LogFile     string `json:"log_file"`
	} `json:"logging"`
}

// LoadConfig reads and parses settings.json
func LoadConfig(filename string) (*Config, error) {
	fmt.Println("[+] Loading configuration...")

	// Open config file
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("[!] Error: Unable to open config file:", err)
		return nil, err
	}
	defer file.Close()

	// Read file contents
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("[!] Error: Unable to read config file:", err)
		return nil, err
	}

	// Parse JSON
	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		fmt.Println("[!] Error: Invalid JSON format:", err)
		return nil, err
	}

	fmt.Println("[+] Configuration loaded successfully!")
	return &config, nil
}
