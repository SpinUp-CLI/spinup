package cmdutils

import (
	"spinup/internal/config"
	"spinup/pkg/fileutils"
)

func ConfigExists() (bool, error) {
	configPath, err := config.GetConfigPath()
	if err != nil {
		return false, err
	}

	configExists, err := fileutils.FileExists(configPath)
	return configExists, err
}
