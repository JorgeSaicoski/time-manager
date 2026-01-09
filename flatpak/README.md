# Flatpak Configuration for Skopii Time Manager

This directory contains the Flatpak manifest and related files for building and distributing the Time Manager application.

## Files

- `dev.sarkis.time-manager.yaml` - Flatpak manifest (current, uses pre-built binary)
- `dev.sarkis.time-manager-source.yaml` - Improved manifest for building from source (Flathub-ready)
- `generate-sources.sh` - Helper script to generate dependency sources for offline builds
- `files/dev.sarkis.time-manager.desktop` - Desktop entry file
- `files/dev.sarkis.time-manager.metainfo.xml` - AppStream metadata
- `files/dev.sarkis.time-manager.png` - Application icon (128x128)
- `files/time-manager` - Pre-built application binary

## Current Status

### ✅ Completed (Critical Blockers Fixed)

1. **Added `<launchable>` tag** - Required by Flathub for application launching
2. **Updated license** - Changed from GPL-3.0-or-later to MIT to match project
3. **Updated release date** - Changed from outdated 2024-01-01 to current date
4. **Improved desktop file** - Added GenericName, Terminal, and StartupNotify fields
5. **Added screenshot placeholder** - Created screenshots directory
6. **Documented finish-args** - Added comments explaining permissions

### ⚠️ Remaining Work for Flathub Submission

1. **Build from Source** - A source-based manifest has been created (`dev.sarkis.time-manager-source.yaml`):
   - Uses Go SDK extension (org.freedesktop.Sdk.Extension.golang)
   - Uses Node.js extension for building the Svelte frontend
   - Requires generating dependency sources using `generate-sources.sh`
   - May need adjustments based on actual Wails build process
   - **Action Required**: Test and refine the source-based build process

2. **Real Screenshots** - Replace the placeholder screenshot with actual application screenshots showing:
   - Main time tracking interface
   - Project management view
   - Settings/configuration screen
   - At least 1280x720 resolution, showing real app usage

3. **Icon Sizes** - Consider providing multiple icon sizes:
   - 48x48, 64x64, 128x128, 256x256, 512x512

4. **Testing** - Validate the manifest using:
   ```bash
   flatpak-builder --install-deps-from=flathub build dev.sarkis.time-manager.yaml
   flatpak run --command=flatpak-builder-lint org.flatpak.Builder manifest dev.sarkis.time-manager.yaml
   flatpak run --command=flatpak-builder-lint org.flatpak.Builder appstream files/dev.sarkis.time-manager.metainfo.xml
   ```

## Building Locally

### Quick Build (Pre-built Binary)

```bash
# Build using the current manifest with pre-built binary
flatpak-builder --force-clean build dev.sarkis.time-manager.yaml

# Install locally
flatpak-builder --user --install --force-clean build dev.sarkis.time-manager.yaml

# Run the application
flatpak run dev.sarkis.time-manager
```

### Source-based Build (Flathub-ready)

```bash
# First, generate dependency sources
./generate-sources.sh

# Build from source
flatpak-builder --force-clean build dev.sarkis.time-manager-source.yaml

# Install locally
flatpak-builder --user --install --force-clean build dev.sarkis.time-manager-source.yaml

# Run the application
flatpak run dev.sarkis.time-manager
```

## References

- [Flathub Requirements](https://docs.flathub.org/docs/for-app-authors/requirements)
- [Flathub MetaInfo Guidelines](https://docs.flathub.org/docs/for-app-authors/metainfo-guidelines)
- [Flathub Submission Process](https://docs.flathub.org/docs/for-app-authors/submission)
- [AppStream Specification](https://www.freedesktop.org/software/appstream/docs/)
