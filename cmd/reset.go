/*
Copyright © 2025 Elouann Hosta ehosta@student.42lyon.fr
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Reset your current settings for SpinUp.",
	Long: `Reset the current configuration file and set
default values everywhere. Also set the comments everywhere
to help understanding the configuration properties.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("[Config reset called!]")
	},
}

func init() {
	configCmd.AddCommand(resetCmd)
}
