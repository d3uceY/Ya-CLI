<div align="center">
    <img src="assets/ya.png" alt="ya logo" height="200px"/>
    <h1>Ya - CLI</h1>
</div>


A lightweight command-line shortcut manager. Execute your frequently used PowerShell commands with simple, memorable shortcuts.

<div align="center">
    <img src="https://github.com/user-attachments/assets/8ec0bed7-a6b0-48ec-a642-e9c0a76154f4" />
</div>



**Name Origin:** "Ya" comes from the Spanish word meaning "right now" - reflecting the instant execution of your commands.

## ⬇️ Download
![Ya CLI Banner](https://img.shields.io/badge/Made%20with-Go-00ADD8?style=for-the-badge&logo=go)
![License](https://img.shields.io/badge/license-MIT-blue?style=for-the-badge)

### Ya CLI — Multi-Platform Support

#### 🪟 Windows

[![Download for Windows](https://img.shields.io/badge/Download-Windows-0078D4?style=for-the-badge\&logo=windows\&logoColor=white)](https://github.com/d3uceY/Ya-CLI/releases/latest)

#### 🍎 macOS

[![Download for macOS](https://img.shields.io/badge/Download-macOS-000000?style=for-the-badge\&logo=apple\&logoColor=white)](https://github.com/d3uceY/Ya-CLI/releases/latest)

#### 🐧 Linux

[![Download for Linux](https://img.shields.io/badge/Download-Linux-FCC624?style=for-the-badge\&logo=linux\&logoColor=black)](https://github.com/d3uceY/Ya-CLI/releases/latest)

> Windows • macOS • Linux | Always the latest version



## Features

- Create custom shortcuts for long or complex commands
- Persistent storage of shortcuts in your user config directory
- Fast command execution via PowerShell
- Easy-to-use CLI interface
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
- **Import Shortcuts:** Import a set of shortcuts from a JSON file.
- **Remove Shortcut:** Delete a shortcut by name.
- **Pass Extra Arguments:** You can now pass additional arguments to your shortcuts at runtime.
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

Shows all shortcuts and their mapped commands.

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

### Remove a Shortcut

```powershell
ya remove <shortcut-name>
```

Deletes the specified shortcut.

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

Imports shortcuts from a JSON file (should be a map of shortcut names to commands).

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
