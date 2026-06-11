package utils

// ShortcutData stores all metadata for a single shortcut.
// This is the canonical type shared between ya-cli and ya-gui.
type ShortcutData struct {
	Command     string   `json:"command"`
	Description string   `json:"description,omitempty"`
	Tags        []string `json:"tags,omitempty"`
	Pinned      bool     `json:"pinned,omitempty"`
	RunCount    int      `json:"runCount,omitempty"`
	LastRun     string   `json:"lastRun,omitempty"`
}

// AppConfig holds application-level settings.
type AppConfig struct {
	DefaultDir       string     `json:"defaultDir,omitempty"`
	SavedDirectories []SavedDir `json:"savedDirectories,omitempty"`
}

// SavedDir is a named workspace directory preset.
type SavedDir struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

// RunHistoryEntry records one execution of a shortcut.
type RunHistoryEntry struct {
	ShortcutName string `json:"shortcutName"`
	Command      string `json:"command"`
	Directory    string `json:"directory"`
	Timestamp    string `json:"timestamp"`
}
