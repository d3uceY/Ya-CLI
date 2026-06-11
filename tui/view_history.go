package tui

import (
	"fmt"
	"strings"
	"time"
)

// viewHistory renders the run history page.
func (m Model) viewHistory() string {
	var b strings.Builder
	w := m.width

	b.WriteString(m.renderHeader())
	b.WriteByte('\n')

	b.WriteByte('\n')
	b.WriteString(sFormTitle.Render("  run history"))
	b.WriteByte('\n')
	b.WriteString(divider(w))
	b.WriteByte('\n')

	visible := m.visibleRows()

	if len(m.history) == 0 {
		b.WriteString(sDim.Render("  no history yet — run a shortcut to get started"))
		b.WriteByte('\n')
		for i := 1; i < visible; i++ {
			b.WriteByte('\n')
		}
	} else {
		end := m.historyOffset + visible
		if end > len(m.history) {
			end = len(m.history)
		}
		shown := end - m.historyOffset
		nameW := m.historyNameColWidth()
		for i := m.historyOffset; i < end; i++ {
			b.WriteString(m.renderHistoryRow(i, nameW))
			b.WriteByte('\n')
		}
		for i := shown; i < visible; i++ {
			b.WriteByte('\n')
		}
	}

	b.WriteString(sCount.Render(fmt.Sprintf("  %d entries", len(m.history))))
	b.WriteByte('\n')
	b.WriteString(m.renderStatus())
	b.WriteByte('\n')
	b.WriteString(sFooter.Width(w).Render(renderKeyHints(
		keyHint("↑↓", "scroll"),
		keyHint("c", "clear history"),
		keyHint("h / esc", "back"),
		keyHint("q", "quit"),
	)))

	return b.String()
}

func (m Model) historyNameColWidth() int {
	w := 8
	for _, e := range m.history {
		if len(e.ShortcutName) > w {
			w = len(e.ShortcutName)
		}
	}
	return w + 2
}

func (m Model) renderHistoryRow(i, nameW int) string {
	e := m.history[i]
	selected := i == m.historyCursor

	// parse timestamp
	ts := e.Timestamp
	if t, err := time.Parse(time.RFC3339, e.Timestamp); err == nil {
		ts = t.Local().Format("01/02 15:04")
	}

	const prefixLen = 6
	// nameW + ts(11) + 2sep + 2sep = nameW + 15
	cmdColW := m.width - prefixLen - nameW - 15
	if cmdColW < 10 {
		cmdColW = 10
	}

	truncCmd := truncateStr(e.Command, cmdColW)
	tsStr := sDim.Render(ts)

	var prefix, nameStr, cmdStr string
	if selected {
		prefix = "  " + sAccent.Render("❯") + "  "
		nameStr = sNameSelected.Render(padRight(e.ShortcutName, nameW))
		cmdStr = renderCommandWithTokens(truncCmd, sCmdSelected)
	} else {
		prefix = "     "
		nameStr = sName.Render(padRight(e.ShortcutName, nameW))
		cmdStr = renderCommandWithTokens(truncCmd, sCmd)
	}

	return prefix + tsStr + "  " + nameStr + "  " + cmdStr
}

// viewConfirmClearHistory is handled via viewConfirm with page=pageConfirmClearHistory.
// We reuse viewConfirm but display different text.
func (m Model) viewConfirm() string {
	var b strings.Builder
	w := m.width

	b.WriteString(m.renderHeader())
	b.WriteByte('\n')
	b.WriteByte('\n')

	if m.page == pageConfirmClearHistory {
		b.WriteString(sConfirmTitle.Render("  clear all history?"))
		b.WriteByte('\n')
		b.WriteByte('\n')
		b.WriteString(sConfirmBody.Render("  This will permanently delete all run history entries."))
	} else {
		b.WriteString(sConfirmTitle.Render("  delete shortcut?"))
		b.WriteByte('\n')
		b.WriteByte('\n')
		b.WriteString(sConfirmBody.Render("  name:     " + sAccent.Render(m.confirmName)))
		b.WriteByte('\n')
		b.WriteString(sConfirmBody.Render("  command:  " + renderCommandWithTokens(m.confirmCommand, sText)))
	}

	b.WriteByte('\n')
	b.WriteByte('\n')
	b.WriteString(
		"  " + sError.Render("[ y ]") + sDim.Render("  yes") +
			"     " +
			sDim.Render("[ n ]  cancel"),
	)
	b.WriteByte('\n')

	remaining := m.height - 10
	if remaining > 0 {
		b.WriteString(m.emptyLines(remaining))
	}

	b.WriteString(m.renderStatus())
	b.WriteByte('\n')
	b.WriteString(sFooter.Width(w).Render(renderKeyHints(
		keyHint("y", "yes"),
		keyHint("n / esc", "cancel"),
	)))

	return b.String()
}

// viewDirs renders the saved directories page.
func (m Model) viewDirs() string {
	var b strings.Builder
	w := m.width

	b.WriteString(m.renderHeader())
	b.WriteByte('\n')
	b.WriteByte('\n')
	b.WriteString(sFormTitle.Render("  saved directories"))
	b.WriteByte('\n')
	b.WriteString(divider(w))
	b.WriteByte('\n')

	visible := m.visibleRows()

	if len(m.dirs) == 0 {
		b.WriteString(sDim.Render("  no saved directories — press ") +
			sAccent.Render("a") + sDim.Render(" to add one"))
		b.WriteByte('\n')
		for i := 1; i < visible; i++ {
			b.WriteByte('\n')
		}
	} else {
		nameW := m.dirsNameColWidth()
		end := len(m.dirs)
		if end > visible {
			end = visible
		}
		shown := end
		for i := 0; i < end; i++ {
			b.WriteString(m.renderDirRow(i, nameW))
			b.WriteByte('\n')
		}
		for i := shown; i < visible; i++ {
			b.WriteByte('\n')
		}
	}

	b.WriteString(sCount.Render(fmt.Sprintf("  %d director(ies)", len(m.dirs))))
	b.WriteByte('\n')
	b.WriteString(m.renderStatus())
	b.WriteByte('\n')
	b.WriteString(sFooter.Width(w).Render(renderKeyHints(
		keyHint("↑↓", "nav"),
		keyHint("a", "add"),
		keyHint("d", "delete"),
		keyHint("D / esc", "back"),
		keyHint("q", "quit"),
	)))

	return b.String()
}

func (m Model) dirsNameColWidth() int {
	w := 8
	for _, d := range m.dirs {
		if len(d.Name) > w {
			w = len(d.Name)
		}
	}
	return w + 2
}

func (m Model) renderDirRow(i, nameW int) string {
	d := m.dirs[i]
	selected := i == m.dirsCursor

	pathW := m.width - 6 - nameW - 2
	if pathW < 10 {
		pathW = 10
	}
	truncPath := truncateStr(d.Path, pathW)

	var prefix, nameStr, pathStr string
	if selected {
		prefix = "  " + sAccent.Render("❯") + "  "
		nameStr = sNameSelected.Render(padRight(d.Name, nameW))
		pathStr = sCmdSelected.Render(truncPath)
	} else {
		prefix = "     "
		nameStr = sName.Render(padRight(d.Name, nameW))
		pathStr = sCmd.Render(truncPath)
	}

	return prefix + nameStr + "  " + pathStr
}

// strings import used above via sDim, etc. — keep it to silence the linter.
var _ = strings.Builder{}
