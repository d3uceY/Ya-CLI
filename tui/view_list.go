package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// viewList renders the main shortcuts list page.
func (m Model) viewList() string {
	var b strings.Builder
	w := m.width

	b.WriteString(sHeader.Width(w).Render(m.renderHeaderContent()))
	b.WriteByte('\n')

	b.WriteString(m.renderSearchBar())
	b.WriteByte('\n')
	b.WriteString(divider(w))
	b.WriteByte('\n')

	visible := m.visibleRows()
	nameWidth := m.calcNameColWidth()

	// The description block is 2 lines (border + text), reserve both.
	selHasDesc := m.selectedKey() != "" && m.shortcuts[m.selectedKey()].Description != ""
	effective := visible
	if selHasDesc {
		effective -= 2
		if effective < 1 {
			effective = 1
		}
	}

	if len(m.filtered) == 0 {
		var emptyMsg string
		if m.searchInput.Value() != "" {
			emptyMsg = sDim.Render(fmt.Sprintf("  no matches for %q", m.searchInput.Value()))
		} else {
			emptyMsg = sDim.Render("  no shortcuts yet — press ") +
				sAccent.Render("a") +
				sDim.Render(" to add one")
		}
		b.WriteString("  " + emptyMsg)
		b.WriteByte('\n')
		for i := 1; i < visible; i++ {
			b.WriteByte('\n')
		}
	} else {
		end := m.offset + effective
		if end > len(m.filtered) {
			end = len(m.filtered)
		}
		shown := end - m.offset
		for i := m.offset; i < end; i++ {
			k := m.filtered[i]
			b.WriteString(m.renderRow(k, i == m.cursor, nameWidth))
			b.WriteByte('\n')
			if i == m.cursor && selHasDesc {
				b.WriteString(m.renderDescLine(nameWidth))
				b.WriteByte('\n')
			}
		}
		used := shown
		if selHasDesc {
			used += 2
		}
		for i := used; i < visible; i++ {
			b.WriteByte('\n')
		}
	}

	// count line
	total := len(m.shortcuts)
	filteredCount := len(m.filtered)
	var countStr string
	if m.searchInput.Value() != "" {
		countStr = fmt.Sprintf("  %d of %d", filteredCount, total)
	} else {
		countStr = fmt.Sprintf("  %d shortcut(s)", total)
	}
	if len(m.filtered) > visible {
		countStr += fmt.Sprintf("   %d/%d", m.cursor+1, len(m.filtered))
	}
	b.WriteString(sCount.Render(countStr))
	b.WriteByte('\n')

	b.WriteString(m.renderStatus())
	b.WriteByte('\n')
	b.WriteString(sFooter.Width(w).Render(m.renderListFooterContent()))

	return b.String()
}

// renderSearchBar renders the search input row.
func (m Model) renderSearchBar() string {
	label := sDim.Render("  /  ")
	var inputContent string
	if m.searchActive {
		inputContent = m.searchInput.View()
	} else if m.searchInput.Value() != "" {
		inputContent = sText.Render(m.searchInput.Value())
	} else {
		inputContent = sDim.Render("search…")
	}
	return sSearchBar.Render(label + inputContent)
}

// calcNameColWidth returns the width for the name column.
func (m Model) calcNameColWidth() int {
	w := 8
	for _, k := range m.filtered {
		if len(k) > w {
			w = len(k)
		}
	}
	return w + 2
}

// renderRow renders a single shortcut row (name + command only).
func (m Model) renderRow(name string, selected bool, nameColWidth int) string {
	s := m.shortcuts[name]
	const prefixLen = 6
	cmdColWidth := m.width - prefixLen - nameColWidth - 2
	if cmdColWidth < 10 {
		cmdColWidth = 10
	}

	cmdPart := truncateStr(s.Command, cmdColWidth)

	pinMark := ""
	if s.Pinned {
		pinMark = sAccent.Render("◆ ")
	}

	var prefix, nameStr, cmdStr string
	if selected {
		prefix = "  " + sAccent.Render("❯") + "  "
		nameStr = sNameSelected.Render(padRight(name, nameColWidth))
		cmdStr = pinMark + renderCommandWithTokens(cmdPart, sCmdSelected)
	} else {
		prefix = "     "
		nameStr = sName.Render(padRight(name, nameColWidth))
		cmdStr = pinMark + renderCommandWithTokens(cmdPart, sCmd)
	}

	return prefix + nameStr + "  " + cmdStr
}

// renderDescLine renders the 2-line description block (border + text) for the
// selected shortcut, aligned to start under the command column.
func (m Model) renderDescLine(nameColWidth int) string {
	s := m.shortcuts[m.selectedKey()]
	// visual prefix width = 2 spaces + ❯ (1 cell) + 2 spaces = 5
	const visualPfx = 5
	indent := strings.Repeat(" ", visualPfx+nameColWidth+2)
	availW := m.width - visualPfx - nameColWidth - 2
	if availW < 10 || s.Description == "" {
		return ""
	}
	borderLine := indent + sDim.Render(strings.Repeat("─", availW))
	desc := truncateStr(s.Description, availW-1)
	contentLine := indent + sDescLine.Render(desc)
	return borderLine + "\n" + contentLine
}

// blankRow is kept for interface compatibility.
func (m Model) blankRow() string {
	return sRowNormal.Width(m.width).Render("")
}

// renderHeaderContent returns the inner content string for the header.
func (m Model) renderHeaderContent() string {
	banner := renderBanner()
	version := sDim.Render(lipgloss.NewStyle().Render("  " + getVersion()))
	gap := m.width - lipgloss.Width(banner) - lipgloss.Width(version) - 2
	if gap < 1 {
		gap = 1
	}
	return banner + strings.Repeat(" ", gap) + version
}

// renderListFooterContent returns the inner content string for the list footer.
func (m Model) renderListFooterContent() string {
	return renderKeyHints(
		keyHint("↑↓", "nav"),
		keyHint("enter", "run"),
		keyHint("/", "search"),
		keyHint("a", "add"),
		keyHint("e", "edit"),
		keyHint("r", "rename"),
		keyHint("d", "delete"),
		keyHint("p", "pin"),
		keyHint("h", "history"),
		keyHint("D", "dirs"),
		keyHint("i/o", "import/export"),
		keyHint("f", "fullscreen"),
		keyHint("?", "help"),
		keyHint("q", "quit"),
	)
}
