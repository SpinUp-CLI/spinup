/*
Copyright © 2025 Elouann Hosta ehosta@student.42lyon.fr
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"spinup/config"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Soon.",
	Long:  `Soon.`,
	Run: func(cmd *cobra.Command, args []string) {
		home, err := os.UserHomeDir()
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			return
		}

		configPath := filepath.Join(home, ".config", "spinup", "defaults.yaml")

		if _, err := os.Stat(configPath); os.IsNotExist(err) {
			fmt.Println("Config file does not exist. Try doing 'spin init' to create it in your home directory.")
			return
		}

		cfg := config.NewDefaultConfig()
		data, err := yaml.Marshal(cfg)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			return
		}

		os.WriteFile(configPath, data, 0644)
	},
}

func init() {
	rootCmd.AddCommand(resetCmd)
}
