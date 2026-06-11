package tui

import (
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/d3uceY/Ya-CLI/utils"
)

// enableKeysMsg is sent after the key-ignore cooldown expires.
type enableKeysMsg struct{}

// Update is the main message dispatcher.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil

	case enableKeysMsg:
		m.ignoreKeys = false
		return m, nil

	case tea.KeyMsg:
		if m.ignoreKeys {
			return m, nil
		}
		return m.handleKey(msg)
	}

	return m.updateInputs(msg)
}

// ── top-level key router ─────────────────────────────────────────────────────

func (m Model) handleKey(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch m.page {
	case pageList:
		return m.handleListKey(msg)
	case pageAdd, pageEdit, pageRename, pageImport, pageExport, pageAddDir:
		return m.handleFormKey(msg)
	case pageTemplate:
		return m.handleTemplateKey(msg)
	case pageConfirmDelete, pageConfirmClearHistory:
		return m.handleConfirmKey(msg)
	case pageHelp:
		return m.handleHelpKey(msg)
	case pageHistory:
		return m.handleHistoryKey(msg)
	case pageDirs:
		return m.handleDirsKey(msg)
	}
	return m, nil
}

// ── list page ────────────────────────────────────────────────────────────────

func (m Model) handleListKey(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	if m.searchActive {
		return m.handleSearchKey(msg)
	}

	switch {
	case key.Matches(msg, keys.Quit):
		return m, tea.Quit

	case key.Matches(msg, keys.Help):
		m.page = pageHelp
		return m, nil

	case key.Matches(msg, keys.FullScreen):
		m.fullscreen = !m.fullscreen
		return m, nil

	case key.Matches(msg, keys.Up):
		m.cursor--
		m.clampCursor()
		return m, nil

	case key.Matches(msg, keys.Down):
		m.cursor++
		m.clampCursor()
		return m, nil

	case key.Matches(msg, keys.Search):
		m.searchActive = true
		m.searchInput.Focus()
		return m, textinput.Blink

	case key.Matches(msg, keys.Back):
		if m.searchInput.Value() != "" {
			m.searchInput.SetValue("")
			m.applyFilter()
			m.cursor = 0
			m.offset = 0
		} else {
			m.searchActive = false
			m.searchInput.Blur()
		}
		return m, nil

	case key.Matches(msg, keys.Enter):
		return m.runSelectedShortcut()

	case key.Matches(msg, keys.Add):
		m.formInputs = []textinput.Model{
			newFormInput("name", "", true),
			newFormInput("command  (use {placeholder} for templates)", "", false),
			newFormInput("description  (optional)", "", false),
		}
		m.formFocus = 0
		m.formMode = formAdd
		m.page = pageAdd
		m.clearStatus()
		return m, textinput.Blink

	case key.Matches(msg, keys.Edit):
		sel := m.selectedKey()
		if sel == "" {
			return m, nil
		}
		s := m.shortcuts[sel]
		m.formInputs = []textinput.Model{
			newFormInput("command", s.Command, true),
			newFormInput("description  (optional)", s.Description, false),
		}
		m.formFocus = 0
		m.formMode = formEdit
		m.page = pageEdit
		m.clearStatus()
		return m, textinput.Blink

	case key.Matches(msg, keys.Rename):
		sel := m.selectedKey()
		if sel == "" {
			return m, nil
		}
		m.formInputs = []textinput.Model{
			newFormInput("new name", sel, true),
		}
		m.formFocus = 0
		m.formMode = formRename
		m.page = pageRename
		m.clearStatus()
		return m, textinput.Blink

	case key.Matches(msg, keys.Delete):
		sel := m.selectedKey()
		if sel == "" {
			return m, nil
		}
		m.confirmName = sel
		m.confirmCommand = m.shortcuts[sel].Command
		m.page = pageConfirmDelete
		m.clearStatus()
		return m, nil

	case key.Matches(msg, keys.Pin):
		sel := m.selectedKey()
		if sel == "" {
			return m, nil
		}
		if err := utils.TogglePin(sel); err != nil {
			m.setStatus("error: "+err.Error(), true)
		} else {
			m.reloadShortcuts()
			m.moveCursorTo(sel)
		}
		return m, nil

	case key.Matches(msg, keys.Import):
		m.formInputs = []textinput.Model{
			newFormInput("path to JSON file", "", true),
		}
		m.formFocus = 0
		m.formMode = formImport
		m.page = pageImport
		m.clearStatus()
		return m, textinput.Blink

	case key.Matches(msg, keys.Export):
		m.formInputs = []textinput.Model{
			newFormInput("output directory", "", true),
			newFormInput("filename", "shortcuts.json", false),
		}
		m.formFocus = 0
		m.formMode = formExport
		m.page = pageExport
		m.clearStatus()
		return m, textinput.Blink

	case key.Matches(msg, keys.History):
		m.reloadHistory()
		m.historyCursor = 0
		m.historyOffset = 0
		m.page = pageHistory
		return m, nil

	case key.Matches(msg, keys.Dirs):
		m.reloadDirs()
		m.dirsCursor = 0
		m.page = pageDirs
		return m, nil
	}
	return m, nil
}

