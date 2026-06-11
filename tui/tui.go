package tui

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/d3uceY/Ya-CLI/utils"
)

// Start launches the Ya TUI. When the user runs a shortcut the TUI quits first,
// then the command is executed directly in the normal terminal.
func Start() {
	m := newModel()
	p := tea.NewProgram(
		m,
		tea.WithAltScreen(),
		tea.WithMouseCellMotion(),
	)
	result, err := p.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "ya: tui error: %v\n", err)
		os.Exit(1)
	}

	final, ok := result.(Model)
	if !ok || final.pendingCmd == "" {
		return
	}

	// Record history and bump run count before handing off to the shell.
	cwd, _ := os.Getwd()
	_ = utils.IncrementRunCount(final.pendingCmdName)
	_ = utils.AddRunHistoryEntry(final.pendingCmdName, final.pendingCmd, cwd)

	runNative(final.pendingCmd)
}

// runNative execs the command in the current terminal (no TUI involved).
func runNative(command string) {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("powershell", "-NoProfile", "-NonInteractive", "-Command", command)
	} else {
		cmd = exec.Command("bash", "-c", command)
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	if err := cmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "command failed: %v\n", err)
		os.Exit(1)
	}
}

