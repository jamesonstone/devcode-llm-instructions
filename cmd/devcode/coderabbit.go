package main

import (
	"fmt"
	"os"

	"github.com/jamesonstone/devcode/internal/sync"
	"github.com/spf13/cobra"
)

var coderabbitCmd = &cobra.Command{
	Use:   "coderabbit",
	Short: "Sync only .coderabbit.yaml",
	Run: func(cmd *cobra.Command, args []string) {
		res := sync.SyncFile(sync.ConfigFiles[3], dryRun)
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
	rootCmd.AddCommand(coderabbitCmd)
}
