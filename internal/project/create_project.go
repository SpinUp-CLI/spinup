package project

import (
	"fmt"
	"spinup/internal/config"
	"spinup/pkg/iostream"
	"spinup/pkg/utils"
	"sync"
)

func CreateProject(projectPath string, cfg config.Config) error {
	proj := NewDefaultProject(cfg, projectPath)
	var wg sync.WaitGroup

	if proj.Frontend.Required {
		wg.Go(func() {
			createService(proj.Frontend, cfg)
		})
	} else {
		iostream.Log("No front-end required for this project.")
	}

	if proj.Backend.Required {
		wg.Go(func() {
			createService(proj.Backend, cfg)
		})
	} else {
		iostream.Log("No front-end required for this project.")
	}

	wg.Wait()
	return nil
}

func createService(stack ServiceStack, cfg config.Config) {
	remotes := cfg.Templates.Remotes
	found := false
	remote := remotes[0]

	for i := range remotes {
		remote = remotes[i]
		stack.URL = remote.URL + stack.Framework

		_, err := TryRemote(stack.Path, stack.URL, remote)
		if err == nil {
			found = true
			break
		}
	}

	if !found {
		iostream.Error(fmt.Errorf("service %s cannot be deployed, '%s' template not found", stack.Name, stack.Framework))
		return
	}

	iostream.Congrats(fmt.Sprintf("Stack has been deployed.\n%s    - Remote: %s\n%s    - Path: %s", utils.TsStep, stack.URL, utils.TsStep, stack.Path))
}
