package cmd

import (
	"fmt"

	"github.com/maddenmanel/taco/pkg/claude"
	"github.com/maddenmanel/taco/pkg/config"
	"github.com/spf13/cobra"
)

var restoreCmd = &cobra.Command{
	Use:   "restore",
	Short: "Restore Claude Code to official Anthropic configuration",
	Long: `Remove all TACO-injected environment variables from Claude Code's
settings.json, restoring it to use the official Anthropic API.

Your other settings (theme, shortcuts, etc.) are preserved.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := claude.Restore(); err != nil {
			return fmt.Errorf("failed to restore: %w", err)
		}

		cfg, err := config.Load()
		if err != nil {
			return err
		}

		cfg.ActiveProvider = ""
		if err := cfg.Save(); err != nil {
			return err
		}

		fmt.Println("🌮 TACO: Restored to official Claude (Anthropic) configuration.")
		fmt.Println("   All injected env vars have been removed.")
		fmt.Println("   Your other Claude settings are untouched.")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(restoreCmd)
}
