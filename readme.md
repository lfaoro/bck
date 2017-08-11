# bck

> The purpose of this tool is to simplify backup/restore operations to multiple destinations.

## Features

- automatically excludes unavailabled paths from the backup job
- backup to local or remote locations
- concurrent backups and restores
- easy restore

## Dependencies
- rsync

## Download

## Usage

```bash
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

## License
- MIT