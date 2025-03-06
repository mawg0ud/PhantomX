package shellcode

import (
	"fmt"
)

// XORKey defines a simple XOR key for obfuscation
var XORKey = byte(0xAA)

// EncryptShellcodeXOR encrypts shellcode using XOR
func EncryptShellcodeXOR(shellcode []byte) []byte {
	fmt.Println("Encrypting shellcode with XOR...")

	encrypted := make([]byte, len(shellcode))
	for i, b := range shellcode {
		encrypted[i] = b ^ XORKey
	}

	fmt.Println("XOR encryption completed.")
	return encrypted
}

// DecryptShellcodeXOR decrypts shellcode using XOR
func DecryptShellcodeXOR(encrypted []byte) []byte {
	fmt.Println("Decrypting shellcode with XOR...")

	decrypted := make([]byte, len(encrypted))
	for i, b := range encrypted {
		decrypted[i] = b ^ XORKey
	}

	fmt.Println("XOR decryption completed.")
	return decrypted
}
