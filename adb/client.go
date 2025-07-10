package adb

import (
	"context"
	"fmt"
	"os/exec"
	"strings"
	"time"
)

// Timeout constants for different operations
const (
	ShortTimeout  = 5 * time.Second  // For quick operations like list devices
	MediumTimeout = 10 * time.Second // For shell commands, uninstall
	LongTimeout   = 30 * time.Second // For install operations
)

// Client represents an ADB client for CLI operations
type Client struct {
	host string
	port int
	path string
}

// Option is a functional option for configuring the Client
type Option func(*Client)

// WithHost sets the ADB server host
func WithHost(host string) Option {
	return func(c *Client) { c.host = host }
}

// WithPort sets the ADB server port
func WithPort(port int) Option {
	return func(c *Client) { c.port = port }
}

// WithPath sets the path to the adb executable
func WithPath(path string) Option {
	return func(c *Client) { c.path = path }
}

// NewClient creates a new Client with the given options
func NewClient(opts ...Option) *Client {
	c := &Client{
		host: "127.0.0.1",
		port: 5037,
		path: "adb",
	}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

// ListDevices returns a list of connected ADB devices
func (c *Client) ListDevices() ([]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), ShortTimeout)
	defer cancel()

	output, err := c.RunCommand(ctx, "devices")
	if err != nil {
		return nil, fmt.Errorf("failed to list devices: %w", err)
	}

	lines := strings.Split(strings.TrimSpace(output), "\n")
	var devices []string

	for _, line := range lines {
		if strings.Contains(line, "\tdevice") {
			deviceID := strings.Split(line, "\t")[0]
			devices = append(devices, deviceID)
		}
	}

	return devices, nil
}

// RunCommand executes an adb command with arguments
func (c *Client) RunCommand(ctx context.Context, args ...string) (string, error) {
	cmd := exec.CommandContext(ctx, c.path, args...)
	out, err := cmd.CombinedOutput()
	return string(out), err
}

// InstallApk installs an APK file to the device
func (c *Client) InstallApk(apkPath string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), LongTimeout)
	defer cancel()

	return c.RunCommand(ctx, "install", apkPath)
}

// UninstallApp uninstalls an app by package name
func (c *Client) UninstallApp(packageName string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), MediumTimeout)
	defer cancel()

	return c.RunCommand(ctx, "uninstall", packageName)
}

// Shell executes a shell command on the device
func (c *Client) Shell(command string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), MediumTimeout)
	defer cancel()

	return c.RunCommand(ctx, "shell", command)
}
