package anti_debugging

import (
	"fmt"
	"syscall"
)

// DetectDebugger checks for attached debuggers
func DetectDebugger() bool {
	fmt.Println("Checking for debuggers...")

	kernel32 := syscall.NewLazyDLL("kernel32.dll")
	isDebuggerPresent := kernel32.NewProc("IsDebuggerPresent")

	ret, _, _ := isDebuggerPresent.Call()
	if ret != 0 {
		fmt.Println("Debugger detected!")
		return true
	}

	fmt.Println("No debugger detected.")
	return false
}
