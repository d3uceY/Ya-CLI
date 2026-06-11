---
sidebar_position: 5
---

# Shell Completion

Ya can tab-complete shortcut names in your shell. After setup, pressing `Tab` after `ya` will suggest your saved shortcuts. This works for `ya show`, `ya remove`, `ya rename`, and `ya search` as well.

---

## PowerShell (Windows)

First, ensure your PowerShell profile file and directory exist:

```powershell
New-Item -ItemType Directory -Force -Path (Split-Path $PROFILE)
```

Then append the completion script:

```powershell
ya completion powershell >> $PROFILE
```

Reload your profile:

```powershell
. $PROFILE
```

---

## Bash (Linux)

```bash
echo 'source <(ya completion bash)' >> ~/.bashrc
source ~/.bashrc
```

---

## Zsh (macOS / Linux)

```zsh
echo 'source <(ya completion zsh)' >> ~/.zshrc
source ~/.zshrc
```

---

## Fish

```fish
ya completion fish > ~/.config/fish/completions/ya.fish
```

---

After setup, `ya <TAB>` will list your shortcuts. No restart needed once the profile is reloaded.
