/*
Copyright © 2025 Elouann Hosta ehosta@student.42lyon.fr
*/
package cmd

import (
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage your SpinUp configuration.",
	Long:  `Manage your SpinUp configuration with some commands.`,
}

func init() {
	rootCmd.AddCommand(configCmd)
}
