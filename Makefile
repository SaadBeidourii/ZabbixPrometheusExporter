# Project name
PROJECT_NAME := zabbixprometheusexporter

# Directories
CMD_DIR := cmd/$(PROJECT_NAME)
INTERNAL_DIR := internal

# Build output
BUILD_OUTPUT := bin/$(PROJECT_NAME)

# Go options
GO := go
GO_BUILD := $(GO) build
GO_RUN := $(GO) run
GO_FMT := $(GO) fmt
GO_TEST := $(GO) test

.PHONY: all build run clean fmt test

# Default target
all: build

# Build the project
build:
	@echo "Building project..."
	$(GO_BUILD) -o $(BUILD_OUTPUT) ./$(CMD_DIR)

# Run the project
run:
	@echo "Running project..."
	$(GO_RUN) ./$(CMD_DIR)

# Format the code
fmt:
	@echo "Formatting code..."
	$(GO_FMT) ./...

# Run tests
test:
	@echo "Running tests..."
	$(GO_TEST) ./...

# Clean build files
clean:
	@echo "Cleaning up..."
	rm -f $(BUILD_OUTPUT)
