package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/thienkb1123/go-adb-cli/adb"
)

// Shell and file operation commands
var shellCmd = &cobra.Command{
	Use:   "shell",
	Short: "Shell and file operation commands",
	Long:  "Commands for executing shell commands and file operations",
}

var executeShellCmd = &cobra.Command{
	Use:   "exec [command]",
	Short: "Execute a shell command on the device",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		command := args[0]
		client := adb.NewClientFromConfig()

		fmt.Printf("Executing: %s\n", command)
		output, err := client.Shell(command)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		fmt.Println("Output:")
		fmt.Println(output)
	},
}

var pushFileCmd = &cobra.Command{
	Use:   "push [local-path] [remote-path]",
	Short: "Push a file to the device",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		localPath := args[0]
		remotePath := args[1]
		client := adb.NewClientFromConfig()

		fmt.Printf("Pushing %s to %s...\n", localPath, remotePath)
		output, err := client.RunCommand(context.Background(), "push", localPath, remotePath)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		fmt.Println("Push result:")
		fmt.Println(output)
	},
}

var pullFileCmd = &cobra.Command{
	Use:   "pull [remote-path] [local-path]",
	Short: "Pull a file from the device",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		remotePath := args[0]
		localPath := args[1]
		client := adb.NewClientFromConfig()

		fmt.Printf("Pulling %s to %s...\n", remotePath, localPath)
		output, err := client.RunCommand(context.Background(), "pull", remotePath, localPath)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		fmt.Println("Pull result:")
		fmt.Println(output)
	},
}

func init() {
	shellCmd.AddCommand(executeShellCmd)
	shellCmd.AddCommand(pushFileCmd)
	shellCmd.AddCommand(pullFileCmd)
	rootCmd.AddCommand(shellCmd)
}
