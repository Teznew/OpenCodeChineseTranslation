# OpenCode Chinese Translation Distribution

[![Release](https://img.shields.io/github/v/release/Teznew/OpenCodeChineseTranslation?label=Latest&style=flat-square&color=blue)](https://github.com/Teznew/OpenCodeChineseTranslation/releases/latest)
[![Platform](https://img.shields.io/badge/platform-Windows%20%7C%20macOS%20%7C%20Linux-lightgrey.svg?style=flat-square)](#)
[![Build Status](https://img.shields.io/github/actions/workflow/status/Teznew/OpenCodeChineseTranslation/release.yml?label=Daily%20Build&style=flat-square)](https://github.com/Teznew/OpenCodeChineseTranslation/actions)
[![License](https://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)](LICENSE)

[中文文档](README.md)

> 🚀 **OpenCode Chinese Distribution | ⚡️ Daily Sync with Official | Automated Builds for Windows x64 / Linux x64 / macOS**

---

## Overview

**OpenCode Chinese Translation** is a fully automated localization project for [OpenCode](https://github.com/anomalyco/opencode). We've built a complete CI/CD pipeline using GitHub Actions that runs daily at **01:16**, pulls the latest source code, applies Chinese translation patches, and builds installation packages for Windows x64, Linux x64, and macOS (Intel/Apple Silicon).

**Key Features:**
*   ⚡️ **Daily Auto-Updates**: Stay up-to-date with the latest official features.
*   📦 **Targeted Platform Support**: Provides Windows x64, Linux x64, and macOS (Intel/Apple Silicon) binaries.
*   🚀 **Zero-Dependency Installation**: New Go-based CLI tool, no Node.js or Bun required.
*   🔧 **Complete Localization**: Covers TUI, dialogs, and core workflows.

---

## Quick Start

### 1. One-Line Installation (Recommended)

The new installation scripts download the **Go-based CLI tool** directly, requiring no runtime dependencies.

**Windows (PowerShell)**
```powershell
powershell -c "irm https://cdn.jsdelivr.net/gh/Teznew/OpenCodeChineseTranslation@main/install.ps1 | iex"
```

**Linux / macOS**
```bash
curl -fsSL https://cdn.jsdelivr.net/gh/Teznew/OpenCodeChineseTranslation@main/install.sh | bash
```

### 2. Usage

After installation, run in your terminal:

```bash
opencode-cli
```

This launches the interactive menu.

### 3. Download Prebuilt Version (New in v8.1+)

If you already have `opencode-cli` installed, use the built-in download feature:

```bash
opencode-cli download
```

This automatically downloads the latest prebuilt Chinese version from GitHub Releases, no local compilation needed.

### 4. Manual Download

You can also visit the [Releases page](https://github.com/Teznew/OpenCodeChineseTranslation/releases/latest) to download binaries directly.

| Platform | CLI Tool |
|----------|----------|
| Windows x64 | `opencode-cli-windows-amd64.exe` |
| macOS Apple Silicon | `opencode-cli-darwin-arm64` |
| macOS Intel | `opencode-cli-darwin-amd64` |
| Linux x64 | `opencode-cli-linux-amd64` |

> Chinese OpenCode binaries are available as ZIP packages on the [Releases page](https://github.com/Teznew/OpenCodeChineseTranslation/releases/latest).

---

## CLI Commands

The CLI tool (v8.6.0) provides comprehensive management capabilities:

| Command | Description |
|---------|-------------|
| `opencode-cli` | Launch interactive menu (default) |
| `opencode-cli download` | Download prebuilt Chinese version (no build env required) |
| `opencode-cli env-install` | **One-click install** build environment (Git/Node.js/Bun) |
| `opencode-cli update` | Update OpenCode source code |
| `opencode-cli apply` | Apply translation patches |
| `opencode-cli verify` | Verify translation configuration |
| `opencode-cli build` | Build OpenCode binary |
| `opencode-cli deploy` | Deploy to system PATH |
| `opencode-cli diagnose` | **Diagnose** and fix conflicts/issues |
| `opencode-cli uninstall` | Uninstall and clean up all files |
| `opencode-cli antigravity` | Configure Antigravity local AI proxy |

---

## Developer Guide

If you want to contribute, please refer to the [Contributing Guide](CONTRIBUTING.md).

*   [📅 Changelog](CHANGELOG.md)
*   [🚀 Antigravity Integration Guide](docs/ANTIGRAVITY_INTEGRATION.md)

---

## FAQ

**Having issues? Run diagnose first!**
```bash
opencode-cli diagnose --fix
```
Auto-detects and fixes version conflicts, missing dependencies, and PATH issues.

**Q: Still shows English after install?**
Run `opencode-cli diagnose` to detect and clean up conflicting versions.

**Q: Build failed?**
```bash
opencode-cli env-install    # Install Git/Node/Bun
opencode-cli download       # Or download prebuilt (no build env needed)
```

**Q: How to completely uninstall?**
```bash
opencode-cli uninstall --all
```

**Q: macOS "cannot verify developer"?**
Run `opencode-cli diagnose --fix` or manually: `xattr -cr /path/to/opencode`

**Q: Where are files installed?**
Unified directory structure: `~/.opencode-i18n/`
- `bin/` - CLI tool and Chinese OpenCode
- `opencode/` - OpenCode source code
- `build/` - Build output

Windows actual path: `%USERPROFILE%\.opencode-i18n\`

**Q: Local development setup?**
Developers can customize paths via environment variables:
```bash
export OPENCODE_SOURCE_DIR=/path/to/opencode   # Source directory
export OPENCODE_BUILD_DIR=/path/to/bin         # Build output
export OPENCODE_PROJECT_DIR=/path/to/project   # Translation project
```
Or create `opencode/` and `bin/` folders in the project directory - CLI auto-detects them.

---

## License

This project is open-sourced under the [MIT License](LICENSE).
The original OpenCode project is copyright [Anomaly Company](https://anomaly.company/).
