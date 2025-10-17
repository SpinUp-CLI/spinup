package config

func NewDefaultConfig() *Config {
	return &Config{
		EnvFile: "none",
		Templates: TemplatesConfig{
			TemplatesPath: "none",
			CancelIfError: true,
			Remotes: []TemplatesRemote{
				{
					Name:   "none",
					URL:    "none",
					Secret: "none",
				},
			},
		},
		Defaults: DefaultsConfig{
			FrontendEnabled: true,
			Frontend:        "vue",
			BackendEnabled:  true,
			Backend:         "go",
			DatabaseEnabled: true,
			Database:        "postgresql",
			ProjectName:     "web-app",
			Docker:          true,
			FrontendPort:    3000,
			BackendPort:     9000,
		},
	}
}
