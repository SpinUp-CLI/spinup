package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"spinup/pkg/fileutils"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
)

func GetConfigPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	configPath := filepath.Join(home, ".config", "spinup", "defaults.yaml")
	return configPath, nil
}

func CreateDefaultConfig(init bool) error {
	configPath, err := GetConfigPath()
	if err != nil {
		return err
	}

	exists, err := fileutils.FileExists(configPath)
	if err != nil {
		return err
	}

	if init {
		if exists {
			return errors.New("Config file already exists. If you want to reset your configuration, run 'spin reset'")
		}

		err = os.MkdirAll(filepath.Dir(configPath), 0755)
		if err != nil {
			return err
		}
	} else {
		if !exists {
			return errors.New("Config file does not exist. Try doing 'spin init' to create it in your home directory")
		}
	}

	cfg := NewDefaultConfig()
	data, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}

	err = os.WriteFile(configPath, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

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

func NewDefaultConfig() *Config {
	return &Config{
		EnvFile: "none",
		Templates: TemplatesConfig{
			TemplatesPath: "none",
			Remotes: []TemplatesRemote{
				{
					Name:   "none",
					URL:    "none",
					Secret: "none",
				},
			},
		},
		Defaults: DefaultsConfig{
			Frontend:    "vue",
			Backend:     "go",
			Database:    "postgresql",
			ProjectName: "web-app",
		},
	}
}
