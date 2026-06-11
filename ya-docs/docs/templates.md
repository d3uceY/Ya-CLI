---
sidebar_position: 4
---

# Template Values

Templates let you build reusable commands with variable parts. Add `{placeholder}` tokens to any command — Ya will prompt you to fill them in each time you run it.

---

## Adding a templated shortcut

```bash
ya add commit 'git commit -m {message}'
ya add deploy 'kubectl set image deployment/{app} {app}={image}:{tag}'
ya add tag 'git tag -d v{tag}; git tag -a v{tag} -m "{metadata}"; git push --tags'
```

---

## Running a templated shortcut

### CLI

```bash
ya commit
```

```
→ This command has template values. Fill them in below:

  [1/1] message: fix null pointer in auth handler
```

Ya substitutes the value and runs:

```
git commit -m fix null pointer in auth handler
```

### TUI

In the TUI, pressing `Enter` on a shortcut with template tokens opens a fill-in form. After you submit, the TUI exits and the resolved command runs natively in your terminal.

---

## Rules

- Token names can contain letters, numbers, hyphens, and underscores: `{my-token}`, `{service_name}`
- The **same token used multiple times** in one command is only prompted once
- Tokens are **case-sensitive** — `{Tag}` and `{tag}` are separate
- Commands with **no tokens** run immediately without any prompts
