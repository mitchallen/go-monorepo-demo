#!/bin/bash

# Build all applications in the cmd directory
echo "Building applications..."

# Create bin directory if it doesn't exist
mkdir -p bin

# Find all cmd directories and build them
for cmd_dir in cmd/*; do
    if [ -d "$cmd_dir" ]; then
        app_name=$(basename "$cmd_dir")
        echo "Building $app_name..."
        go build -o "bin/$app_name" "./$cmd_dir"

        if [ $? -eq 0 ]; then
            echo "✓ Successfully built bin/$app_name"
        else
            echo "✗ Failed to build $app_name"
            exit 1
        fi
    fi
done

echo ""
echo "All applications built successfully!"
echo "Binaries are in the bin/ directory"
