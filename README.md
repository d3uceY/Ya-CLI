<div align="center">
    <img src="assets/ya.png" alt="ya logo" />
    <h1>Ya</h1>
</div>


A lightweight command-line shortcut manager for Windows. Execute your frequently used PowerShell commands with simple, memorable shortcuts.

**Name Origin:** "Ya" comes from the Spanish word meaning "right now" - reflecting the instant execution of your commands.

## Features

- 🚀 Create custom shortcuts for long or complex commands
- 💾 Persistent storage of shortcuts in your user config directory
- ⚡ Fast command execution via PowerShell
- 📝 Easy-to-use CLI interface

## Platform Support

Supports **Windows**, **Linux**, and **macOS**.

- Windows: Uses PowerShell for command execution
- Linux/macOS: Uses Bash for command execution

## Installation

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

## Usage

### View All Shortcuts

Display all available shortcuts:

```powershell
ya help
```

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
