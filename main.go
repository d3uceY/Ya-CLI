package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"

	"pronto/utils"
)

func loadShortcuts() (map[string]string, error) {
	data, err := os.ReadFile("shortcuts.json")
	if err != nil {
		return nil, err
	}

	var shortcuts map[string]string
	err = json.Unmarshal(data, &shortcuts)
	if err != nil {
		return nil, err
	}

	return shortcuts, nil
}

func addShortcut(name, command string) error {
	shortcuts, err := loadShortcuts()

	if err != nil {
		return err
	}

	shortcuts[name] = command

	data, err := json.MarshalIndent(shortcuts, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile("shortcuts.json", data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	shortcuts, err := loadShortcuts()
	if err != nil {
		fmt.Println("Error loading shortcuts:", err)
		os.Exit(1)
	}

	if len(os.Args) < 2 {
		fmt.Println("Usage: pronto <shortcut> \n for shortcuts use: pronto help")
		os.Exit(1)
	}

	shortcut := os.Args[1]

	switch shortcut {
	case "help":
		fmt.Println("Available shortcuts:")
		for key := range shortcuts {
			fmt.Println(" -", key, ":", shortcuts[key])
		}
		return
	case "add":
		if len(os.Args) < 4 {
			fmt.Println("Usage: pronto add <shortcut> <command>")
			os.Exit(1)
		}
		shortcutName := os.Args[2]
		command := os.Args[3]
		if utils.IsInvalidString(shortcutName) || utils.IsInvalidString(command) {
			fmt.Println("Usage: pronto add <shortcut> <command>")
			os.Exit(1)
		}
		addShortcut(shortcutName, command)
		return
	}

	command, exists := shortcuts[shortcut]

	if !exists {
		fmt.Printf("Unknown shortcut: %s\n", shortcut)
		os.Exit(1)
	}

	cmd := exec.Command("powershell", "-Command", command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	cmdError := cmd.Run()
	if cmdError != nil {
		fmt.Println("Command failed:", cmdError)
		os.Exit(1)
	}
}
