# Makefile for go-adb-cli
# Build configuration for Go ADB CLI tool

# Variables
BINARY_NAME=go-adb-cli
MAIN_PATH=main.go
BUILD_DIR=build
VERSION?=$(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
COMMIT_HASH?=$(shell git rev-parse --short HEAD 2>/dev/null || echo "unknown")
BUILD_TIME?=$(shell date -u '+%Y-%m-%d_%H:%M:%S')
LDFLAGS=-ldflags "-X main.Version=${VERSION} -X main.CommitHash=${COMMIT_HASH} -X main.BuildTime=${BUILD_TIME}"

# Go build flags
GOOS?=$(shell go env GOOS)
GOARCH?=$(shell go env GOARCH)
CGO_ENABLED?=0

# Default target
.PHONY: all
all: clean build

# Build the application for current platform
.PHONY: build
build:
	@echo "Building $(BINARY_NAME) for $(GOOS)/$(GOARCH)..."
	@mkdir -p $(BUILD_DIR)
	CGO_ENABLED=$(CGO_ENABLED) GOOS=$(GOOS) GOARCH=$(GOARCH) go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_PATH)
	chmod +x $(BUILD_DIR)/$(BINARY_NAME)
	@echo "Build completed: $(BUILD_DIR)/$(BINARY_NAME)"

# Build for multiple platforms
.PHONY: build-all
build-all: clean
	@echo "Building for multiple platforms..."
	@mkdir -p $(BUILD_DIR)
	
	# Linux
	@echo "Building for Linux..."
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-linux-amd64 $(MAIN_PATH)
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-linux-arm64 $(MAIN_PATH)
	
	# macOS
	@echo "Building for macOS..."
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-amd64 $(MAIN_PATH)
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-arm64 $(MAIN_PATH)
	
	# Windows
	@echo "Building for Windows..."
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-windows-amd64.exe $(MAIN_PATH)
	CGO_ENABLED=0 GOOS=windows GOARCH=arm64 go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-windows-arm64.exe $(MAIN_PATH)
	
	@echo "Multi-platform build completed!"

# Build with debug information
.PHONY: build-debug
build-debug:
	@echo "Building $(BINARY_NAME) with debug information..."
	@mkdir -p $(BUILD_DIR)
	CGO_ENABLED=$(CGO_ENABLED) GOOS=$(GOOS) GOARCH=$(GOARCH) go build -gcflags="all=-N -l" $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-debug $(MAIN_PATH)
	chmod +x $(BUILD_DIR)/$(BINARY_NAME)-debug
	@echo "Debug build completed: $(BUILD_DIR)/$(BINARY_NAME)-debug"

# Install the binary to GOPATH/bin
.PHONY: install
install: build
	@echo "Installing $(BINARY_NAME)..."
	@cp $(BUILD_DIR)/$(BINARY_NAME) $(shell go env GOPATH)/bin/
	@echo "Installed to $(shell go env GOPATH)/bin/$(BINARY_NAME)"

# Run the application
.PHONY: run
run:
	@echo "Running $(BINARY_NAME)..."
	@go run $(MAIN_PATH)

# Run with specific command
.PHONY: run-list
run-list:
	@echo "Running list command..."
	@go run $(MAIN_PATH) list

# Test the application
.PHONY: test
test:
	@echo "Running tests..."
	@go test -v ./...

# Test with coverage
.PHONY: test-coverage
test-coverage:
	@echo "Running tests with coverage..."
	@go test -v -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

# Clean build artifacts
.PHONY: clean
clean:
	@echo "Cleaning build artifacts..."
	@rm -rf $(BUILD_DIR)
	@rm -f coverage.out coverage.html
	@echo "Clean completed"

# Format code
.PHONY: fmt
fmt:
	@echo "Formatting code..."
	@go fmt ./...

# Lint code
.PHONY: lint
lint:
	@echo "Linting code..."
	@if command -v golangci-lint >/dev/null 2>&1; then \
		golangci-lint run; \
	else \
		echo "golangci-lint not found. Install with: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest"; \
	fi

# Vet code
.PHONY: vet
vet:
	@echo "Vetting code..."
	@go vet ./...

# Check code quality
.PHONY: check
check: fmt vet lint

# Download dependencies
.PHONY: deps
deps:
	@echo "Downloading dependencies..."
	@go mod download
	@go mod tidy

# Update dependencies
.PHONY: deps-update
deps-update:
	@echo "Updating dependencies..."
	@go get -u ./...
	@go mod tidy

# Show help
.PHONY: help
help:
	@echo "Available targets:"
	@echo "  build         - Build for current platform"
	@echo "  build-all     - Build for multiple platforms (Linux, macOS, Windows)"
	@echo "  build-debug   - Build with debug information"
	@echo "  install       - Install binary to GOPATH/bin"
	@echo "  run           - Run the application"
	@echo "  run-list      - Run with list command"
	@echo "  test          - Run tests"
	@echo "  test-coverage - Run tests with coverage report"
	@echo "  clean         - Clean build artifacts"
	@echo "  fmt           - Format code"
	@echo "  lint          - Lint code (requires golangci-lint)"
	@echo "  vet           - Vet code"
	@echo "  check         - Run fmt, vet, and lint"
	@echo "  deps          - Download dependencies"
	@echo "  deps-update   - Update dependencies"
	@echo "  help          - Show this help message"
	@echo ""
	@echo "Environment variables:"
	@echo "  VERSION       - Version string (default: git tag or 'dev')"
	@echo "  GOOS          - Target OS (default: current OS)"
	@echo "  GOARCH        - Target architecture (default: current arch)"
	@echo "  CGO_ENABLED   - Enable CGO (default: 0)"

# Default target
.DEFAULT_GOAL := build 