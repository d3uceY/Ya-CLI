package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// shortcutMeta is the GUI-compatible metadata file structure stored in
// shortcuts-meta.json alongside the CLI-compatible shortcuts.json.
type shortcutMeta struct {
	Description string   `json:"description,omitempty"`
	Tags        []string `json:"tags,omitempty"`
	Pinned      bool     `json:"pinned,omitempty"`
	RunCount    int      `json:"runCount,omitempty"`
	LastRun     string   `json:"lastRun,omitempty"`
}

// ── file paths ────────────────────────────────────────────────────────────────

func shortcutFilePath() (string, error) {
	appDir, err := getAppDataDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(appDir, "shortcuts.json"), nil
}

func metaFilePath() (string, error) {
	appDir, err := getAppDataDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(appDir, "shortcuts-meta.json"), nil
}

// ── commands (CLI-compatible map[string]string) ───────────────────────────────

func loadCommands() (map[string]string, error) {
	path, err := shortcutFilePath()
	if err != nil {
		return nil, err
	}
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			empty := map[string]string{}
			_ = saveCommands(empty)
			return empty, nil
		}
		return nil, err
	}

	data = bytes.TrimPrefix(data, []byte{0xEF, 0xBB, 0xBF})

	var cmds map[string]string
	if json.Unmarshal(data, &cmds) == nil {
		return cmds, nil
	}

	// Recovery: old rich format
	var rich map[string]ShortcutData
	if err := json.Unmarshal(data, &rich); err != nil {
		return nil, fmt.Errorf("shortcuts.json is in an unrecognised format: %w", err)
	}
	cmds = make(map[string]string, len(rich))
	meta := make(map[string]shortcutMeta, len(rich))
	for name, s := range rich {
		cmds[name] = s.Command
		if s.Description != "" || len(s.Tags) > 0 || s.Pinned || s.RunCount > 0 {
			meta[name] = shortcutMeta{
				Description: s.Description, Tags: s.Tags,
				Pinned: s.Pinned, RunCount: s.RunCount, LastRun: s.LastRun,
			}
		}
	}
	_ = saveCommands(cmds)
	_ = saveMeta(meta)
	return cmds, nil
}

func saveCommands(cmds map[string]string) error {
	path, err := shortcutFilePath()
	if err != nil {
		return err
	}
	data, err := json.MarshalIndent(cmds, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}

// ── metadata ──────────────────────────────────────────────────────────────────

func loadMeta() (map[string]shortcutMeta, error) {
	path, err := metaFilePath()
	if err != nil {
		return nil, err
	}
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return map[string]shortcutMeta{}, nil
		}
		return nil, err
	}
	data = bytes.TrimPrefix(data, []byte{0xEF, 0xBB, 0xBF})
	var meta map[string]shortcutMeta
	if err := json.Unmarshal(data, &meta); err != nil {
		return map[string]shortcutMeta{}, nil
	}
	return meta, nil
}

