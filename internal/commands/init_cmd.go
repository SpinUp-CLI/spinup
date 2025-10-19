package commands

import (
	"errors"
	"os"
	"spinup/internal/config"
	"spinup/pkg/fileutils"
	"spinup/pkg/iostream"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

func InitCmd(cmd *cobra.Command, args []string) {
	configPath, err := config.GetConfigPath()
	iostream.Error(iostream.OutString{}, iostream.OutString{String: "Ho... Something went wrong with config file path"}, err)

	configExists, err := fileutils.FileExists(configPath)
	iostream.Error(iostream.OutString{}, iostream.OutString{String: "Ho... Something went wrong with config file"}, err)

	if configExists {
		iostream.Warning(iostream.OutString{}, iostream.OutString{String: "Oops! It is impossible"}, errors.New("it seems you already have a configuration"))
		return
	}

	cfg := config.NewDefaultConfig()
	cfgYamlized, err := yaml.Marshal(cfg)
	iostream.Error(iostream.OutString{}, iostream.OutString{String: "Ho... Something went wrong with YAML generation"}, err)

	err = os.WriteFile(configPath, cfgYamlized, 0644)
	iostream.Error(iostream.OutString{}, iostream.OutString{String: "Ho... Something went wrong with YAML file"}, err)

	iostream.Log("Your configuration is now created! Just right there: %s", configPath)
}
