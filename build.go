package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

// main orchestrates the cross-platform build process
// It builds optimized binaries for Windows and Linux systems
func main() {
	// Get the current working directory
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting current directory: %v\n", err)
		return
	}

	// Create the build output directory if it doesn't exist
	buildDir := filepath.Join(currentDir, "build")
	if err := os.MkdirAll(buildDir, 0755); err != nil {
		fmt.Printf("Error creating build directory: %v\n", err)
		return
	}

	// Define the executable name (without extension)
	executableName := "hash-calculator"

	// Source file to compile
	filename := "main.go"

	// Build Windows executable (amd64)
	fmt.Println("Building for Windows...")
	if err := buildForWindows(filepath.Join(buildDir, executableName+".exe"), filename, currentDir); err != nil {
		fmt.Printf("Error building for Windows: %v\n", err)
	}

	// Build Linux executable (amd64)
	fmt.Println("Building for Linux...")
	if err := buildForOS("linux", filepath.Join(buildDir, executableName), filename, currentDir); err != nil {
		fmt.Printf("Error building for Linux: %v\n", err)
	}

	fmt.Println("Build completed. Binaries are located in the 'build' directory.")
}

// buildForWindows compiles the application for Windows (amd64)
// Parameters:
//   - outputName: path where the executable will be saved
//   - filename: Go source file to compile
//   - currentDir: current working directory
//
// Returns: error if build fails
func buildForWindows(outputName, filename, currentDir string) error {
	// Create build command with optimization flags
	cmd := exec.Command("go", "build",
		"-ldflags=-s -w", // Strip debug symbols and DWARF info to reduce file size
		"-trimpath",      // Remove local file paths from binary for reproducibility
		"-o", outputName,
		filename,
	)
	// Set environment variables for cross-platform compilation
	cmd.Env = append(os.Environ(),
		"GOOS=windows", // Target operating system: Windows
		"GOARCH=amd64", // Target architecture: x86-64
	)
	cmd.Dir = currentDir

	// Execute build command and capture all output
	buildOutput, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("build error: %v\nOutput: %s", err, buildOutput)
	}

	fmt.Printf("✓ Windows build completed.\n")

	// Verify that the executable was successfully created
	if _, err := os.Stat(outputName); os.IsNotExist(err) {
		return fmt.Errorf("Windows executable was not generated correctly")
	}

	return nil
}

// buildForOS compiles the application for a specified operating system
// Parameters:
//   - goos: target operating system (linux, darwin, etc.)
//   - outputName: path where the executable will be saved
//   - filename: Go source file to compile
//   - currentDir: current working directory
//
// Returns: error if build fails
func buildForOS(goos, outputName, filename string, currentDir string) error {
	// Create build command with optimization flags
	cmd := exec.Command("go", "build",
		"-ldflags=-s -w", // Strip debug symbols and DWARF info to reduce file size
		"-trimpath",      // Remove local file paths from binary for reproducibility
		"-o", outputName,
		filename,
	)
	// Set environment variables for cross-platform compilation
	cmd.Env = append(os.Environ(),
		"GOOS="+goos,   // Target operating system
		"GOARCH=amd64", // Target architecture: x86-64
	)
	cmd.Dir = currentDir
	cmd.Stdout = os.Stdout // Display build output to console
	cmd.Stderr = os.Stderr // Display build errors to console

	// Execute the build command
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("build error for %s: %v", goos, err)
	}

	fmt.Printf("✓ Build for %s completed.\n", goos)

	// Verify that the executable was successfully created
	if _, err := os.Stat(outputName); os.IsNotExist(err) {
		return fmt.Errorf("executable for %s was not generated correctly", goos)
	}

	return nil
}
