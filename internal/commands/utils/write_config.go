package cmdutils

import (
	"os"
	"spinup/internal/config"
	"spinup/pkg/iostream"

	"gopkg.in/yaml.v3"
)

func WriteConfig() string {
	cfg := config.NewDefaultConfig()
	cfgYamlized, err := yaml.Marshal(cfg)
	iostream.Error(
		iostream.OutString{String: "Creating a new configuration"},
		iostream.OutString{String: "Ho... Something went wrong with the YAML generation"},
		err,
	)

	configPath, err := config.GetConfigPath()
	if err != nil {
		iostream.Error(
			iostream.OutString{String: "Fetching your configuration file"},
			iostream.OutString{String: "Ho... Something went wrong with the YAML file"},
			err,
		)
	}

	err = os.WriteFile(configPath, cfgYamlized, 0644)
	iostream.Error(
		iostream.OutString{String: "Writing your configuration"},
		iostream.OutString{String: "Ho... Something went wrong with the YAML file writing"},
		err,
	)

	return configPath
}
