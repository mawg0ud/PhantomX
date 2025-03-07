package build

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

// BuildConfig holds compilation settings
type BuildConfig struct {
	OutputName string
	GOOS       string
	GOARCH     string
	Flags      []string
}

// DefaultConfig defines default build settings
var DefaultConfig = BuildConfig{
	OutputName: "phantomx",
	GOOS:       runtime.GOOS,   // Use system's OS
	GOARCH:     runtime.GOARCH, // Use system's architecture
	Flags:      []string{"-ldflags", "-s -w"}, // Stripped binary for stealth
}

// CompileProject compiles the PhantomX project with the given settings
func CompileProject(config BuildConfig) error {
	fmt.Println("Starting PhantomX compilation...")

	// Set environment variables for cross-compilation
	os.Setenv("GOOS", config.GOOS)
	os.Setenv("GOARCH", config.GOARCH)

	// Construct the Go build command
	args := []string{"build", "-o", config.OutputName}
	args = append(args, config.Flags...)

	cmd := exec.Command("go", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println("Compilation failed:", err)
		return err
	}

	fmt.Println("Compilation successful. Output:", config.OutputName)
	return nil
}

func main() {
	// Execute the default build process
	if err := CompileProject(DefaultConfig); err != nil {
		os.Exit(1)
	}
}
