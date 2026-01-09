# Contributing to Flatpak Packaging

This guide helps contributors improve the Flatpak packaging for Time Manager.

## Overview

The Time Manager application is being prepared for publication on Flathub. Critical blockers have been addressed, but some work remains for full Flathub compliance.

## What's Been Done

✅ **Critical Blockers Fixed:**
- Added required `<launchable>` tag to metainfo.xml
- Fixed license inconsistency (now MIT)
- Updated release date to current
- Improved desktop file with required fields
- Created source-based build manifest template
- Documented all changes and remaining work

## What Still Needs Work

### 1. Real Application Screenshots

**Priority: High**

The current screenshot is a placeholder. Flathub requires real screenshots showing:
- Minimum 1280x720 resolution
- Real application usage (not logos or generic images)
- Multiple screenshots showing different features

**How to help:**
1. Build and run the application
2. Take screenshots of:
   - Main time tracking interface
   - Project management view
   - Settings/configuration
   - Reports or analytics views
3. Save screenshots as PNG in the `screenshots/` directory
4. Update `flatpak/files/dev.sarkis.time-manager.metainfo.xml` with proper image URLs
5. Commit to repository so they can be referenced

### 2. Test and Refine Source Build

**Priority: High**

The `dev.sarkis.time-manager-source.yaml` manifest needs testing:

```bash
# Install required tools
pip install flatpak-node-generator
git clone https://github.com/flatpak/flatpak-builder-tools.git

# Generate dependency sources
cd flatpak
./generate-sources.sh

# Test the build
flatpak-builder --force-clean build dev.sarkis.time-manager-source.yaml
```

**Known issues to address:**
- Wails may require specific build steps not captured in the manifest
- Dependencies may need adjustments
- Build flags may need tuning

### 3. Run Flatpak Linter

**Priority: Medium**

Validate all files with Flathub's official linter:

```bash
# Install the linter
flatpak install flathub org.flatpak.Builder

# Lint the manifest
flatpak run --command=flatpak-builder-lint org.flatpak.Builder manifest flatpak/dev.sarkis.time-manager-source.yaml

# Lint the appstream metadata
flatpak run --command=flatpak-builder-lint org.flatpak.Builder appstream flatpak/files/dev.sarkis.time-manager.metainfo.xml

# Lint the built flatpak (after building)
flatpak run --command=flatpak-builder-lint org.flatpak.Builder flatpak flatpak/build
```

Fix any errors or warnings reported by the linter.

### 4. Additional Icon Sizes

**Priority: Low**

Provide multiple icon sizes for better desktop integration:
- 48x48
- 64x64
- 128x128 (already exists)
- 256x256
- 512x512

These should be placed in `flatpak/files/` and installed to appropriate paths.

### 5. Enhanced Metadata

**Priority: Low**

Consider adding to metainfo.xml:
- Brand colors for software center
- More detailed release notes
- Links to bug tracker
- More screenshots
- Video demonstration (optional)

## Testing Checklist

Before submitting to Flathub:

- [ ] Application builds successfully from source
- [ ] Application runs and all features work
- [ ] All linter checks pass with no errors
- [ ] Real screenshots are included
- [ ] Desktop file is valid
- [ ] MetaInfo file is valid
- [ ] Icons display correctly
- [ ] Application appears in desktop menu
- [ ] Application has correct permissions

## Submitting to Flathub

Once all blockers are resolved:

1. Fork the Flathub repository: https://github.com/flathub/flathub
2. Create a new repository for the app in Flathub organization
3. Submit the manifest files via pull request
4. Work with Flathub reviewers to address any feedback

See: https://docs.flathub.org/docs/for-app-authors/submission

## Resources

- [Flathub Documentation](https://docs.flathub.org/)
- [AppStream Specification](https://www.freedesktop.org/software/appstream/docs/)
- [Flatpak Builder Tools](https://github.com/flatpak/flatpak-builder-tools)
- [Desktop Entry Specification](https://specifications.freedesktop.org/desktop-entry-spec/latest/)

## Questions?

Open an issue in the repository if you need help or have questions about Flatpak packaging.
