package cmd

import (
	"fmt"

	"github.com/maddenmanel/taco/pkg/claude"
	"github.com/maddenmanel/taco/pkg/config"
	"github.com/spf13/cobra"
)

var currentCmd = &cobra.Command{
	Use:   "current",
	Short: "Show the currently active provider",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.Load()
		if err != nil {
			return err
		}

		if cfg.ActiveProvider == "" {
			fmt.Println("🌮 TACO: Using official Claude (Anthropic) — no provider override active.")
			return nil
		}

		p, ok := cfg.GetProvider(cfg.ActiveProvider)
		if !ok {
			fmt.Printf("🌮 TACO: Active provider %q not found in config.\n", cfg.ActiveProvider)
			return nil
		}

		fmt.Printf("🌮 TACO: Currently using %s\n", p.DisplayName)
		fmt.Printf("   Base URL:  %s\n", p.BaseURL)
		if p.SonnetModel != "" {
			fmt.Printf("   Sonnet ->  %s\n", p.SonnetModel)
		}
		if p.OpusModel != "" {
			fmt.Printf("   Opus   ->  %s\n", p.OpusModel)
		}

		// Also show what's actually in settings.json
		env, err := claude.GetCurrentEnv()
		if err == nil && env != nil {
			if url, ok := env["ANTHROPIC_BASE_URL"]; ok {
				fmt.Printf("\n   Verified in settings.json: %v\n", url)
			}
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(currentCmd)
}
