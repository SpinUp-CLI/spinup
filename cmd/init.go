package cmd

import (
	"spinup/internal/commands"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Soon.",
	Long:  `Soon.`,
	Run:   commands.InitCmd,
}

func init() {
	rootCmd.AddCommand(initCmd)
}
