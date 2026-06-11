# Contributing to Ya CLI

Thanks for your interest in contributing! This file covers everything you need to build, run, and submit changes.

---

## Prerequisites

- [Go 1.25.5+](https://go.dev/dl/)
- Windows, macOS, or Linux

---

## Build from Source

```bash
git clone https://github.com/d3uceY/Ya-CLI
cd Ya-CLI/ya-cli
```

**Windows:**
```powershell
go build -o ya.exe .
```

**Linux / macOS:**
```bash
go build -o ya .
chmod +x ya
```

---

## Adding to PATH

### Windows (PowerShell, as Administrator)

```powershell
New-Item -Path "C:\Program Files\Ya" -ItemType Directory -Force
Move-Item .\ya.exe "C:\Program Files\Ya\ya.exe"
[Environment]::SetEnvironmentVariable("Path", $env:Path + ";C:\Program Files\Ya", "User")
```

Restart your terminal, then: `ya version`

### Linux / macOS

```bash
sudo mv ya /usr/local/bin/
ya version
```

---

## Project Structure

```
ya-cli/
  main.go           # CLI entry point — all cobra subcommands
  tui/
    model.go        # TUI model struct + data helpers
    update.go       # Message handlers / state transitions
    view.go         # View router
    view_list.go    # Main shortcuts list page
    view_form.go    # Shared add/edit/rename/import/export form
    view_history.go # History, saved dirs, confirm pages
    view_help.go    # Help/keybindings page
    view_template.go# Template fill-in page
    tui.go          # Entry point — starts bubbletea, runs command after quit
    keys.go         # All key bindings
    styles.go       # All lipgloss styles and shared renderers
  utils/
    utilities.go    # Version, data dir, ResolveTemplates
    types.go        # Shared type definitions (matches ya-gui)
    shortcut.go     # Shortcut CRUD (two-file format: shortcuts.json + shortcuts-meta.json)
    history.go      # Run history persistence
    config.go       # App config (saved directories)
    colors.go       # Shared lipgloss colour palette for CLI output
ya-gui/             # Separate Wails v2 GUI app (shares the same data files)
ya-docs/            # Docusaurus documentation site
```

---

## Key Design Decisions

**Quit-then-exec pattern** — The TUI exits before running a command. This avoids the Windows console issue where buffered subprocess output replays as keystrokes into bubbletea after the process exits.

**Shared data directory** — `os.UserConfigDir()/ya/data/` is used by both the CLI/TUI and Ya-GUI. The two-file format (`shortcuts.json` for CLI compat, `shortcuts-meta.json` for rich metadata) keeps both apps in sync without breaking the simple key→command format.

**Colour palette** — All CLI output uses lipgloss with the same hex colours as the TUI (`#5F9EA0` accent, `#7EC8A4` success, `#CC6666` error, `#666666` dim). The palette is defined once in `utils/colors.go`.

---

## Running Tests

```bash
cd ya-cli
go test ./...
```

---

## Submitting Changes

1. Fork the repo and create a branch from `main`
2. Make your changes — keep them focused on one thing
3. Run `go build ./...` and confirm there are no errors
4. Open a pull request with a clear description of what changed and why

For larger changes or new features, open an issue first to discuss direction before writing code.