func saveMeta(meta map[string]shortcutMeta) error {
	path, err := metaFilePath()
	if err != nil {
		return err
	}
	data, err := json.MarshalIndent(meta, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}

// ── combined load/save ────────────────────────────────────────────────────────

func loadAllShortcuts() (map[string]ShortcutData, error) {
	cmds, err := loadCommands()
	if err != nil {
		return nil, err
	}
	meta, err := loadMeta()
	if err != nil {
		return nil, err
	}
	result := make(map[string]ShortcutData, len(cmds))
	for name, cmd := range cmds {
		m := meta[name]
		result[name] = ShortcutData{
			Command: cmd, Description: m.Description, Tags: m.Tags,
			Pinned: m.Pinned, RunCount: m.RunCount, LastRun: m.LastRun,
		}
	}
	return result, nil
}

func saveAllShortcuts(shortcuts map[string]ShortcutData) error {
	cmds := make(map[string]string, len(shortcuts))
	meta := make(map[string]shortcutMeta, len(shortcuts))
	for name, s := range shortcuts {
		cmds[name] = s.Command
		if s.Description != "" || len(s.Tags) > 0 || s.Pinned || s.RunCount > 0 {
			meta[name] = shortcutMeta{
				Description: s.Description, Tags: s.Tags,
				Pinned: s.Pinned, RunCount: s.RunCount, LastRun: s.LastRun,
			}
		}
	}
	if err := saveCommands(cmds); err != nil {
		return err
	}
	return saveMeta(meta)
}

// ── public API ────────────────────────────────────────────────────────────────

// GetShortcuts returns all shortcuts merged with their metadata.
func GetShortcuts() (map[string]ShortcutData, error) {
	return loadAllShortcuts()
}

// LoadShortcuts returns shortcuts as a plain map (CLI-compatible).
func LoadShortcuts() (map[string]string, error) {
	return loadCommands()
}

// GetShortcut returns the command string for the given shortcut name.
func GetShortcut(name string) (string, error) {
	cmds, err := loadCommands()
	if err != nil {
		return "", err
	}
	cmd, ok := cmds[name]
	if !ok {
		return "", fmt.Errorf("shortcut `%s` not found", name)
	}
	return cmd, nil
}

// SearchShortcut returns shortcuts whose name or command contains the query.
func SearchShortcut(query string) (map[string]string, error) {
	cmds, err := loadCommands()
	if err != nil {
		return nil, err
	}
	q := strings.ToLower(query)
	out := map[string]string{}
	for k, v := range cmds {
		if strings.Contains(strings.ToLower(k), q) || strings.Contains(strings.ToLower(v), q) {
			out[k] = v
		}
	}
	return out, nil
}

// AddShortcut creates or replaces a shortcut (CLI-compatible, no description).
func AddShortcut(name, command string) error {
	return AddShortcutFull(name, command, "")
}

// AddShortcutFull creates or replaces a shortcut with an optional description.
// Existing RunCount and LastRun are preserved.
func AddShortcutFull(name, command, description string) error {
	shortcuts, err := loadAllShortcuts()
	if err != nil {
		return err
	}
	existing := shortcuts[name]
	shortcuts[name] = ShortcutData{
		Command:     command,
		Description: description,
		Tags:        existing.Tags,
		Pinned:      existing.Pinned,
		RunCount:    existing.RunCount,
		LastRun:     existing.LastRun,
	}
	return saveAllShortcuts(shortcuts)
}

// RemoveShortcut deletes a shortcut and its metadata.
func RemoveShortcut(name string) error {
	shortcuts, err := loadAllShortcuts()
	if err != nil {
		return err
	}
	delete(shortcuts, name)
	return saveAllShortcuts(shortcuts)
}

// RenameShortcut renames a shortcut, preserving all metadata.
func RenameShortcut(oldName, newName string) error {
	newName = strings.TrimSpace(newName)
	if newName == "" {
		return fmt.Errorf("shortcut name cannot be empty")
	}
	shortcuts, err := loadAllShortcuts()
	if err != nil {
		return err
	}
	src, ok := shortcuts[oldName]
	if !ok {
		return fmt.Errorf("shortcut `%s` not found", oldName)
	}
	if oldName != newName {
		if _, exists := shortcuts[newName]; exists {
			return fmt.Errorf("a shortcut named %q already exists", newName)
		}
		delete(shortcuts, oldName)
	}
	shortcuts[newName] = src
	return saveAllShortcuts(shortcuts)
}

// TogglePin flips the Pinned flag on a shortcut.
func TogglePin(name string) error {
	shortcuts, err := loadAllShortcuts()
	if err != nil {
		return err
	}
	s, ok := shortcuts[name]
	if !ok {
		return nil
	}
	s.Pinned = !s.Pinned
	shortcuts[name] = s
	return saveAllShortcuts(shortcuts)
}

// IncrementRunCount bumps RunCount and records LastRun for the named shortcut.
func IncrementRunCount(name string) error {
	shortcuts, err := loadAllShortcuts()
	if err != nil {
		return err
	}
	s, ok := shortcuts[name]
	if !ok {
		return nil
	}
	s.RunCount++
	s.LastRun = time.Now().UTC().Format(time.RFC3339)
	shortcuts[name] = s
	return saveAllShortcuts(shortcuts)
}

// ImportShortcuts merges shortcuts from a JSON file (CLI or GUI format).
func ImportShortcuts(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	data = bytes.TrimPrefix(data, []byte{0xEF, 0xBB, 0xBF})

	// Try CLI map format first.
	var imported map[string]string
	if json.Unmarshal(data, &imported) != nil {
		// Fall back to GUI rich format.
		var rich map[string]ShortcutData
		if err := json.Unmarshal(data, &rich); err != nil {
			return fmt.Errorf("unrecognised shortcuts file format: %w", err)
		}
		imported = make(map[string]string, len(rich))
		for n, s := range rich {
			imported[n] = s.Command
		}
	}

	current, err := loadAllShortcuts()
	if err != nil {
		return err
	}
	for name, cmd := range imported {
		existing := current[name]
		existing.Command = cmd
		current[name] = existing
	}
	return saveAllShortcuts(current)
}

// ExportShortcuts writes the CLI-compatible shortcuts.json to the given path.
func ExportShortcuts(filePath string) error {
	cmds, err := loadCommands()
	if err != nil {
		return err
	}
	data, err := json.MarshalIndent(cmds, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filePath, data, 0644)
}
