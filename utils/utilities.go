package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
)

// i go need explain myself? 😭
func GetAppVersion() string {
	return "v0.3.2"
}

// i think you should know what this one does
func getAppDataDir() (string, error) {
	dir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	appDir := filepath.Join(dir, "ya/data")
	err = os.MkdirAll(appDir, 0755)

	return appDir, err
}

// this shpould load the shortcuts from the JSON file and return them as a map.
func LoadShortcuts() (map[string]string, error) {

	appDir, err := getAppDataDir()
	if err != nil {
		panic(err)
	}

	shortCutpath := filepath.Join(appDir, "shortcuts.json")

	data, err := os.ReadFile(shortCutpath)
	if err != nil {
		if os.IsNotExist(err) {
			// Create file with empty JSON object
			emptyShortcuts := map[string]string{}
			data, err := json.MarshalIndent(emptyShortcuts, "", "  ")
			if err != nil {
				return nil, err
			}
			err = os.WriteFile(shortCutpath, data, 0644)
			if err != nil {
				return nil, err
			}
			return emptyShortcuts, nil
		}
		return nil, err
	}

	var shortcuts map[string]string
	err = json.Unmarshal(data, &shortcuts)
	if err != nil {
		return nil, err
	}

	return shortcuts, nil
}

// this function retrieves the command associated with a given shortcut name.
func GetShortcut(shortcut string) (string, error) {
	shortcuts, err := LoadShortcuts()

	if err != nil {
		return "", err
	}

	command, exists := shortcuts[shortcut]

	if !exists {
		return "", fmt.Errorf("shortcut `%s` not found", shortcut)
	}
	return command, nil
}

// this function searches for the shortcut
func SearchShortcut(searchParam string) (map[string]string, error) {
	shortcuts, err := LoadShortcuts()
	if err != nil {
		return nil, err
	}

	filteredShortcuts := map[string]string{}
	search := strings.ToLower(searchParam)

	for key, command := range shortcuts {
		if strings.Contains(strings.ToLower(key), search) ||
			strings.Contains(strings.ToLower(command), search) {
			filteredShortcuts[key] = command
		}
	}
	return filteredShortcuts, nil
}

func AddShortcut(name, command string) error {
	shortcuts, err := LoadShortcuts()

	if err != nil {
		return err
	}

	shortcuts[name] = command

	data, err := json.MarshalIndent(shortcuts, "", "  ")
	if err != nil {
		return err
	}

	appDir, err := getAppDataDir()
	if err != nil {
		return err
	}

	shortCutpath := filepath.Join(appDir, "shortcuts.json")
	err = os.WriteFile(shortCutpath, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func RemoveShortcut(name string) error {
	shortcuts, err := LoadShortcuts()
	if err != nil {
		return err
	}
	delete(shortcuts, name)

	data, err := json.MarshalIndent(shortcuts, "", "  ")
	if err != nil {
		return err
	}

	appDir, err := getAppDataDir()
	if err != nil {
		return err
	}

	shortCutpath := filepath.Join(appDir, "shortcuts.json")
	err = os.WriteFile(shortCutpath, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func IsInvalidString(s string) bool {
	if len(s) == 0 {
		return true
	}
	return false
}

func ImportShortcuts(path string) error {

	data, err := os.ReadFile(path)

	if err != nil {
		color.Red(err.Error())
		return err
	}

	appDir, err := getAppDataDir()
	if err != nil {
		color.Red(err.Error())
		return err
	}

	shortCutpath := filepath.Join(appDir, "shortcuts.json")

	currentShortcutData, err := os.ReadFile(shortCutpath)

	// If file does not exist, create it with the imported json, umm, hopefully it's valid lmao
	if err != nil {
		if os.IsNotExist(err) {
			err := os.WriteFile(shortCutpath, data, 0644)
			if err != nil {
				color.Red(err.Error())
				return err
			}
		}
	}

	// for some reason, i forgot maps only store unique keys 😂
	var currentShortcuts map[string]string
	err = json.Unmarshal(currentShortcutData, &currentShortcuts)

	if err != nil {
		return err
	}

	// store imported shortcuts in map
	var importedShortcuts map[string]string
	err = json.Unmarshal(data, &importedShortcuts)

	if err != nil {
		return err
	}

	// merge imported shortcuts into current shortcuts
	for key, value := range importedShortcuts {
		currentShortcuts[key] = value
	}

	// write merged shortcuts back to file
	data, err = json.MarshalIndent(currentShortcuts, "", "  ")
	if err != nil {
		return err
	}

	// this permission param lowkey threw me off ngl
	return os.WriteFile(shortCutpath, data, 0644)
}
