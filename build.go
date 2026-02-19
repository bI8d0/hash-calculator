package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	// Get the current directory
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting current directory: %v\n", err)
		return
	}

	// Create the build directory if it doesn't exist
	buildDir := filepath.Join(currentDir, "build")
	if err := os.MkdirAll(buildDir, 0755); err != nil {
		fmt.Printf("Error creating build directory: %v\n", err)
		return
	}

	// Executable name
	executableName := "hash-calculator"

	// Name of the file to compile
	filename := "main.go"

	// Build for Windows
	fmt.Println("Building for Windows...")
	if err := buildForWindows(filepath.Join(buildDir, executableName+".exe"), filename, currentDir); err != nil {
		fmt.Printf("Error building for Windows: %v\n", err)
	}

	// Build for Linux
	fmt.Println("Building for Linux...")
	if err := buildForOS("linux", filepath.Join(buildDir, executableName), filename, currentDir); err != nil {
		fmt.Printf("Error building for Linux: %v\n", err)
	}

	fmt.Println("Build completed. Binaries are located in the 'build' directory.")
}

func buildForWindows(outputName, filename, currentDir string) error {
	// Build for Windows with optimized flags to reduce antivirus false positives
	cmd := exec.Command("go", "build",
		"-ldflags=-s -w", // Remove debug symbols
		"-trimpath",      // Remove compiler paths
		"-o", outputName,
		filename,
	)
	cmd.Env = append(os.Environ(),
		"GOOS=windows",
		"GOARCH=amd64",
	)
	cmd.Dir = currentDir

	buildOutput, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("build error: %v\nOutput: %s", err, buildOutput)
	}

	fmt.Printf("✓ Windows build completed.\n")

	// Verify that the executable was created
	if _, err := os.Stat(outputName); os.IsNotExist(err) {
		return fmt.Errorf("Windows executable was not generated correctly")
	}

	return nil
}

func buildForOS(goos, outputName, filename string, currentDir string) error {
	cmd := exec.Command("go", "build",
		"-ldflags=-s -w", // Remove debug symbols
		"-trimpath",      // Remove compiler paths
		"-o", outputName,
		filename,
	)
	cmd.Env = append(os.Environ(),
		"GOOS="+goos,
		"GOARCH=amd64",
	)
	cmd.Dir = currentDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("build error for %s: %v", goos, err)
	}

	fmt.Printf("✓ Build for %s completed.\n", goos)

	// Verify that the executable was created
	if _, err := os.Stat(outputName); os.IsNotExist(err) {
		return fmt.Errorf("executable for %s was not generated correctly", goos)
	}

	return nil
}
