package main

import (
	"github.com/thienkb1123/go-adb-cli/cmd"
	_ "github.com/thienkb1123/go-adb-cli/cmd/adb"
	_ "github.com/thienkb1123/go-adb-cli/cmd/ldplayer"
)

func main() {
	cmd.Execute()
}
