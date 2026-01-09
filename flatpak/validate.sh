#!/bin/bash
# Validation script for Flatpak files
# Checks common issues before submission to Flathub

set -e

echo "🔍 Validating Flatpak configuration..."
echo ""

ERRORS=0
WARNINGS=0

# Check if required files exist
echo "Checking required files..."
REQUIRED_FILES=(
    "flatpak/files/dev.sarkis.time-manager.desktop"
    "flatpak/files/dev.sarkis.time-manager.metainfo.xml"
    "flatpak/files/dev.sarkis.time-manager.png"
)

for file in "${REQUIRED_FILES[@]}"; do
    if [ -f "$file" ]; then
        echo "  ✓ $file exists"
    else
        echo "  ✗ $file missing"
        ((ERRORS++))
    fi
done

echo ""

# Check desktop file format
echo "Checking desktop file format..."
DESKTOP_FILE="flatpak/files/dev.sarkis.time-manager.desktop"
if [ -f "$DESKTOP_FILE" ]; then
    if grep -q "\[Desktop Entry\]" "$DESKTOP_FILE"; then
        echo "  ✓ Desktop Entry section found"
    else
        echo "  ✗ Missing [Desktop Entry] section"
        ((ERRORS++))
    fi
    
    REQUIRED_KEYS=("Name" "Exec" "Icon" "Type")
    for key in "${REQUIRED_KEYS[@]}"; do
        if grep -q "^$key=" "$DESKTOP_FILE"; then
            echo "  ✓ $key field present"
        else
            echo "  ✗ Missing $key field"
            ((ERRORS++))
        fi
    done
fi

echo ""

# Check metainfo file format
echo "Checking metainfo.xml format..."
METAINFO_FILE="flatpak/files/dev.sarkis.time-manager.metainfo.xml"
if [ -f "$METAINFO_FILE" ]; then
    if grep -q "<component" "$METAINFO_FILE"; then
        echo "  ✓ Component tag found"
    else
        echo "  ✗ Missing component tag"
        ((ERRORS++))
    fi
    
    REQUIRED_TAGS=("id" "name" "summary" "description" "launchable" "metadata_license" "project_license")
    for tag in "${REQUIRED_TAGS[@]}"; do
        if grep -q "<$tag" "$METAINFO_FILE"; then
            echo "  ✓ <$tag> tag present"
        else
            echo "  ✗ Missing <$tag> tag"
            ((ERRORS++))
        fi
    done
    
    # Check for screenshots
    if grep -q "<screenshots>" "$METAINFO_FILE"; then
        echo "  ✓ Screenshots section present"
        if grep -q "favicon.png" "$METAINFO_FILE"; then
            echo "  ⚠ Screenshot uses favicon.png - should use real app screenshot"
            ((WARNINGS++))
        fi
    else
        echo "  ✗ Missing screenshots section"
        ((ERRORS++))
    fi
fi

echo ""

# Check icon file
echo "Checking icon file..."
ICON_FILE="flatpak/files/dev.sarkis.time-manager.png"
if [ -f "$ICON_FILE" ]; then
    if file "$ICON_FILE" | grep -q "PNG"; then
        echo "  ✓ Icon is PNG format"
        # Try to extract size, fallback gracefully if not available
        if command -v file > /dev/null && file "$ICON_FILE" | grep -qE '[0-9]+ x [0-9]+'; then
            SIZE=$(file "$ICON_FILE" | grep -oE '[0-9]+ x [0-9]+' | head -1)
            echo "  ℹ Icon size: $SIZE"
            if [[ "$SIZE" == "128 x 128" ]]; then
                echo "  ✓ Icon is 128x128"
            else
                echo "  ⚠ Icon should be 128x128"
                ((WARNINGS++))
            fi
        else
            echo "  ℹ Icon size: Could not determine (file command may not support image info)"
        fi
    else
        echo "  ✗ Icon is not PNG format"
        ((ERRORS++))
    fi
fi

echo ""

# Check manifests
echo "Checking Flatpak manifests..."
if [ -f "flatpak/dev.sarkis.time-manager.yaml" ]; then
    echo "  ✓ Main manifest exists"
    if grep -q "id: dev.sarkis.time-manager" "flatpak/dev.sarkis.time-manager.yaml"; then
        echo "  ✓ Manifest has correct ID"
    fi
fi

if [ -f "flatpak/dev.sarkis.time-manager-source.yaml" ]; then
    echo "  ✓ Source-based manifest exists"
else
    echo "  ⚠ Source-based manifest not found (recommended for Flathub)"
    ((WARNINGS++))
fi

echo ""
echo "============================================"
echo "Validation Summary:"
echo "  Errors: $ERRORS"
echo "  Warnings: $WARNINGS"
echo "============================================"

if [ $ERRORS -eq 0 ] && [ $WARNINGS -eq 0 ]; then
    echo "✅ All checks passed!"
    exit 0
elif [ $ERRORS -eq 0 ]; then
    echo "⚠️  Passed with warnings"
    exit 0
else
    echo "❌ Validation failed - please fix errors"
    exit 1
fi
