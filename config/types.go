/*
Copyright © 2025 Elouann Hosta ehosta@student.42lyon.fr
*/
package config

type Config struct {
	EnvFile   string          `yaml:"env_file"`
	Templates TemplatesConfig `yaml:"templates"`
	Defaults  DefaultsConfig  `yaml:"defaults"`
	CLI       CLIConfig       `yaml:"cli"`
}
