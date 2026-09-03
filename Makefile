# ==============================================================================
# Variables & Build Flags
# ==============================================================================
APP_NAME      := service-evertale-engine
BIN_DIR       := bin
GO            := go
GO_FLAGS      := -trimpath
GIT_COMMIT    := $(shell git rev-parse --short HEAD 2>/dev/null || echo "unknown")
BUILD_TIME    := $(shell date -u '+%Y-%m-%dT%H:%M:%SZ')
LDFLAGS       := -s -w -X 'main.GitCommit=$(GIT_COMMIT)' -X 'main.BuildTime=$(BUILD_TIME)'

# Default target when running just `make`
.DEFAULT_GOAL := help

# ==============================================================================
# Help Menu (Self-Documenting)
# ==============================================================================
.PHONY: help
help: ## Display this help screen
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n\nTargets:\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-18s\033[0m %s\n", $$1, $$2 }' $(MAKEFILE_LIST)

# ==============================================================================
# Development & Running
# ==============================================================================
.PHONY: run-engine run-tui
run-engine: ## Run the game engine server (Godot gateway)
	$(GO) run $(GO_FLAGS) ./cmd/service-evertale-engine

run-tui: ## Run the 2D terminal client
	$(GO) run $(GO_FLAGS) ./cmd/tui

# ==============================================================================
# Building Binaries
# ==============================================================================
.PHONY: build build-engine build-tui build-aigen
build: build-engine build-tui build-aigen ## Build all binaries into bin/

build-engine: ## Build game engine server binary
	@mkdir -p $(BIN_DIR)
	$(GO) build $(GO_FLAGS) -ldflags "$(LDFLAGS)" -o $(BIN_DIR)/engine ./cmd/service-evertale-engine

build-tui: ## Build terminal client binary
	@mkdir -p $(BIN_DIR)
	$(GO) build $(GO_FLAGS) -ldflags "$(LDFLAGS)" -o $(BIN_DIR)/tui ./cmd/tui
build-aigen: ## Build standalone AI generation service binary
	@mkdir -p $(BIN_DIR)
	$(GO) build $(GO_FLAGS) -ldflags "$(LDFLAGS)" -o $(BIN_DIR)/aigen ./cmd/aigen

# ==============================================================================
# Quality, Testing & Formatting
# ==============================================================================
.PHONY: test test-coverage fmt vet lint tidy

test: ## Run all unit tests
	$(GO) test -v -race ./...

test-coverage: ## Run tests and generate HTML coverage report
	@mkdir -p coverage
	$(GO) test -race -coverprofile=coverage/coverage.out ./...
	$(GO) tool cover -html=coverage/coverage.out -o coverage/coverage.html
	@echo "Coverage report written to coverage/coverage.html"

fmt: ## Format Go source code
	$(GO) fmt ./...

vet: ## Run Go static analysis (vet)
	$(GO) vet ./...

tidy: ## Clean up and verify go.mod dependencies
	$(GO) mod tidy
	$(GO) mod verify

# ==============================================================================
# Cleanup
# ==============================================================================
.PHONY: clean
clean: ## Remove build artifacts, binaries, and test coverage
	rm -rf $(BIN_DIR) coverage/