// ── search input mode ────────────────────────────────────────────────────────

func (m Model) handleSearchKey(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch {
	case key.Matches(msg, keys.Back):
		m.searchActive = false
		m.searchInput.Blur()
		return m, nil

	case key.Matches(msg, keys.Enter):
		m.searchActive = false
		m.searchInput.Blur()
		return m, nil

	// Arrow-only navigation while typing — do NOT intercept j/k so they
	// go into the search box as regular characters.
	case msg.String() == "up":
		m.cursor--
		m.clampCursor()
		return m, nil

	case msg.String() == "down":
		m.cursor++
		m.clampCursor()
		return m, nil

	case msg.String() == "ctrl+c":
		return m, tea.Quit
	}

	var cmd tea.Cmd
	m.searchInput, cmd = m.searchInput.Update(msg)
	m.applyFilter()
	m.cursor = 0
	m.offset = 0
	return m, cmd
}

// ── form pages ───────────────────────────────────────────────────────────────

func (m Model) handleFormKey(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch {
	case key.Matches(msg, keys.Back):
		if m.formMode == formAddDir {
			m.page = pageDirs
		} else {
			m.page = pageList
		}
		m.clearStatus()
		return m, nil

	case key.Matches(msg, keys.Quit):
		return m, tea.Quit

	case key.Matches(msg, keys.Tab):
		m.formInputs[m.formFocus].Blur()
		m.formFocus = (m.formFocus + 1) % len(m.formInputs)
		m.formInputs[m.formFocus].Focus()
		return m, textinput.Blink

	case key.Matches(msg, keys.Enter):
		if m.formFocus < len(m.formInputs)-1 {
			m.formInputs[m.formFocus].Blur()
			m.formFocus++
			m.formInputs[m.formFocus].Focus()
			return m, textinput.Blink
		}
		return m.submitForm()
	}

	var cmd tea.Cmd
	m.formInputs[m.formFocus], cmd = m.formInputs[m.formFocus].Update(msg)
	return m, cmd
}

