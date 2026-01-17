package main

import (
	"fmt"
	"os"

	"github.com/jamesonstone/devcode/internal/sync"
	"github.com/spf13/cobra"
)

var (
	dryRun bool
	version = "v0.1.0"
)

var rootCmd = &cobra.Command{
	Use:   "devcode",
	Short: "Sync LLM instruction files to your project",
	Run: func(cmd *cobra.Command, args []string) {
		results := sync.SyncAll(dryRun)
		hasError := false
		for _, r := range results {
			if r.Status == "error" {
				fmt.Printf("❌ %s: %s\n", r.File.Name, r.ErrorMsg)
				hasError = true
			} else if r.Status == "updated" {
				fmt.Printf("✅ Updated %s\n", r.File.Name)
			} else {
				fmt.Printf("✅ %s is up-to-date\n", r.File.Name)
			}
		}
		if hasError {
			os.Exit(1)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVar(&dryRun, "dry-run", false, "Show what would change without writing files")
	rootCmd.Version = version
	rootCmd.SetVersionTemplate("devcode version {{.Version}}\n")
}
