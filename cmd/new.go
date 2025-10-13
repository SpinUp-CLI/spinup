/*
Copyright © 2025 Elouann Hosta ehosta@student.42lyon.fr
*/
package cmd

import (
	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Soon.",
	Long:  `Soon.`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	rootCmd.AddCommand(newCmd)
}
