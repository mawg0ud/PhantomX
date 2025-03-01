package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

// EncryptShellcode encrypts shellcode using AES
func EncryptShellcode(shellcode []byte) []byte {
	fmt.Println("Encrypting shellcode...")
	key := []byte("0123456789abcdef") // 16-byte AES key
	block, _ := aes.NewCipher(key)

	ciphertext := make([]byte, len(shellcode))
	mode := cipher.NewCFBEncrypter(block, key[:block.BlockSize()])
	mode.XORKeyStream(ciphertext, shellcode)

	fmt.Println("Encryption completed.")
	return ciphertext
}

// DecryptShellcode decrypts encrypted shellcode
func DecryptShellcode(encrypted []byte) []byte {
	fmt.Println("Decrypting shellcode...")
	key := []byte("0123456789abcdef")
	block, _ := aes.NewCipher(key)

	plaintext := make([]byte, len(encrypted))
	mode := cipher.NewCFBDecrypter(block, key[:block.BlockSize()])
	mode.XORKeyStream(plaintext, encrypted)

	fmt.Println("Decryption completed.")
	return plaintext
}
