# Binary name
BINARY_NAME=delete-cf-videos

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

# Optimization flags
OPTFLAGS=-ldflags="-s -w" -trimpath

# Build the project
all: build

# Build for the current platform
build:
	$(GOBUILD) -o $(BINARY_NAME) -v

# Clean the binary
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -rf dist/

# Run tests
test:
	$(GOTEST) -v ./...

# Cross compilation
build-all: build-linux build-windows build-macos

build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) $(OPTFLAGS) -o dist/$(BINARY_NAME)-linux-amd64 -v
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 $(GOBUILD) $(OPTFLAGS) -o dist/$(BINARY_NAME)-linux-arm64 -v

build-windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GOBUILD) $(OPTFLAGS) -o dist/$(BINARY_NAME)-windows-amd64.exe -v

build-macos:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 $(GOBUILD) $(OPTFLAGS) -o dist/$(BINARY_NAME)-macos-intel -v
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 $(GOBUILD) $(OPTFLAGS) -o dist/$(BINARY_NAME)-macos-apple-chip -v

# Build and run
run: build
	./$(BINARY_NAME)

.PHONY: all build clean test build-all build-linux build-windows build-darwin run
