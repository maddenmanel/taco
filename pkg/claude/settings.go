package claude

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/maddenmanel/taco/pkg/config"
)

const backupFileName = ".settings.taco-backup.json"

// SettingsPath returns the Claude Code settings.json path.
func SettingsPath() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".claude", "settings.json")
}

// BackupPath returns the path for the backup file.
func BackupPath() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".claude", backupFileName)
}

// readSettings reads and parses the Claude settings.json.
func readSettings() (map[string]interface{}, error) {
	data, err := os.ReadFile(SettingsPath())
	if err != nil {
		if os.IsNotExist(err) {
			return make(map[string]interface{}), nil
		}
		return nil, fmt.Errorf("failed to read settings: %w", err)
	}

	var settings map[string]interface{}
	if err := json.Unmarshal(data, &settings); err != nil {
		return nil, fmt.Errorf("failed to parse settings: %w", err)
	}

	return settings, nil
}

// writeSettings writes the settings map back to settings.json.
func writeSettings(settings map[string]interface{}) error {
	dir := filepath.Dir(SettingsPath())
	if err := os.MkdirAll(dir, 0700); err != nil {
		return fmt.Errorf("failed to create .claude dir: %w", err)
	}

	data, err := json.MarshalIndent(settings, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal settings: %w", err)
	}

	return os.WriteFile(SettingsPath(), data, 0644)
}

// Backup creates a backup of the current settings.json before modification.
func Backup() error {
	data, err := os.ReadFile(SettingsPath())
	if err != nil {
		if os.IsNotExist(err) {
			return nil // nothing to back up
		}
		return err
	}
	return os.WriteFile(BackupPath(), data, 0644)
}

// InjectProvider modifies ~/.claude/settings.json to route through the given provider.
// It only touches the "env" field, preserving all other user settings.
func InjectProvider(p config.Provider) error {
	if err := Backup(); err != nil {
		return fmt.Errorf("failed to backup settings: %w", err)
	}

	settings, err := readSettings()
	if err != nil {
		return err
	}

	envObj, ok := settings["env"].(map[string]interface{})
	if !ok {
		envObj = make(map[string]interface{})
	}

	envObj["ANTHROPIC_BASE_URL"] = p.BaseURL
	envObj["ANTHROPIC_AUTH_TOKEN"] = p.APIKey
	envObj["CLAUDE_CODE_DISABLE_NONESSENTIAL_TRAFFIC"] = "1"
	envObj["API_TIMEOUT_MS"] = "600000"

	if p.OpusModel != "" {
		envObj["ANTHROPIC_DEFAULT_OPUS_MODEL"] = p.OpusModel
	}
	if p.SonnetModel != "" {
		envObj["ANTHROPIC_DEFAULT_SONNET_MODEL"] = p.SonnetModel
	}
	if p.HaikuModel != "" {
		envObj["ANTHROPIC_DEFAULT_HAIKU_MODEL"] = p.HaikuModel
	}

	settings["env"] = envObj

	return writeSettings(settings)
}

// Restore removes all TACO-injected env vars from settings.json,
// returning Claude Code to its original (official) configuration.
func Restore() error {
	settings, err := readSettings()
	if err != nil {
		return err
	}

	envObj, ok := settings["env"].(map[string]interface{})
	if !ok {
		return nil // nothing to restore
	}

	tacoKeys := []string{
		"ANTHROPIC_BASE_URL",
		"ANTHROPIC_AUTH_TOKEN",
		"ANTHROPIC_DEFAULT_OPUS_MODEL",
		"ANTHROPIC_DEFAULT_SONNET_MODEL",
		"ANTHROPIC_DEFAULT_HAIKU_MODEL",
		"CLAUDE_CODE_DISABLE_NONESSENTIAL_TRAFFIC",
		"API_TIMEOUT_MS",
	}

	for _, key := range tacoKeys {
		delete(envObj, key)
	}

	if len(envObj) == 0 {
		delete(settings, "env")
	} else {
		settings["env"] = envObj
	}

	return writeSettings(settings)
}

// GetCurrentEnv returns the currently injected env vars, if any.
func GetCurrentEnv() (map[string]interface{}, error) {
	settings, err := readSettings()
	if err != nil {
		return nil, err
	}

	envObj, ok := settings["env"].(map[string]interface{})
	if !ok {
		return nil, nil
	}

	return envObj, nil
}
