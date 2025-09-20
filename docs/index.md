# Gitzy CLI Documentation

Welcome to the official documentation for **Gitzy**, your cross-platform CLI for managing and switching between multiple GitHub identities. This guide will walk you through installation, setup, usage, and troubleshooting, ensuring you get the most out of Gitzy.

---

## 🚀 What is Gitzy?

Gitzy helps developers maintain separate work, personal, and alternate GitHub profiles, and switch between them with a single command. It works on Windows, Linux, and macOS, and is designed for security, portability, and ease of use.

---

## 📦 Installation

### 1. Download from GitHub Releases (Recommended)

1. Go to the [Gitzy Releases page](https://github.com/gh0stfrk/gitzy/releases)
2. Download the binary for your platform:
   - **Windows**: `gitzy-windows-amd64.zip`

   - **Linux**: `gitzy-linux-amd64.zip`
   - **macOS**: `gitzy-darwin-amd64.zip`
3. Extract and run the installer for your OS:

   - **Windows**: Open PowerShell in the extracted folder and run `.\install_windows.ps1`
   - **Linux/macOS**: Run `chmod +x install_linux.sh && ./install_linux.sh`

### 2. Install with Go

```bash
go install github.com/gh0stfrk/gitzy@latest
```

### 3. Build from Source

```bash
git clone https://github.com/gh0stfrk/gitzy.git
cd gitzy
make build
```

### 4. Verify Installation

```bash
gitzy --help
```

---

## ⚡ Quick Start

1. **Add a new profile**
   ```bash
   gitzy add work
   ```

2. **Switch to a profile**
   ```bash
   gitzy switch work
   ```

3. **List all profiles**
   ```bash
   gitzy list
   ```

4. **Remove a profile**
   ```bash
   gitzy wipe work
   ```

---

## 📚 Command Reference

- [Main Command](cli/gitzy.md)
- [Add Profile](cli/gitzy_add.md)
- [Switch Profile](cli/gitzy_switch.md)
- [List Profiles](cli/gitzy_list.md)
- [Remove Profile](cli/gitzy_wipe.md)
- [Troubleshooting](cli/gitzy_doctor.md)

---

## 🔧 Configuration & Storage

Profiles and tokens are stored in:
- **Windows**: `%USERPROFILE%\.gitzy`
- **Linux/macOS**: `~/.gitzy` (or `XDG_CONFIG_HOME`)
- **Custom**: Set `GITZY_HOME` environment variable

**Security Note:** Tokens are stored unencrypted. Use minimal scopes and follow [GitHub token best practices](https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/creating-a-personal-access-token).

---

## 🆘 Troubleshooting

If you encounter issues, use:
```bash
gitzy doctor
```
Or consult the [Troubleshooting Guide](cli/gitzy_doctor.md).

---

## 📖 More Resources

- [Credential Management](credential-helper.md)
- [README](../README.md)
- [GitHub Repository](https://github.com/gh0stfrk/gitzy#readme)
# gitzy CLI Documentation

Welcome to the gitzy CLI documentation!

## About

gitzy is a cross-platform command-line tool for managing and switching between multiple GitHub identities. It helps developers easily maintain separate work, personal, and alternate GitHub profiles, and switch between them with a single command. gitzy is open source, works on Windows, Linux, and macOS, and is designed for security, portability, and ease of use.

## 🚀 Quick Navigation

- [📥 Installation Guide](#installation)
- [⚡ Quick Start](#quick-start)
- [📚 Command Reference](#command-reference)
- [🔧 Configuration](#configuration)

## 📥 Installation

### Option 1: Install from GitHub Releases (Recommended)

1. Go to the [Releases page](https://github.com/gh0stfrk/gitzy/releases)
2. Download the appropriate binary for your platform:

| Platform | File to Download |
|----------|------------------|
| Windows  | `gitzy-windows-amd64.zip` |
| Linux    | `gitzy-linux-amd64.zip` |
| macOS    | `gitzy-darwin-amd64.zip` |

3. Extract and install:

#### 🪟 Windows
```powershell
# Extract the ZIP file and run in PowerShell:
.\install_windows.ps1
```

#### 🐧 Linux
```bash
wget https://github.com/gh0stfrk/gitzy/releases/latest/download/gitzy-linux-amd64.zip
unzip gitzy-linux-amd64.zip -d gitzy
cd gitzy
chmod +x install_linux.sh
./install_linux.sh
```

#### 🍎 macOS
```bash
# Extract the ZIP file and run:
chmod +x install_linux.sh
./install_linux.sh
```

### Option 2: Install with Go

```bash
go install github.com/gh0stfrk/gitzy@latest
```

### Option 3: Build from Source

```bash
git clone https://github.com/gh0stfrk/gitzy.git
cd gitzy
make build
```

### Verify Installation

```bash
gitzy --help
```

## ⚡ Quick Start

1. **Add a new profile**:
   ```bash
   gitzy add work
   ```

2. **Switch to a profile**:
   ```bash
   gitzy switch work
   ```

3. **List all profiles**:
   ```bash
   gitzy list
   ```

4. **Remove a profile**:
   ```bash
   gitzy wipe work
   ```

## 📚 Command Reference

- [gitzy](cli/gitzy.md) - Main command overview
- [gitzy add](cli/gitzy_add.md) - Add a new profile
- [gitzy switch](cli/gitzy_switch.md) - Switch between profiles
- [gitzy list](cli/gitzy_list.md) - List all profiles
- [gitzy wipe](cli/gitzy_wipe.md) - Remove a profile
- [gitzy doctor](cli/gitzy_doctor.md) - Troubleshooting and diagnostics

## 🔧 Configuration

### Storage Locations

- **Windows**: `%USERPROFILE%\.gitzy`
- **Linux/macOS**: `~/.gitzy` (respects `XDG_CONFIG_HOME` if set)
- **Custom**: Set `GITZY_HOME` environment variable

### Security Notes

- Tokens are stored unencrypted in the configuration directory
- Follow [GitHub token best practices](https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/creating-a-personal-access-token)
- Consider using tokens with minimal required scopes

---

📖 For more details, see the [README](../README.md) or [view on GitHub](https://github.com/gh0stfrk/gitzy#readme).
