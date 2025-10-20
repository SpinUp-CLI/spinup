package project

import (
	"path/filepath"
	"spinup/internal/config"
)

type Project struct {
	Name     string
	Path     string
	Frontend ServiceStack
	Backend  ServiceStack
}

const (
	RemoteTemplate int = iota
	LocalTemplate
)

type ServiceStack struct {
	Name      string
	Path      string
	Framework string
	From      int
	URI       string
}

func CreateProject(cfg config.Config, path string) *Project {
	frontName := cfg.Defaults.ProjectName + "-front"
	backName := cfg.Defaults.ProjectName + "-back"

	return &Project{
		Name: cfg.Defaults.ProjectName,
		Path: path,
		Frontend: ServiceStack{
			Name:      frontName,
			Path:      filepath.Join(path, frontName),
			Framework: cfg.Defaults.Frontend,
			From:      RemoteTemplate,
			URI:       "undefined",
		},
		Backend: ServiceStack{
			Name:      backName,
			Path:      filepath.Join(path, backName),
			Framework: cfg.Defaults.Backend,
			From:      RemoteTemplate,
			URI:       "undefined",
		},
	}
}
