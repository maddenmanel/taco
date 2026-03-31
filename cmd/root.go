package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "taco",
	Short: "TACO - Terminal AI Configuration Organizer",
	Long: `TACO 🌮 - Terminal AI Configuration Organizer

A lightweight CLI tool that lets you seamlessly switch between
AI providers (DeepSeek, OpenRouter, SiliconFlow, etc.) while
continuing to use the "claude" command as usual.

No GUI, no database, no background processes.
Just pure text config manipulation — the Unix way.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
