package project

import (
	"fmt"
	"os"
	"path/filepath"
	"spinup/config"
)

func CreateProject(conf config.Config) error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

  projectPath := filepath.Join(cwd, conf.Defaults.ProjectName)
  err = os.Mkdir(projectPath, 0755)
  if err != nil {
  	return err
  }

	frontend := ServiceStack{
		Stack: "front",
		Name: conf.Defaults.Frontend,
		Path: projectPath,
		ProjectName: conf.Defaults.ProjectName,
		TemplatesPath: conf.Templates.TemplatesPath,
		Remotes: conf.Templates.Remotes,
	}
	backend := ServiceStack{
		Stack: "back",
		Name: conf.Defaults.Backend,
		Path: projectPath,
		ProjectName: conf.Defaults.ProjectName,
		TemplatesPath: conf.Templates.TemplatesPath,
		Remotes: conf.Templates.Remotes,
	}

	err = setupStack(frontend)
	err = setupStack(backend)

	return err
}

func setupStack(stack ServiceStack) error {
	fmt.Printf("Building stack.\n")

	err := os.Mkdir(filepath.Join(stack.Path, fmt.Sprintf("%s-%s", stack.ProjectName, stack.Stack)), 0755)
	if err != nil {
		return err
	}

  for i := 0; i < len(stack.Remotes); i++ {
  }

	return nil
}
