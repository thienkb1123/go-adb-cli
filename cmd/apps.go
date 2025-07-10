package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/thienkb1123/go-adb-cli/adb"
)

// App-related commands
var appsCmd = &cobra.Command{
	Use:   "apps",
	Short: "App management commands",
	Long:  "Commands for managing Android apps (install, uninstall, list)",
}

var installAppCmd = &cobra.Command{
	Use:   "install [apk-path]",
	Short: "Install an APK file to the device",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apkPath := args[0]
		client := adb.NewClient()

		fmt.Printf("Installing %s...\n", apkPath)
		output, err := client.InstallApk(apkPath)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		fmt.Println("Installation output:")
		fmt.Println(output)
	},
}

var uninstallAppCmd = &cobra.Command{
	Use:   "uninstall [package-name]",
	Short: "Uninstall an app by package name",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		packageName := args[0]
		client := adb.NewClient()

		fmt.Printf("Uninstalling %s...\n", packageName)
		output, err := client.UninstallApp(packageName)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		fmt.Println("Uninstall output:")
		fmt.Println(output)
	},
}

var listAppsCmd = &cobra.Command{
	Use:   "list",
	Short: "List installed apps",
	Run: func(cmd *cobra.Command, args []string) {
		client := adb.NewClient()

		output, err := client.Shell("pm list packages")
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		fmt.Println("Installed packages:")
		fmt.Println(output)
	},
}

func init() {
	appsCmd.AddCommand(installAppCmd)
	appsCmd.AddCommand(uninstallAppCmd)
	appsCmd.AddCommand(listAppsCmd)
	rootCmd.AddCommand(appsCmd)
}
