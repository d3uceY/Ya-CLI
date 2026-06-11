---
sidebar_position: 2
---

# Installation

## Download a pre-built binary

Pre-built binaries are available for Windows, macOS, and Linux on the [Releases page](https://github.com/d3uceY/Ya-CLI/releases/latest).

Download the binary for your platform, place it somewhere on your `PATH`, and you're done.

---

## Homebrew (macOS & Linux)

```bash
brew tap d3uceY/homebrew-ya
brew install ya
```

---

## Build from source

**Prerequisites:** [Go 1.25.5+](https://go.dev/dl/)

```bash
git clone https://github.com/d3uceY/Ya-CLI
cd Ya-CLI/ya-cli
```

### Windows
```powershell
go build -o ya.exe .
```

### Linux / macOS
```bash
go build -o ya .
chmod +x ya
```

---

## Adding to PATH

### Windows (PowerShell, run as Administrator)

```powershell
New-Item -Path "C:\Program Files\Ya" -ItemType Directory -Force
Move-Item .\ya.exe "C:\Program Files\Ya\ya.exe"
[Environment]::SetEnvironmentVariable("Path", $env:Path + ";C:\Program Files\Ya", "User")
```

Restart your terminal, then verify:

```powershell
ya version
```

### Linux / macOS

```bash
sudo mv ya /usr/local/bin/
ya version
```

---

## Data directory

Ya stores all data in your user config directory. The path is the same for the CLI, the TUI, and [Ya-GUI](https://github.com/d3uceY/Ya-GUI) — all three share the same files.

| Platform | Path |
|----------|------|
| Windows | `%APPDATA%\ya\data\` |
| macOS | `~/Library/Application Support/ya/data/` |
| Linux | `~/.config/ya/data/` |

The directory is created automatically on first run.
