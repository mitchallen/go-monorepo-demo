#!/bin/bash

# Run golangci-lint across all modules
echo "Running linters..."

# Check if golangci-lint is installed
if ! command -v golangci-lint &> /dev/null; then
    echo "golangci-lint not found. Installing..."
    go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
fi

# Run linter for each module
go list -f '{{.Dir}}' -m | while read dir; do
    echo ""
    echo "Linting $dir..."
    (cd "$dir" && golangci-lint run ./...)
done

echo ""
echo "Linting complete!"
