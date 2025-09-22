# Testkit - Go Testing Utility Library

.PHONY: help build test clean lint fmt vet

# Default target
help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-15s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

# Build the project
build: ## Build all packages
	@echo "Building packages..."
	@go build ./...

# Run tests
test: ## Run all tests
	@echo "Running tests..."
	@go test ./... -v

# Run tests with coverage
test-coverage: ## Run tests with coverage report
	@echo "Running tests with coverage..."
	@go test ./... -cover -coverprofile=coverage.out
	@go tool cover -html=coverage.out -o coverage.html

# Clean build artifacts
clean: ## Clean build artifacts and temporary files
	@echo "Cleaning..."
	@rm -f coverage.out coverage.html
	@go clean ./...

# Format code
fmt: ## Format Go code
	@echo "Formatting code..."
	@go fmt ./...

# Vet code
vet: ## Run go vet
	@echo "Running go vet..."
	@go vet ./...

# Lint code (requires golangci-lint)
lint: ## Run golangci-lint
	@echo "Running linter..."
	@golangci-lint run ./...

# Run all checks
check: fmt vet lint test ## Run all checks (format, vet, lint, test)

# Install dependencies
deps: ## Download and install dependencies
	@echo "Installing dependencies..."
	@go mod download
	@go mod tidy

# Update dependencies
update: ## Update dependencies
	@echo "Updating dependencies..."
	@go get -u ./...
	@go mod tidy