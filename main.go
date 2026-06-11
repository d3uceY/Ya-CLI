package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/d3uceY/Ya-CLI/tui"
	"github.com/d3uceY/Ya-CLI/utils"
	"github.com/spf13/cobra"
)

func main() {

	// colored output — uses TUI palette via lipgloss
	accent := utils.CAccent.Render
	success := utils.CSuccess.Render

	// completion function — reads shortcut names from storage and returns them for shell tab-completion
	shortcutCompletions := func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		if len(args) > 0 {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}
		shortcuts, err := utils.LoadShortcuts()
		if err != nil {
			return nil, cobra.ShellCompDirectiveError
		}
		names := make([]string, 0, len(shortcuts))
		for name := range shortcuts {
			names = append(names, name)
		}
		return names, cobra.ShellCompDirectiveNoFileComp
	}

	// root command — handles running shortcuts directly via `ya <shortcut> [args...]`
	// or launches the TUI when called with no arguments.
	var rootCmd = &cobra.Command{
		Use:   "ya",
		Short: "Ya - a CLI shortcut runner",
		// DisableFlagParsing allows extra args like -m to be passed through to the shortcut command
		DisableFlagParsing: true,
		Args:               cobra.ArbitraryArgs,
		Run: func(cmd *cobra.Command, args []string) {
			// No arguments → launch the interactive TUI
			if len(args) == 0 {
				tui.Start()
				return
			}

			shortcut := args[0]

			shortcuts, err := utils.LoadShortcuts()
			if err != nil {
				utils.PrintError("Error loading shortcuts: %v", err)
				os.Exit(1)
			}

			command, exists := shortcuts[shortcut]

			if !exists {
				fmt.Printf("Unknown shortcut: %s\n to add a new shortcut use: ya add <shortcut> '<command>'", accent(shortcut))
				os.Exit(1)
			}

			// if the command has {placeholder} tokens, prompt the user to fill them in
			var templateErr error
			command, templateErr = utils.ResolveTemplates(command)
			if templateErr != nil {
				utils.PrintError("Template error: %v", templateErr)
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
				utils.PrintError("Command failed: %v", cmdError)
				os.Exit(1)
			}
		},
		ValidArgsFunction: shortcutCompletions,
	}

	// Version command
	var versionCmd = &cobra.Command{
		Use:     "version",
		Short:   "Show the current Ya version",
		Aliases: []string{"-v", "--version"},
		Args:    cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			version := utils.GetAppVersion()
			utils.PrintSuccess("Ya version: %s", version)
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
				utils.PrintError("Error loading shortcuts: %v", err)
				os.Exit(1)
			}
			utils.PrintSuccess("Available shortcuts:")
			for key, command := range shortcuts {
				fmt.Printf("- %s : %s\n", accent(key), success(command))
			}
			fmt.Printf("\n%s\n", utils.CDim.Render(fmt.Sprintf("%d shortcut(s)", len(shortcuts))))
		},
	}

	// Help command
	var helpCmd = &cobra.Command{
		Use:     "help",
		Short:   "Show usage information",
		Aliases: []string{"--help", "-h"},
		Args:    cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("\n" + utils.CAccent.Bold(true).Render("ya") + utils.CDim.Render("  usage"))
			fmt.Println()

			fmt.Printf("%s %s\n", utils.CDim.Render("To add a new shortcut use:"), accent("ya add <shortcut> <command>"))
			fmt.Printf("%s %s\n", utils.CDim.Render("To remove a shortcut use:"), accent("ya remove <shortcut>"))
			fmt.Printf("%s %s\n", utils.CDim.Render("To list all shortcuts:"), accent("ya list"))
			fmt.Printf("%s %s\n", utils.CDim.Render("To show version:"), accent("ya version"))
			fmt.Printf("%s %s\n", utils.CDim.Render("To import shortcuts use:"), accent("ya import <file-path>"))
			fmt.Printf("%s %s\n", utils.CDim.Render("To export shortcuts use:"), accent("ya export <dir> [--name <filename>]"))
			fmt.Printf("%s %s\n", utils.CDim.Render("To search shortcuts use:"), accent("ya search <shortcut>"))
			fmt.Printf("%s %s\n", utils.CDim.Render("To show a shortcut use:"), accent("ya show <shortcut>"))
			fmt.Printf("%s %s\n", utils.CDim.Render("To run a shortcut use:"), accent("ya <shortcut> [extra args...]"))
			fmt.Printf("%s %s\n", utils.CDim.Render("To use template values in a command:"), accent("ya add commit 'git commit -m {message}'"))

			fmt.Printf("%s %s\n", utils.CDim.Render("To enable shell autocomplete:"), accent("ya completion <shell>"))
			fmt.Printf("%s\n", utils.CDim.Render("  (see README: https://github.com/d3uceY/Ya-CLI#shell-tab-completion)"))

			fmt.Println()
		},
	}

	// Search command
	var searchCmd = &cobra.Command{
		Use:               "search <shortcut>",
		Short:             "Search for shortcuts by name or command",
		Aliases:           []string{"--search"},
		Args:              cobra.ExactArgs(1),
		ValidArgsFunction: shortcutCompletions,
		Run: func(cmd *cobra.Command, args []string) {
			shortcuts, err := utils.SearchShortcut(args[0])
			if err != nil {
				utils.PrintError("%s", err.Error())
			}
			if !(len(shortcuts) >= 1) {
				utils.PrintError("Shortcuts with `%s` not found", args[0])
			}
			utils.PrintSuccess("Search results:")
			for key, command := range shortcuts {
				fmt.Printf(" - %s : %s\n", accent(key), success(command))
			}
			fmt.Printf("\n%s\n", utils.CDim.Render(fmt.Sprintf("%d shortcut(s)", len(shortcuts))))
		},
	}

	// Show command
	var showCmd = &cobra.Command{
		Use:               "show <shortcut>",
		Short:             "Show the command mapped to a shortcut",
		Args:              cobra.ExactArgs(1),
		ValidArgsFunction: shortcutCompletions,
		Run: func(cmd *cobra.Command, args []string) {
			command, err := utils.GetShortcut(args[0])
			if err != nil {
				utils.PrintError("%s", err.Error())
				os.Exit(1)
			}
			fmt.Printf("Shortcut `%s` maps to command: %s\n", accent(args[0]), success(command))
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
				utils.PrintError("Usage: ya add <shortcut> '<command>'")
				os.Exit(1)
			}

			existing, err := utils.GetShortcut(shortcutName)
			if err == nil {
				fmt.Printf("Shortcut %s already exists: %s\n", accent(shortcutName), success(existing))
				fmt.Printf("Overwrite? [y/N]: ")
				var input string
				fmt.Scanln(&input)
				if input != "y" && input != "Y" {
					utils.PrintDim("Aborted.")
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
				utils.PrintError("%s", err.Error())
				os.Exit(1)
			}
			utils.PrintSuccess("Shortcut Imported `%s`", args[0])
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
				utils.PrintError("%s", err.Error())
				os.Exit(1)
			}
			utils.PrintSuccess("Shortcuts exported to `%s`", filePath)
		},
	}
	// --name flag to set the output filename, defaults to shortcuts.json
	exportCmd.Flags().StringVarP(&exportName, "name", "n", "shortcuts.json", "Name of the exported file")

	// Remove command
	var removeCmd = &cobra.Command{
		Use:               "remove <shortcut>",
		Short:             "Remove an existing shortcut",
		Args:              cobra.ExactArgs(1),
		ValidArgsFunction: shortcutCompletions,
		Run: func(cmd *cobra.Command, args []string) {
			shortcutName := args[0]
			existing, err := utils.GetShortcut(shortcutName)
			if err != nil {
				utils.PrintError("%s", err.Error())
				os.Exit(1)
			}

			fmt.Printf("This will remove the shortcut %s: %s\n", accent(shortcutName), success(existing))
			fmt.Printf("Are you sure? [y/N]: ")
			var input string
			fmt.Scanln(&input)
			if input != "y" && input != "Y" {
				utils.PrintDim("Aborted.")
				return
			}
			utils.RemoveShortcut(args[0])
		},
	}

	var renameCmd = &cobra.Command{
		Use:               "rename <shortcut> <new-name>",
		Short:             "Rename an existing shortcut",
		Args:              cobra.ExactArgs(2),
		ValidArgsFunction: shortcutCompletions,
		Run: func(cmd *cobra.Command, args []string) {
			shortcutName := args[0]
			newName := args[1]
			existing, err := utils.GetShortcut(shortcutName)
			if err != nil {
				utils.PrintError("%s", err.Error())
				os.Exit(1)
			}

			fmt.Printf("This will rename the shortcut %s: %s to %s\n", accent(shortcutName), success(existing), accent(newName))
			fmt.Printf("Are you sure? [y/N]: ")
			var input string
			fmt.Scanln(&input)
			if input != "y" && input != "Y" {
				utils.PrintDim("Aborted.")
				return
			}

			err = utils.RenameShortcut(shortcutName, newName)
			if err != nil {
				utils.PrintError("%s", err.Error())
				os.Exit(1)
			}
			utils.PrintSuccess("Renamed `%s` to `%s`", shortcutName, newName)
		},
	}

	// register all subcommands onto the root
	rootCmd.AddCommand(versionCmd, listCmd, helpCmd, searchCmd, showCmd, addCmd, importCmd, exportCmd, removeCmd, renameCmd)

	// disable cobra's default help so our custom help command takes over cleanly
	rootCmd.SetHelpCommand(helpCmd)

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
