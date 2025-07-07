package cmd

import (
	"fmt"
	"os"

	"github.com/gh0stfrk/gitzy/internal/ui"
	"github.com/spf13/cobra"
)

var wipeCmd = &cobra.Command{
	Use:   "wipe <profile>",
	Short: "Irreversibly delete a profile (double-confirm!)",
	Long: `Irreversibly delete a gitzy profile.

This command will prompt you twice for confirmation before deleting the profile directory and removing it from the index. This action cannot be undone.`,
	Example: `
  # Wipe the work profile
  gitzy wipe work
`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		profile := args[0]
		if err := ui.WipeProfileFlow(profile); err != nil {
			fmt.Fprintf(os.Stderr, "Wipe failed: %v\n", err)
			os.Exit(1)
		}
	},
}
