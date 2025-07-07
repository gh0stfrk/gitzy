package git

import (
    "crypto/sha256"
    "encoding/hex"
    "fmt"
    "io"
    "os"
    "path/filepath"
    "time"
)

func HashFile(path string) (string, error) {
    f, err := os.Open(path)
    if err != nil {
        return "", err
    }
    defer f.Close()
    h := sha256.New()
    if _, err := io.Copy(h, f); err != nil {
        return "", err
    }
    return hex.EncodeToString(h.Sum(nil)), nil
}

func backupCurrentConfigs() (string, error) {
    home, err := os.UserHomeDir()
    if err != nil {
        return "", err
    }
    configDir, err := GetConfigDir()
    if err != nil {
        return "", err
    }
    backupDir := filepath.Join(configDir, "backups", time.Now().Format("20060102-150405"))
    if err := os.MkdirAll(backupDir, 0700); err != nil {
        return "", err
    }
    files := []string{".gitconfig", ".git-credentials"}
    for _, f := range files {
        src := filepath.Join(home, f)
        dst := filepath.Join(backupDir, f)
        if _, err := os.Stat(src); err == nil {
            in, _ := os.Open(src)
            out, _ := os.Create(dst)
            io.Copy(out, in)
            in.Close()
            out.Close()
        }
    }
    return backupDir, nil
}

func SwitchProfile(profile string) error {
    idx, err := LoadIndex()
    if err != nil {
        return err
    }
    p, ok := idx.Profiles[profile]
    if !ok {
        return fmt.Errorf("profile %q not found", profile)
    }
    if _, err := backupCurrentConfigs(); err != nil {
        return err
    }
    home, err := os.UserHomeDir()
    if err != nil {
        return err
    }
    for _, pair := range [][2]string{
        {p.GitConfig, filepath.Join(home, ".gitconfig")},
        {p.Credentials, filepath.Join(home, ".git-credentials")},
    } {
        src, dst := pair[0], pair[1]
        in, err := os.Open(src)
        if err != nil {
            return err
        }
        out, err := os.Create(dst)
        if err != nil {
            in.Close()
            return err
        }
        io.Copy(out, in)
        in.Close()
        out.Close()
    }
    return nil
}

func ListProfiles() ([]string, string, error) {
    idx, err := LoadIndex()
    if err != nil {
        return nil, "", err
    }
    var profiles []string
    for k := range idx.Profiles {
        profiles = append(profiles, k)
    }
    home, err := os.UserHomeDir()
    if err != nil {
        return profiles, "", nil
    }
    gitconfig := filepath.Join(home, ".gitconfig")
    creds := filepath.Join(home, ".git-credentials")
    gitSHA, _ := HashFile(gitconfig)
    credsSHA, _ := HashFile(creds)
    for _, p := range idx.Profiles {
        if p.GitConfigSHA == gitSHA && p.CredsSHA == credsSHA {
            return profiles, p.Name, nil
        }
    }
    return profiles, "", nil
}
