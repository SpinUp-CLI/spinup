package config

func NewDefaultConfig() *Config {
	return &Config{
		EnvFile: "none",
		Templates: TemplatesConfig{
			TemplatePaths: "templates/",
			TemplateURL: "https://github.com/SpinUp-CLI/",
			CancelIfError: true,
			RemoteTemplates: []RemoteTemplate{
				{
					Name: "none",
					URL:  "none",
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
		CLI: CLIConfig{
			Verbose: true,
		},
	}
}

