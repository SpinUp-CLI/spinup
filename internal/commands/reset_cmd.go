package commands

import (
	"errors"
	cmdutils "spinup/internal/commands/utils"
	"spinup/pkg/iostream"

	"github.com/spf13/cobra"
)

func ResetCmd(cmd *cobra.Command, args []string) {
	configExists := cmdutils.ConfigExists()
	if !configExists {
		iostream.Error(
			iostream.OutString{},
			iostream.OutString{String: "Configuration file not found"},
			errors.New("it seems you don't have a configuration yet"),
		)
		return
	}

	configPath := cmdutils.WriteConfig()
	iostream.Success("Configuration reset to defaults. Configuration file path: %s", configPath)
}
