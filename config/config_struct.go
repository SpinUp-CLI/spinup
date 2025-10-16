package config

type Config struct {
	EnvFile   string          `yaml:"env_file"`
	Templates TemplatesConfig `yaml:"templates"`
	Defaults  DefaultsConfig  `yaml:"defaults"`
}

type MinimalConfig struct {
	EnvFile string `yaml:"env_file"`
}

type TemplatesConfig struct {
	TemplatesPath   string           `yaml:"templates_path"`
	CancelIfError   bool             `yaml:"cancel_if_error"`
	Remotes []TemplatesRemote `yaml:"remote_templates"`
}

type TemplatesRemote struct {
	Name string `yaml:"name"`
	URL  string `yaml:"url"`
	Secret string `yaml:"secret"`
}

type DefaultsConfig struct {
	FrontendEnabled bool   `yaml:"frontend_enabled"`
	Frontend        string `yaml:"frontend"`
	BackendEnabled  bool   `yaml:"backend_enabled"`
	Backend         string `yaml:"backend"`
	DatabaseEnabled bool   `yaml:"database_enabled"`
	Database        string `yaml:"database"`
	ProjectName     string `yaml:"project_name"`
	Docker          bool   `yaml:"docker"`
	FrontendPort    int    `yaml:"frontend_port"`
	BackendPort     int    `yaml:"backend_port"`
}
