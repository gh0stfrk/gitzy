package ui

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/AlecAivazis/survey/v2"
	"github.com/fatih/color"
	"github.com/gh0stfrk/gitzy/internal/git"
)

func AddProfileFlow(profileName string) error {
	configDir, err := git.GetConfigDir()
	if err != nil {
		return err
	}
	idx, _ := git.LoadIndex()
	if _, exists := idx.Profiles[profileName]; exists {
		return fmt.Errorf("profile %q already exists", profileName)
	}
	if !isSlug(profileName) {
		return fmt.Errorf("profile name must be a slug (letters, numbers, dashes, underscores)")
	}
	var useCurrent bool
	survey.AskOne(&survey.Confirm{
		Message: "Copy current ~/.gitconfig as base?",
		Default: true,
	}, &useCurrent)
	var userName, userEmail, token string
	survey.AskOne(&survey.Input{Message: "Git user.name:"}, &userName)
	survey.AskOne(&survey.Input{Message: "Git user.email:"}, &userEmail)
	survey.AskOne(&survey.Password{Message: "GitHub PAT/token (stored unencrypted):"}, &token)
	color.New(color.FgRed, color.Bold).Println("WARNING: Token will be stored unencrypted. See https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/creating-a-personal-access-token")
	profDir := filepath.Join(configDir, profileName)
	os.MkdirAll(profDir, 0700)
	gitconfigPath := filepath.Join(profDir, "gitconfig")
	credsPath := filepath.Join(profDir, "git-credentials")
	if useCurrent {
		home, _ := os.UserHomeDir()
		src := filepath.Join(home, ".gitconfig")
		if data, err := os.ReadFile(src); err == nil {
			os.WriteFile(gitconfigPath, data, 0600)
		}
	} else {
		os.WriteFile(gitconfigPath, []byte(fmt.Sprintf("[user]\n\tname = %s\n\temail = %s\n[credential]\n\thelper = store", userName, userEmail)), 0600)
	}
	os.WriteFile(credsPath, []byte(token), 0600)
	gitSHA, _ := git.HashFile(gitconfigPath)
	credsSHA, _ := git.HashFile(credsPath)
	idx.Profiles[profileName] = &git.Profile{
		Name:         profileName,
		UserName:     userName,
		UserEmail:    userEmail,
		Token:        token,
		GitConfig:    gitconfigPath,
		Credentials:  credsPath,
		GitConfigSHA: gitSHA,
		CredsSHA:     credsSHA,
	}
	git.SaveIndex(idx)
	var switchNow bool
	survey.AskOne(&survey.Confirm{
		Message: "Switch to this profile now?",
		Default: true,
	}, &switchNow)
	if switchNow {
		return git.SwitchProfile(profileName)
	}
	return nil
}

func isSlug(s string) bool {
	for _, c := range s {
		if !(c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' || c >= '0' && c <= '9' || c == '-' || c == '_') {
			return false
		}
	}
	return true
}

func WipeProfileFlow(profile string) error {
	var confirm1, confirm2 bool
	survey.AskOne(&survey.Confirm{
		Message: fmt.Sprintf("Are you sure you want to irreversibly delete profile %q?", profile),
		Default: false,
	}, &confirm1)
	if !confirm1 {
		return fmt.Errorf("aborted")
	}
	survey.AskOne(&survey.Confirm{
		Message: "This cannot be undone. Confirm again?",
		Default: false,
	}, &confirm2)
	if !confirm2 {
		return fmt.Errorf("aborted")
	}
	configDir, _ := git.GetConfigDir()
	profDir := filepath.Join(configDir, profile)
	os.RemoveAll(profDir)
	idx, _ := git.LoadIndex()
	delete(idx.Profiles, profile)
	git.SaveIndex(idx)
	color.New(color.FgRed, color.Bold).Printf("Profile %q deleted.\n", profile)
	return nil
}
