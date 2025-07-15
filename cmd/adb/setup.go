package adb

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/thienkb1123/go-adb-cli/adb"
)

// setupCmd defines the CLI command for setting up the ADB path, host, and port.
var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Setup default adb path, host, and port",
	Run: func(cmd *cobra.Command, args []string) {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter adb path (default: adb): ")
		adbPath, _ := reader.ReadString('\n')
		adbPath = trimNL(adbPath)
		if adbPath == "" {
			adbPath = "adb"
		}

		fmt.Print("Enter adb host (default: 127.0.0.1): ")
		host, _ := reader.ReadString('\n')
		host = trimNL(host)
		if host == "" {
			host = "127.0.0.1"
		}

		fmt.Print("Enter adb port (default: 5037): ")
		var port int
		_, err := fmt.Scanf("%d", &port)
		if err != nil || port == 0 {
			port = 5037
		}

		cfg := adb.Config{Path: adbPath, Host: host, Port: port}
		if err := adb.SaveConfig(cfg); err != nil {
			fmt.Println("Could not write config:", err)
			return
		}
		fmt.Println("Config saved to", adb.ConfigPath())
	},
}

// trimNL removes trailing carriage return and newline characters from a string.
func trimNL(s string) string {
	return strings.TrimRight(s, "\r\n")
} 