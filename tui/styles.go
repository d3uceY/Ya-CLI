package tui

import (
	"regexp"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// templateRe matches {placeholder} tokens in command strings.
var templateRe = regexp.MustCompile(`\{([^}]+)\}`)

// Minimal, muted palette. No background fills — let the terminal breathe.
// The logo's cyan is kept as a single, reserved accent used sparingly.
const (
	colAccent   = "#5F9EA0" // cadet blue — muted cyan, the only brand colour
	colText     = "#CCCCCC" // soft white for body text
	colDim      = "#666666" // dim gray — secondary info, hints
	colBorder   = "#3A3A3A" // subtle dark border / divider
	colSelected = "#FFFFFF" // bold white for the selected row name
	colSuccess  = "#7EC8A4" // muted sage green
	colError    = "#CC6666" // muted rose red
	colTemplate = "#C8A96E" // muted amber — highlights {tokens}
)

var (
	// ── base text ─────────────────────────────────────────────────────────
	sAccent   = lipgloss.NewStyle().Foreground(lipgloss.Color(colAccent))
	sText     = lipgloss.NewStyle().Foreground(lipgloss.Color(colText))
	sDim      = lipgloss.NewStyle().Foreground(lipgloss.Color(colDim))
	sSuccess  = lipgloss.NewStyle().Foreground(lipgloss.Color(colSuccess))
	sError    = lipgloss.NewStyle().Foreground(lipgloss.Color(colError))
	sTemplate = lipgloss.NewStyle().Foreground(lipgloss.Color(colTemplate))

	// keep sMuted as an alias for backward compat with other files
	sMuted  = sDim
	sPrimary = sAccent

	// ── header ────────────────────────────────────────────────────────────
	// single bottom border, no background fill
	sHeader = lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).
		BorderBottom(true).
		BorderForeground(lipgloss.Color(colBorder)).
		Padding(0, 1)

	// ── footer ────────────────────────────────────────────────────────────
	sFooter = lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).
		BorderTop(true).
		BorderForeground(lipgloss.Color(colBorder)).
		Foreground(lipgloss.Color(colDim)).
		Padding(0, 1)

	// ── search bar ────────────────────────────────────────────────────────
	sSearchBar = lipgloss.NewStyle().Padding(0, 1)

	// ── list rows — no background fills ──────────────────────────────────
	sRowNormal   = lipgloss.NewStyle()
	sRowSelected = lipgloss.NewStyle()

	// ── shortcut name ────────────────────────────────────────────────────
	sName         = lipgloss.NewStyle().Foreground(lipgloss.Color(colDim))
	sNameSelected = lipgloss.NewStyle().Foreground(lipgloss.Color(colSelected)).Bold(true)

	// ── command text ─────────────────────────────────────────────────────
	sCmd         = lipgloss.NewStyle().Foreground(lipgloss.Color(colDim))
	sCmdSelected = lipgloss.NewStyle().Foreground(lipgloss.Color(colText))

	// ── divider ──────────────────────────────────────────────────────────
	sDivider = lipgloss.NewStyle().Foreground(lipgloss.Color(colBorder))

	// ── count / info ─────────────────────────────────────────────────────
	sCount = lipgloss.NewStyle().Foreground(lipgloss.Color(colDim))

	// ── status bar ───────────────────────────────────────────────────────
	sStatus = lipgloss.NewStyle().Padding(0, 1)

	// ── form ─────────────────────────────────────────────────────────────
	sFormTitle = lipgloss.NewStyle().
		Foreground(lipgloss.Color(colText)).
		Bold(true).
		Padding(0, 2)

	sFormLabel = lipgloss.NewStyle().
		Foreground(lipgloss.Color(colAccent)).
		Padding(0, 2)

	sFormHint = lipgloss.NewStyle().
		Foreground(lipgloss.Color(colDim)).
		Italic(true).
		Padding(0, 2)

	// ── confirm dialog ───────────────────────────────────────────────────
	sConfirmTitle = lipgloss.NewStyle().
		Foreground(lipgloss.Color(colError)).
		Bold(true).
		Padding(0, 2)

	sConfirmBody = lipgloss.NewStyle().
		Foreground(lipgloss.Color(colText)).
		Padding(0, 2)

	// ── help view ────────────────────────────────────────────────────────
	sHelpTitle = lipgloss.NewStyle().
		Foreground(lipgloss.Color(colText)).
		Bold(true).
		Padding(0, 2)

	sHelpSection = lipgloss.NewStyle().
		Foreground(lipgloss.Color(colAccent)).
		Padding(0, 4)

	sHelpKey  = lipgloss.NewStyle().Foreground(lipgloss.Color(colText)).Bold(true)
	sHelpDesc = lipgloss.NewStyle().Foreground(lipgloss.Color(colDim))
)

// renderBanner returns the header title text.
func renderBanner() string {
	return sAccent.Bold(true).Render("ya") +
		sDim.Render("  shortcut runner")
}

// keyHint renders a key+description pair for use in footers.
func keyHint(k, desc string) string {
	return sAccent.Render(k) + sDim.Render(" "+desc)
}

// renderKeyHints joins hints with muted bullet separators.
func renderKeyHints(hints ...string) string {
	sep := sDim.Render("  ·  ")
	return strings.Join(hints, sep)
}

// divider renders a horizontal rule of `width` characters.
func divider(width int) string {
	if width <= 0 {
		return ""
	}
	return sDivider.Render(strings.Repeat("─", width))
}

// renderCommandWithTokens renders a command string with {tokens} styled in amber.
func renderCommandWithTokens(cmd string, baseStyle lipgloss.Style) string {
	parts := templateRe.Split(cmd, -1)
	tokens := templateRe.FindAllString(cmd, -1)
	var sb strings.Builder
	for i, part := range parts {
		if part != "" {
			sb.WriteString(baseStyle.Render(part))
		}
		if i < len(tokens) {
			sb.WriteString(sTemplate.Render(tokens[i]))
		}
	}
	return sb.String()
}

// truncateStr truncates s to maxLen visible characters, appending "…" if needed.
func truncateStr(s string, maxLen int) string {
	if maxLen <= 0 {
		return ""
	}
	if len(s) <= maxLen {
		return s
	}
	if maxLen == 1 {
		return "…"
	}
	return s[:maxLen-1] + "…"
}

// padRight pads s with spaces to reach the given width.
func padRight(s string, width int) string {
	n := width - len(s)
	if n <= 0 {
		return s
	}
	return s + strings.Repeat(" ", n)
}

// fillBg is kept for compatibility but now just returns content unchanged
// (no background fills in the minimal design).
func fillBg(content string, _ int, _ string) string {
	return content
}
