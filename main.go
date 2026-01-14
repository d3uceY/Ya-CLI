package main

import (
	"os"
	"os/exec"

	"github.com/fatih/color"

	"ya/utils"
)

func main() {
	shortcuts, err := utils.LoadShortcuts()
	if err != nil {
		color.Red("Error loading shortcuts:", err)
		os.Exit(1)
	}

	if len(os.Args) < 2 {
		color.Red("Usage: ya <shortcut> \n for shortcuts use: ya help")
		os.Exit(1)
	}

	shortcut := os.Args[1]

	switch shortcut {
	case "help":
		color.Green("Available shortcuts:")
		for key := range shortcuts {
			color.Green(" - %s : %s", key, shortcuts[key])
		}
		color.Green("\nTo add a new shortcut use: ya add <shortcut> '<command>'")
		color.Green("To remove a shortcut use: ya remove <shortcut>")
		return
	case "add":
		if len(os.Args) < 4 {
			color.Red("Usage: ya add <shortcut> '<command>'")
			os.Exit(1)
		}
		shortcutName := os.Args[2]
		command := os.Args[3]
		if utils.IsInvalidString(shortcutName) || utils.IsInvalidString(command) {
			color.Red("Usage: ya add <shortcut> '<command>'")
			os.Exit(1)
		}
		utils.AddShortcut(shortcutName, command)
		return
	case "remove":
		if len(os.Args) < 3 {
			color.Red("Usage: ya remove <shortcut>")
			os.Exit(1)
		}
		utils.RemoveShortcut(os.Args[2])
		return
	}

	command, exists := shortcuts[shortcut]

	if !exists {
		color.Red("Unknown shortcut: %s\n to add a new shortcut use: ya add <shortcut> '<command>'", shortcut)
		os.Exit(1)
	}

	cmd := exec.Command("powershell", "-NoProfile", "-NonInteractive", "-Command", command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	cmdError := cmd.Run()
	if cmdError != nil {
		color.Red("Command failed: %v", cmdError)
		os.Exit(1)
	}
}
