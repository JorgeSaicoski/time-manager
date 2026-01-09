# Flatpak/Flathub Publishing Blockers - Status Report

This document tracks all critical blockers that must be resolved for successful Flathub publication.

**Last Updated:** 2026-01-09

## Critical Blockers (Must Fix)

### ✅ FIXED: Missing `<launchable>` Tag
- **Status:** RESOLVED
- **Priority:** CRITICAL
- **Issue:** The metainfo.xml was missing the required `<launchable>` tag
- **Solution:** Added `<launchable type="desktop-id">dev.sarkis.time-manager.desktop</launchable>`
- **Impact:** Flathub requires this tag for proper application launching
- **Reference:** [Flathub MetaInfo Guidelines](https://docs.flathub.org/docs/for-app-authors/metainfo-guidelines)

### ✅ FIXED: License Inconsistency
- **Status:** RESOLVED
- **Priority:** CRITICAL
- **Issue:** README.md stated MIT license, but metainfo.xml stated GPL-3.0-or-later
- **Solution:** Updated metainfo.xml to use MIT license
- **Impact:** License mismatches can block Flathub submission
- **Files Changed:**
  - `flatpak/files/dev.sarkis.time-manager.metainfo.xml` (line 39)

### ✅ FIXED: Outdated Release Date
- **Status:** RESOLVED
- **Priority:** MEDIUM
- **Issue:** Release date was set to 2024-01-01 (outdated)
- **Solution:** Updated to current date (2026-01-09)
- **Impact:** Helps users understand when the application was released
- **Files Changed:**
  - `flatpak/files/dev.sarkis.time-manager.metainfo.xml` (line 30)

### ✅ FIXED: Incomplete Desktop File
- **Status:** RESOLVED
- **Priority:** MEDIUM
- **Issue:** Desktop file was missing recommended fields
- **Solution:** Added GenericName, Terminal=false, and StartupNotify=true
- **Impact:** Better desktop integration and user experience
- **Files Changed:**
  - `flatpak/files/dev.sarkis.time-manager.desktop`

### ✅ FIXED: Undocumented Manifest Permissions
- **Status:** RESOLVED
- **Priority:** LOW
- **Issue:** Flatpak manifest finish-args lacked explanatory comments
- **Solution:** Added comments explaining each permission
- **Impact:** Makes it easier for reviewers to understand permissions
- **Files Changed:**
  - `flatpak/dev.sarkis.time-manager.yaml`

## High Priority (Recommended for Submission)

### ⚠️ PARTIAL: Build from Source
- **Status:** IN PROGRESS
- **Priority:** HIGH
- **Issue:** Current manifest uses pre-built binary; Flathub requires building from source
- **Solution in Progress:**
  - Created `dev.sarkis.time-manager-source.yaml` with proper Go and Node.js SDK
  - Created `generate-sources.sh` script for dependency generation
  - Added comprehensive documentation
- **Remaining Work:**
  - Test the source-based build process
  - Generate dependency sources (requires flatpak-builder-tools)
  - Adjust build commands if Wails requires specific steps
  - Verify the build produces a working binary
- **How to Complete:**
  ```bash
  # Install tools
  pip install flatpak-node-generator
  
  # Generate sources
  cd flatpak && ./generate-sources.sh
  
  # Test build
  flatpak-builder --force-clean build dev.sarkis.time-manager-source.yaml
  ```

### ⚠️ TODO: Real Application Screenshots
- **Status:** NOT STARTED
- **Priority:** HIGH
- **Issue:** Current screenshot is a placeholder (favicon)
- **Requirements:**
  - Minimum 1280x720 resolution
  - Show actual application in use
  - Multiple screenshots recommended (3-5)
  - PNG format
- **Remaining Work:**
  1. Build and run the application
  2. Take screenshots of:
     - Main time tracking interface
     - Project management view
     - Settings/configuration
     - Reports or analytics
  3. Save to `screenshots/` directory
  4. Update metainfo.xml with proper URLs
  5. Push to repository
- **Current File:** `screenshots/main-window.png` (placeholder)

## Medium Priority (Quality Improvements)

### ⚠️ TODO: Validate with Flatpak Linter
- **Status:** NOT STARTED
- **Priority:** MEDIUM
- **Issue:** Files haven't been validated with official Flathub linter
- **Requirements:**
  - All errors must be fixed
  - All warnings should be addressed
- **How to Complete:**
  ```bash
  flatpak install flathub org.flatpak.Builder
  flatpak run --command=flatpak-builder-lint org.flatpak.Builder \
    manifest flatpak/dev.sarkis.time-manager-source.yaml
  flatpak run --command=flatpak-builder-lint org.flatpak.Builder \
    appstream flatpak/files/dev.sarkis.time-manager.metainfo.xml
  ```

### ⚠️ TODO: Multiple Icon Sizes
- **Status:** NOT STARTED
- **Priority:** LOW
- **Issue:** Only 128x128 icon provided
- **Recommended:** Provide icons in multiple sizes
  - 48x48
  - 64x64
  - 128x128 (✓ exists)
  - 256x256
  - 512x512
- **Impact:** Better appearance across different display contexts

## Documentation and Tooling

### ✅ COMPLETE: Documentation
- **Status:** COMPLETE
- Created comprehensive documentation:
  - `flatpak/README.md` - Overview and build instructions
  - `flatpak/CONTRIBUTING.md` - Guide for contributors
  - `FLATPAK_BLOCKERS.md` - This status report

### ✅ COMPLETE: Validation Tools
- **Status:** COMPLETE
- Created validation and helper scripts:
  - `flatpak/validate.sh` - Validates Flatpak files
  - `flatpak/generate-sources.sh` - Generates dependency sources
- **Usage:**
  ```bash
  # Validate files
  ./flatpak/validate.sh
  
  # Generate dependency sources
  ./flatpak/generate-sources.sh
  ```

## Summary

### Fixed (6 items)
- ✅ Missing `<launchable>` tag
- ✅ License inconsistency
- ✅ Outdated release date
- ✅ Incomplete desktop file
- ✅ Undocumented manifest permissions
- ✅ Documentation and tooling

### In Progress (1 item)
- ⚠️ Build from source (manifest created, needs testing)

### TODO (3 items)
- ⚠️ Real application screenshots
- ⚠️ Flatpak linter validation
- ⚠️ Multiple icon sizes

## Next Steps

1. **Test source-based build** - Highest priority
   - Install required tools
   - Generate dependency sources
   - Build and test the application

2. **Add real screenshots** - Required for quality submission
   - Run the application
   - Capture high-quality screenshots
   - Update metainfo.xml

3. **Run linter** - Final validation
   - Install Flatpak linter
   - Fix any reported issues
   - Re-run until clean

4. **Submit to Flathub** - Once all above are complete
   - Follow submission process at https://docs.flathub.org/docs/for-app-authors/submission

## Resources

- [Flathub Requirements](https://docs.flathub.org/docs/for-app-authors/requirements)
- [Flathub Submission Process](https://docs.flathub.org/docs/for-app-authors/submission)
- [AppStream Specification](https://www.freedesktop.org/software/appstream/docs/)
- [Flatpak Builder Tools](https://github.com/flatpak/flatpak-builder-tools)
