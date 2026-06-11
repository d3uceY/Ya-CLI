---
sidebar_position: 3
---

# Run History

Every shortcut you run — from the TUI or the CLI — is recorded in `history.json`. The TUI lets you browse that history.

---

## Opening history

Press `h` from the shortcuts list.

---

## History page keys

| Key | Action |
|-----|--------|
| `↑` / `k` | Move up |
| `↓` / `j` | Move down |
| `c` | Clear all history (with confirmation) |
| `Esc` / `h` | Go back to the shortcuts list |

---

## What's recorded

Each entry shows:

- **Shortcut name** — the key that was used
- **Command** — the resolved command (with templates already filled in)
- **Timestamp** — when it was run (RFC 3339, UTC)

The history is capped at **500 entries**. Older entries are dropped automatically when the cap is reached.

---

## Storage

History is saved to `history.json` in the shared data directory and is readable by [Ya-GUI](https://github.com/d3uceY/Ya-GUI) as well.
