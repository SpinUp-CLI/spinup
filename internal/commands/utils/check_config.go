package cmdutils

import (
	"spinup/internal/config"
	"spinup/pkg/fileutils"
	"spinup/pkg/iostream"
)

func ConfigExists() bool {
	configPath, err := config.GetConfigPath()
	iostream.Error(
		iostream.OutString{String: "Fetching your configuration file"},
		iostream.OutString{String: "Ho... Something went wrong with configuration file path"},
		err,
	)

	configExists, err := fileutils.FileExists(configPath)
	iostream.Error(
		iostream.OutString{String: "Checking if it exists"},
		iostream.OutString{String: "Ho... Something went wrong with configuration file"},
		err,
	)

	return configExists
}
