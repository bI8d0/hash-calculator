# Hash Calculator 🔐
![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)
![Go Version](https://img.shields.io/badge/Go-1.23-00ADD8?logo=go&logoColor=white)
![Platform](https://img.shields.io/badge/Platform-Windows%20%7C%20Linux-blue?logo=windows&logoColor=white)
![Build](https://img.shields.io/badge/Build-Passing-brightgreen)

A command-line tool written in Go to calculate and verify file integrity using MD5, SHA1, and SHA256 hashes.

**Author:** bI8d0  
**Version:** 1.0  
**License:** See LICENSE file

## 📋 Features

- ✅ Simultaneous calculation of **MD5**, **SHA1**, and **SHA256** hashes
- ✅ Automatic hash verification
- ✅ User-friendly interactive interface with ANSI color codes
- ✅ Hash validation and formatting
- ✅ Cross-platform support (Windows and Linux)
- ✅ No external dependencies (Go standard library only)
- ✅ Optimized binary builds with reduced file size

## 🚀 Quick Start

### Prerequisites

- **Go 1.24** or higher
- Operating System: Windows or Linux

### Installation

#### Option 1: Download Pre-built Binaries

Visit the [Releases](https://github.com/bI8d0/hash-calculator/releases) page to download:
- `hash-calculator-windows-amd64.zip` for Windows
- `hash-calculator-linux-amd64.zip` for Linux
- `hash-calculator.deb` package for Debian/Ubuntu systems

#### Option 2: Build from Source

1. **Clone the repository:**
```bash
git clone https://github.com/bI8d0/hash-calculator.git
cd hash-calculator
```

2. **Build the project:**

#### Option A: Build script (Recommended)
```bash
go run build.go
```

This will generate optimized executables for:
- Windows (hash-calculator.exe)
- Linux (hash-calculator)

Binaries will be saved in the `build/` folder

#### Option B: Manual compilation

**For Windows:**
```bash
go build -ldflags="-s -w" -trimpath -o build/hash-calculator.exe main.go
```

**For Linux:**
```bash
GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -trimpath -o build/hash-calculator main.go
```

**For macOS:**
```bash
GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -trimpath -o build/hash-calculator main.go
```

### Building on Linux (Creating .deb Package)

#### Prerequisites:
```bash
sudo apt-get update
sudo apt-get install build-essential devscripts debhelper dh-golang golang-go
```

#### Build the .deb package:
```bash
git clone https://github.com/bI8d0/hash-calculator.git
cd hash-calculator
dpkg-buildpackage -us -uc
```

The `.deb` package will be created in the parent directory.

#### Install the package:
```bash
sudo dpkg -i hash-calculator_1.0_amd64.deb
```

#### Verify installation:
```bash
hash-calculator -h
```

## 💻 Usage

### Basic Syntax

```bash
hash-calculator -f <file_path>
```

### Arguments

| Argument | Description | Required |
|----------|-------------|----------|
| `-f` | Full path of the file to analyze | ✅ Yes |
| `-h` | Show help message | ❌ No |

### Usage Examples

#### 1. Display help:
```bash
hash-calculator -h
```

#### 2. Calculate hashes of a file:
```bash
hash-calculator -f "file.exe"
```

**Expected output:**
```
    ╔═══════════════════════════════╗
    ║        HASH CALCULATOR        ║
    ║  File Integrity Verification  ║
    ║         by bI8d0 vx.x         ║
    ╚═══════════════════════════════╝

Calculating hashes...

File: file.exe

MD5:    5d41402abc4b2a76b9719d911017c592
SHA1:   aaf4c61ddcc5e8a2dabede0f3b482cd9aea9434d
SHA256: 2c26b46911185131006cba356cb9c25e12bef956258dafd62b57dad280e7214f

HASH: 
```

#### 3. Verify a hash:
After viewing the hashes, you can enter one for verification:
```
HASH: 2c26b46911185131006cba356cb9c25e12bef956258dafd62b57dad280e7214f
```

If it matches:
```
✓ VERIFICATION SUCCESSFUL: 🔐 Hash matches SHA256 🔐
```

If it doesn't match:
```
✗ VERIFICATION FAILED: ⚠️ The provided hash does not match any of the calculated ones ⚠️
```

## 🔧 Internal Operation

### Architecture

The project consists of two main files:

#### `main.go` - Core Logic
- **`calculateHashes(filePath string)`**: Simultaneously calculates MD5, SHA1, and SHA256 using `io.MultiWriter` for maximum efficiency
- **`clearScreen()`**: Clears the console in a cross-platform way (Windows/Linux)
- **`printLogo()`**: Displays the program's ASCII logo
- **`printUsage()`**: Shows the usage guide
- **`printSuccess()` and `printError()`**: Helper functions to print with ANSI colors
- **`main()`**: Orchestrates the program's main flow

#### `build.go` - Build Script
Automates compilation for multiple platforms with optimizations:
- Debug symbols removal (`-ldflags="-s -w"`)
- Compiler paths removal (`-trimpath`)
- Generates smaller and safer binaries

### Execution Flow

```
┌─────────────────────────────┐
│  Program startup            │
│  Display logo               │
└────────────┬────────────────┘
             │
             ▼
┌─────────────────────────────┐
│  Parse CLI arguments        │
│  (-f file, -h help)         │
└────────────┬────────────────┘
             │
             ├─ -h requested? ──► Show help ──► Exit
             │
             ├─ -f empty? ──► Show help ──► Exit with error
             │
             ▼
┌─────────────────────────────┐
│  File exists?               │
└────────────┬────────────────┘
             │
             ├─ NO ──► Show error ──► Exit
             │
             ▼
┌─────────────────────────────┐
│  Calculate hashes           │
│  (MD5, SHA1, SHA256)        │
└────────────┬────────────────┘
             │
             ▼
┌─────────────────────────────┐
│  Display results            │
│  Request user input         │
└────────────┬────────────────┘
             │
             ▼
┌─────────────────────────────┐
│  Hash entered?              │
└────────────┬────────────────┘
             │
             ├─ NO ──► Finish
             │
             ▼
┌─────────────────────────────┐
│  Verify hash                │
│  (Compare with calculated)  │
└────────────┬────────────────┘
             │
             ├─ Match? ──► Show success ✓
             │
             └─ NO ──► Show failure ✗
```

### Supported Algorithms

| Algorithm | Characters | Security | Use |
|-----------|-----------|----------|-----|
| **MD5** | 32 hex | ❌ Weak | Compatibility only |
| **SHA1** | 40 hex | ⚠️ Deprecated | Avoid |
| **SHA256** | 64 hex | ✅ Recommended | Production |

### Implemented Optimizations

1. **MultiWriter**: Calculates all three hashes in a single pass through the file
2. **Streaming**: Processes large files without loading them completely into memory
3. **Optimized compilation**: Flags to minimize size and antivirus false positives

## 📦 Project Structure

```
hash-calculator/
├── main.go              # Main code
├── build.go             # Build script
├── go.mod              # Dependencies (no external dependencies)
├── go.sum              # Checksums (empty)
├── README.md           # This file
├── LICENSE             # License information
├── debian/             # Debian packaging files
│   ├── changelog
│   ├── compat
│   ├── control
│   ├── copyright
│   └── rules
└── build/              # Compiled binaries
 
```

## 🛠️ Development

### Clone and prepare the environment

```bash
git clone https://github.com/your-username/hash-calculator.git
cd hash-calculator
```

### Build for development

```bash
go build -o hash-calculator main.go
```

### Run tests

```bash
# Create a test file
echo "test" > test.txt

# Run the program
./hash-calculator -f test.txt
```

## 📋 Use Cases

### 1. Verify Download Integrity
```bash
hash-calculator -f "download.iso"
# Compare the SHA256 with the one published on the web
```

### 2. Validate Backups
```bash
hash-calculator -f "backup.zip"
# Save hashes for future verification
```

### 3. Detect File Modifications
```bash
hash-calculator -f "file"
# Compare with previous hash to detect changes
```

## ⚠️ Important Notes

- Hashes are automatically formatted to **UPPERCASE**
- Whitespace in provided hashes is ignored
- The program **always calculates all three hash types** (MD5, SHA1, SHA256)
- For maximum security, use **SHA256**
- You can press Enter without entering a hash if you just want to see the values

## 🔒 Security

- No external dependencies (only Go standard library)
- Compiled with optimization flags to minimize antivirus false positives
- Native ANSI codes, no external color libraries
- Robust user input validation

## 📝 Changelog

### Version 1.0 (2026-02-20)
- Initial release
- Hash calculation: MD5, SHA1, SHA256
- Interactive verification
- Cross-platform support (Windows, Linux, macOS)
- Debian package support

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🐛 Troubleshooting

### Issue: "File not found"
- **Solution:** Use the absolute file path. On Windows, use double backslashes or quotes:
  ```bash
  hash-calculator -f "file.exe"
  ```

### Issue: Permission denied on Linux
- **Solution:** Add execute permissions:
  ```bash
  chmod +x hash-calculator
  ```

### Issue: ".deb package installation failed"
- **Solution:** Ensure Go 1.21+ is installed:
  ```bash
  go version
  ```

## 📞 Support

For issues, questions, or suggestions, please open an [Issue](https://github.com/bI8d0/hash-calculator/issues) on GitHub.

---

**Made with ❤️ by bI8d0**

