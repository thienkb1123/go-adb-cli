package adb

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Config struct {
	Path string `json:"path"`
	Host string `json:"host"`
	Port int    `json:"port"`
}

func ConfigPath() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".go-adb-cli.json")
}

func LoadConfig() (Config, error) {
	f, err := os.Open(ConfigPath())
	if err != nil {
		return Config{}, err
	}
	defer f.Close()
	var cfg Config
	err = json.NewDecoder(f).Decode(&cfg)
	return cfg, err
}

func SaveConfig(cfg Config) error {
	fmt.Println("Saving config:", ConfigPath())
	f, err := os.Create(ConfigPath())
	if err != nil {
		return err
	}
	defer f.Close()
	return json.NewEncoder(f).Encode(cfg)
}
