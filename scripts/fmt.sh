#!/bin/bash

# Format all Go code in the repository
echo "Formatting Go code..."

# Run gofmt on all Go files
go fmt ./...

# Also run goimports if available
if command -v goimports &> /dev/null; then
    echo "Running goimports..."
    goimports -w .
fi

echo "Formatting complete!"
