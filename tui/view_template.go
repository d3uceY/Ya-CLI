package tui

import (
	"fmt"
	"strings"
)

// viewTemplate renders the template fill-in page.
func (m Model) viewTemplate() string {
	var b strings.Builder

	b.WriteString(m.renderHeader())
	b.WriteByte('\n')

	b.WriteByte('\n')
	b.WriteString(sFormTitle.Render("fill in values — " + sAccent.Render(m.templateTarget)))
	b.WriteByte('\n')
	b.WriteByte('\n')

	b.WriteString(sFormHint.Render("command: ") + renderCommandWithTokens(m.templateCommand, sDim))
	b.WriteByte('\n')
	b.WriteByte('\n')

	for i, ti := range m.templateInputs {
		label := fmt.Sprintf("[%d/%d]  %s",
			i+1, len(m.templateInputs),
			sAccent.Render(m.templateKeys[i]),
		)
		b.WriteString(sFormLabel.Render(label))
		b.WriteByte('\n')
		b.WriteString("  ")
		b.WriteString(ti.View())
		b.WriteByte('\n')
		b.WriteByte('\n')
	}

	usedLines := 6 + len(m.templateInputs)*3
	remaining := m.height - usedLines - 3
	if remaining > 0 {
		b.WriteString(m.emptyLines(remaining))
	}

	b.WriteString(m.renderStatus())
	b.WriteByte('\n')
	b.WriteString(m.renderTemplateFooter())

	return b.String()
}

// keep strings import live
var _ = strings.Builder{}


