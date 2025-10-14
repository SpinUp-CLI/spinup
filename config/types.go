package config

type Config struct {
	EnvFile   string          `yaml:"env_file"`
	Templates TemplatesConfig `yaml:"templates"`
	Defaults  DefaultsConfig  `yaml:"defaults"`
	CLI       CLIConfig       `yaml:"cli"`
}

type MinimalConfig struct {
	EnvFile string `yaml:"env_file"`
}