func (m Model) submitForm() (tea.Model, tea.Cmd) {
	switch m.formMode {

	case formAdd:
		name := strings.TrimSpace(m.formInputs[0].Value())
		command := strings.TrimSpace(m.formInputs[1].Value())
		desc := strings.TrimSpace(m.formInputs[2].Value())
		if name == "" || command == "" {
			m.setStatus("name and command cannot be empty.", true)
			return m, nil
		}
		if _, exists := m.shortcuts[name]; exists {
			m.setStatus(fmt.Sprintf("'%s' already exists — use edit (e) to change it.", name), true)
			return m, nil
		}
		if err := utils.AddShortcutFull(name, command, desc); err != nil {
			m.setStatus("error: "+err.Error(), true)
			return m, nil
		}
		m.reloadShortcuts()
		m.moveCursorTo(name)
		m.setStatus(fmt.Sprintf("added '%s'.", name), false)
		m.page = pageList

	case formEdit:
		sel := m.selectedKey()
		if sel == "" {
			m.page = pageList
			return m, nil
		}
		command := strings.TrimSpace(m.formInputs[0].Value())
		desc := strings.TrimSpace(m.formInputs[1].Value())
		if command == "" {
			m.setStatus("command cannot be empty.", true)
			return m, nil
		}
		if err := utils.AddShortcutFull(sel, command, desc); err != nil {
			m.setStatus("error: "+err.Error(), true)
			return m, nil
		}
		m.reloadShortcuts()
		m.moveCursorTo(sel)
		m.setStatus(fmt.Sprintf("updated '%s'.", sel), false)
		m.page = pageList

	case formRename:
		oldName := m.selectedKey()
		if oldName == "" {
			m.page = pageList
			return m, nil
		}
		newName := strings.TrimSpace(m.formInputs[0].Value())
		if newName == "" {
			m.setStatus("name cannot be empty.", true)
			return m, nil
		}
		if newName == oldName {
			m.page = pageList
			return m, nil
		}
		if err := utils.RenameShortcut(oldName, newName); err != nil {
			m.setStatus("error: "+err.Error(), true)
			return m, nil
		}
		m.reloadShortcuts()
		m.moveCursorTo(newName)
		m.setStatus(fmt.Sprintf("renamed '%s' → '%s'.", oldName, newName), false)
		m.page = pageList

	case formImport:
		path := strings.TrimSpace(m.formInputs[0].Value())
		if path == "" {
			m.setStatus("path cannot be empty.", true)
			return m, nil
		}
		if err := utils.ImportShortcuts(path); err != nil {
			m.setStatus("import failed: "+err.Error(), true)
			return m, nil
		}
		m.reloadShortcuts()
		m.setStatus(fmt.Sprintf("imported from '%s'.", path), false)
		m.page = pageList

	case formExport:
		dir := strings.TrimSpace(m.formInputs[0].Value())
		filename := strings.TrimSpace(m.formInputs[1].Value())
		if dir == "" {
			m.setStatus("directory cannot be empty.", true)
			return m, nil
		}
		if filename == "" {
			filename = "shortcuts.json"
		}
		filePath := dir + "/" + filename
		if err := utils.ExportShortcuts(filePath); err != nil {
			m.setStatus("export failed: "+err.Error(), true)
			return m, nil
		}
		m.setStatus(fmt.Sprintf("exported to '%s'.", filePath), false)
		m.page = pageList

	case formAddDir:
		name := strings.TrimSpace(m.formInputs[0].Value())
		path := strings.TrimSpace(m.formInputs[1].Value())
		if name == "" || path == "" {
			m.setStatus("name and path cannot be empty.", true)
			return m, nil
		}
		if err := utils.AddSavedDirectory(name, path); err != nil {
			m.setStatus("error: "+err.Error(), true)
			return m, nil
		}
		m.reloadDirs()
		m.setStatus(fmt.Sprintf("saved directory '%s'.", name), false)
		m.page = pageDirs
	}

	return m, nil
}

// ── template fill-in page ────────────────────────────────────────────────────

func (m Model) handleTemplateKey(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch {
	case key.Matches(msg, keys.Back):
		m.page = pageList
		m.clearStatus()
		return m, nil

	case key.Matches(msg, keys.Quit):
		return m, tea.Quit

	case key.Matches(msg, keys.Tab):
		m.templateInputs[m.templateFocus].Blur()
		m.templateFocus = (m.templateFocus + 1) % len(m.templateInputs)
		m.templateInputs[m.templateFocus].Focus()
		return m, textinput.Blink

	case key.Matches(msg, keys.Enter):
		if m.templateFocus < len(m.templateInputs)-1 {
			m.templateInputs[m.templateFocus].Blur()
			m.templateFocus++
			m.templateInputs[m.templateFocus].Focus()
			return m, textinput.Blink
		}
		// All filled — resolve command and quit TUI to run natively.
		resolved := resolveTemplateCommand(m.templateCommand, m.templateKeys, m.templateInputs)
		if len(m.templateExtraArgs) > 0 {
			resolved += " " + strings.Join(m.templateExtraArgs, " ")
		}
		m.pendingCmd = resolved
		m.pendingCmdName = m.templateTarget
		return m, tea.Quit
	}

	var cmd tea.Cmd
	m.templateInputs[m.templateFocus], cmd = m.templateInputs[m.templateFocus].Update(msg)
	return m, cmd
}

// ── confirm pages ────────────────────────────────────────────────────────────

func (m Model) handleConfirmKey(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch {
	case key.Matches(msg, keys.Yes):
		if m.page == pageConfirmClearHistory {
			if err := utils.ClearRunHistory(); err != nil {
				m.setStatus("error: "+err.Error(), true)
			} else {
				m.history = nil
				m.historyCursor = 0
				m.historyOffset = 0
				m.setStatus("history cleared.", false)
			}
			m.page = pageHistory
		} else {
			// confirm delete
			if err := utils.RemoveShortcut(m.confirmName); err != nil {
				m.setStatus("error: "+err.Error(), true)
			} else {
				m.setStatus(fmt.Sprintf("deleted '%s'.", m.confirmName), false)
			}
			if m.cursor >= len(m.filtered)-1 && m.cursor > 0 {
				m.cursor--
			}
			m.reloadShortcuts()
			m.page = pageList
		}

	case key.Matches(msg, keys.No), key.Matches(msg, keys.Back):
		if m.page == pageConfirmClearHistory {
			m.page = pageHistory
		} else {
			m.page = pageList
		}
		m.clearStatus()

	case key.Matches(msg, keys.Quit):
		return m, tea.Quit
	}
	return m, nil
}

