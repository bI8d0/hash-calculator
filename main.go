package main

import (
	"bufio"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

// ANSI color codes
const (
	ColorGreen = "\033[92m" // Green
	ColorRed   = "\033[91m" // Red
	ColorReset = "\033[0m"  // Reset to normal color
)

// clearScreen clears the console on Windows and Linux
func clearScreen() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func calculateHashes(filePath string) (md5Hash, sha1Hash, sha256Hash string, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", "", "", err
	}
	defer func() {
		if closeErr := file.Close(); closeErr != nil && err == nil {
			err = closeErr
		}
	}()

	// Create three hashers
	md5Hasher := md5.New()
	sha1Hasher := sha1.New()
	sha256Hasher := sha256.New()

	// Create a multi-writer to write to all three hashers simultaneously
	multiWriter := io.MultiWriter(md5Hasher, sha1Hasher, sha256Hasher)

	// Copy the file content to the hashers
	if _, err := io.Copy(multiWriter, file); err != nil {
		return "", "", "", err
	}

	// Get hashes in hexadecimal format
	md5Hash = fmt.Sprintf("%x", md5Hasher.Sum(nil))
	sha1Hash = fmt.Sprintf("%x", sha1Hasher.Sum(nil))
	sha256Hash = fmt.Sprintf("%x", sha256Hasher.Sum(nil))

	return md5Hash, sha1Hash, sha256Hash, nil
}

func printLogo() {
	fmt.Println(`
    ╔═══════════════════════════════╗
    ║        HASH CALCULATOR        ║
    ║  File Integrity Verification  ║
    ║         by bI8d0 v1.0         ║
    ╚═══════════════════════════════╝`)
}

func printUsage() {
	clearScreen()
	fmt.Println(`
╔════════════════════════════════════════════════════════════════╗
║                         USAGE GUIDE                            ║
╚════════════════════════════════════════════════════════════════╝

DESCRIPTION:
  Calculates and verifies file integrity using hashes
  MD5, SHA1 and SHA256.

SYNTAX:
  hash.exe -f <file_path>

REQUIRED ARGUMENTS:
  -f    Full path of the file to analyze
        Example: -f "file_path"

OPTIONAL ARGUMENTS:
  -h    Show this help message
        Example: -h

USAGE EXAMPLES:
  1. Show help:
     hash.exe -h

  2. Calculate file hashes:
     hash.exe -f "file_path"
     Then enter the hash in the prompt to verify (optional)

IMPORTANT NOTES:
  • Hashes are automatically formatted to uppercase
  • Whitespace in the provided hashes is ignored
  • The program always calculates all three hash types (MD5, SHA1, SHA256)
  • For maximum security, use SHA256
  • You can press Enter without entering a hash if you only want to see the values

SUPPORTED ALGORITHMS:
  • MD5 (32 hexadecimal characters) - Weak, compatibility only
  • SHA1 (40 hexadecimal characters) - Deprecated, avoid
  • SHA256 (64 hexadecimal characters) - Recommended ⭐

════════════════════════════════════════════════════════════════
`)
}

func printSuccess(message string, args ...interface{}) {
	fmt.Printf(ColorGreen+message+ColorReset+"\n", args...)
}

func printError(message string, args ...interface{}) {
	fmt.Printf(ColorRed+message+ColorReset+"\n", args...)
}

func main() {
	printLogo()

	// Define command-line arguments
	filePath := flag.String("f", "", "Full path of the file to calculate hashes")
	showHelp := flag.Bool("h", false, "Show help message")

	// Parse command-line arguments
	flag.Parse()

	// If help is requested, show usage guide and exit
	if *showHelp {
		printUsage()
		os.Exit(0)
	}

	// Validate that file path argument was provided
	if *filePath == "" {
		printUsage()
		os.Exit(1)
	}

	// Check if the file exists before processing
	if _, err := os.Stat(*filePath); os.IsNotExist(err) {
		fmt.Printf("\nError: The file '%s' does not exist\n", *filePath)
		os.Exit(1)
	}

	fmt.Printf("\nCalculating hashes...\n")

	// Calculate hashes for the file using all three algorithms
	md5Hash, sha1Hash, sha256Hash, err := calculateHashes(*filePath)
	if err != nil {
		fmt.Printf("\nError calculating hashes: %v\n", err)
		os.Exit(1)
	}

	// Display the calculated hashes
	fmt.Printf("\nFile: %s\n\n", *filePath)
	fmt.Printf("MD5:    %s\n", md5Hash)
	fmt.Printf("SHA1:   %s\n", sha1Hash)
	fmt.Printf("SHA256: %s\n\n", sha256Hash)

	// Prompt user to enter a hash for verification (optional)
	fmt.Print("HASH: ")
	scanner := bufio.NewScanner(os.Stdin)
	var hashToVerify string
	if scanner.Scan() {
		hashToVerify = scanner.Text()
	}

	// If user entered a hash, attempt to verify it against calculated hashes
	if hashToVerify != "" {
		// Normalize hashes: remove whitespace and convert to uppercase for comparison
		hashToVerifyFormatted := strings.ToUpper(strings.TrimSpace(hashToVerify))
		hashToVerifyFormatted = strings.ReplaceAll(hashToVerifyFormatted, " ", "")

		md5HashFormatted := strings.ToUpper(strings.TrimSpace(md5Hash))
		md5HashFormatted = strings.ReplaceAll(md5HashFormatted, " ", "")

		sha1HashFormatted := strings.ToUpper(strings.TrimSpace(sha1Hash))
		sha1HashFormatted = strings.ReplaceAll(sha1HashFormatted, " ", "")

		sha256HashFormatted := strings.ToUpper(strings.TrimSpace(sha256Hash))
		sha256HashFormatted = strings.ReplaceAll(sha256HashFormatted, " ", "")

		// Compare provided hash with all calculated hashes
		matches := false
		matchType := ""

		if hashToVerifyFormatted == md5HashFormatted {
			matches = true
			matchType = "MD5"
		} else if hashToVerifyFormatted == sha1HashFormatted {
			matches = true
			matchType = "SHA1"
		} else if hashToVerifyFormatted == sha256HashFormatted {
			matches = true
			matchType = "SHA256"
		}

		// Display verification result
		if matches {
			printSuccess("\n✓ VERIFICATION SUCCESSFUL: 🔐 Hash matches %s 🔐", matchType)
		} else {
			printError("\n✗ VERIFICATION FAILED: ⚠️ The provided hash does not match any of the calculated ones ⚠️")
		}
	}
}
