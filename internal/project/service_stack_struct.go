package project

import "spinup/internal/config"

type ServiceStack struct {
	Name          string
	Stack         string
	ProjectName   string
	TemplatesPath string
	Path          string
	Remotes       []config.TemplatesRemote
}
