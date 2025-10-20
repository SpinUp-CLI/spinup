package cmd

import (
	"spinup/internal/commands"

	"github.com/spf13/cobra"
)

var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Soon.",
	Long:  `Soon.`,
	Run:   commands.ResetCmd,
	Args: cobra.MaximumNArgs(0),
}

func init() {
	rootCmd.AddCommand(resetCmd)
}
