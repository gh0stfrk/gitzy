# 🛠️ Gitzy CLI Installation Guide

Gitzy is a cross-platform CLI tool for switching between multiple Git credentials easily. This guide will help you download and install the latest release on **Windows**, **Linux**, or **macOS**.

---

## 📥 Step 1: Visit the GitHub Releases Page

1. Go to the [Gitzy GitHub Releases page](https://github.com/gh0stfrk/gitzy/releases).
2. Find the latest release at the top of the page.
3. Download the appropriate ZIP file for your OS and architecture:

| Platform | File to Download |
|----------|------------------|
| Windows  | `gitzy-windows-amd64.zip` |
| Linux    | `gitzy-linux-amd64.zip` |
| macOS    | `gitzy-darwin-amd64.zip` |

---

## 🪟 Windows Installation

1. **Download & Extract**:
   - Download `gitzy-windows-amd64.zip` from the [Releases page](https://github.com/yourusername/gitzy/releases).
   - Right-click the ZIP file and choose **Extract All**.

2. **Run the Installer**:
   - Inside the extracted folder, right-click and select "Open PowerShell window here"
   - Run: `.\install_windows.ps1`
   - This will copy `gitzy.exe` to a directory in your user `PATH`

3. **Verify Installation**:
   Open **Command Prompt** and run:
   ```cmd
   gitzy --help
   ```

## 🐧 Linux Installation

### ✅ Manual Installation

1. **Download & Extract**
   ```bash
   wget https://github.com/yourusername/gitzy/releases/latest/download/gitzy-linux-amd64.zip
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