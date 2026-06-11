package utils

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func configFilePath() (string, error) {
	appDir, err := getAppDataDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(appDir, "config.json"), nil
}

// GetConfig reads config.json, bootstrapping it if absent.
func GetConfig() (AppConfig, error) {
	path, err := configFilePath()
	if err != nil {
		return AppConfig{}, err
	}
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			cfg := AppConfig{}
			_ = saveConfig(cfg)
			return cfg, nil
		}
		return AppConfig{}, err
	}
	var cfg AppConfig
	if json.Unmarshal(data, &cfg) == nil {
		return cfg, nil
	}
	return AppConfig{}, nil
}

func saveConfig(cfg AppConfig) error {
	path, err := configFilePath()
	if err != nil {
		return err
	}
	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}

// AddSavedDirectory adds or updates a named directory preset.
func AddSavedDirectory(name, path string) error {
	cfg, err := GetConfig()
	if err != nil {
		return err
	}
	for i, d := range cfg.SavedDirectories {
		if d.Name == name {
			cfg.SavedDirectories[i].Path = path
			return saveConfig(cfg)
		}
	}
	cfg.SavedDirectories = append(cfg.SavedDirectories, SavedDir{Name: name, Path: path})
	return saveConfig(cfg)
}

// RemoveSavedDirectory removes a named directory preset.
func RemoveSavedDirectory(name string) error {
	cfg, err := GetConfig()
	if err != nil {
		return err
	}
	filtered := cfg.SavedDirectories[:0]
	for _, d := range cfg.SavedDirectories {
		if d.Name != name {
			filtered = append(filtered, d)
		}
	}
	cfg.SavedDirectories = filtered
	return saveConfig(cfg)
}
