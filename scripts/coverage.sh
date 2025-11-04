#!/bin/bash

# Run tests with coverage across all modules
echo "Running tests with coverage..."

# Create coverage directory if it doesn't exist
mkdir -p coverage

# Run tests for each module and collect coverage
go list -f '{{.Dir}}' -m | while read dir; do
    echo "Testing $dir..."
    go test -C "$dir" -coverprofile=coverage/$(basename "$dir").out -covermode=atomic ./...
done

echo ""
echo "=== Coverage Summary ==="
go list -f '{{.Dir}}' -m | while read dir; do
    coverage_file="coverage/$(basename "$dir").out"
    if [ -f "$coverage_file" ]; then
        echo ""
        echo "Module: $(basename "$dir")"
        go tool cover -func="$coverage_file" | tail -1
    fi
done

echo ""
echo "Coverage reports saved to coverage/ directory"
echo "To view HTML coverage report for a module, run:"
echo "  go tool cover -html=coverage/<module-name>.out"
