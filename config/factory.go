/*
Copyright © 2025 Elouann Hosta ehosta@student.42lyon.fr
*/
package config

func NewDefaultConfig() *Config {
	return &Config{
		EnvFile: "none",
		Templates: TemplatesConfig{
			TemplatePaths: "templates/",
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

