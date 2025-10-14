package config

import (
	"errors"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

/*
Return the path of the configuration file if it exists.
Does not handle error logging.
*/
func getConfigPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	configPath := filepath.Join(home, ".config", "spinup", "defaults.yaml")
	return configPath, nil
}

/*
(Re)initialize the configuration file. If the init parameter
is set to true, then the file will be created. Otherwise, it will be overwitten.
Does not handle error logging.
*/
func CreateDefaultConfig(init bool) error {
	// Get the path of defaults.yaml
	configPath, err := getConfigPath()
	if err != nil {
		return err
	}

	if init {
		// Check if the file exists
		_, err := os.Stat(configPath);
		if err == nil {
			return errors.New("Config file already exists. If you want to reset your configuration, run 'spin reset'")
		}
		if os.IsExist(err) {
			return err
		}

		// Create the dir if it does not exist.
		err = os.MkdirAll(filepath.Dir(configPath), 0755)
		if err != nil {
			return err
		}
	} else {
		// Check if the file does not exist.
		_, err := os.Stat(configPath);
		if os.IsNotExist(err) {
			return errors.New("Config file does not exist. Try doing 'spin init' to create it in your home directory")
		}
	}

	// Create a new instance of configuration
	cfg := NewDefaultConfig()
	data, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}

	// Write the YAML data to the configuration file
	err = os.WriteFile(configPath, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

/*
Read the configuration file and return the struct containing all
the values.
Does not handle error logging.
*/
func ParseConfig() (Config, error) {
	var config = Config{}

	// Get the path of defaults.yaml
	configPath, err := getConfigPath()
	if err != nil {
		return config, err
	}

	// Read the content of the configuration file
	yamlContent, err := os.ReadFile(configPath)
	if err != nil {
		return config, err
	}

	// Get the configuration structure from the file content
	err = yaml.Unmarshal(yamlContent, &config)
	if err != nil {
		return config, err
	}

	return config, nil
}
