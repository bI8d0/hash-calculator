# Hash Calculator 🔐

A command-line tool written in Go to calculate and verify file integrity using MD5, SHA1, and SHA256 hashes.

**Author:** bI8d0

## 📋 Features

- ✅ Simultaneous calculation of **MD5**, **SHA1**, and **SHA256**
- ✅ Automatic hash verification provided by the user
- ✅ User-friendly interface with ANSI color codes
- ✅ Automatic hash validation and formatting
- ✅ Compatible with Windows and Linux
- ✅ No external dependencies (only Go standard library)

## 🚀 Quick Start

### Prerequisites

- **Go 1.24** or higher
- Operating System: Windows or Linux

### Installation

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
    ║           by bI8d0            ║
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
├── debian/             # Debian packaging files
│   ├── changelog
│   ├── compat
│   ├── control
│   ├── copyright
│   └── rules
└── build/              # Compiled binaries
    ├── hash-calculator
    └── hash-calculator.exe
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
hash-calculator -f "file.exe"
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

## 📄 License

This project is provided as-is. See the `debian/copyright` file for more details.

## 👨‍💻 Contributions

Contributions are welcome. For major changes:

1. Fork the repository
2. Create a branch for your feature (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## 📞 Support

If you encounter issues:

1. Verify that Go 1.24+ is installed: `go version`
2. Make sure the file exists and is accessible
3. On Windows, use quoted paths if they contain spaces
4. Check file read permissions

## 🎯 Roadmap

- [ ] Support for more algorithms (SHA512, BLAKE3)
- [ ] Export results to JSON/CSV
- [ ] REST API
- [ ] Graphical interface (GUI)
- [ ] Batch processing of multiple files

---

**Last Updated:** February 19, 2026  
**Version:** 1.0  
**Status:** Production Ready ✅

