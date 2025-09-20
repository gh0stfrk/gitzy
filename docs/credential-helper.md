
# Gitzy Credential Management Guide

Welcome to the Gitzy credential management documentation. This guide explains how Git handles credentials, the challenges of switching profiles across platforms, and how Gitzy solves these problems for you.

---

## What Are Git Credentials?

Git credentials are used to authenticate you with remote repositories (e.g., GitHub, GitLab). Credentials can be stored and managed in different ways depending on your operating system and Git configuration.

### Credential Helpers

Git uses "credential helpers" to store and retrieve credentials:

- **store**: Saves credentials in plaintext in `~/.git-credentials`.
- **cache**: Temporarily stores credentials in memory for the session.
- **manager/manager-core**: Secure credential managers (Windows Credential Manager, macOS Keychain, Linux Secret Service).
- **libsecret/gnome-keyring/kwallet**: Desktop environment-specific helpers on Linux.

### Configuration Layers & Precedence

Credential helpers can be set at different configuration layers:

1. **Local**: `<repo>/.git/config` (highest precedence)
2. **Global**: `~/.gitconfig` or `%USERPROFILE%\.gitconfig`
3. **System**: `/etc/gitconfig` or `C:\ProgramData\Git\config` (lowest precedence)

To check which credential helper is active, run:

```bash
git config --list --show-origin | grep credential.helper
```

---

## The Cross-Platform Problem

On **Windows**, Git often uses the Windows Credential Manager (WCM) by default. If you update `.git-credentials`, Git may still use credentials from WCM, causing confusion and failed authentication. On Linux/macOS, the default may be `store`, `cache`, or a desktop-specific manager. This makes switching profiles unreliable if you only update one credential source.

---

## How Gitzy Solves Credential Switching

Gitzy is a cross-platform CLI tool written in Go that:

1. **Detects your OS at runtime**
2. **Configures the correct credential helper** for your platform
3. **Switches credentials** by updating the relevant config and credential store

This ensures Git always uses the intended credentials, regardless of platform or default helper.

---

## Usage: Switching Profiles with Gitzy

1. **Install Gitzy** (see the main install guide)
2. **Run Gitzy to switch profiles**:
   - Gitzy will detect your OS and set the appropriate credential helper
   - It will update your credentials in the correct store (e.g., WCM on Windows, `.git-credentials` on Linux)

Example Go code (used internally by Gitzy):

```go
package credentials

import (
	"os"
	"os/exec"
	"runtime"
	"fmt"
)

func SetCredentialHelper(profile string) error {
	var helper string
	switch runtime.GOOS {
	case "windows":
		helper = "manager-core"
	case "darwin":
		helper = "osxkeychain"
	case "linux":
		helper = "store" // or "libsecret", "gnome-keyring", etc.
	default:
		return fmt.Errorf("unsupported OS: %s", runtime.GOOS)
	}
	cmd := exec.Command("git", "config", "--global", "credential.helper", helper)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to set credential helper: %w", err)
	}
	// Optionally, update credentials for the selected profile
	return nil
}
```

---

## Setup Checklist

1. **Install Git** on your system
2. **Ensure your credential manager is available**:
   - Windows: `manager-core` (comes with Git for Windows)
   - macOS: `osxkeychain`
   - Linux: `store`, `libsecret`, etc.
3. **Install Gitzy** and use it to switch profiles

---

## Troubleshooting

- If Git is not using the expected credentials, check the active credential helper:
  ```bash
  git config --list --show-origin | grep credential.helper
  ```
- Make sure only one helper is set at the desired config layer
- If switching fails, verify your credential manager is installed and available

---

## Advanced: Testing Environments

- **Linux**: Use Vagrant boxes (Ubuntu, Debian)
- **Windows**: Use community Vagrant boxes (licensing required)
- **macOS**: Manual or cloud-based VMs

---

## Summary

Gitzy solves the cross-platform credential switching problem by dynamically configuring the correct Git credential helper for your OS, ensuring seamless profile management for developers who need to switch between multiple Git identities.

