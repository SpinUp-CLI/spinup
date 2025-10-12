/*
Copyright © 2025 Elouann Hosta ehosta@student.42lyon.fr
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "spin",
	Short: "SpinUp — Build customizable and full web apps in an instant.",
	Long: `SpinUp is a CLI tools that build fully configurable web apps,
based on the frameworks and templates of your choice.

Tap 'spin --premium' to discover how to become a premium member, and
the full list of advantages.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
