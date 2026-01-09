#!/bin/bash
# Helper script to generate Flatpak dependency sources
# This script generates the necessary JSON files for offline builds

set -e

echo "Generating Flatpak dependency sources for Time Manager..."
echo ""

# Check if flatpak-node-generator is available
if ! command -v flatpak-node-generator &> /dev/null; then
    echo "⚠️  flatpak-node-generator not found"
    echo "   Install with: pip install flatpak-node-generator"
    echo "   Or from: https://github.com/flatpak/flatpak-builder-tools"
    MISSING_TOOLS=1
fi

# Check if flatpak-go-generator is available
if ! command -v flatpak-go-generator &> /dev/null; then
    echo "⚠️  flatpak-go-generator not found"
    echo "   Install from: https://github.com/flatpak/flatpak-builder-tools"
    MISSING_TOOLS=1
fi

if [ -n "$MISSING_TOOLS" ]; then
    echo ""
    echo "Please install the missing tools and try again."
    exit 1
fi

echo "✓ All tools found"
echo ""

# Generate Node.js dependencies
echo "Generating Node.js dependencies..."
if [ -f "frontend/package-lock.json" ]; then
    flatpak-node-generator npm frontend/package-lock.json -o flatpak/flatpak-node-sources.json
    echo "✓ Generated flatpak/flatpak-node-sources.json"
else
    echo "⚠️  frontend/package-lock.json not found"
fi

echo ""

# Generate Go dependencies
echo "Generating Go dependencies..."
if [ -f "go.mod" ]; then
    flatpak-go-generator go.mod -o flatpak/flatpak-go-sources.json
    echo "✓ Generated flatpak/flatpak-go-sources.json"
else
    echo "⚠️  go.mod not found"
fi

echo ""
echo "✅ Done! Dependency sources generated."
echo ""
echo "Next steps:"
echo "1. Update flatpak/dev.sarkis.time-manager-source.yaml to include the generated sources"
echo "2. Test the build with: flatpak-builder --force-clean build flatpak/dev.sarkis.time-manager-source.yaml"
