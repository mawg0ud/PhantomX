package shellcode

import (
	"fmt"
)

// ReverseBytes reverses the order of bytes in shellcode
func ReverseBytes(shellcode []byte) []byte {
	fmt.Println("Reversing shellcode bytes...")

	reversed := make([]byte, len(shellcode))
	for i := 0; i < len(shellcode); i++ {
		reversed[i] = shellcode[len(shellcode)-1-i]
	}

	fmt.Println("Shellcode successfully reversed.")
	return reversed
}
