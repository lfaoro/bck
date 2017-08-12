# bck

> The purpose of this tool is to simplify backup/restore operations to/from multiple origins/destinations.

## Features

- automatically excludes unavailabled paths from the backup job
- backup to local or remote locations
- concurrent backups and restores
- easy restore

## Use cases
- backup configuration files to multiple locations e.g. multiple drives, remote servers, cloud storage
- restore all your configurations into the correct system paths on a new machine

## Problems solved
- setting up a new computer
- keeping multiple computer settings in sync
- control over your backups (no abstractions)
- security of your backups (encryption)
- automation (daemon) backup and sync automatically 
- privacy (there's no cloud in the middle)

## Dependencies
- rsync

## Download

## Usage

```
Usage of bck:
  -debug
    	activate debug mode.
  -restore
    	restore files in destinations to their origins.
```

## Contributing

All dependencies are vendored, so there's nothing to download except Go and build.

### Install `Go` 
- macOS `brew install go`
    - Package Manager: [brew.sh](https://brew.sh)
- Windows `choco install golang`
    - Package Manager: [Chocolatey](https://chocolatey.org/install)

### Build
- make
