package ldplayer

import (
	"github.com/spf13/cobra"
	"github.com/thienkb1123/go-adb-cli/cmd/root"
)

// ldplayerCmd is the parent command for all LDPlayer related subcommands.
var ldplayerCmd = &cobra.Command{
	Use:   "ldplayer",
	Short: "Interact with LDPlayer emulator",
}

func init() {
	// Register LDPlayer subcommands here
	ldplayerCmd.AddCommand(setupCmd)
	// Register the LDPlayer group command to the root command
	root.Cmd.AddCommand(ldplayerCmd)
}
