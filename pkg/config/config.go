package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// Provider represents a user-configured AI provider.
type Provider struct {
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
	BaseURL     string `json:"base_url"`
	APIKey      string `json:"api_key"`
	OpusModel   string `json:"opus_model"`
	SonnetModel string `json:"sonnet_model"`
	HaikuModel  string `json:"haiku_model"`
}

// TacoConfig holds all user-configured providers.
type TacoConfig struct {
	ActiveProvider string              `json:"active_provider"`
	Providers      map[string]Provider `json:"providers"`
}

// ConfigDir returns the taco config directory path.
func ConfigDir() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".taco")
}

// ConfigPath returns the taco config file path.
func ConfigPath() string {
	return filepath.Join(ConfigDir(), "config.json")
}

// Load reads the taco config from disk.
func Load() (*TacoConfig, error) {
	cfg := &TacoConfig{
		Providers: make(map[string]Provider),
	}

	data, err := os.ReadFile(ConfigPath())
	if err != nil {
		if os.IsNotExist(err) {
			return cfg, nil
		}
		return nil, fmt.Errorf("failed to read config: %w", err)
	}

	if err := json.Unmarshal(data, cfg); err != nil {
		return nil, fmt.Errorf("failed to parse config: %w", err)
	}

	if cfg.Providers == nil {
		cfg.Providers = make(map[string]Provider)
	}

	return cfg, nil
}

// Save writes the taco config to disk.
func (c *TacoConfig) Save() error {
	dir := ConfigDir()
	if err := os.MkdirAll(dir, 0700); err != nil {
		return fmt.Errorf("failed to create config dir: %w", err)
	}

	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal config: %w", err)
	}

	if err := os.WriteFile(ConfigPath(), data, 0600); err != nil {
		return fmt.Errorf("failed to write config: %w", err)
	}

	return nil
}

// AddProvider adds or updates a provider in the config.
func (c *TacoConfig) AddProvider(p Provider) {
	c.Providers[p.Name] = p
}

// GetProvider returns a provider by name.
func (c *TacoConfig) GetProvider(name string) (Provider, bool) {
	p, ok := c.Providers[name]
	return p, ok
}

// RemoveProvider removes a provider from the config.
func (c *TacoConfig) RemoveProvider(name string) {
	delete(c.Providers, name)
	if c.ActiveProvider == name {
		c.ActiveProvider = ""
	}
}
