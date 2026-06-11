---
sidebar_position: 2
---

# Shortcuts List

The shortcuts list is the main page of the TUI. It shows all your saved shortcuts, sorted with pinned items first, then alphabetically.

---

## Navigation

| Key | Action |
|-----|--------|
| `↑` / `k` | Move cursor up |
| `↓` / `j` | Move cursor down |
| `Enter` | Run selected shortcut |

---

## Managing shortcuts

| Key | Action |
|-----|--------|
| `a` | Add a new shortcut |
| `e` | Edit selected shortcut (command + description) |
| `r` | Rename selected shortcut |
| `d` / `x` | Delete selected shortcut (with confirmation) |
| `p` | Pin / unpin selected shortcut |
| `i` | Import shortcuts from a JSON file |
| `o` | Export shortcuts to a JSON file |

---

## The add / edit form

**Add** (`a`) opens a 3-field form:

- **name** — the shortcut key
- **command** — the shell command to run (supports `{placeholder}` tokens)
- **description** *(optional)* — a short note shown beneath the row when selected

**Edit** (`e`) opens a 2-field form:

- **command**
- **description**

Press `Tab` to move between fields, `Enter` to submit, `Esc` to cancel.

---

## Descriptions

When a selected shortcut has a description, it appears as a sub-line beneath the row:

```
  ❯  gt    git tag -d v{tag}; git tag -a v{tag} -m "{metadata}"
             ──────────────────────────────────────────────────
             template for pushing git tags
```

Descriptions are also searched when you use the `/` search.

---

## Pinning

Press `p` on any shortcut to pin it. Pinned shortcuts:

- Float to the **top** of the list
- Show a `◆` indicator next to the command

Press `p` again to unpin.

---

## Search

Press `/` to activate search. Type to filter the list in real time — matches against the shortcut name, command, and description.

While search is active:
- `↑` / `↓` navigate the filtered results
- Letter keys (including `j` and `k`) type into the search box
- `Enter` or `Esc` closes search (keeping or clearing the filter)
