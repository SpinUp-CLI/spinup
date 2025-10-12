/*
Copyright © 2025 Elouann Hosta ehosta@student.42lyon.fr
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Init the configuration file for spin.",
	Long: `If the configuration file does not exist,
it will be created in the "normal" place, in
'$HOME/.spinup/.config.yaml'.

Default values will be set, with commands.
Example:
	spin init`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("[Init called!]")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
