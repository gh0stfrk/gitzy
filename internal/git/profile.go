package git

import (
    "encoding/json"
    "os"
    "path/filepath"
)

type Profile struct {
    Name         string `json:"name"`
    UserName     string `json:"user_name"`
    UserEmail    string `json:"user_email"`
    Token        string `json:"token"`
    GitConfig    string `json:"gitconfig_path"`
    Credentials  string `json:"credentials_path"`
    GitConfigSHA string `json:"gitconfig_sha"`
    CredsSHA     string `json:"credentials_sha"`
}

type Index struct {
    Profiles map[string]*Profile `json:"profiles"`
}

func GetConfigDir() (string, error) {
    if home := os.Getenv("GITZY_HOME"); home != "" {
        return home, nil
    }
    if xdg := os.Getenv("XDG_CONFIG_HOME"); xdg != "" {
        return filepath.Join(xdg, "gitzy"), nil
    }
    home, err := os.UserHomeDir()
    if err != nil {
        return "", err
    }
    return filepath.Join(home, ".gitzy"), nil
}

func LoadIndex() (*Index, error) {
    dir, err := GetConfigDir()
    if err != nil {
        return nil, err
    }
    idxPath := filepath.Join(dir, "index.json")
    f, err := os.Open(idxPath)
    if err != nil {
        if os.IsNotExist(err) {
            return &Index{Profiles: map[string]*Profile{}}, nil
        }
        return nil, err
    }
    defer f.Close()
    var idx Index
    if err := json.NewDecoder(f).Decode(&idx); err != nil {
        return nil, err
    }
    return &idx, nil
}

func SaveIndex(idx *Index) error {
    dir, err := GetConfigDir()
    if err != nil {
        return err
    }
    idxPath := filepath.Join(dir, "index.json")
    f, err := os.Create(idxPath)
    if err != nil {
        return err
    }
    defer f.Close()
    enc := json.NewEncoder(f)
    enc.SetIndent("", "  ")
    return enc.Encode(idx)
}
