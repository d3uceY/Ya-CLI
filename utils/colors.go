package utils

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

// Shared palette — mirrors the TUI exactly so CLI and TUI output look the same.
const (
	ColAccent   = "#5F9EA0" // cadet blue — names, keys, prompts
	ColText     = "#CCCCCC" // soft white — body text
	ColDim      = "#666666" // dim gray — secondary info
	ColSuccess  = "#7EC8A4" // muted sage green — success / commands
	ColError    = "#CC6666" // muted rose red — errors
	ColTemplate = "#C8A96E" // muted amber — {placeholder} tokens
)

var (
	CAccent   = lipgloss.NewStyle().Foreground(lipgloss.Color(ColAccent))
	CText     = lipgloss.NewStyle().Foreground(lipgloss.Color(ColText))
	CDim      = lipgloss.NewStyle().Foreground(lipgloss.Color(ColDim))
	CSuccess  = lipgloss.NewStyle().Foreground(lipgloss.Color(ColSuccess))
	CError    = lipgloss.NewStyle().Foreground(lipgloss.Color(ColError))
	CTemplate = lipgloss.NewStyle().Foreground(lipgloss.Color(ColTemplate))
)

// Printlnf helpers — drop-in replacements for color.Red / color.Green etc.

func PrintError(format string, a ...any) {
	fmt.Println(CError.Render(fmt.Sprintf(format, a...)))
}

func PrintSuccess(format string, a ...any) {
	fmt.Println(CSuccess.Render(fmt.Sprintf(format, a...)))
}

func PrintAccent(format string, a ...any) {
	fmt.Println(CAccent.Render(fmt.Sprintf(format, a...)))
}

func PrintDim(format string, a ...any) {
	fmt.Println(CDim.Render(fmt.Sprintf(format, a...)))
}
