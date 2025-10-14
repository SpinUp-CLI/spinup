package cmd

import (
	"spinup/config"
	"log"

	"github.com/spf13/cobra"
)

var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Soon.",
	Long:  `Soon.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := config.CreateDefaultConfig(false)
		if err != nil {
			log.Fatalf("Oops! An error occurred: %s\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(resetCmd)
}
