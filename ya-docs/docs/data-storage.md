---
sidebar_position: 6
---

# Data & Storage

All Ya data lives in a single shared directory. The CLI, TUI, and [Ya-GUI](https://github.com/d3uceY/Ya-GUI) all read and write the same files — changes made in any one of them are immediately visible in the others.

---

## Data directory path

| Platform | Path |
|----------|------|
| Windows | `%APPDATA%\ya\data\` |
| macOS | `~/Library/Application Support/ya/data/` |
| Linux | `~/.config/ya/data/` |

The directory is created automatically on first run.

---

## Files

### `shortcuts.json`

A flat JSON object mapping shortcut names to commands. This is the CLI-compatible format.

```json
{
  "gs": "git status",
  "gt": "git tag -d v{tag}; git tag -a v{tag} -m \"{metadata}\"; git push --tags"
}
```

### `shortcuts-meta.json`

Metadata for each shortcut — descriptions, pin status, run counts, and last-run timestamps.

```json
{
  "gt": {
    "description": "template for pushing git tags",
    "pinned": true,
    "runCount": 12,
    "lastRun": "2026-06-11T14:32:00Z"
  }
}
```

### `history.json`

An array of run history entries, newest first. Capped at 500 entries.

```json
[
  {
    "shortcutName": "gt",
    "command": "git tag -d v1.2.0; git tag -a v1.2.0 -m \"release\"",
    "directory": "C:\\projects\\myapp",
    "timestamp": "2026-06-11T14:32:00Z"
  }
]
```

### `config.json`

Application config — currently stores saved directories.

```json
{
  "defaultDir": "",
  "savedDirectories": [
    { "name": "api", "path": "C:\\projects\\myapp\\api" }
  ]
}
```

---

## Backup & portability

To back up all your shortcuts, history, and config, just copy the entire data directory. To restore, copy it back. To share shortcuts between machines, copy `shortcuts.json` and `shortcuts-meta.json`.

You can also use `ya export` to get a plain `shortcuts.json` copy anywhere you like.
