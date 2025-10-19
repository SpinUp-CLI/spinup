package config

import (
	"os"
	"path/filepath"
)

type Config struct {
	EnvFile   string          `yaml:"env_file"`
	Templates TemplatesConfig `yaml:"templates"`
	Defaults  DefaultsConfig  `yaml:"defaults"`
}

type MinimalConfig struct {
	EnvFile string `yaml:"env_file"`
}

type TemplatesConfig struct {
	TemplatesPath string            `yaml:"templates_path"`
	Remotes       []TemplatesRemote `yaml:"remote_templates"`
}

type TemplatesRemote struct {
	Name   string `yaml:"name"`
	URL    string `yaml:"url"`
	Secret string `yaml:"secret"`
}

type DefaultsConfig struct {
	Frontend    string `yaml:"frontend"`
	Backend     string `yaml:"backend"`
	Database    string `yaml:"database"`
	ProjectName string `yaml:"project_name"`
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

func GetConfigPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	configPath := filepath.Join(home, ".config", "spinup", "defaults.yaml")
	return configPath, nil
}
