package commands

import (
	"errors"
	"os"
	cmdutils "spinup/internal/commands/utils"
	"spinup/pkg/iostream"

	"github.com/spf13/cobra"
)

func InitCmd(cmd *cobra.Command, args []string) {
	// Check if configuration exists. If not, a default one is created.
	iostream.Log("Checking current configuration.")
	configExists, err := cmdutils.ConfigExists()
	if err != nil {
		iostream.Error(err)
		os.Exit(1)
	}
	if configExists {
		iostream.Warning(errors.New("configuration already exists"))
		os.Exit(0)
	}

	// Create the configuration file.
	iostream.Info("Configuration does not exist yet. Creating one.")
	_, err = cmdutils.WriteConfig()
	if err != nil {
		iostream.Error(err)
		os.Exit(1)
	}
	iostream.Success()
}
