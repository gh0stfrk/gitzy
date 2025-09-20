# Gitzy Installation Guide

Follow these steps to install Gitzy on Windows, Linux, or macOS.

---

## 📥 Download from GitHub Releases

1. Visit the [Gitzy Releases page](https://github.com/gh0stfrk/gitzy/releases)
2. Download the ZIP file for your platform:
   - **Windows**: `gitzy-windows-amd64.zip`
   - **Linux**: `gitzy-linux-amd64.zip`
   - **macOS**: `gitzy-darwin-amd64.zip`

---

## 🖥️ Windows Installation

1. **Extract the ZIP file**
   - Right-click and choose **Extract All**
2. **Run the installer**
   - Open PowerShell in the extracted folder
   - Run:
     ```powershell
     .\install_windows.ps1
     ```
   - This will copy `gitzy.exe` to a directory in your user `PATH`
3. **Verify installation**
   - Open Command Prompt and run:
     ```cmd
     gitzy --help
     ```

---

## 🐧 Linux Installation

1. **Extract the ZIP file**
   ```bash
   unzip gitzy-linux-amd64.zip -d gitzy
   cd gitzy
   ```
2. **Run the installer**
   ```bash
   chmod +x install_linux.sh
   ./install_linux.sh
   ```
3. **Verify installation**
   ```bash
   gitzy --help
   ```

---

## 🍏 macOS Installation

1. **Extract the ZIP file**
   ```bash
   unzip gitzy-darwin-amd64.zip -d gitzy
   cd gitzy
   ```
2. **Run the installer**
   ```bash
   chmod +x install_linux.sh
   ./install_linux.sh
   ```
3. **Verify installation**
   ```bash
   gitzy --help
   ```

---

## 🛠️ Alternative Installation Methods

### Install with Go
```bash
go install github.com/gh0stfrk/gitzy@latest
```

### Build from Source
```bash
git clone https://github.com/gh0stfrk/gitzy.git
cd gitzy
make build
```

---

For troubleshooting or advanced configuration, see the [main documentation](index.md) and [credential management guide](credential-helper.md).