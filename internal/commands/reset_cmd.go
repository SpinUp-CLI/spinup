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
		iostream.Warning(
			iostream.OutString{},
			iostream.OutString{String: "Oops! It is impossible"},
			errors.New("it seems you don't have a configuration yet"),
		)
		return
	}

	configPath := cmdutils.WriteConfig()
	iostream.Success("Your configuration is now reset to defaults! Just right there: %s", configPath)
}
