package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Version is set at build time via:
//
//	go build -ldflags "-X github.com/maddenmanel/taco/cmd.Version=v1.0.0"
var Version = "dev"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print TACO version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("taco %s\n", Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
	// Also wire --version flag on root
	rootCmd.Version = Version
}
