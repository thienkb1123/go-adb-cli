package cmd

import (
	"github.com/thienkb1123/go-adb-cli/cmd/root"
)

// Execute runs the root command.
func Execute() error {
	return root.Cmd.Execute()
}
