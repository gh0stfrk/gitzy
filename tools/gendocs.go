//go:build tools
// +build tools

package main

import (
	"log"
	"os"

	"github.com/gh0stfrk/gitzy/cmd"
	"github.com/spf13/cobra/doc"
)

func main() {
	outDir := "docs/cli"
	if err := os.MkdirAll(outDir, 0755); err != nil {
		log.Fatalf("failed to create docs dir: %v", err)
	}
	// Ensure RootCmd is exported in cmd/root.go
	if err := doc.GenMarkdownTree(cmd.RootCmd, outDir); err != nil {
		log.Fatalf("failed to generate CLI docs: %v", err)
	}
	log.Printf("CLI docs generated in %s", outDir)
}
