---
sidebar_position: 3
---

# CLI Reference

All commands are available directly from your terminal. Run `ya help` for a quick summary.

---

## `ya` — run a shortcut

```bash
ya <shortcut>
```

Looks up `<shortcut>` and runs its command. If the command contains `{placeholder}` tokens, Ya prompts you to fill them in first.

**Passing extra arguments:**

Extra arguments are appended to the command at runtime:

```bash
ya gcm -m "fix login bug"
# runs: git commit -m 'fix login bug'
```

---

## `ya add`

```bash
ya add <name> '<command>'
```

Creates a new shortcut. If a shortcut with that name already exists, Ya asks before overwriting.

**Examples:**

```bash
ya add gs 'git status'
ya add deploy 'kubectl rollout restart deployment/{service}'
ya add dev 'cd ~/projects/myapp && npm run dev'
```

---

## `ya remove`

```bash
ya remove <name>
```

Deletes a shortcut after confirmation.

---

## `ya rename`

```bash
ya rename <old-name> <new-name>
```

Renames a shortcut key without changing its command.

---

## `ya list`

```bash
ya list
```

Prints all shortcuts with their commands and a total count.

---

## `ya show`

```bash
ya show <name>
```

Prints the command mapped to a single shortcut.

---

## `ya search`

```bash
ya search <term>
```

Finds shortcuts whose name **or** command contains the search term (case-insensitive).

---

## `ya import`

```bash
ya import <file-path>
```

Imports shortcuts from a JSON file. The file should be a flat object mapping names to commands:

```json
{
  "gs": "git status",
  "gp": "git pull"
}
```

Duplicate keys are overwritten.

---

## `ya export`

```bash
ya export <directory>
ya export <directory> --name <filename>
ya export <directory> -n <filename>
```

Exports all shortcuts to a JSON file. Defaults to `shortcuts.json` if `--name` is not provided.

```bash
ya export ./backups --name jan-2026.json
```

---

## `ya version`

```bash
ya version
```

Prints the current Ya version.

---

## `ya completion`

```bash
ya completion <shell>
```

Generates a shell completion script. See [Shell Completion](./shell-completion) for setup instructions.

---

## `ya` (no arguments)

Launches the interactive [TUI](./tui/overview).
