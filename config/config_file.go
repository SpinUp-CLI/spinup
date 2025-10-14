package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/joho/godotenv"
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

func ConfigFileExists() (bool, error) {
	configPath, err := getConfigPath()

	_, err = os.Stat(configPath)
	if err == nil {
		return true, nil
	}
	if os.IsExist(err) {
		return true, err
	}

	return false, nil
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

	// Check if the file exists
	exists, err := ConfigFileExists()
	if err != nil {
		return err
	}

	if init {
		if exists {
			return errors.New("Config file already exists. If you want to reset your configuration, run 'spin reset'")
		}

		// Create the dir if it does not exist.
		err = os.MkdirAll(filepath.Dir(configPath), 0755)
		if err != nil {
			return err
		}
	} else {
		// Check if the file does not exist.
		if !exists {
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
	var config Config

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

	processedYaml, err := processYamlEnv(yamlContent, filepath.Dir(configPath))
	if err != nil {
		return config, err
	}

	// Get the configuration structure from the file content
	err = yaml.Unmarshal(processedYaml, &config)
	if err != nil {
		return config, err
	}

	return config, nil
}

func processYamlEnv(yamlContent []byte, configDir string) ([]byte, error) {
	envPath, err := extractEnvPath(yamlContent)
	if err != nil {
		return nil, err
	}

	if envPath != "" {
		absEnvPath := resolveEnvPath(envPath, configDir)
		if err := godotenv.Load(absEnvPath); err != nil && envPath != "none" {
			fmt.Printf("Cannot load environment file: %s, because: %s\n", absEnvPath, err)
		}
	}

	processed := replaceEnvVars(yamlContent)
	return processed, nil
}

func extractEnvPath(yamlContent []byte) (string, error) {
	var minimal MinimalConfig
	if err := yaml.Unmarshal(yamlContent, &minimal); err != nil {
		return "", nil
	}
	return minimal.EnvFile, nil
}

func resolveEnvPath(envPath, configDir string) string {
	if filepath.IsAbs(envPath) {
		return envPath
	}
	return filepath.Join(configDir, envPath)
}

func replaceEnvVars(content []byte) []byte {
	pattern := regexp.MustCompile(`\$\{([^}]+)\}`)
	result := pattern.ReplaceAllStringFunc(string(content), func(match string) string {
		inner := match[2 : len(match)-1]

		parts := strings.SplitN(inner, ":", 2)
		varName := parts[0]

		defaultValue := ""
		if len(parts) > 1 {
			defaultValue = parts[1]
		}

		if value := os.Getenv(varName); value != "" {
			return value
		}
		return defaultValue
	})

	return []byte(result)
}
