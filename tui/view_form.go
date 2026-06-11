package tui

import (
	"strings"
)

// viewForm renders the shared form page (add / edit / rename / import / export).
func (m Model) viewForm() string {
	var b strings.Builder
	w := m.width

	b.WriteString(m.renderHeader())
	b.WriteByte('\n')

	// ── title ─────────────────────────────────────────────────────────────
	b.WriteByte('\n')
	b.WriteString(sFormTitle.Render(m.formTitle()))
	b.WriteByte('\n')
	b.WriteByte('\n')

	// ── fields ────────────────────────────────────────────────────────────
	for i, ti := range m.formInputs {
		label := m.formFieldLabel(i)
		b.WriteString(sFormLabel.Render(label))
		b.WriteByte('\n')
		b.WriteString("  ")
		b.WriteString(ti.View())
		b.WriteByte('\n')
		b.WriteByte('\n')
	}

	// ── hint ──────────────────────────────────────────────────────────────
	if hint := m.formHint(); hint != "" {
		b.WriteString(sFormHint.Render(hint))
		b.WriteByte('\n')
	}

	// fill remaining space
	usedLines := 4 + len(m.formInputs)*3 + 1
	remaining := m.height - usedLines - 3
	if remaining > 0 {
		b.WriteString(m.emptyLines(remaining))
	}

	b.WriteString(m.renderStatus())
	b.WriteByte('\n')
	b.WriteString(sFooter.Width(w).Render(renderKeyHints(
		keyHint("tab", "next field"),
		keyHint("enter", "submit"),
		keyHint("esc", "cancel"),
	)))

	return b.String()
}

func (m Model) formTitle() string {
	switch m.formMode {
	case formAdd:
		return "add shortcut"
	case formEdit:
		return "edit — " + sAccent.Render(m.selectedKey())
	case formRename:
		return "rename — " + sAccent.Render(m.selectedKey())
	case formImport:
		return "import shortcuts"
	case formExport:
		return "export shortcuts"
	case formAddDir:
		return "add saved directory"
	}
	return ""
}

func (m Model) formFieldLabel(i int) string {
	switch m.formMode {
	case formAdd:
		switch i {
		case 0:
			return "name"
		case 1:
			return "command"
		case 2:
			return "description  (optional)"
		}
	case formEdit:
		if i == 0 {
			return "command"
		}
		return "description  (optional)"
	case formRename:
		return "new name"
	case formImport:
		return "file path"
	case formExport:
		if i == 0 {
			return "directory"
		}
		return "filename"
	case formAddDir:
		if i == 0 {
			return "name"
		}
		return "path"
	}
	return ""
}

func (m Model) formHint() string {
	switch m.formMode {
	case formAdd, formEdit:
		return "use {placeholder} for interactive prompts at run time"
	}
	return ""
}

