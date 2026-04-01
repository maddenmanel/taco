package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/maddenmanel/taco/pkg/claude"
	"github.com/maddenmanel/taco/pkg/config"
	"github.com/spf13/cobra"
)

var uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "Remove TACO and restore Claude to official configuration",
	Long: `Uninstall TACO completely:
  1. Restores Claude Code to the official Anthropic configuration
  2. Removes all TACO config and provider data (~/.taco)
  3. Removes the taco binary from your system`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("🌮 Uninstalling TACO...")
		fmt.Println()

		// Step 1: restore Claude settings
		fmt.Print("  Restoring Claude Code to official config... ")
		if err := claude.Restore(); err != nil {
			fmt.Println("skipped (no settings found)")
		} else {
			fmt.Println("done")
		}

		// Step 2: remove taco config dir
		fmt.Print("  Removing ~/.taco config directory... ")
		if err := os.RemoveAll(config.ConfigDir()); err != nil {
			fmt.Printf("warning: %v\n", err)
		} else {
			fmt.Println("done")
		}

		// Step 3: remove the binary itself
		exePath, err := os.Executable()
		if err != nil {
			fmt.Println()
			fmt.Println("  Could not determine binary path.")
			fmt.Println("  Please remove the taco binary manually.")
			return nil
		}
		exePath, err = filepath.EvalSymlinks(exePath)
		if err != nil {
			exePath, _ = os.Executable()
		}

		fmt.Printf("  Removing binary (%s)... ", exePath)

		if runtime.GOOS == "windows" {
			// On Windows a running process cannot delete itself.
			// Schedule deletion via a cmd /c script that runs after we exit.
			batPath := filepath.Join(os.TempDir(), "taco_uninstall.bat")
			bat := fmt.Sprintf(`@echo off
ping -n 2 127.0.0.1 >nul
del /f /q "%s"
`, exePath)
			if writeErr := os.WriteFile(batPath, []byte(bat), 0600); writeErr == nil {
				exec.Command("cmd", "/c", "start", "/min", batPath).Start() //nolint
				fmt.Println("scheduled (runs after exit)")
			} else {
				fmt.Printf("\n  Could not schedule removal. Delete manually: %s\n", exePath)
			}
		} else {
			if err := os.Remove(exePath); err != nil {
				fmt.Printf("\n  Could not remove binary: %v\n", err)
				fmt.Printf("  Delete manually: %s\n", exePath)
			} else {
				fmt.Println("done")
			}
		}

		fmt.Println()
		fmt.Println("  TACO uninstalled. Claude Code is back to official Anthropic.")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(uninstallCmd)
}
