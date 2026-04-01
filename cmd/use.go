package cmd

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/maddenmanel/taco/pkg/claude"
	"github.com/maddenmanel/taco/pkg/config"
	"github.com/spf13/cobra"
)

var useCmd = &cobra.Command{
	Use:   "use [provider-name]",
	Short: "Switch Claude Code to use the specified provider",
	Long: `Inject the specified provider's configuration into Claude Code's
settings.json. After switching, just type "claude" as usual.

If no provider name is given, an interactive picker is shown.`,
	Example: `  taco use deepseek
  taco use          # interactive picker`,
	Args: cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.Load()
		if err != nil {
			return err
		}

		var name string

		if len(args) == 1 {
			name = args[0]
		} else {
			// Interactive picker
			name, err = pickProvider(cfg)
			if err != nil {
				return err
			}
			if name == "" {
				return nil
			}
		}

		p, ok := cfg.GetProvider(name)
		if !ok {
			return fmt.Errorf("provider %q not found.\n\n  Add it first: taco add %s --key=YOUR_KEY\n  Or see available presets: taco list", name, name)
		}

		if err := claude.InjectProvider(p); err != nil {
			return fmt.Errorf("failed to switch provider: %w", err)
		}

		cfg.ActiveProvider = name
		if err := cfg.Save(); err != nil {
			return err
		}

		fmt.Printf("🌮 Switched to %s\n", p.DisplayName)
		fmt.Printf("   Base URL:  %s\n", p.BaseURL)
		if p.SonnetModel != "" {
			fmt.Printf("   Sonnet  →  %s\n", p.SonnetModel)
		}
		if p.OpusModel != "" {
			fmt.Printf("   Opus    →  %s\n", p.OpusModel)
		}
		fmt.Println()
		fmt.Println("   Just type `claude` to start coding!")
		return nil
	},
}

// pickProvider shows a numbered list of configured providers and reads a choice.
func pickProvider(cfg *config.TacoConfig) (string, error) {
	if len(cfg.Providers) == 0 {
		return "", fmt.Errorf("no providers configured yet.\n\n  Add one: taco add deepseek --key=YOUR_KEY\n  See presets: taco list")
	}

	names := make([]string, 0, len(cfg.Providers))
	for name := range cfg.Providers {
		names = append(names, name)
	}
	sort.Strings(names)

	fmt.Println("Select a provider:")
	fmt.Println()
	for i, name := range names {
		p := cfg.Providers[name]
		active := "  "
		if name == cfg.ActiveProvider {
			active = "* "
		}
		fmt.Printf("  %s%d) %-15s %s\n", active, i+1, name, p.DisplayName)
	}
	fmt.Println()
	fmt.Print("Enter number or name: ")

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	input = strings.TrimSpace(input)

	if input == "" {
		return "", nil
	}

	// Try as number first
	if n, err := strconv.Atoi(input); err == nil {
		if n >= 1 && n <= len(names) {
			return names[n-1], nil
		}
		return "", fmt.Errorf("invalid selection: %s", input)
	}

	// Try as name
	if _, ok := cfg.GetProvider(input); ok {
		return input, nil
	}
	return "", fmt.Errorf("provider %q not found", input)
}

func init() {
	rootCmd.AddCommand(useCmd)
}
