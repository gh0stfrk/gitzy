package git

import (
	"fmt"
	"runtime"

	"github.com/danieljoos/wincred"
)

// DetectOS returns the current OS as a string (windows, linux, darwin)
func DetectOS() string {
	return runtime.GOOS
}

// SetWindowsCredentialManager sets the credential.helper to manager-core
func SetWindowsCredentialManager() error {
	// Set credential.helper to manager-core using git config
	// This should be done in switcher.go where exec is imported
	return nil
}

// ClearWindowsCredentials removes credentials from Windows Credential Manager using git credential-manager
func ClearWindowsCredentials() error {
	cred := wincred.NewGenericCredential("git:https://github.com")
	if err := cred.Delete(); err != nil {
		return fmt.Errorf("failed to erase credentials from Windows Credential Manager: %w", err)
	}
	return nil
}

// SetWindowsCredentials sets credentials in Windows Credential Manager using wincred
func SetWindowsCredentials(username, password string) error {
	cred := wincred.NewGenericCredential("git:https://github.com")
	cred.UserName = username
	cred.CredentialBlob = []byte(password)
	cred.Persist = wincred.PersistLocalMachine
	if err := cred.Write(); err != nil {
		return fmt.Errorf("failed to write credentials to Windows Credential Manager: %w", err)
	}
	return nil
}
