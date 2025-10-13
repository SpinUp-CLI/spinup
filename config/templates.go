/*
Copyright © 2025 Elouann Hosta ehosta@student.42lyon.fr
*/
package config

type TemplatesConfig struct {
	TemplatePaths   string            `yaml:"template_paths"`
	CancelIfError   bool              `yaml:"cancel_if_error"`
	RemoteTemplates []RemoteTemplate `yaml:"remote_templates"`
}

type RemoteTemplate struct {
	Name string `yaml:"name"`
	URL  string `yaml:"url"`
}
