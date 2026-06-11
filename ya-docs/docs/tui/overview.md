---
sidebar_position: 1
---

# TUI Overview

Run `ya` with no arguments to open the interactive full-screen terminal UI:

```bash
ya
```

![Ya TUI — main shortcuts list](/img/tui-screenshot.png)

The TUI lets you browse, search, run, and manage all your shortcuts without typing subcommands. Every feature available in the CLI is also available here, plus extras like run history, saved directories, and pinning.

---

## Pages

| Page | How to get there |
|------|-----------------|
| **Shortcuts list** | Default view on launch |
| **Search** | Press `/` |
| **Add / Edit / Rename** | `a`, `e`, `r` on the list |
| **Run history** | `h` |
| **Saved directories** | `D` |
| **Template fill-in** | Triggered automatically when running a shortcut with `{tokens}` |
| **Help** | `?` |

---

## Global keys

| Key | Action |
|-----|--------|
| `q` / `Ctrl+C` | Quit |
| `Esc` | Go back / cancel |
| `?` | Toggle help page |
| `f` | Toggle fullscreen |

---

## How commands run

When you press `Enter` on a shortcut, the TUI **exits cleanly first**, then runs the command natively in your terminal. This means:

- Full terminal access — interactive commands, colors, pagers all work normally
- No output buffering or keystroke-replay issues
- The run is logged to [history](./history) after it completes
