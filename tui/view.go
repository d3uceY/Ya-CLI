package tui

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/d3uceY/Ya-CLI/utils"
)

// getVersion is a thin wrapper so view_list.go can call it without importing utils.
func getVersion() string { return utils.GetAppVersion() }

// View routes rendering to the appropriate sub-view.
func (m Model) View() string {
	switch m.page {
	case pageList:
		return m.viewList()
	case pageAdd, pageEdit, pageRename, pageImport, pageExport, pageAddDir:
		return m.viewForm()
	case pageTemplate:
		return m.viewTemplate()
	case pageConfirmDelete, pageConfirmClearHistory:
		return m.viewConfirm()
	case pageHelp:
		return m.viewHelp()
	case pageHistory:
		return m.viewHistory()
	case pageDirs:
		return m.viewDirs()
	}
	return m.viewList()
}

// ── shared layout blocks ─────────────────────────────────────────────────────

// renderHeader renders the top bar with a bottom border.
func (m Model) renderHeader() string {
	return sHeader.Width(m.width).Render(m.renderHeaderContent())
}

// renderStatus renders the status bar line (no border).
func (m Model) renderStatus() string {
	var content string
	switch {
	case m.statusMsg == "":
		content = ""
	case m.statusIsError:
		content = sError.Render("  ✗  " + m.statusMsg)
	default:
		content = sSuccess.Render("  ✓  " + m.statusMsg)
	}
	return sStatus.Render(content)
}

// renderFormFooter renders the key-hint bar for form pages.
func (m Model) renderFormFooter() string {
	return sFooter.Width(m.width).Render(renderKeyHints(
		keyHint("tab", "next field"),
		keyHint("enter", "submit"),
		keyHint("esc", "cancel"),
	))
}

// renderTemplateFooter renders the key-hint bar for the template page.
func (m Model) renderTemplateFooter() string {
	return sFooter.Width(m.width).Render(renderKeyHints(
		keyHint("tab", "next"),
		keyHint("enter", "next / run"),
		keyHint("esc", "cancel"),
	))
}

// renderConfirmFooter renders the key-hint bar for the confirm pages.
func (m Model) renderConfirmFooter() string {
	return sFooter.Width(m.width).Render(renderKeyHints(
		keyHint("y", "yes"),
		keyHint("n / esc", "cancel"),
	))
}

// renderHelpFooter renders the key-hint bar for the help page.
func (m Model) renderHelpFooter() string {
	return sFooter.Width(m.width).Render(renderKeyHints(
		keyHint("esc / ?", "close"),
		keyHint("q", "quit"),
	))
}

// emptyLines returns n blank lines.
func (m Model) emptyLines(n int) string {
	if n <= 0 {
		return ""
	}
	return strings.Repeat("\n", n)
}

// renderHeaderContent is defined in view_list.go.
var _ = lipgloss.Width // keep lipgloss import alive


