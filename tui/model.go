package tui

import (
	"os/exec"
	"runtime"
	"sort"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/d3uceY/Ya-CLI/utils"
)

// page represents which view is currently displayed.
type page int

const (
	pageList page = iota
	pageAdd
	pageEdit
	pageRename
	pageConfirmDelete
	pageTemplate
	pageHelp
	pageImport
	pageExport
	pageHistory
	pageDirs
	pageAddDir
	pageConfirmClearHistory
)

// formMode distinguishes the purpose of the shared form view.
type formMode int

const (
	formAdd formMode = iota
	formEdit
	formRename
	formImport
	formExport
	formAddDir
)

// Model is the root TUI model.
type Model struct {
	// terminal dimensions
	width  int
	height int

	// current view
	page page

	// shortcut data — full metadata-rich map
	shortcuts  map[string]utils.ShortcutData
	sortedKeys []string // alphabetically sorted names
	filtered   []string // filtered by active search query

	// list navigation
	cursor int
	offset int // first visible row index (scroll)

	// search bar
	searchInput  textinput.Model
	searchActive bool

	// shared form (add / edit / rename / import / export / add-dir)
	formInputs []textinput.Model
	formFocus  int
	formMode   formMode

	// template fill-in form
	templateInputs    []textinput.Model
	templateFocus     int
	templateKeys      []string // placeholder names in order
	templateCommand   string   // original command with {tokens}
	templateTarget    string   // shortcut name (display only)
	templateExtraArgs []string // extra CLI args to forward

	// confirm-delete state
	confirmName    string
	confirmCommand string

	// run history
	history      []utils.RunHistoryEntry
	historyCursor int
	historyOffset int

	// saved directories
	dirs       []utils.SavedDir
	dirsCursor int

	// status bar
	statusMsg     string
	statusIsError bool

	// UI flags
	fullscreen bool
	ignoreKeys bool // true briefly after a subprocess exits to drain buffered input

	// pending command — set before tea.Quit so Start() can exec it
	pendingCmd     string
	pendingCmdName string
}

// newModel creates a fresh Model and loads shortcuts from disk.
func newModel() Model {
	si := textinput.New()
	si.Placeholder = "type to filter…"
	si.PromptStyle = sPrimary.Bold(true)
	si.TextStyle = sText
	si.PlaceholderStyle = sMuted
	si.Prompt = ""
	si.CharLimit = 100

	m := Model{
		searchInput: si,
		width:       80,
		height:      24,
	}

	shortcuts, err := utils.GetShortcuts()
	if err != nil {
		m.shortcuts = map[string]utils.ShortcutData{}
	} else {
		m.shortcuts = shortcuts
	}
	m.rebuildSorted()
	m.filtered = append([]string(nil), m.sortedKeys...)
	return m
}

// ── data helpers ─────────────────────────────────────────────────────────────

func (m *Model) rebuildSorted() {
	m.sortedKeys = make([]string, 0, len(m.shortcuts))
	for k := range m.shortcuts {
		m.sortedKeys = append(m.sortedKeys, k)
	}
	// pinned items first, then alphabetical within each group
	sort.Slice(m.sortedKeys, func(i, j int) bool {
		pi := m.shortcuts[m.sortedKeys[i]].Pinned
		pj := m.shortcuts[m.sortedKeys[j]].Pinned
		if pi != pj {
			return pi
		}
		return m.sortedKeys[i] < m.sortedKeys[j]
	})
	m.applyFilter()
}

func (m *Model) applyFilter() {
	q := strings.ToLower(m.searchInput.Value())
	if q == "" {
		out := make([]string, len(m.sortedKeys))
		copy(out, m.sortedKeys)
		m.filtered = out
		return
	}
	out := make([]string, 0, len(m.sortedKeys))
	for _, k := range m.sortedKeys {
		s := m.shortcuts[k]
		if strings.Contains(strings.ToLower(k), q) ||
			strings.Contains(strings.ToLower(s.Command), q) ||
			strings.Contains(strings.ToLower(s.Description), q) {
			out = append(out, k)
		}
	}
	m.filtered = out
}

func (m *Model) reloadShortcuts() {
	shortcuts, err := utils.GetShortcuts()
	if err != nil {
		m.setStatus("error loading shortcuts: "+err.Error(), true)
		return
	}
	m.shortcuts = shortcuts
	m.rebuildSorted()
	m.clampCursor()
}

func (m *Model) reloadHistory() {
	h, err := utils.GetRunHistory()
	if err != nil {
		m.history = nil
		return
	}
	m.history = h
}

func (m *Model) reloadDirs() {
	cfg, err := utils.GetConfig()
	if err != nil {
		m.dirs = nil
		return
	}
	m.dirs = cfg.SavedDirectories
}

