package cmd

import (
    "fmt"
    "os"

    "github.com/spf13/cobra"
)

var doctorCmd = &cobra.Command{
    Use:   "doctor",
    Short: "Show troubleshooting info and embedded docs",
    Run: func(cmd *cobra.Command, args []string) {
        if docsFlag {
            data, err := os.ReadFile("README.md")
            if err != nil {
                fmt.Println("Could not read embedded docs.")
                return
            }
            fmt.Println(string(data))
        } else {
            fmt.Println("Run with --docs to see embedded documentation.")
        }
    },
}
