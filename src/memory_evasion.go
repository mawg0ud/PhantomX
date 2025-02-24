package memory_evasion

import (
	"fmt"
	"syscall"
	"unsafe"
)

// ExecuteEvasion runs memory evasion techniques
func ExecuteEvasion() {
	fmt.Println("Executing memory evasion techniques...")

	// Example technique: Unhooking ntdll.dll
	UnhookNtdll()
}

// UnhookNtdll reloads a fresh copy of ntdll.dll to evade hooked APIs
func UnhookNtdll() {
	kernel32 := syscall.NewLazyDLL("kernel32.dll")
	loadLibrary := kernel32.NewProc("LoadLibraryA")

	ntdll := "C:\\Windows\\System32\\ntdll.dll"
	_, _, _ = loadLibrary.Call(uintptr(unsafe.Pointer(syscall.StringBytePtr(ntdll))))

	fmt.Println("Successfully unhooked ntdll.dll")
}
