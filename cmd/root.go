package cmd

import (
	"embed"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	//go:embed README.md
	readmeFS embed.FS
	docsFlag bool
	RootCmd  = &cobra.Command{
		Use:   "gitzy",
		Short: "Manage multiple GitHub identities and switch between them easily.",
		Long:  "gitzy is a cross-platform CLI for managing and switching between multiple GitHub identities.",
	}
)

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize()
	RootCmd.PersistentFlags().BoolVar(&docsFlag, "docs", false, "Show embedded documentation")
	RootCmd.AddCommand(addCmd)
	RootCmd.AddCommand(switchCmd)
	RootCmd.AddCommand(listCmd)
	RootCmd.AddCommand(wipeCmd)
	RootCmd.AddCommand(doctorCmd)
}
