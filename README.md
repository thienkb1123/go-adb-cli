# Go ADB CLI

A modern, Go-powered command-line interface for Android Debug Bridge (ADB) operations. This CLI tool provides a clean and organized way to interact with Android devices through ADB commands, similar to [AdvancedSharpAdbClient](https://github.com/SharpAdb/AdvancedSharpAdbClient) but built with Go.

## Features

- **Device Management**: List connected devices, connect to devices via TCP/IP
- **App Management**: Install/uninstall APK files, list installed applications  
- **Shell Operations**: Execute shell commands on devices
- **File Operations**: Push/pull files between host and device
- **Organized Commands**: Commands are grouped by functionality for better usability
- **Cross-platform**: Works on Linux, macOS, and Windows
- **No Installation Required**: Just download and run

## Download

### Prerequisites

- ADB (Android Debug Bridge) installed and in your PATH
- No Go installation required

### Download Pre-built Binaries

Download the latest release for your platform:

#### macOS
```bash
# Apple Silicon (M1/M2)
wget https://github.com/thienkb1123/go-adb-cli/releases/latest/download/go-adb-cli-darwin-arm64
chmod +x go-adb-cli-darwin-arm64
./go-adb-cli-darwin-arm64 --help

# Intel
wget https://github.com/thienkb1123/go-adb-cli/releases/latest/download/go-adb-cli-darwin-amd64
chmod +x go-adb-cli-darwin-amd64
./go-adb-cli-darwin-amd64 --help
```

#### Linux
```bash
# AMD64
wget https://github.com/thienkb1123/go-adb-cli/releases/latest/download/go-adb-cli-linux-amd64
chmod +x go-adb-cli-linux-amd64
./go-adb-cli-linux-amd64 --help

# ARM64
wget https://github.com/thienkb1123/go-adb-cli/releases/latest/download/go-adb-cli-linux-arm64
chmod +x go-adb-cli-linux-arm64
./go-adb-cli-linux-arm64 --help
```

#### Windows
```bash
# Download from browser or use PowerShell
# AMD64: go-adb-cli-windows-amd64.exe
# ARM64: go-adb-cli-windows-arm64.exe
```

### Quick Start

1. Download the binary for your platform
2. Make it executable: `chmod +x go-adb-cli-*`
3. Run: `./go-adb-cli-* --help`

## Usage

### Device Management

```bash
# List all connected devices
./go-adb-cli devices list

# Connect to a device via TCP/IP
./go-adb-cli devices connect 192.168.1.100:5555
```

### App Management

```bash
# Install an APK file
./go-adb-cli apps install /path/to/your/app.apk

# Uninstall an app by package name
./go-adb-cli apps uninstall com.example.app

# List installed applications
./go-adb-cli apps list
```

### Shell and File Operations

```bash
# Execute a shell command
./go-adb-cli shell exec "ls /sdcard"

# Push a file to the device
./go-adb-cli shell push local_file.txt /sdcard/

# Pull a file from the device
./go-adb-cli shell pull /sdcard/remote_file.txt ./
```

### Getting Help

```bash
# Show general help
./go-adb-cli --help

# Show help for specific command group
./go-adb-cli devices --help
./go-adb-cli apps --help
./go-adb-cli shell --help

# Show help for specific command
./go-adb-cli devices list --help
./go-adb-cli apps install --help
```

## Command Structure

The CLI organizes commands into logical groups:

### Devices Group
- `list` - List connected ADB devices
- `connect [host:port]` - Connect to device via TCP/IP

### Apps Group  
- `install [apk-path]` - Install APK file to device
- `uninstall [package-name]` - Uninstall app by package name
- `list` - List installed applications

### Shell Group
- `exec [command]` - Execute shell command on device
- `push [local] [remote]` - Push file to device
- `pull [remote] [local]` - Pull file from device

## Examples

### Basic Device Operations
```bash
# Check connected devices
./go-adb-cli devices list

# Connect to device over network
./go-adb-cli devices connect 192.168.1.100:5555
```

### App Management Workflow
```bash
# Install new app
./go-adb-cli apps install /path/to/app.apk

# List installed apps
./go-adb-cli apps list

# Uninstall app
./go-adb-cli apps uninstall com.example.app
```

### File and Shell Operations
```bash
# Check device storage
./go-adb-cli shell exec "df /sdcard"

# Copy file to device
./go-adb-cli shell push myfile.txt /sdcard/

# Copy file from device
./go-adb-cli shell pull /sdcard/myfile.txt ./
```

## Configuration

The ADB client uses default configuration:
- Host: `127.0.0.1`
- Port: `5037`
- ADB Path: `adb` (from PATH)

## Timeout Configuration

Different operations have different timeout values:
- **ShortTimeout** (5s): Quick operations like listing devices
- **MediumTimeout** (10s): Shell commands, uninstall operations
- **LongTimeout** (30s): Install operations, large file transfers

## Comparison with AdvancedSharpAdbClient

This project is similar to [AdvancedSharpAdbClient](https://github.com/SharpAdb/AdvancedSharpAdbClient) but built with Go:

| Feature | AdvancedSharpAdbClient | go-adb-cli |
|---------|------------------------|------------|
| Language | C# | Go |
| Installation | NuGet package | Single binary download |
| Platform | .NET/Mono/Unity | Cross-platform binary |
| Architecture | Library + CLI | Standalone CLI tool |
| Configuration | Constructor parameters | Default settings |
| Timeout | Manual management | Built-in constants |
| Error Handling | Exceptions | Error returns |

## Troubleshooting

### Common Issues

1. **Permission denied**: Make sure the binary is executable
   ```bash
   chmod +x go-adb-cli-*
   ```

2. **ADB not found**: Ensure ADB is installed and in your PATH
   ```bash
   which adb
   ```

3. **No devices found**: Check if your device is connected and USB debugging is enabled

### Platform-specific Notes

- **macOS**: May need to allow execution in Security & Privacy settings
- **Windows**: Run as Administrator if needed for device access
- **Linux**: May need to add udev rules for device access

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Built with [Cobra](https://github.com/spf13/cobra) for CLI functionality
- Uses Go's standard library for ADB command execution
- Inspired by [AdvancedSharpAdbClient](https://github.com/SharpAdb/AdvancedSharpAdbClient)
- Designed for clean, organized ADB CLI operations
