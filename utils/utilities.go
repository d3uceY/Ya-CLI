package utils

import (
    "os"
   "encoding/json"
)


func LoadShortcuts() (map[string]string, error) {
	data, err := os.ReadFile("shortcuts.json")
	if err != nil {
		return nil, err
	}

	var shortcuts map[string]string
	err = json.Unmarshal(data, &shortcuts)
	if err != nil {
		return nil, err
	}

	return shortcuts, nil
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

	err = os.WriteFile("shortcuts.json", data, 0644)
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
