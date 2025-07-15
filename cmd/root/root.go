package root

import "github.com/spf13/cobra"

// RootCmd is the root command for the CLI application.
var Cmd = &cobra.Command{
	Use:   "go-adb-cli",
	Short: "ADB and LDPlayer CLI",
}
