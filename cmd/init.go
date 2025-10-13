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

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Soon.",
	Long:  `Soon.`,
	Run: func(cmd *cobra.Command, args []string) {
		home, err := os.UserHomeDir()
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			return
		}

		configPath := filepath.Join(home, ".config", "spinup", "defaults.yaml")

		if _, err := os.Stat(configPath); err == nil {
			fmt.Println("Config file already exists. If you want to reset your configuration, run 'spin reset'.")
		}
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			return
		}

		err = os.MkdirAll(filepath.Dir(configPath), 0755)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
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
	rootCmd.AddCommand(initCmd)
}
