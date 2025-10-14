package cmd

import (
	"log"
	"spinup/config"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Soon.",
	Long:  `Soon.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := config.CreateDefaultConfig(true)
		if err != nil {
			log.Fatalf("Oops! An error occurred: %s\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
