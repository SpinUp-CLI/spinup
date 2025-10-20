package project

import (
	"fmt"
	"spinup/internal/config"
	"spinup/pkg/iostream"
	"sync"
	"time"
)

func CreateProject(projectPath string, cfg config.Config) error {
	proj := NewDefaultProject(cfg, projectPath)
	var wg sync.WaitGroup

	iostream.Log("Starting services deployment")

	if proj.Frontend.Required {
		wg.Go(func() {
			createService(proj.Frontend)
		})
	} else {
		iostream.Log("No front-end required for this project.")
	}

	if proj.Backend.Required {
		wg.Go(func() {
			createService(proj.Backend)
		})
	} else {
		iostream.Log("No front-end required for this project.")
	}

	wg.Wait()
	return nil
}

func createService(stack ServiceStack) {
	time.Sleep(4 * time.Second)
	iostream.Congrats(fmt.Sprintf("%s: ready!", stack.Name))
}
