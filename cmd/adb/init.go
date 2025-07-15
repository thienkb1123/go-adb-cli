package adb

import (
	"github.com/spf13/cobra"
	"github.com/thienkb1123/go-adb-cli/cmd/root"
)

// AdbCmd is the parent command for all ADB related subcommands.
var AdbCmd = &cobra.Command{
	Use:   "adb",
	Short: "Interact with Android Debug Bridge (ADB)",
}

func init() {
	// Register ADB subcommands here
	AdbCmd.AddCommand(setupCmd)
	AdbCmd.AddCommand(devicesCmd)
	AdbCmd.AddCommand(appsCmd)
	AdbCmd.AddCommand(shellCmd)
	// Register the ADB group command to the root command
	root.Cmd.AddCommand(AdbCmd)
}
