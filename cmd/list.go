package cmd

import (
	"fmt"
	"sort"

	"github.com/maddenmanel/taco/pkg/config"
	"github.com/maddenmanel/taco/pkg/provider"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "List all configured and available providers",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.Load()
		if err != nil {
			return err
		}

		// Show configured providers
		if len(cfg.Providers) > 0 {
			fmt.Println("Configured providers:")
			fmt.Println()

			names := make([]string, 0, len(cfg.Providers))
			for name := range cfg.Providers {
				names = append(names, name)
			}
			sort.Strings(names)

			for _, name := range names {
				p := cfg.Providers[name]
				active := "  "
				if name == cfg.ActiveProvider {
					active = "* "
				}
				fmt.Printf("  %s%-15s %s\n", active, name, p.BaseURL)
			}
			fmt.Println()
		} else {
			fmt.Println("No providers configured yet.")
			fmt.Println()
		}

		// Show available presets
		fmt.Println("Built-in presets (use with `taco add <name> --key=...`):")
		fmt.Println()

		presets := provider.ListPresets()
		sort.Strings(presets)
		for _, name := range presets {
			p := provider.BuiltinPresets[name]
			_, configured := cfg.Providers[name]
			status := ""
			if configured {
				status = " (configured)"
			}
			fmt.Printf("  %-15s %s%s\n", name, p.DisplayName, status)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
