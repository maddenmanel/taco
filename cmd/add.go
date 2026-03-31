package cmd

import (
	"fmt"

	"github.com/maddenmanel/taco/pkg/config"
	"github.com/maddenmanel/taco/pkg/provider"
	"github.com/spf13/cobra"
)

var (
	addKey     string
	addURL     string
	addOpus    string
	addSonnet  string
	addHaiku   string
)

var addCmd = &cobra.Command{
	Use:   "add <provider-name>",
	Short: "Add or update a provider configuration",
	Long: `Add a new AI provider or update an existing one.

If the provider name matches a built-in preset (deepseek, openrouter,
siliconflow, zhipu, volcengine), the base URL and model mappings are
auto-filled. You only need to supply your API key.

For custom providers, use --url, --opus, --sonnet, --haiku flags.`,
	Example: `  taco add deepseek --key="sk-your-deepseek-key"
  taco add openrouter --key="sk-or-your-key"
  taco add my-proxy --key="sk-xxx" --url="https://my-proxy.com/v1"`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]

		if addKey == "" {
			return fmt.Errorf("--key is required")
		}

		cfg, err := config.Load()
		if err != nil {
			return err
		}

		p := config.Provider{
			Name:   name,
			APIKey: addKey,
		}

		// Check if there's a built-in preset
		if preset, ok := provider.GetPreset(name); ok {
			p.DisplayName = preset.DisplayName
			p.BaseURL = preset.BaseURL
			p.OpusModel = preset.OpusModel
			p.SonnetModel = preset.SonnetModel
			p.HaikuModel = preset.HaikuModel
		}

		// Override with user-provided flags
		if addURL != "" {
			p.BaseURL = addURL
		}
		if addOpus != "" {
			p.OpusModel = addOpus
		}
		if addSonnet != "" {
			p.SonnetModel = addSonnet
		}
		if addHaiku != "" {
			p.HaikuModel = addHaiku
		}

		if p.BaseURL == "" {
			return fmt.Errorf("--url is required for custom providers (no built-in preset for %q)", name)
		}

		if p.DisplayName == "" {
			p.DisplayName = name
		}

		cfg.AddProvider(p)
		if err := cfg.Save(); err != nil {
			return err
		}

		fmt.Printf("🌮 Provider %q added successfully.\n", p.DisplayName)
		fmt.Printf("   Base URL: %s\n", p.BaseURL)
		fmt.Printf("   Run: taco use %s\n", name)
		return nil
	},
}

func init() {
	addCmd.Flags().StringVar(&addKey, "key", "", "API key for the provider (required)")
	addCmd.Flags().StringVar(&addURL, "url", "", "Base URL (auto-filled for built-in presets)")
	addCmd.Flags().StringVar(&addOpus, "opus", "", "Model name for opus tier")
	addCmd.Flags().StringVar(&addSonnet, "sonnet", "", "Model name for sonnet tier")
	addCmd.Flags().StringVar(&addHaiku, "haiku", "", "Model name for haiku tier")
	rootCmd.AddCommand(addCmd)
}
