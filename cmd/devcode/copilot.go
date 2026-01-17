package main

import (
	"fmt"
	"os"

	"github.com/jamesonstone/devcode/internal/sync"
	"github.com/spf13/cobra"
)

var copilotCmd = &cobra.Command{
	Use:   "copilot",
	Short: "Sync only copilot-instructions.md",
	Run: func(cmd *cobra.Command, args []string) {
		res := sync.SyncFile(sync.ConfigFiles[0], dryRun)
		if res.Status == "error" {
			fmt.Printf("❌ %s: %s\n", res.File.Name, res.ErrorMsg)
			os.Exit(1)
		} else if res.Status == "updated" {
			fmt.Printf("✅ Updated %s\n", res.File.Name)
		} else {
			fmt.Printf("✅ %s is up-to-date\n", res.File.Name)
		}
	},
}

func init() {
	rootCmd.AddCommand(copilotCmd)
}
