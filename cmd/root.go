package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "go-adb-cli",
	Short: "CLI tool remote ADB",
	Long:  "Go-powered CLI tool to run adb commands like list-devices, install, shell, etc.",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
