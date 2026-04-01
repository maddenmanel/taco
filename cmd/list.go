package cmd

import (
	"fmt"
	"sort"
	"strings"

	"github.com/maddenmanel/taco/pkg/config"
	"github.com/maddenmanel/taco/pkg/provider"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "List configured providers and available presets",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.Load()
		if err != nil {
			return err
		}

		// ── Configured providers ──────────────────────────────────────────────
		if len(cfg.Providers) > 0 {
			names := sortedKeys(cfg.Providers)
			fmt.Println("Configured providers:")
			fmt.Println()
			printDivider(52)
			fmt.Printf("  %-2s %-14s %-18s %s\n", "", "NAME", "PROVIDER", "BASE URL")
			printDivider(52)
			for _, name := range names {
				p := cfg.Providers[name]
				marker := "  "
				if name == cfg.ActiveProvider {
					marker = "▶ "
				}
				displayName := p.DisplayName
				if len(displayName) > 17 {
					displayName = displayName[:14] + "..."
				}
				fmt.Printf("  %s%-14s %-18s %s\n", marker, name, displayName, p.BaseURL)
			}
			printDivider(52)
			fmt.Println()
			if cfg.ActiveProvider != "" {
				p := cfg.Providers[cfg.ActiveProvider]
				fmt.Printf("  Active: %s\n", p.DisplayName)
				if p.SonnetModel != "" {
					fmt.Printf("    Sonnet  →  %s\n", p.SonnetModel)
				}
				if p.OpusModel != "" {
					fmt.Printf("    Opus    →  %s\n", p.OpusModel)
				}
			}
			fmt.Println()
		} else {
			fmt.Println("No providers configured yet.")
			fmt.Println()
		}

		// ── Built-in presets ──────────────────────────────────────────────────
		fmt.Println("Available presets (taco add <name> --key=YOUR_KEY):")
		fmt.Println()

		allPresets := provider.ListPresets()
		sort.Strings(allPresets)

		// Group into international and China
		intl := []string{}
		china := []string{}
		chinaKeywords := []string{"siliconflow", "zhipu", "volcengine", "moonshot", "qwen", "yi", "baichuan", "minimax", "stepfun", "infini"}
		chinaSet := map[string]bool{}
		for _, k := range chinaKeywords {
			chinaSet[k] = true
		}
		for _, name := range allPresets {
			if chinaSet[name] {
				china = append(china, name)
			} else {
				intl = append(intl, name)
			}
		}

		fmt.Println("  International:")
		for _, name := range intl {
			p := provider.BuiltinPresets[name]
			_, configured := cfg.Providers[name]
			suffix := configuredSuffix(configured)
			fmt.Printf("    %-14s %s%s\n", name, p.DisplayName, suffix)
		}
		fmt.Println()
		fmt.Println("  China / 国内:")
		for _, name := range china {
			p := provider.BuiltinPresets[name]
			_, configured := cfg.Providers[name]
			suffix := configuredSuffix(configured)
			fmt.Printf("    %-14s %s%s\n", name, p.DisplayName, suffix)
		}
		fmt.Println()
		fmt.Println("  Quick add:  taco add deepseek --key=\"sk-...\"")
		fmt.Println("  Quick test: taco test")

		return nil
	},
}

func sortedKeys(m map[string]config.Provider) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func printDivider(n int) {
	fmt.Println("  " + strings.Repeat("─", n))
}

func configuredSuffix(configured bool) string {
	if configured {
		return "  ✓ configured"
	}
	return ""
}

func init() {
	rootCmd.AddCommand(listCmd)
}
