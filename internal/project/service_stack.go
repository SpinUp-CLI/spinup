package project

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"spinup/internal/config"

	"github.com/go-git/go-git/v5"
)

type ServiceStack struct {
	Name          string
	Stack         string
	ProjectName   string
	TemplatesPath string
	Path          string
	Remotes       []config.TemplatesRemote
}

func CreateProject(conf config.Config) error {
	projectPath := filepath.Join(cwd, conf.Defaults.ProjectName)

	frontend := ServiceStack{
		Stack:         "front",
		Name:          conf.Defaults.Frontend,
		Path:          projectPath,
		ProjectName:   conf.Defaults.ProjectName,
		TemplatesPath: conf.Templates.TemplatesPath,
		Remotes:       conf.Templates.Remotes,
	}

	err = setupStack(frontend)

	return err
}

func setupStack(stack ServiceStack) error {

	stackDir := filepath.Join(stack.Path, fmt.Sprintf("%s-%s", stack.ProjectName, stack.Stack))
	err := os.Mkdir(stackDir, 0755)
	if err != nil {
		return err
	}

	cloned := false
	for i := 0; i < len(stack.Remotes); i++ {
		remote := stack.Remotes[i]

		if remote.URL != "none" {
			// cloned = true
		}
	}

	if !cloned {
		if stack.TemplatesPath != "none" {
			// cloned true
			// moved the directory
		} else {
			templateUrl := fmt.Sprintf("https://github.com/SpinUp-CLI/%s-template", stack.Name)
			_, err := git.PlainClone(stackDir, false, &git.CloneOptions{
				URL: templateUrl,
			})
			if err != nil {
				return err
			}

			install := exec.Command("npm", "install")
			install.Dir = stackDir
			out, err := install.Output()
			if err != nil {
				fmt.Printf("%+v\n", err)
				return err
			}
			fmt.Printf("%s\nWebsite URL: http://localhost:5173/", out)
			run := exec.Command("npm", "run", "dev")
			run.Dir = stackDir
			out, err = run.Output()
			if err != nil {
				fmt.Printf("%+v\n", err)
				return err
			}
			fmt.Printf("%s\n", out)
		}
	}

	return nil
}
