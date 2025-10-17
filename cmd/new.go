package cmd

import (
	"log"
	"spinup/internal/config"
	"spinup/internal/project"
	"spinup/pkg/fileutils"

	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Soon.",
	Long:  `Soon.`,
	Run: func(cmd *cobra.Command, args []string) {
		path, err := config.GetConfigPath()
		if err != nil {
			log.Fatalf("Oops! An error occurred: %+v\n", err)
			return
		}

		exists, err := fileutils.FileExists(path)
		if err != nil {
			log.Fatalf("Oops! An error occurred: %+v\n", err)
			return
		}

		if !exists {
			err = config.CreateDefaultConfig(true)
			if err != nil {
				log.Fatalf("Oops! An error occurred: %+v\n", err)
				return
			}
		}

		config, err := config.ParseConfig()
		if err != nil {
			log.Fatalf("Oops! An error occurred: %+v\n", err)
			return
		}

		if err = project.CreateProject(config); err != nil {
			log.Fatalf("Oops! An error occurred: %+v\n", err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
}
