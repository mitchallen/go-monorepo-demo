.PHONY: help setup test lint fmt build coverage clean run-demo run-server all

help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Available targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-15s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

setup: ## Install development tools and sync workspace
	@echo "Installing development tools..."
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	@echo "Syncing workspace..."
	@go work sync
	@go list -f '{{.Dir}}' -m | xargs -L1 go mod tidy -C
	@echo "Setup complete!"

test: ## Run all tests
	@./scripts/test.sh

lint: ## Run linters
	@./scripts/lint.sh

fmt: ## Format all Go code
	@./scripts/fmt.sh

build: ## Build all applications
	@./scripts/build.sh

coverage: ## Run tests with coverage reports
	@./scripts/coverage.sh

clean: ## Clean build artifacts and coverage reports
	@echo "Cleaning build artifacts..."
	@rm -rf bin/
	@rm -rf coverage/
	@echo "Clean complete!"

run-demo: build ## Build and run the demo-app
	@echo "Running demo-app..."
	@./bin/demo-app

run-server: build ## Build and run the web-server
	@echo "Running web-server..."
	@./bin/web-server

all: fmt lint test build ## Run format, lint, test, and build

# Workspace management
workspace-sync: ## Sync workspace and tidy all modules
	@go work sync
	@go list -f '{{.Dir}}' -m | xargs -L1 go mod tidy -C
	@echo "Workspace synced!"

# Development helpers
dev-demo: ## Run demo-app without building (go run)
	@go run ./cmd/demo-app

dev-server: ## Run web-server without building (go run)
	@go run ./cmd/web-server

# Quick checks
check: fmt lint test ## Quick check before committing (format, lint, test)
	@echo "All checks passed!"
