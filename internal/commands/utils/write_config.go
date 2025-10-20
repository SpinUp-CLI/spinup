package cmdutils

import (
	"os"
	"spinup/internal/config"

	"gopkg.in/yaml.v3"
)

func WriteConfig() (string, error) {
	cfg := config.NewDefaultConfig()
	cfgYamlized, err := yaml.Marshal(cfg)
	if err != nil {
		return "", err
	}

	configPath, err := config.GetConfigPath()
	if err != nil {
		return configPath, err
	}

	err = os.WriteFile(configPath, cfgYamlized, 0644)
	return configPath, err
}
