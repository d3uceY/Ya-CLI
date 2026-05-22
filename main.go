package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/d3uceY/Ya-CLI/utils"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func main() {

	// colored output
	green := color.New(color.FgGreen).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()

	// root command — handles running shortcuts directly via `ya <shortcut> [args...]`
	var rootCmd = &cobra.Command{
		Use:   "ya",
		Short: "Ya - a CLI shortcut runner",
		// DisableFlagParsing allows extra args like -m to be passed through to the shortcut command
		DisableFlagParsing: true,
		Args:               cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			shortcut := args[0]

			shortcuts, err := utils.LoadShortcuts()
			if err != nil {
				color.Red("Error loading shortcuts: %v", err)
				os.Exit(1)
			}

			command, exists := shortcuts[shortcut]

			if !exists {
				fmt.Printf("Unknown shortcut: %s\n to add a new shortcut use: ya add <shortcut> '<command>'", yellow(shortcut))
				os.Exit(1)
			}

			// if the command has {placeholder} tokens, prompt the user to fill them in
			var templateErr error
			command, templateErr = utils.ResolveTemplates(command)
			if templateErr != nil {
				color.Red("Template error: %v", templateErr)
				os.Exit(1)
			}

			// i added this because i was wondering how i would have been using this
			// this allows arguments passing the shortcut commands to also pass messages like git commit -m 'message'
			if len(args) > 1 {
				for index, value := range args {
					if index >= 1 {
						if value != "-m" {
							if args[index-1] == "-m" {
								command += " " + fmt.Sprintf("'%s'", value)
							} else {
								command += " " + value
							}
						}
					}
				}
			}

			// if we are here, detect if it is windows, you get?
			var execCmd *exec.Cmd
			if runtime.GOOS == "windows" {
				execCmd = exec.Command("powershell", "-NoProfile", "-NonInteractive", "-Command", command)
			} else {
				// for linux and macOS typeshit
				execCmd = exec.Command("bash", "-c", command)
			}

			execCmd.Stdout = os.Stdout
			execCmd.Stderr = os.Stderr
			execCmd.Stdin = os.Stdin

			cmdError := execCmd.Run()
			if cmdError != nil {
				color.Red("Command failed: %v", cmdError)
				os.Exit(1)
			}
		},
	}

	// Version command
	var versionCmd = &cobra.Command{
		Use:     "version",
		Short:   "Show the current Ya version",
		Aliases: []string{"-v", "--version"},
		Args:    cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			version := utils.GetAppVersion()
			color.Green("Ya version: %s", version)
		},
	}

	// list command
	var listCmd = &cobra.Command{
		Use:     "list",
		Short:   "List all available shortcuts",
		Aliases: []string{"-l", "--list"},
		Args:    cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			shortcuts, err := utils.LoadShortcuts()
			if err != nil {
				color.Red("Error loading shortcuts: %v", err)
				os.Exit(1)
			}
			color.Green("Available shortcuts:")
			for key, command := range shortcuts {
				fmt.Printf("- %s : %s\n", yellow(key), green(command))
			}
		},
	}

	// Help command
	var helpCmd = &cobra.Command{
		Use:     "help",
		Short:   "Show usage information",
		Aliases: []string{"--help", "-h"},
		Args:    cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("\n--- Ya Usage ---")

			fmt.Printf("%s %s\n", yellow("To add a new shortcut use:"), green("ya add <shortcut> <command>"))
			fmt.Printf("%s %s\n", yellow("To remove a shortcut use:"), green("ya remove <shortcut>"))
			fmt.Printf("%s %s\n", yellow("To list all shortcuts:"), green("ya list"))
			fmt.Printf("%s %s\n", yellow("To show version:"), green("ya version"))
			fmt.Printf("%s %s\n", yellow("To import shortcuts use:"), green("ya import <file-path>"))
			fmt.Printf("%s %s\n", yellow("To export shortcuts use:"), green("ya export <dir> [--name <filename>]"))
			fmt.Printf("%s %s\n", yellow("To search shortcuts use:"), green("ya search <shortcut>"))
			fmt.Printf("%s %s\n", yellow("To show a shortcut use:"), green("ya show <shortcut>"))
			fmt.Printf("%s %s\n", yellow("To run a shortcut use:"), green("ya <shortcut> [extra args...]"))
			fmt.Printf("%s %s\n", yellow("To use template values in a command:"), green("ya add commit 'git commit -m {message}'"))

			fmt.Println()
		},
	}

	// Search command
	var searchCmd = &cobra.Command{
		Use:     "search <shortcut>",
		Short:   "Search for shortcuts by name or command",
		Aliases: []string{"--search"},
		Args:    cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			shortcuts, err := utils.SearchShortcut(args[0])
			if err != nil {
				color.Red(err.Error())
			}
			if !(len(shortcuts) >= 1) {
				color.Red("Shortcuts with `%s` not found", args[0])
			}
			color.Green("Search results:")
			for key, command := range shortcuts {
				color.Yellow(" - %s :", key)
				color.Green(" %s", command)
			}
		},
	}

	// Show command
	var showCmd = &cobra.Command{
		Use:   "show <shortcut>",
		Short: "Show the command mapped to a shortcut",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			command, err := utils.GetShortcut(args[0])
			if err != nil {
				color.Red(err.Error())
				os.Exit(1)
			}
			fmt.Printf("Shortcut `%s` maps to command: %s\n", yellow(args[0]), green(command))
		},
	}

	// Add command
	var addCmd = &cobra.Command{
		Use:   "add <shortcut> <command>",
		Short: "Add a new shortcut",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			shortcutName := args[0]
			command := args[1]
			if utils.IsInvalidString(shortcutName) || utils.IsInvalidString(command) {
				color.Red("Usage: ya add <shortcut> '<command>'")
				os.Exit(1)
			}

			existing, err := utils.GetShortcut(shortcutName)
			if err == nil {
				fmt.Printf("Shortcut %s already exists: %s\n", yellow(shortcutName), green(existing))
				fmt.Printf("Overwrite? [y/N]: ")
				var input string
				fmt.Scanln(&input)
				if input != "y" && input != "Y" {
					color.Yellow("Aborted.")
					return
				}
			}

			utils.AddShortcut(shortcutName, command)
		},
	}

	// import shortcuts
	var importCmd = &cobra.Command{
		Use:   "import <file-path>",
		Short: "Import shortcuts from a JSON file",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			err := utils.ImportShortcuts(args[0])
			if err != nil {
				color.Red(err.Error())
				os.Exit(1)
			}
			color.Green("Shortcut Imported `%s`", args[0])
		},
	}

	// export shortcuts
	// usage: ya export <dir> [--name <filename>]
	// if --name is not provided, defaults to shortcuts.json
	var exportName string
	var exportCmd = &cobra.Command{
		Use:   "export <dir>",
		Short: "Export shortcuts to a JSON file in the given directory",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			filePath := args[0] + "/" + exportName
			err := utils.ExportShortcuts(filePath)
			if err != nil {
				color.Red(err.Error())
				os.Exit(1)
			}
			color.Green("Shortcuts exported to `%s`", filePath)
		},
	}
	// --name flag to set the output filename, defaults to shortcuts.json
	exportCmd.Flags().StringVarP(&exportName, "name", "n", "shortcuts.json", "Name of the exported file")

	// Remove command
	var removeCmd = &cobra.Command{
		Use:   "remove <shortcut>",
		Short: "Remove an existing shortcut",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			shortcutName := args[0]
			existing, err := utils.GetShortcut(shortcutName)
			if err != nil {
				color.Red(err.Error())
				os.Exit(1)
			}

			fmt.Printf("This will remove the shortcut %s: %s\n", yellow(shortcutName), green(existing))
			fmt.Printf("Are you sure? [y/N]: ")
			var input string
			fmt.Scanln(&input)
			if input != "y" && input != "Y" {
				color.Yellow("Aborted.")
				return
			}
			utils.RemoveShortcut(args[0])
		},
	}

	var renameCmd = &cobra.Command{
		Use:   "rename <shortcut> <new-name>",
		Short: "Rename an existing shortcut",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			shortcutName := args[0]
			newName := args[1]
			existing, err := utils.GetShortcut(shortcutName)
			if err != nil {
				color.Red(err.Error())
				os.Exit(1)
			}

			fmt.Printf("This will rename the shortcut %s: %s to %s\n", yellow(shortcutName), green(existing), yellow(newName))
			fmt.Printf("Are you sure? [y/N]: ")
			var input string
			fmt.Scanln(&input)
			if input != "y" && input != "Y" {
				color.Yellow("Aborted.")
				return
			}

			err = utils.RenameShortcut(shortcutName, newName)
			if err != nil {
				color.Red(err.Error())
				os.Exit(1)
			}
			color.Green("Renamed `%s` to `%s`", shortcutName, newName)
		},
	}

	// register all subcommands onto the root
	rootCmd.AddCommand(versionCmd, listCmd, helpCmd, searchCmd, showCmd, addCmd, importCmd, exportCmd, removeCmd, renameCmd)

	// disable cobra's default help/completion so our custom help command takes over cleanly
	rootCmd.SetHelpCommand(helpCmd)
	rootCmd.CompletionOptions.DisableDefaultCmd = true

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
