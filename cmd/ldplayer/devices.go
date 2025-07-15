package ldplayer

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/thienkb1123/go-adb-cli/ldplayer"
)

var devicesCmd = &cobra.Command{
	Use:   "devices",
	Short: "Device management commands",
}

var listDevicesCmd = &cobra.Command{
	Use:   "list",
	Short: "List all connected LDPlayer devices",
	Run: func(cmd *cobra.Command, args []string) {
		client := ldplayer.NewClientFromConfig()
		devices, err := client.List()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		if len(devices) == 0 {
			fmt.Println("No devices connected")
			return
		}

		fmt.Println("Connected devices:")
		for i, device := range devices {
			fmt.Printf("%d. %s\n", i+1, device)
		}
	},
}

func init() {
	devicesCmd.AddCommand(listDevicesCmd)
	ldplayerCmd.AddCommand(devicesCmd)
}
