
# PhantomX - Advanced EDR Evasion Framework

PhantomX is an advanced **EDR (Endpoint Detection and Response) evasion tool** built in **Golang**, designed to bypass modern security defenses using **memory evasion, shellcode encryption, direct syscalls, and anti-debugging techniques**.

## 1. Features

↦ **Memory Evasion** – Unhooks API functions and hides malicious execution.  
↦ **Shellcode Encryption** – Encrypts shellcode with **AES & XOR** to evade detection.  
↦ **Direct Syscalls** – Executes system calls **without API hooks**.  
↦ **Anti-Debugging** – Detects debuggers and sandboxes before execution.  
↦ **Cross-Platform** – Works on **Windows, Linux, and macOS**.  
↦ **Automated Compilation** – Supports **cross-compilation** with optimized binaries.  


## 2. Project Structure

```
PhantomX/
│── src/
│   │── main.go              # Main entry file
│   │── memory_evasion.go    # Handles memory evasion techniques
│   │── shellcode_loader.go  # Loads and executes shellcode
│   │── encryption.go        # Encrypts/decrypts shellcode
│   │── anti_debugging.go    # Implements anti-debugging & sandbox checks
│   │── syscalls.go          # Implements direct system calls
│   └── utils.go             # Utility functions (config loading, file handling)
│
│── shellcode/
│   │── reverse_shellcode.go  # Reverses shellcode before execution
│   └── shellcode_xor.go      # Shellcode encryption using XOR
│
│── build/
│   └── compile.go           # Compilation script
│
│── docs/
│   └── README.md            # Documentation
│
└── config/
    └── settings.json        # Configuration file for customization
```

## 3. Installation & Setup

### **🔹 Prerequisites**
- Install **Golang** (version 1.20+ recommended)
- Windows users: Install **MinGW** for cross-compilation
- Linux/macOS users: Ensure **gcc** is installed

### **🔹 Clone the Repository**
```sh
git clone https://github.com/mawg0ud/PhantomX.git
cd PhantomX
```

### **🔹 Compile PhantomX**
To compile the project for your current OS:
```sh
go build -o phantomx src/main.go
```

For cross-compilation (Windows binary from Linux/macOS):
```sh
GOOS=windows GOARCH=amd64 go build -o phantomx.exe src/main.go
```

For Linux binary from Windows:
```sh
GOOS=linux GOARCH=amd64 go build -o phantomx src/main.go
```

## 4. Usage

Run PhantomX with default settings:
```sh
./phantomx
```

Run with a custom configuration file:
```sh
./phantomx -config config/settings.json
```

To execute encrypted shellcode manually:
```sh
go run src/shellcode_loader.go -file shellcode/payload.bin
```


## 5. Legal Disclaimer
This tool is for **educational purposes** only. **Unauthorized use of this tool on third-party systems is illegal**. The developers are **not responsible** for any misuse.


## 6. License
PhantomX is released under the **MIT License**.


## 7 Future Enhancements
-  **Process Hollowing** – Injecting payload into legitimate processes.
-  **Polymorphic Shellcode** – Generate dynamic payloads on execution.
-  **Kernel-Level Evasion** – More advanced techniques to bypass monitoring.
