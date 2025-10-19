package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
)

func ParseConfig() (Config, error) {
	config := *NewDefaultConfig()

	configPath, err := GetConfigPath()
	if err != nil {
		return config, err
	}

	yamlContent, err := os.ReadFile(configPath)
	if err != nil {
		return config, err
	}

	processedYaml, err := processYamlEnv(yamlContent, filepath.Dir(configPath))
	if err != nil {
		return config, err
	}

	err = yaml.Unmarshal(processedYaml, &config)
	if err != nil {
		return config, err
	}

	return config, nil
}

func processYamlEnv(yamlContent []byte, configDir string) ([]byte, error) {
	envPath, err := ExtractEnvPath(yamlContent)
	if err != nil {
		return nil, err
	}

	if envPath != "" {
		absEnvPath := ResolveEnvPath(envPath, configDir)
		if err := godotenv.Load(absEnvPath); err != nil && envPath != "none" {
			fmt.Printf("Cannot load environment file: %s, because: %s\n", absEnvPath, err)
		}
	}

	processed := ReplaceEnvVars(yamlContent)
	return processed, nil
}
