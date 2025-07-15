package ldplayer

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

// Timeout constants for different LDPlayer operations
const (
	ShortTimeout  = 5 * time.Second
	MediumTimeout = 10 * time.Second
	LongTimeout   = 30 * time.Second
)

// Client manages LDPlayer configuration and command execution
type Client struct {
	path string // Path to ldconsole.exe
}

// Option is a functional option for configuring the Client
type Option func(*Client)

// WithPath sets the path to ldconsole.exe for the Client
func WithPath(path string) Option {
	return func(c *Client) { c.path = path }
}

// NewClient creates a new Client with the given options
func NewClient(opts ...Option) *Client {
	c := &Client{
		path: "C:\\LDPlayer\\LDPlayer9\\ldconsole.exe", // default path
	}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

// NewClientFromConfig creates a Client from the config file ~/.go-ldplayer-cli.json
func NewClientFromConfig() *Client {
	home, err := os.UserHomeDir()
	if err != nil {
		return NewClient()
	}
	configPath := filepath.Join(home, ".go-ldplayer-cli.json")
	f, err := os.Open(configPath)
	if err != nil {
		return NewClient()
	}
	defer f.Close()
	var cfg struct {
		ConsolePath string `json:"console_path"`
	}
	if err := json.NewDecoder(f).Decode(&cfg); err != nil {
		return NewClient()
	}
	return NewClient(WithPath(cfg.ConsolePath))
}

// RunCommand executes ldconsole.exe with the provided arguments
func (c *Client) RunCommand(ctx context.Context, args ...string) (string, error) {
	cmd := exec.CommandContext(ctx, c.path, args...)
	out, err := cmd.CombinedOutput()
	return string(out), err
}

// List returns a list of connected LDPlayer devices
func (c *Client) List() ([]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), ShortTimeout)
	defer cancel()

	output, err := c.RunCommand(ctx, "list")
	if err != nil {
		return nil, fmt.Errorf("failed to list devices: %w", err)
	}

	devices := strings.Split(strings.TrimSpace(output), "\n")
	return devices, nil
}
