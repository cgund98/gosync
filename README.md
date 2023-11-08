# gosync
A file syncing tool based on Go and rsync.

## Installation

Install the latest binary from the releases section.

```bash
sudo mv gosync /usr/local/bin/gosync
```

## Usage

### Configuration
`gosync` expects the user to configure their backups via `yaml`.

See `gosync.example.yaml` for an example configuration file. 

There are two places that `gosync` will look for this configuration file.
1. A path specified with `--config` flag.
2. `$HOME/.gosync.yaml`, this is the default location.

### Running
```bash
# Validate the default configuration file. This won't touch any actual files
gosync parse

# Validate a specific configuration file.
gosync parse --config ./gosync.example.yaml

# Sync a file with a specific configuration path
gosync sync
```