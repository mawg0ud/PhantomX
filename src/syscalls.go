package syscalls

import (
	"fmt"
	"syscall"
	"unsafe"
)

// NtAllocateVirtualMemory defines system call for memory allocation
var NtAllocateVirtualMemory = syscall.NewLazyDLL("ntdll.dll").NewProc("NtAllocateVirtualMemory")

// AllocateMemory allocates memory using direct system calls
func AllocateMemory(size uintptr) unsafe.Pointer {
	fmt.Println("Allocating memory using direct system call...")

	var baseAddr uintptr
	status, _, _ := NtAllocateVirtualMemory.Call(
		uintptr(syscall.Handle(0xFFFFFFFF)), // Current process handle
		uintptr(unsafe.Pointer(&baseAddr)),
		0,
		uintptr(unsafe.Pointer(&size)),
		0x3000, // MEM_COMMIT | MEM_RESERVE
		0x40,   // PAGE_EXECUTE_READWRITE
	)

	if status != 0 {
		fmt.Println("Memory allocation failed.")
		return nil
	}

	fmt.Println("Memory allocated at:", baseAddr)
	return unsafe.Pointer(baseAddr)
}
