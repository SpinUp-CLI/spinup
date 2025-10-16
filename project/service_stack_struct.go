package project

import "spinup/config"

type ServiceStack struct {
	Name string
	Stack string
	ProjectName string
	TemplatesPath string
	Path string
	Remotes []config.TemplatesRemote
}

