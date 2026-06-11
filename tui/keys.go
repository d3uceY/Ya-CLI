package tui

import "github.com/charmbracelet/bubbles/key"

// keyMap holds all the key bindings used across the TUI.
type keyMap struct {
	Up           key.Binding
	Down         key.Binding
	Enter        key.Binding
	Search       key.Binding
	Add          key.Binding
	Edit         key.Binding
	Rename       key.Binding
	Delete       key.Binding
	Pin          key.Binding
	Import       key.Binding
	Export       key.Binding
	History      key.Binding
	Dirs         key.Binding
	ClearHistory key.Binding
	Help         key.Binding
	FullScreen   key.Binding
	Quit         key.Binding
	Back         key.Binding
	Tab          key.Binding
	Yes          key.Binding
	No           key.Binding
}

var keys = keyMap{
	Up: key.NewBinding(
		key.WithKeys("up", "k"),
		key.WithHelp("↑/k", "up"),
	),
	Down: key.NewBinding(
		key.WithKeys("down", "j"),
		key.WithHelp("↓/j", "down"),
	),
	Enter: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "run"),
	),
	Search: key.NewBinding(
		key.WithKeys("/"),
		key.WithHelp("/", "search"),
	),
	Add: key.NewBinding(
		key.WithKeys("a"),
		key.WithHelp("a", "add"),
	),
	Edit: key.NewBinding(
		key.WithKeys("e"),
		key.WithHelp("e", "edit"),
	),
	Rename: key.NewBinding(
		key.WithKeys("r"),
		key.WithHelp("r", "rename"),
	),
	Delete: key.NewBinding(
		key.WithKeys("d", "x"),
		key.WithHelp("d", "delete"),
	),
	Pin: key.NewBinding(
		key.WithKeys("p"),
		key.WithHelp("p", "pin"),
	),
	Import: key.NewBinding(
		key.WithKeys("i"),
		key.WithHelp("i", "import"),
	),
	Export: key.NewBinding(
		key.WithKeys("o"),
		key.WithHelp("o", "export"),
	),
	History: key.NewBinding(
		key.WithKeys("h"),
		key.WithHelp("h", "history"),
	),
	Dirs: key.NewBinding(
		key.WithKeys("D"),
		key.WithHelp("D", "saved dirs"),
	),
	ClearHistory: key.NewBinding(
		key.WithKeys("c"),
		key.WithHelp("c", "clear history"),
	),
	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "help"),
	),
	FullScreen: key.NewBinding(
		key.WithKeys("f"),
		key.WithHelp("f", "fullscreen"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q", "ctrl+c"),
		key.WithHelp("q", "quit"),
	),
	Back: key.NewBinding(
		key.WithKeys("esc"),
		key.WithHelp("esc", "back"),
	),
	Tab: key.NewBinding(
		key.WithKeys("tab"),
		key.WithHelp("tab", "next field"),
	),
	Yes: key.NewBinding(
		key.WithKeys("y"),
		key.WithHelp("y", "yes"),
	),
	No: key.NewBinding(
		key.WithKeys("n"),
		key.WithHelp("n", "no"),
	),
}

