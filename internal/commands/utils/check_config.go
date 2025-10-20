package cmdutils

import (
	"spinup/internal/config"
	"spinup/pkg/fileutils"
	"spinup/pkg/iostream"
)

func ConfigExists() bool {
	configPath, err := config.GetConfigPath()
	iostream.Error(
		iostream.OutString{String: "Fetched the configuration file path"},
		iostream.OutString{String: "Something went wrong with configuration file path"},
		err,
	)

	configExists, err := fileutils.FileExists(configPath)
	iostream.Error(
		iostream.OutString{String: "Checked if it exists"},
		iostream.OutString{String: "Something went wrong with configuration file"},
		err,
	)

	return configExists
}
