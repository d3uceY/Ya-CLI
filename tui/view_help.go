package tui

import (
	"fmt"
	"strings"
)

// viewHelp renders the keyboard shortcuts reference page.
func (m Model) viewHelp() string {
	var b strings.Builder
	w := m.width

	b.WriteString(m.renderHeader())
	b.WriteByte('\n')

	b.WriteByte('\n')
	b.WriteString(sHelpTitle.Render("  keyboard shortcuts"))
	b.WriteByte('\n')
	b.WriteByte('\n')

	type entry struct{ key, desc string }
	sections := []struct {
		title   string
		entries []entry
	}{
		{
			"navigation",
			[]entry{
				{"↑ / k", "move up"},
				{"↓ / j", "move down"},
			},
		},
		{
			"actions",
			[]entry{
				{"enter", "run selected shortcut (TUI closes, command runs natively)"},
				{"a", "add new shortcut"},
				{"e", "edit shortcut command / description"},
				{"r", "rename shortcut"},
				{"d / x", "delete shortcut"},
				{"p", "pin / unpin shortcut"},
			},
		},
		{
			"search",
			[]entry{
				{"/", "activate search bar (filters by name, command, description)"},
				{"esc", "clear / exit search"},
			},
		},
		{
			"views",
			[]entry{
				{"h", "run history"},
				{"D", "saved directories"},
			},
		},
		{
			"import / export",
			[]entry{
				{"i", "import shortcuts from JSON file"},
				{"o", "export shortcuts to JSON file"},
			},
		},
		{
			"general",
			[]entry{
				{"f", "toggle fullscreen"},
				{"?", "toggle this help"},
				{"q / ctrl+c", "quit"},
				{"esc", "cancel / go back"},
			},
		},
	}

	for _, sec := range sections {
		b.WriteString(sHelpSection.Render(sec.title))
		b.WriteByte('\n')
		for _, e := range sec.entries {
			keyW := 14
			k := sHelpKey.Render(fmt.Sprintf("%-*s", keyW, e.key))
			d := sHelpDesc.Render(e.desc)
			b.WriteString("        " + k + "  " + d)
			b.WriteByte('\n')
		}
		b.WriteByte('\n')
	}

	lineCount := 4 + len(sections)
	for _, s := range sections {
		lineCount += len(s.entries) + 1
	}
	remaining := m.height - lineCount - 3
	if remaining > 0 {
		b.WriteString(strings.Repeat("\n", remaining))
	}

	b.WriteString(m.renderStatus())
	b.WriteByte('\n')
	b.WriteString(sFooter.Width(w).Render(renderKeyHints(
		keyHint("esc / ?", "close help"),
		keyHint("q", "quit"),
	)))

	return b.String()
}


