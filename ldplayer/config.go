package ldplayer

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// Config holds the configuration for LDPlayer.
type Config struct {
	ConsolePath string `json:"console_path"`
}

// ConfigPath returns the path to the LDPlayer config file in the user's home directory.
func ConfigPath() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".go-ldplayer-cli.json")
}

// SaveConfig writes the given configuration to the config file.
func SaveConfig(cfg Config) error {
	f, err := os.Create(ConfigPath())
	if err != nil {
		return err
	}
	defer f.Close()
	return json.NewEncoder(f).Encode(cfg)
} 