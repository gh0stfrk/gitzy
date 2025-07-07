package git

import (
    "os"
    "path/filepath"
    "testing"
)

func TestBackupCurrentConfigs(t *testing.T) {
    dir := t.TempDir()
    os.Setenv("GITZY_HOME", dir)
    home := t.TempDir()
    os.Setenv("HOME", home)
    f1 := filepath.Join(home, ".gitconfig")
    f2 := filepath.Join(home, ".git-credentials")
    os.WriteFile(f1, []byte("testconfig"), 0600)
    os.WriteFile(f2, []byte("testcreds"), 0600)
    backupDir, err := backupCurrentConfigs()
    if err != nil {
        t.Fatal(err)
    }
    if _, err := os.Stat(filepath.Join(backupDir, ".gitconfig")); err != nil {
        t.Error("gitconfig not backed up")
    }
    if _, err := os.Stat(filepath.Join(backupDir, ".git-credentials")); err != nil {
        t.Error("git-credentials not backed up")
    }
}
