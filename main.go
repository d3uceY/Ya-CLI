package main

import (
	"fmt"
	"github.com/fatih/color"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"ya/utils"
)

func main() {

	// colored output
	green := color.New(color.FgGreen).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()

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
	// Version command
	case "version", "-v":
		if len(os.Args) > 2 {
			color.Red("usage: ya version")
			os.Exit(1)
		}
		version := utils.GetAppVersion()
		color.Green("Ya version: %s", version)
		return

	// list command
	case "list", "-l", "--list":
		if len(os.Args) > 2 {
			color.Red("usage: ya list")
			os.Exit(1)
		}
		color.Green("Available shortcuts:")
		for key, cmd := range shortcuts {
			fmt.Printf("%s %s\n", yellow(" - %s :", key), green(" %s", cmd))
		}
		return

	// Help command
	case "help", "--help", "-h":
		if len(os.Args) > 2 {
			color.Red("usage: ya help")
			os.Exit(1)
		}

		fmt.Println("\n--- Ya Usage ---")

		fmt.Printf("%s %s\n", yellow("To add a new shortcut use:"), green("ya add <shortcut> <command>"))
		fmt.Printf("%s %s\n", yellow("To remove a shortcut use:"), green("ya remove <shortcut>"))
		fmt.Printf("%s %s\n", yellow("To list all shortcuts:"), green("ya list"))
		fmt.Printf("%s %s\n", yellow("To show version:"), green("ya version"))
		fmt.Printf("%s %s\n", yellow("To import shortcuts use:"), green("ya import <file-path>"))
		fmt.Printf("%s %s\n", yellow("To search shortcuts use:"), green("ya search <shortcut>"))
		fmt.Printf("%s %s\n", yellow("To show shortcuts use:"), green("ya show <shortcut>"))

		fmt.Println()
		return

	// Search command
	case "search", "--search":
		if len(os.Args) > 3 {
			color.Red("usage: ya search <shortcut>")
			os.Exit(1)
		}
		shortcuts, err := utils.SearchShortcut(os.Args[2])
		if err != nil {
			color.Red(err.Error())
		}
		if !(len(shortcuts) >= 1) {
			color.Red("Shortcuts with `%s` not found", os.Args[2])
		}
		color.Green("Search results:")
		for key, cmd := range shortcuts {
			color.Yellow(" - %s :", key)
			color.Green(" %s", cmd)
		}
		return

	// Show command
	case "show":
		if len(os.Args) > 3 {
			color.Red("usage: ya show <shortcut>")
			os.Exit(1)
		}
		command, err := utils.GetShortcut(os.Args[2])
		if err != nil {
			color.Red(err.Error())
			os.Exit(1)
		}
		fmt.Printf("Shortcut `%s` maps to command: %s\n", yellow(os.Args[2]), green(command))
		return

	// Add command
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

		// import shortcuts
	case "import":
		if len(os.Args) < 3 {
			color.Red("Usage: ya import '<file-path>'")
			os.Exit(1)
		}

		err := utils.ImportShortcuts(os.Args[2])
		if err != nil {
			color.Red(err.Error())
			os.Exit(1)
		}
		color.Green("Shortcut Imported `%s`", os.Args[2])

	// Remove command
	case "remove":
		if len(os.Args) < 3 {
			color.Red("Usage: ya remove <shortcut>")
			os.Exit(1)
		}
		utils.RemoveShortcut(os.Args[2])
		return
	}

	command, exists := shortcuts[shortcut]

	// i added this because i was wondering how i would have been using this
	// if i cannot pass more arguments to the shortcut e.g git commit -m (extra args)
	// so hear it is
	if len(os.Args) > 2 {
		args := os.Args[2:]
		command += " " + strings.Join(args, " ")
	}

	if !exists {
		color.Red("Unknown shortcut: %s\n to add a new shortcut use: ya add <shortcut> '<command>'", shortcut)
		os.Exit(1)
	}

	// if we are here, detect if it is windows, you get?
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("powershell", "-NoProfile", "-NonInteractive", "-Command", command)
	} else {
		// for linux and macOS typeshit
		cmd = exec.Command("bash", "-c", command)
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	cmdError := cmd.Run()
	if cmdError != nil {
		color.Red("Command failed: %v", cmdError)
		os.Exit(1)
	}
}
