package project

import (
	"errors"
	"fmt"
	"spinup/internal/config"
	"spinup/pkg/iostream"
	"sync"
)

func CreateProject(projectPath string, cfg config.Config) error {
	proj := NewDefaultProject(cfg, projectPath)
	var wg sync.WaitGroup

	iostream.Log("Starting services deployment")

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
		iostream.Danger(errors.New("could not find any repository for the given templates"))
		return
	}
	iostream.Congrats(fmt.Sprintf("Stack %s has been deployed with this URL: %s", stack.Name, stack.URL))
}
