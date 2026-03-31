package cmd

import (
	"fmt"

	"github.com/maddenmanel/taco/pkg/claude"
	"github.com/maddenmanel/taco/pkg/config"
	"github.com/spf13/cobra"
)

var useCmd = &cobra.Command{
	Use:   "use <provider-name>",
	Short: "Switch Claude Code to use the specified provider",
	Long: `Inject the specified provider's configuration into Claude Code's
settings.json. After switching, just type "claude" as usual — requests
will be routed through the selected provider transparently.

A backup of your original settings is saved automatically.`,
	Example: `  taco use deepseek
  taco use openrouter
  taco use my-proxy`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]

		cfg, err := config.Load()
		if err != nil {
			return err
		}

		p, ok := cfg.GetProvider(name)
		if !ok {
			return fmt.Errorf("provider %q not found. Run: taco add %s --key=YOUR_KEY", name, name)
		}

		if err := claude.InjectProvider(p); err != nil {
			return fmt.Errorf("failed to switch provider: %w", err)
		}

		cfg.ActiveProvider = name
		if err := cfg.Save(); err != nil {
			return err
		}

		fmt.Printf("🌮 TACO: Successfully switched to %s.\n", p.DisplayName)
		fmt.Printf("   Base URL:  %s\n", p.BaseURL)
		if p.SonnetModel != "" {
			fmt.Printf("   Sonnet ->  %s\n", p.SonnetModel)
		}
		if p.OpusModel != "" {
			fmt.Printf("   Opus   ->  %s\n", p.OpusModel)
		}
		fmt.Println()
		fmt.Println("💡 Just type `claude` to start coding!")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(useCmd)
}
