package commands

import (
	"errors"
	cmdutils "spinup/internal/commands/utils"
	"spinup/pkg/iostream"

	"github.com/spf13/cobra"
)

func InitCmd(cmd *cobra.Command, args []string) {
	configExists := cmdutils.ConfigExists()
	if configExists {
		iostream.Error(
			iostream.OutString{},
			iostream.OutString{String: "Configuration file found"},
			errors.New("it seems a configuration already exists"),
		)
		return
	}

	configPath := cmdutils.WriteConfig()
	iostream.Success("Configuration file created. Configuration file path: %s", configPath)
}
