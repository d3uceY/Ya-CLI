<div align="center">
    <img src="assets/ya.png" alt="ya logo" height="200px"/>
    <h1>Ya - CLI</h1>
</div>

A lightweight command-line shortcut manager. Save, run, and organize your frequently used commands with simple, memorable shortcuts — or explore them visually inside the built-in interactive TUI.

<div align="center">
    <img src="https://github.com/user-attachments/assets/8ec0bed7-a6b0-48ec-a642-e9c0a76154f4" />
</div>

**Name Origin:** "Ya" comes from the Spanish word meaning "right now" — reflecting the instant execution of your commands.

## ⬇️ Download

![Ya CLI Banner](https://img.shields.io/badge/Made%20with-Go-00ADD8?style=for-the-badge&logo=go)
![License](https://img.shields.io/badge/license-MIT-blue?style=for-the-badge)

| Platform | Link |
|----------|------|
| 🪟 Windows | [![Download](https://img.shields.io/badge/Download-Windows-0078D4?style=for-the-badge&logo=windows&logoColor=white)](https://github.com/d3uceY/Ya-CLI/releases/latest) |
| 🍎 macOS | [![Download](https://img.shields.io/badge/Download-macOS-000000?style=for-the-badge&logo=apple&logoColor=white)](https://github.com/d3uceY/Ya-CLI/releases/latest) |
| 🐧 Linux | [![Download](https://img.shields.io/badge/Download-Linux-FCC624?style=for-the-badge&logo=linux&logoColor=black)](https://github.com/d3uceY/Ya-CLI/releases/latest) |

---

## Features

### CLI
- Create, run, rename, and delete shortcuts
- Template values — add `{placeholders}` and Ya prompts you to fill them in at runtime
- Import / export shortcuts to JSON
- Search shortcuts by name or command
- Pass extra arguments to shortcuts at runtime
- Overwrite protection and safe-delete confirmation prompts
- Shell tab-completion (Bash, Zsh, Fish, PowerShell)

### Interactive TUI (`ya` with no arguments)
- Full-screen terminal UI — browse, run, and manage shortcuts without typing subcommands
- **Pin shortcuts** to keep your most-used ones at the top of the list
- **Descriptions** — add an optional description to any shortcut, shown inline when selected
- **Run history** — every shortcut run is recorded; browse it from the TUI
- **Saved directories** — save frequently-used paths and reference them easily
- **Template fill-in page** — interactive form for `{placeholder}` values inside the TUI
- **Search as you type** — filter the list instantly; arrow keys still navigate while typing
- **Import / export** from inside the TUI
- Runs commands natively (TUI exits cleanly before executing)
- Shared data directory with [Ya-GUI](https://github.com/d3uceY/Ya-GUI) — changes in either app are immediately visible in the other

---

## Installation

### 🍺 Homebrew (macOS & Linux)

```bash
brew tap d3uceY/homebrew-ya
brew install ya
```

### Build from Source

**Prerequisites:** Go 1.25.5+

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

### Adding to PATH

#### Windows (PowerShell, as Administrator)
```powershell
New-Item -Path "C:\Program Files\Ya" -ItemType Directory -Force
Move-Item .\ya.exe "C:\Program Files\Ya\ya.exe"
[Environment]::SetEnvironmentVariable("Path", $env:Path + ";C:\Program Files\Ya", "User")
```

#### Linux / macOS
```bash
sudo mv ya /usr/local/bin/
```

---

## CLI Usage

### Run a shortcut
```bash
ya <shortcut>
```

### Add a shortcut
```bash
ya add <name> '<command>'
```

If a shortcut with that name already exists, Ya warns you before overwriting.

### Remove a shortcut
```bash
ya remove <name>
```
Prompts for confirmation before deleting.

### Rename a shortcut
```bash
ya rename <old-name> <new-name>
```

### List all shortcuts
```bash
ya list
```

### Show a shortcut's command
```bash
ya show <name>
```

### Search shortcuts
```bash
ya search <term>
```
Matches against both the shortcut name and its command.

### Import from a JSON file
```bash
ya import <file-path>
```
Merges with existing shortcuts — duplicate keys are overwritten.

### Export to a JSON file
```bash
ya export <directory>
ya export <directory> --name <filename>   # default: shortcuts.json
```

### Template values
Add `{placeholder}` tokens to a command. Ya prompts you to fill them in at runtime:

```bash
ya add commit 'git commit -m {message}'
ya add deploy 'kubectl set image deployment/{app} {app}={image}:{tag}'
```

Running `ya commit`:
```
→ This command has template values. Fill them in below:

  [1/1] message: fix login bug
```

The same `{placeholder}` used multiple times is only prompted once.

### Pass extra arguments
Extra arguments are appended to the shortcut command at runtime:

```bash
ya gcm -m "Initial commit"
# runs: git commit -m "Initial commit"
```

---

## TUI — Interactive Mode

Launch the TUI by running `ya` with no arguments:

```bash
ya
```

### Navigation

| Key | Action |
|-----|--------|
| `↑` / `k` | Move up |
| `↓` / `j` | Move down |
| `Enter` | Run selected shortcut |
| `q` / `Ctrl+C` | Quit |
| `?` | Toggle help page |
| `f` | Toggle fullscreen |
| `Esc` | Go back / cancel |

### Search

| Key | Action |
|-----|--------|
| `/` | Open search |
| Type anything | Filter list in real time |
| `↑` / `↓` | Navigate results while typing |
| `Enter` / `Esc` | Close search |

> Letter keys like `j` and `k` type normally into the search box and do **not** trigger navigation.

### Shortcut Management

| Key | Action |
|-----|--------|
| `a` | Add a new shortcut (name, command, optional description) |
| `e` | Edit selected shortcut (command + description) |
| `r` | Rename selected shortcut |
| `d` / `x` | Delete selected shortcut (with confirmation) |
| `p` | Pin / unpin selected shortcut |
| `i` | Import shortcuts from a JSON file |
| `o` | Export shortcuts to a JSON file |

### Pinned shortcuts

Press `p` on any shortcut to pin it. Pinned shortcuts float to the top of the list and are marked with a `◆` indicator. Press `p` again to unpin.

### Descriptions

When adding (`a`) or editing (`e`) a shortcut, fill in an optional description. When a shortcut with a description is selected, it appears as a sub-line beneath the row:

```
  ❯  gt    git tag -d v{tag}; git tag -a v{tag} -m "{metadata}"
             ──────────────────────────────────────────────────
             template for pushing git tags
```

### Run History

| Key | Action |
|-----|--------|
| `h` | Open history page |
| `c` | Clear all history (with confirmation) |

The history page shows every shortcut you've run, with the shortcut name, command, and timestamp.

### Saved Directories

| Key | Action |
|-----|--------|
| `D` | Open saved directories page |
| `a` (on dirs page) | Add a new saved directory |
| `d` (on dirs page) | Remove a saved directory |

### Template prompts in the TUI

If a shortcut contains `{placeholder}` tokens, the TUI opens a fill-in form before running. After submitting, the TUI exits and the resolved command runs natively in your terminal.

---

## Data Storage

All data is stored in your user config directory and **shared between Ya CLI and Ya GUI**:

| File | Contents |
|------|----------|
| `shortcuts.json` | Shortcut name → command map (CLI-compatible) |
| `shortcuts-meta.json` | Descriptions, pin status, run counts, last-run timestamps |
| `history.json` | Run history (up to 500 entries) |
| `config.json` | App config — saved directories |

**Default paths:**
- Windows: `%APPDATA%\ya\data\`
- macOS: `~/Library/Application Support/ya/data/`
- Linux: `~/.config/ya/data/`

---

## Shell Tab-Completion

Tab-complete shortcut names after `ya`. Works for `show`, `remove`, `rename`, and `search` too.

#### 🪟 PowerShell
```powershell
# Create profile directory if needed
New-Item -ItemType Directory -Force -Path (Split-Path $PROFILE)

ya completion powershell >> $PROFILE
. $PROFILE
```

#### 🐧 Bash
```bash
echo 'source <(ya completion bash)' >> ~/.bashrc
source ~/.bashrc
```

#### 🍎 Zsh
```zsh
echo 'source <(ya completion zsh)' >> ~/.zshrc
source ~/.zshrc
```

#### 🐟 Fish
```fish
ya completion fish > ~/.config/fish/completions/ya.fish
```

---

## GUI Companion

Prefer a graphical interface? [Ya-GUI](https://github.com/d3uceY/Ya-GUI) is a desktop app that reads and writes the same data files. Anything created in the CLI or TUI is instantly available in the GUI, and vice versa.

<img alt="Ya GUI" src="https://github.com/user-attachments/assets/258c0012-6944-43e7-95ee-a5f147ceca2b" width="100%"/>

---

## Troubleshooting

**"Unknown shortcut"** — the shortcut doesn't exist. Use `ya list` to see what's saved, or `ya add` to create one.

**Commands not working** — test the raw command in your terminal first before adding it as a shortcut.

**TUI looks garbled** — ensure your terminal supports UTF-8 and 256-colour output (Windows Terminal, iTerm2, and most modern terminals do).

---

## License

MIT — see [LICENSE](LICENSE) for details.

## Contributing

Contributions are welcome! Feel free to open an issue or pull request.


[![Download for macOS](https://img.shields.io/badge/Download-macOS-000000?style=for-the-badge\&logo=apple\&logoColor=white)](https://github.com/d3uceY/Ya-CLI/releases/latest)

#### 🐧 Linux

[![Download for Linux](https://img.shields.io/badge/Download-Linux-FCC624?style=for-the-badge\&logo=linux\&logoColor=black)](https://github.com/d3uceY/Ya-CLI/releases/latest)

> Windows • macOS • Linux | Always the latest version



## Features

- Create custom shortcuts for long or complex commands
- Persistent storage of shortcuts in your user config directory
- Fast command execution via PowerShell (Windows) or Bash (Linux/macOS)
- Easy-to-use CLI interface built with Cobra
- **Template values** — add `{placeholders}` to a command and Ya will prompt you to fill them in at runtime
- Export your shortcuts to a JSON file for backup or sharing
- Import shortcuts from a JSON file — merges with existing shortcuts
- Search shortcuts by name or command
- Pass extra arguments to shortcuts at runtime (e.g. `-m "message"`)
- **Rename shortcuts** — rename a shortcut key without changing its command
- **Overwrite protection** — warns and asks for confirmation before overwriting an existing shortcut
- **Safe delete** — confirmation prompt before removing a shortcut
- **Shell tab-completion** — tab-complete shortcut names in Bash, Zsh, Fish, and PowerShell
- **GUI Available:** Prefer a graphical interface? Check out [Ya-GUI](https://github.com/d3uceY/Ya-GUI) - a modern desktop application for managing your shortcuts visually
<img  alt="image" src="https://github.com/user-attachments/assets/258c0012-6944-43e7-95ee-a5f147ceca2b" width="100%"/>

## Platform Support

Supports **Windows**, **Linux**, and **macOS**.

- Windows: Uses PowerShell for command execution
- Linux/macOS: Uses Bash for command execution

## Installation

## 🍺 Install with Homebrew (macOS & Linux)

If you have **Homebrew** installed, you can install **Ya** with:

```bash
brew tap d3uceY/homebrew-ya
brew install ya
```

### Prerequisites

- Go 1.25.5 or higher
- Windows (PowerShell), Linux, or macOS (Bash)

### Build from Source

1. Clone or download this repository
2. Navigate to the project directory
3. Build the executable:

**Windows:**
```powershell
go build -o ya.exe .
```

**Linux/macOS:**
```bash
go build -o ya .
chmod +x ya
```

4. (Optional) Add the executable to your PATH for system-wide access

### Adding to PATH (Recommended)

To use `ya` from any directory in your terminal, you need to add it to your system's PATH environment variable.

#### Windows

1. **Move the executable** to a permanent location (e.g., `C:\Program Files\Ya\ya.exe`)
2. **Add to PATH:**
   - Press `Win + X` and select "System"
   - Click "Advanced system settings" on the right
   - Click "Environment Variables"
   - Under "User variables" or "System variables", find and select "Path"
   - Click "Edit"
   - Click "New" and add the directory path (e.g., `C:\Program Files\Ya`)
   - Click "OK" on all windows
3. **Restart your terminal** for changes to take effect
4. Verify by running: `ya help`

**Quick method (PowerShell as Administrator):**
```powershell
# Create a directory for Ya
New-Item -Path "C:\Program Files\Ya" -ItemType Directory -Force
# Move ya.exe to this directory
Move-Item .\ya.exe "C:\Program Files\Ya\ya.exe"
# Add to PATH
[Environment]::SetEnvironmentVariable("Path", $env:Path + ";C:\Program Files\Ya", "User")
```

#### Linux/macOS

1. **Move the executable** to a directory in your PATH:
   ```bash
   # Option 1: System-wide (requires sudo)
   sudo mv ya /usr/local/bin/
   
   # Option 2: User-only
   mkdir -p ~/.local/bin
   mv ya ~/.local/bin/
   ```

2. **Ensure the directory is in your PATH** (if using `~/.local/bin`):
   ```bash
   # For Bash (~/.bashrc)
   echo 'export PATH="$HOME/.local/bin:$PATH"' >> ~/.bashrc
   source ~/.bashrc
   
   # For Zsh (~/.zshrc)
   echo 'export PATH="$HOME/.local/bin:$PATH"' >> ~/.zshrc
   source ~/.zshrc
   ```

3. **Verify installation:**
   ```bash
   ya help
   ```

## Usage

## New Features (v0.3.0+)

- **Search Shortcuts:** Quickly find shortcuts by name or command.
- **Show Shortcut:** Display the command mapped to a specific shortcut.
- **Import Shortcuts:** Import a set of shortcuts from a JSON file — merges with existing ones.
- **Export Shortcuts:** Export your shortcuts to a JSON file for backup or sharing.
- **Remove Shortcut:** Delete a shortcut by name (prompts for confirmation).
- **Rename Shortcut:** Rename a shortcut key without changing its command.
- **Pass Extra Arguments:** You can now pass additional arguments to your shortcuts at runtime.
- **Template Values:** Use `{placeholder}` syntax in commands — Ya will prompt you to fill each one in before running.
- **Shell Tab-Completion:** Tab-complete shortcut names in your shell.
- **Overwrite Protection:** Ya warns you before overwriting an existing shortcut.
<img width="550" height="198" alt="image" src="https://github.com/user-attachments/assets/666e368d-b43c-4723-9027-8a26ff9b371e" />

See below for usage examples!


### Managing Shortcuts with GUI

For a more visual and user-friendly experience, you can use [Ya-GUI](https://github.com/d3uceY/Ya-GUI) - a desktop application that provides:
- Visual shortcut management
- Real-time search and filtering
- Inline editing capabilities
- Modern, intuitive interface

All shortcuts created in Ya-GUI are fully compatible with the Ya CLI and vice versa.

### View All Shortcuts

Display all available shortcuts:

```powershell
ya help
```

### List All Shortcuts

```powershell
ya list
```

Shows all shortcuts, their mapped commands, and a total count at the bottom.

### Add a New Shortcut

Create a new shortcut with a name and PowerShell command:

```powershell
ya add <shortcut-name> '<powershell-command>'
```

**Example:**

```powershell
ya add gs 'git status'
ya add dev 'cd C:\Projects\MyApp; npm run dev'
ya add ports 'netstat -ano | findstr LISTENING'
```

If a shortcut with that name already exists, Ya will warn you and ask before overwriting:

```
Shortcut gs already exists: git status
Overwrite? [y/N]:
```

### Remove a Shortcut

```powershell
ya remove <shortcut-name>
```

Deletes the specified shortcut. Ya will show the shortcut's current command and ask for confirmation before deleting.

```
This will remove the shortcut gs: git status
Are you sure? [y/N]:
```

### Rename a Shortcut

```powershell
ya rename <shortcut-name> <new-name>
```

Renames a shortcut key without changing the command it maps to.

**Example:**

```powershell
ya rename gs gitstatus   # renames the key, keeps the same command
```

Ya will show a confirmation before renaming:

```
This will rename the shortcut gs: git status to gitstatus
Are you sure? [y/N]:
```

### Execute a Shortcut

Run a saved shortcut:

```powershell
ya <shortcut-name>
```

**Example:**

```powershell
ya gs       # Runs: git status
ya dev      # Runs: cd C:\Projects\MyApp; npm run dev
ya ports    # Runs: netstat -ano | findstr LISTENING
```

#### Pass Extra Arguments to Shortcuts

You can append extra arguments when running a shortcut. These will be added to the end of the command.

```powershell
ya gcm -m "Initial commit"
# Runs: git commit -m "Initial commit"
```

### Show a Shortcut's Command

```powershell
ya show <shortcut-name>
```

Displays the command mapped to the shortcut.

### Search for Shortcuts

```powershell
ya search <search-term>
```

Finds shortcuts whose name or command contains the search term.

### Import Shortcuts from a File

```powershell
ya import <file-path>
```

Imports shortcuts from a JSON file (should be a map of shortcut names to commands). Merges with your existing shortcuts — duplicate keys are overwritten.

### Export Shortcuts to a File

```powershell
ya export <dir>
ya export <dir> --name <filename>
ya export <dir> -n <filename>
```

Exports all your shortcuts to a JSON file in the given directory. If `--name` is not provided, the file will be named `shortcuts.json`.

**Examples:**

```powershell
ya export ./                        # exports to ./shortcuts.json
ya export ./backups --name jan.json # exports to ./backups/jan.json
```

### Template Values in Commands

You can add `{placeholder}` tokens anywhere in a command. When you run the shortcut, Ya will prompt you to fill in each value before executing.

**Add a templated shortcut:**

```powershell
ya add commit 'git commit -m {message}'
ya add deploy 'kubectl set image deployment/{app} {app}={image}:{tag}'
```

**Run it:**

```powershell
ya commit
```

```
→ This command has template values. Fill them in below:

  [1/1] message: fix login bug
```

Then runs `git commit -m fix login bug`.

- The same `{placeholder}` used multiple times in one command is only prompted once
- Commands without any `{}` tokens run immediately as normal

### Shell Tab-Completion

Ya supports tab-completion of shortcut names — press `Tab` after `ya` and your shell will suggest your saved shortcuts.

**Setup is a one-time step.** Follow the instructions for your shell below.

---

#### 🪟 PowerShell (Windows)

> **First-time only:** If you have never set up a PowerShell profile before, the profile file and its folder may not exist yet. Run this command to create them:
> ```powershell
> New-Item -ItemType Directory -Force -Path (Split-Path $PROFILE)
> ```

Then append the completion script to your profile:
```powershell
ya completion powershell >> $PROFILE
```

Reload your profile to activate it immediately (or just restart your terminal):
```powershell
. $PROFILE
```

---

#### 🐧 Bash (Linux)

```bash
echo 'source <(ya completion bash)' >> ~/.bashrc
source ~/.bashrc
```

---

#### 🍎 Zsh (macOS / Linux)

```zsh
echo 'source <(ya completion zsh)' >> ~/.zshrc
source ~/.zshrc
```

---

#### 🐟 Fish

```fish
ya completion fish | source
```

To make it permanent:
```fish
ya completion fish > ~/.config/fish/completions/ya.fish
```

---

After setup, `ya <TAB>` will suggest your saved shortcut names. This works for `ya show`, `ya remove`, `ya rename`, and `ya search` as well.

## How It Works

- Shortcuts are stored in JSON format at: `%APPDATA%\ya\data\shortcuts.json`
- Commands are executed via PowerShell using the `-Command` flag
- The shortcuts file is automatically created on first use

## Examples

### Common Use Cases

```powershell
# Quick navigation shortcuts
ya add proj 'cd C:\Users\YourName\Projects'
ya add docs 'cd C:\Users\YourName\Documents'

# Git shortcuts
ya add gp 'git pull'
ya add gps 'git push'
ya add gcm 'git commit -m'

# Development shortcuts
ya add start-api 'cd C:\MyApp\API; dotnet run'
ya add build 'cd C:\MyApp; npm run build'

# System maintenance
ya add clean 'Remove-Item -Path .\temp\* -Recurse -Force'
```

## Troubleshooting

### "Unknown shortcut" Error

If you see this message, the shortcut doesn't exist. Use `ya help` to see all available shortcuts or `ya add` to create a new one.

### Commands Not Working

Ensure your command works in PowerShell before adding it as a shortcut. Test it directly in a PowerShell terminal first.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contributing

Contributions are welcome! Feel free to submit issues or pull requests.
