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

type ServiceStack struct {
	Name      string
	Path      string
	Framework string
	URL       string
	Required  bool
}

func NewDefaultProject(cfg config.Config, path string) *Project {
	frontName := cfg.Defaults.ProjectName + "-front"
	backName := cfg.Defaults.ProjectName + "-back"

	return &Project{
		Name: cfg.Defaults.ProjectName,
		Path: path,
		Frontend: ServiceStack{
			Name:      frontName,
			Path:      filepath.Join(path, frontName),
			Framework: cfg.Defaults.Frontend,
			URL:       "undefined",
			Required:  cfg.Defaults.Frontend != "none",
		},
		Backend: ServiceStack{
			Name:      backName,
			Path:      filepath.Join(path, backName),
			Framework: cfg.Defaults.Backend,
			URL:       "undefined",
			Required:  cfg.Defaults.Backend != "none",
		},
	}
}
