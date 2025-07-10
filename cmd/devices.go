package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/thienkb1123/go-adb-cli/adb"
)

// Device-related commands
var devicesCmd = &cobra.Command{
	Use:   "devices",
	Short: "Device management commands",
	Long:  "Commands for managing ADB devices (list, connect, disconnect)",
}

var listDevicesCmd = &cobra.Command{
	Use:   "list",
	Short: "List all connected ADB devices",
	Run: func(cmd *cobra.Command, args []string) {
		client := adb.NewClientFromConfig()
		devices, err := client.ListDevices()
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

var connectDeviceCmd = &cobra.Command{
	Use:   "connect [host:port]",
	Short: "Connect to a device via TCP/IP",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		address := args[0]
		client := adb.NewClientFromConfig()

		fmt.Printf("Connecting to %s...\n", address)
		output, err := client.RunCommand(context.Background(), "connect", address)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		fmt.Println("Connection result:")
		fmt.Println(output)
	},
}

func init() {
	devicesCmd.AddCommand(listDevicesCmd)
	devicesCmd.AddCommand(connectDeviceCmd)
	rootCmd.AddCommand(devicesCmd)
}
