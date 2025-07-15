package ldplayer

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/thienkb1123/go-adb-cli/ldplayer"
)

// setupCmd defines the CLI command for setting up the LDPlayer console path.
var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Setup default LDPlayer console path",
	Run: func(cmd *cobra.Command, args []string) {
		// Prompt the user to enter the path to ldconsole.exe
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter ldconsole.exe path (default: C:\\LDPlayer\\LDPlayer9\\ldconsole.exe): ")
		ldPath, _ := reader.ReadString('\n')
		ldPath = strings.TrimRight(ldPath, "\r\n")
		if ldPath == "" {
			ldPath = "C:\\LDPlayer\\LDPlayer9\\ldconsole.exe"
		}

		// Save the configuration to the config file
		cfg := ldplayer.Config{ConsolePath: ldPath}
		if err := ldplayer.SaveConfig(cfg); err != nil {
			fmt.Println("Could not write config:", err)
			return
		}
		fmt.Println("Config saved to", ldplayer.ConfigPath())
	},
} 