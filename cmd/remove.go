package cmd

import (
	"fmt"

	"github.com/maddenmanel/taco/pkg/claude"
	"github.com/maddenmanel/taco/pkg/config"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:     "remove <provider-name>",
	Aliases: []string{"rm"},
	Short:   "Remove a configured provider",
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]

		cfg, err := config.Load()
		if err != nil {
			return err
		}

		if _, ok := cfg.GetProvider(name); !ok {
			return fmt.Errorf("provider %q not found", name)
		}

		wasActive := cfg.ActiveProvider == name
		cfg.RemoveProvider(name)

		if wasActive {
			if err := claude.Restore(); err != nil {
				return fmt.Errorf("failed to restore settings: %w", err)
			}
		}

		if err := cfg.Save(); err != nil {
			return err
		}

		fmt.Printf("🌮 Provider %q removed.\n", name)
		if wasActive {
			fmt.Println("   Claude Code restored to official Anthropic configuration.")
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