// ── help page ────────────────────────────────────────────────────────────────

func (m Model) handleHelpKey(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch {
	case key.Matches(msg, keys.Back), key.Matches(msg, keys.Help):
		m.page = pageList
	case key.Matches(msg, keys.Quit):
		return m, tea.Quit
	}
	return m, nil
}

// ── history page ─────────────────────────────────────────────────────────────

func (m Model) handleHistoryKey(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch {
	case key.Matches(msg, keys.Back), key.Matches(msg, keys.History):
		m.page = pageList
		return m, nil

	case key.Matches(msg, keys.Up):
		m.historyCursor--
		m.clampHistoryCursor()

	case key.Matches(msg, keys.Down):
		m.historyCursor++
		m.clampHistoryCursor()

	case key.Matches(msg, keys.ClearHistory):
		m.page = pageConfirmClearHistory
		return m, nil

	case key.Matches(msg, keys.Quit):
		return m, tea.Quit
	}
	return m, nil
}

// ── saved dirs page ───────────────────────────────────────────────────────────

func (m Model) handleDirsKey(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch {
	case key.Matches(msg, keys.Back), key.Matches(msg, keys.Dirs):
		m.page = pageList
		return m, nil

	case key.Matches(msg, keys.Up):
		m.dirsCursor--
		m.clampDirsCursor()

	case key.Matches(msg, keys.Down):
		m.dirsCursor++
		m.clampDirsCursor()

	case key.Matches(msg, keys.Add):
		m.formInputs = []textinput.Model{
			newFormInput("name", "", true),
			newFormInput("path  (e.g. ~/projects/myapp)", "", false),
		}
		m.formFocus = 0
		m.formMode = formAddDir
		m.page = pageAddDir
		m.clearStatus()
		return m, textinput.Blink

	case key.Matches(msg, keys.Delete):
		if m.dirsCursor < len(m.dirs) {
			name := m.dirs[m.dirsCursor].Name
			if err := utils.RemoveSavedDirectory(name); err != nil {
				m.setStatus("error: "+err.Error(), true)
			} else {
				m.setStatus(fmt.Sprintf("removed '%s'.", name), false)
				m.reloadDirs()
				m.clampDirsCursor()
			}
		}

	case key.Matches(msg, keys.Quit):
		return m, tea.Quit
	}
	return m, nil
}

// ── run shortcut ─────────────────────────────────────────────────────────────

func (m Model) runSelectedShortcut() (tea.Model, tea.Cmd) {
	sel := m.selectedKey()
	if sel == "" {
		return m, nil
	}
	s := m.shortcuts[sel]
	command := s.Command
	placeholders := extractPlaceholders(command)
	if len(placeholders) > 0 {
		m.templateCommand = command
		m.templateKeys = placeholders
		m.templateInputs = newTemplateInputs(placeholders)
		m.templateFocus = 0
		m.templateTarget = sel
		m.templateExtraArgs = nil
		m.page = pageTemplate
		return m, textinput.Blink
	}
	// No templates — quit TUI immediately and let Start() run the command.
	m.pendingCmd = command
	m.pendingCmdName = sel
	return m, tea.Quit
}

// ── helpers ──────────────────────────────────────────────────────────────────

// updateInputs propagates non-key tea messages to whichever inputs are active.
func (m Model) updateInputs(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	if m.searchActive {
		var c tea.Cmd
		m.searchInput, c = m.searchInput.Update(msg)
		cmds = append(cmds, c)
	}
	for i := range m.formInputs {
		var c tea.Cmd
		m.formInputs[i], c = m.formInputs[i].Update(msg)
		cmds = append(cmds, c)
	}
	for i := range m.templateInputs {
		var c tea.Cmd
		m.templateInputs[i], c = m.templateInputs[i].Update(msg)
		cmds = append(cmds, c)
	}
	return m, tea.Batch(cmds...)
}

// ignoreKeyFor returns a command that re-enables key processing after d.
func ignoreKeyFor(d time.Duration) tea.Cmd {
	return tea.Tick(d, func(time.Time) tea.Msg { return enableKeysMsg{} })
}