// ── cursor / scroll ──────────────────────────────────────────────────────────

func (m *Model) clampCursor() {
	n := len(m.filtered)
	if n == 0 {
		m.cursor = 0
		m.offset = 0
		return
	}
	if m.cursor < 0 {
		m.cursor = n - 1
	}
	if m.cursor >= n {
		m.cursor = 0
	}
	vis := m.visibleRows()
	if m.cursor < m.offset {
		m.offset = m.cursor
	}
	if m.cursor >= m.offset+vis {
		m.offset = m.cursor - vis + 1
	}
	if m.offset < 0 {
		m.offset = 0
	}
}

func (m *Model) clampHistoryCursor() {
	n := len(m.history)
	if n == 0 {
		m.historyCursor = 0
		m.historyOffset = 0
		return
	}
	if m.historyCursor < 0 {
		m.historyCursor = n - 1
	}
	if m.historyCursor >= n {
		m.historyCursor = 0
	}
	vis := m.visibleRows()
	if m.historyCursor < m.historyOffset {
		m.historyOffset = m.historyCursor
	}
	if m.historyCursor >= m.historyOffset+vis {
		m.historyOffset = m.historyCursor - vis + 1
	}
}

func (m *Model) clampDirsCursor() {
	n := len(m.dirs)
	if n == 0 {
		m.dirsCursor = 0
		return
	}
	if m.dirsCursor < 0 {
		m.dirsCursor = n - 1
	}
	if m.dirsCursor >= n {
		m.dirsCursor = 0
	}
}

// visibleRows returns how many list rows can fit on screen.
func (m *Model) visibleRows() int {
	v := m.height - 8
	if v < 1 {
		v = 1
	}
	return v
}

// selectedKey returns the shortcut name at the cursor, or "".
func (m *Model) selectedKey() string {
	if m.cursor < 0 || m.cursor >= len(m.filtered) {
		return ""
	}
	return m.filtered[m.cursor]
}

// ── status bar ───────────────────────────────────────────────────────────────

func (m *Model) setStatus(msg string, isError bool) {
	m.statusMsg = msg
	m.statusIsError = isError
}

func (m *Model) clearStatus() {
	m.statusMsg = ""
	m.statusIsError = false
}

// ── exec helpers ─────────────────────────────────────────────────────────────

// buildExecCmd wraps a shell command for the current OS.
func buildExecCmd(command string) *exec.Cmd {
	if runtime.GOOS == "windows" {
		return exec.Command("powershell", "-NoProfile", "-NonInteractive", "-Command", command)
	}
	return exec.Command("bash", "-c", command)
}

// resolveTemplateCommand substitutes {keys} with values from inputs.
func resolveTemplateCommand(command string, keys []string, inputs []textinput.Model) string {
	result := command
	for i, k := range keys {
		if i < len(inputs) {
			result = strings.ReplaceAll(result, "{"+k+"}", inputs[i].Value())
		}
	}
	return result
}

// extractPlaceholders returns unique, ordered placeholder names from a command.
func extractPlaceholders(command string) []string {
	matches := templateRe.FindAllStringSubmatch(command, -1)
	seen := map[string]bool{}
	var out []string
	for _, m := range matches {
		if !seen[m[1]] {
			seen[m[1]] = true
			out = append(out, m[1])
		}
	}
	return out
}

// ── input constructors ───────────────────────────────────────────────────────

func newFormInput(placeholder, value string, focused bool) textinput.Model {
	ti := textinput.New()
	ti.Placeholder = placeholder
	ti.SetValue(value)
	ti.PromptStyle = sAccent.Bold(true)
	ti.TextStyle = sText
	ti.PlaceholderStyle = sMuted
	ti.Prompt = "  "
	ti.CharLimit = 300
	if focused {
		ti.Focus()
	}
	return ti
}

func newTemplateInputs(placeholders []string) []textinput.Model {
	inputs := make([]textinput.Model, len(placeholders))
	for i, p := range placeholders {
		ti := textinput.New()
		ti.Placeholder = p
		ti.PromptStyle = sAccent.Bold(true)
		ti.TextStyle = sText
		ti.PlaceholderStyle = sMuted
		ti.Prompt = "  › "
		ti.CharLimit = 200
		if i == 0 {
			ti.Focus()
		}
		inputs[i] = ti
	}
	return inputs
}

// moveCursorTo positions the cursor on the named shortcut in the filtered list.
func (m *Model) moveCursorTo(name string) {
	for i, k := range m.filtered {
		if k == name {
			m.cursor = i
			m.clampCursor()
			return
		}
	}
}

// ── Init ─────────────────────────────────────────────────────────────────────

func (m Model) Init() tea.Cmd {
	return textinput.Blink
}

