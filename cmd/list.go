package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/gh0stfrk/gitzy/internal/git"
	"github.com/spf13/cobra"
)

var jsonOut bool

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List available profiles and show the active one",
	Long: `List all available gitzy profiles and indicate which one is currently active.

The active profile is detected by comparing the hashes of your current ~/.gitconfig and ~/.git-credentials with those stored in your profiles.`,
	Example: `
  # List all profiles
  gitzy list

  # List all profiles in JSON format
  gitzy list --json
`,
	Run: func(cmd *cobra.Command, args []string) {
		profiles, active, err := git.ListProfiles()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		if jsonOut {
			out := map[string]interface{}{
				"profiles": profiles,
				"active":   active,
			}
			enc := json.NewEncoder(os.Stdout)
			enc.SetIndent("", "  ")
			enc.Encode(out)
			return
		}
		for _, p := range profiles {
			if p == active {
				color.New(color.FgGreen, color.Bold).Printf("* %s (active)\n", p)
			} else {
				fmt.Printf("  %s\n", p)
			}
		}
	},
}

func init() {
	listCmd.Flags().BoolVar(&jsonOut, "json", false, "Output as JSON")
}
