package cmd

import (
	"log"
	"spinup/config"
	"spinup/project"

	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Soon.",
	Long:  `Soon.`,
	Run: func(cmd *cobra.Command, args []string) {
		exists, err := config.ConfigFileExists()
		if err != nil {
			log.Fatalf("Oops! An error occurred: %s\n", err)
			return
		}

		if !exists {
			err = config.CreateDefaultConfig(true)
			if err != nil {
				log.Fatalf("Oops! An error occurred: %s\n", err)
				return
			}
		}

		config, err := config.ParseConfig()
		if err != nil {
			log.Fatalf("Oops! An error occurred: %s\n", err)
			return
		}

		if project.CreateProject(config) != nil {
			log.Fatalf("Oops! An error occurred: %s\n", err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
}
