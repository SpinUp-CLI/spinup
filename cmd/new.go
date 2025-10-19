package cmd

import (
	"spinup/internal/commands"

	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Soon.",
	Long:  `Soon.`,
	Run:   commands.NewCmd,
}

func init() {
	rootCmd.AddCommand(newCmd)
}
