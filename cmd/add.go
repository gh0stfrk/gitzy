package cmd

import (
	"fmt"
	"os"

	"github.com/gh0stfrk/gitzy/internal/ui"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add <profile>",
	Short: "Add a new GitHub profile",
	Long: `Add a new GitHub profile to gitzy.

This command will interactively prompt you for user.name, user.email, and a GitHub token.
It will store these in a new profile directory and optionally switch to it immediately.

Security note: The token is stored unencrypted. See GitHub token best practices.`,
	Example: `
  # Add a new work profile
  gitzy add work

  # Add a personal profile
  gitzy add personal
`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		profileName := args[0]
		if err := ui.AddProfileFlow(profileName); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to add profile: %v\n", err)
			os.Exit(1)
		}
	},
}
