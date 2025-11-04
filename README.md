# go-monorepo-demo

[![Test](https://github.com/mitchallen/go-monorepo-demo/actions/workflows/test.yml/badge.svg)](https://github.com/mitchallen/go-monorepo-demo/actions/workflows/test.yml)
[![Lint](https://github.com/mitchallen/go-monorepo-demo/actions/workflows/lint.yml/badge.svg)](https://github.com/mitchallen/go-monorepo-demo/actions/workflows/lint.yml)
[![Build](https://github.com/mitchallen/go-monorepo-demo/actions/workflows/build.yml/badge.svg)](https://github.com/mitchallen/go-monorepo-demo/actions/workflows/build.yml)

A comprehensive demonstration of Go monorepo patterns using Go workspaces. This project showcases how to structure, develop, and maintain multiple packages and applications in a single repository.

## Documentation

See the detailed article: [How to Create and Use a Go Monorepo (Golang, Workspaces)](/how-to-create-and-use-a-go-monorepo.md)

## Project Structure

```
go-monorepo-demo/
├── pkg/                    # Shared packages
│   ├── alpha/             # Example package with external dependency
│   ├── beta/              # Package depending on alpha and shared
│   └── shared/            # Common utilities (logger, math)
├── cmd/                    # Applications
│   ├── demo-app/          # CLI demo application
│   └── web-server/        # Web service with REST API
├── scripts/               # Development scripts
│   ├── test.sh           # Run all tests
│   ├── coverage.sh       # Generate coverage reports
│   ├── lint.sh           # Run linters
│   ├── build.sh          # Build all applications
│   └── fmt.sh            # Format code
├── .github/workflows/     # CI/CD pipelines
├── go.work               # Workspace configuration
└── Makefile              # Common development tasks
```

## Quick Start

### Setup

```sh
# Clone the repository
git clone https://github.com/mitchallen/go-monorepo-demo.git
cd go-monorepo-demo

# Setup workspace and install tools
make setup

# Or manually:
go work sync
go mod download
```

### Running Tests

```sh
# Run all tests
make test

# Run with coverage
make coverage

# Run linters
make lint
```

### Building Applications

```sh
# Build all applications
make build

# Binaries will be in bin/
./bin/demo-app --help
./bin/web-server --help
```

### Running Applications

**CLI Demo App:**

```sh
# Run in hello mode
./bin/demo-app --mode=hello

# Run coin flip analysis
./bin/demo-app --mode=analyze --flips=1000

# Compare multiple trials
./bin/demo-app --mode=compare

# Or run directly without building
go run ./cmd/demo-app --mode=analyze
```

**Web Server:**

```sh
# Start the web server
./bin/web-server --port=8080

# Or run directly
go run ./cmd/web-server

# Try the endpoints:
curl http://localhost:8080/health
curl http://localhost:8080/api/flip?count=100
curl http://localhost:8080/api/analyze?count=500
```

## Package Overview

### pkg/alpha

Basic package demonstrating external dependencies and simple functionality.

```go
import "github.com/mitchallen/go-monorepo-demo/pkg/alpha"

alpha.Hello()
results := alpha.CoinCount(100)
```

### pkg/beta

Demonstrates internal package dependencies (uses both `alpha` and `shared`).

```go
import "github.com/mitchallen/go-monorepo-demo/pkg/beta"

beta.Hello()
analysis := beta.AnalyzeCoinFlips(100)
```

### pkg/shared

Common utilities used across packages.

```go
import "github.com/mitchallen/go-monorepo-demo/pkg/shared"

logger := shared.NewLogger("myapp")
logger.Info("Starting application")

max := shared.Max(5, 10)
sum := shared.Sum([]int{1, 2, 3, 4, 5})
```

## Development

### Available Make Targets

```sh
make help              # Show all available targets
make setup             # Install tools and sync workspace
make test              # Run all tests
make lint              # Run linters
make fmt               # Format code
make build             # Build all applications
make coverage          # Run tests with coverage
make clean             # Clean build artifacts
make run-demo          # Build and run demo-app
make run-server        # Build and run web-server
make check             # Run fmt, lint, and test
make all               # Run fmt, lint, test, and build
```

### Development Workflow

```sh
# Before committing
make check

# Full build pipeline
make all
```

### Using the Devcontainer

This project includes a devcontainer configuration for VS Code:

1. Install Docker and VS Code with Remote-Containers extension
2. Open project in VS Code
3. Click "Reopen in Container" when prompted
4. All tools will be automatically installed

The devcontainer includes:
- Go 1.21
- golangci-lint
- Docker-in-Docker support
- GitHub CLI
- Automatic port forwarding for web-server (port 8080)

## CI/CD

This project uses GitHub Actions for continuous integration:

- **Test**: Runs all tests and generates coverage reports
- **Lint**: Runs golangci-lint across all modules
- **Build**: Builds all applications and uploads artifacts

## Workspace Management

The project uses Go workspaces to manage multiple modules:

```sh
# Sync workspace after adding new modules
go work sync

# Add a new module to workspace
go work use ./pkg/newmodule

# Recursively add all modules
go work use -r .

# Tidy all modules
go list -f '{{.Dir}}' -m | xargs -L1 go mod tidy -C
```

## Testing

### Run Tests for Specific Module

```sh
go test -C ./pkg/alpha
go test -C ./pkg/beta
go test -C ./pkg/shared
```

### Run Benchmarks

```sh
go test -bench=. ./pkg/shared
go test -bench=. ./pkg/alpha
```

### View Coverage HTML Report

```sh
make coverage
go tool cover -html=coverage/alpha.out
```

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Run `make check` to ensure tests pass
5. Submit a pull request

## License

See [LICENSE](LICENSE) for details.

## Author

Mitch Allen

## References

- [Go Workspaces Documentation](https://go.dev/doc/tutorial/workspaces)
- [How to Create and Use a Go Monorepo](/how-to-create-and-use-a-go-monorepo.md)

