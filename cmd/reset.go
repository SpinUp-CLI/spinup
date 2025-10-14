package cmd

import (
	"spinup/config"

	"github.com/spf13/cobra"
)

var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Soon.",
	Long:  `Soon.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := config.CreateDefaultConfig(false)
		if err != nil {
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(resetCmd)
}
