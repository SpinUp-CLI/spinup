package config

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
