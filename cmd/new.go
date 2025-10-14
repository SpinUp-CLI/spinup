package cmd

import (
	"fmt"
	"log"
	"spinup/config"

	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Soon.",
	Long:  `Soon.`,
	Run: func(cmd *cobra.Command, args []string) {
		config, err := config.ParseConfig()
		if err != nil {
			log.Fatalf("Oops! An error occurred: %s\n", err)
		}

		fmt.Printf("Front-end framework will be: %s\n", config.Defaults.Frontend)
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
}
