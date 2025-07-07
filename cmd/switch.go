package cmd

import (
	"fmt"
	"os"

	"github.com/gh0stfrk/gitzy/internal/git"
	"github.com/spf13/cobra"
)

var switchCmd = &cobra.Command{
	Use:   "switch <profile>",
	Short: "Switch to a different GitHub profile",
	Long: `Switch to a different GitHub profile.

This command will backup your current ~/.gitconfig and ~/.git-credentials, then replace them with the files from the selected profile.

Backups are stored in ~/.gitzy/backups/<timestamp>/ for safety.`,
	Example: `
  # Switch to the work profile
  gitzy switch work

  # Switch to the personal profile
  gitzy switch personal
`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		profile := args[0]
		if err := git.SwitchProfile(profile); err != nil {
			fmt.Fprintf(os.Stderr, "Switch failed: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Switched to profile: %s\n", profile)
	},
}
