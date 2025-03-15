package payloadgenerator

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

// Generate random AES key
func generateAESKey() []byte {
	key := make([]byte, 32)
	_, err := io.ReadFull(rand.Reader, key)
	if err != nil {
		fmt.Println("[!] Error generating AES key:", err)
	}
	return key
}

// Encrypt shellcode using AES
func encryptShellcode(shellcode []byte, key []byte) ([]byte, []byte) {
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("[!] AES Cipher creation failed:", err)
		return nil, nil
	}

	iv := make([]byte, aes.BlockSize)
	_, err = io.ReadFull(rand.Reader, iv)
	if err != nil {
		fmt.Println("[!] IV generation failed:", err)
		return nil, nil
	}

	ciphertext := make([]byte, len(shellcode))
	mode := cipher.NewCFBEncrypter(block, iv)
	mode.XORKeyStream(ciphertext, shellcode)

	return ciphertext, iv
}

// Save encrypted payload to file
func savePayloadToFile(encryptedShellcode, iv, key []byte) {
	file, err := os.Create("encrypted_payload.bin")
	if err != nil {
		fmt.Println("[!] Error creating payload file:", err)
		return
	}
	defer file.Close()

	file.Write(append(iv, encryptedShellcode...))
	fmt.Println("[+] Encrypted payload saved successfully!")

	// Save AES key separately
	keyFile, err := os.Create("aes_key.txt")
	if err != nil {
		fmt.Println("[!] Error saving AES key:", err)
		return
	}
	defer keyFile.Close()

	keyFile.WriteString(hex.EncodeToString(key))
	fmt.Println("[+] AES key saved successfully!")
}

// Main payload generation function
func GenerateEncryptedPayload(shellcode []byte) {
	fmt.Println("[+] Generating encrypted payload...")

	key := generateAESKey()
	encryptedShellcode, iv := encryptShellcode(shellcode, key)
	if encryptedShellcode == nil {
		fmt.Println("[!] Encryption failed!")
		return
	}

	savePayloadToFile(encryptedShellcode, iv, key)
	fmt.Println("[+] Payload generation completed!")
}
