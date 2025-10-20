package cmd

import (
	"spinup/internal/commands"

	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Soon.",
	Long:  `Soon.`,
	Run:   commands.CreateCmd,
	Args:  cobra.MaximumNArgs(1),
}

func init() {
	rootCmd.AddCommand(createCmd)
}
