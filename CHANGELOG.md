# Changelog

All notable changes to Hash Calculator are documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/),
and this project adheres to [Semantic Versioning](https://semver.org/).

---

## [Unreleased]

### Planned Features
- Additional hash algorithms (BLAKE3, CRC32)
- Batch file processing
- Configuration file support
- Integration with file managers (context menu)

---

## [1.0] - 2026-02-20

### Added
- Initial release of Hash Calculator
- Simultaneous MD5, SHA1, and SHA256 hash calculation
- Interactive hash verification mode
- User-friendly command-line interface with ANSI colors
- Cross-platform support:
  - Windows (32-bit and 64-bit)
  - Linux (amd64)
  - macOS (amd64)
- Debian package (.deb) for easy Linux installation
- Comprehensive documentation:
  - README
  - Build guide for creating .deb packages
  - GitHub setup and contribution guidelines
- Optimized binary builds with size reduction (~50%)
- Error handling for missing and corrupted files
- Hash validation with automatic format normalization
- Support for files with spaces in names
- Cross-platform shell support (Windows CMD, Linux Bash)

### Technical Details
- Written in Go 1.24
- No external dependencies (Go standard library only)
- Optimized with `-ldflags="-s -w"` and `-trimpath` flags
- Reproducible builds
- Source code in English with detailed comments

### Known Limitations
- MD5 and SHA1 included for compatibility (not recommended for security-critical applications)
- Windows binary may trigger antivirus alerts (false positive due to packing)
- Linux build requires Go 1.21 or higher

---

## Version History

### Why Only Version 1.0?

This is the initial release of Hash Calculator. Future versions will include:

#### Version 1.1 (Planned)
- BLAKE3 hash algorithm support
- Configuration file for default settings
- Batch file processing capability
- Performance improvements
- Additional error handling for edge cases
---

## Release Notes Template for New Releases

Use this template when creating new releases:

```
## [X.X] - YYYY-MM-DD

### Added
- New features added

### Changed
- Modifications to existing features

### Fixed
- Bug fixes

### Removed
- Deprecated features removed

### Security
- Security-related fixes
```

---

## How to Report Issues

Found a bug? Have a feature request? 

1. Check [existing issues](https://github.com/bI8d0/hash-calculator/issues)
2. Click "New Issue"
3. Choose "Bug Report" or "Feature Request"
4. Provide detailed information:
   - Operating system and version
   - Go version (if building from source)
   - Steps to reproduce
   - Expected vs actual behavior
   - Screenshots or error messages

---

## Deprecated & Removed

### Version 1.0 Notes
- MD5 and SHA1 are included for backward compatibility but are considered cryptographically weak
- For security-critical applications, use SHA256 instead
- MD5 and SHA1 support may be removed in a future major version

---

## Migration Guide

### From Pre-Release to 1.0

No migration needed! This is the initial public release.

---

## Support

For support, visit:
- [GitHub Issues](https://github.com/bI8d0/hash-calculator/issues)
---

**Last Updated:** 2026-02-27
**Maintained by:** bI8d0

