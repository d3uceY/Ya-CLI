---
sidebar_position: 1
---

# Introduction

**Ya** is a lightweight command-line shortcut manager. Save long or complex commands under short, memorable names and run them instantly — from the CLI or the built-in full-screen TUI.

```bash
# Save a command once
ya add deploy 'kubectl rollout restart deployment/{service}'

# Run it any time
ya deploy
```

Ya prompts you to fill in any `{placeholder}` tokens before executing.

![Ya TUI](/img/tui-screenshot.png)

---

## What Ya gives you

| Feature | Description |
|---------|-------------|
| **Shortcuts** | Map any shell command to a short name |
| **TUI** | Full-screen interactive interface — browse, search, run, manage |
| **Descriptions** | Add an optional note to each shortcut |
| **Pin** | Float your most-used shortcuts to the top |
| **Run history** | Every execution is logged; browse it in the TUI |
| **Saved directories** | Keep frequently-used paths one keypress away |
| **Templates** | `{placeholder}` tokens prompt you for values at runtime |
| **Import / Export** | Backup and share shortcuts as JSON |
| **Shell completion** | Tab-complete shortcut names in Bash, Zsh, Fish, PowerShell |
| **GUI companion** | [Ya-GUI](https://github.com/d3uceY/Ya-GUI) shares the same data files |

---

## Pick your path

- **New to Ya?** Start with [Installation](./installation).
- **Want to use the CLI directly?** See [CLI Reference](./cli-reference).
- **Prefer a visual interface?** Open [TUI Guide](./tui/overview).
- **Using templates?** Read [Template Values](./templates).
