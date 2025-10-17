package project

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"spinup/config"

	"github.com/go-git/go-git/v5"
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
		Stack:         "front",
		Name:          conf.Defaults.Frontend,
		Path:          projectPath,
		ProjectName:   conf.Defaults.ProjectName,
		TemplatesPath: conf.Templates.TemplatesPath,
		Remotes:       conf.Templates.Remotes,
	}
	// backend := ServiceStack{
	// 	Stack:         "back",
	// 	Name:          conf.Defaults.Backend,
	// 	Path:          projectPath,
	// 	ProjectName:   conf.Defaults.ProjectName,
	// 	TemplatesPath: conf.Templates.TemplatesPath,
	// 	Remotes:       conf.Templates.Remotes,
	// }

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
			// moved the folder
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
			fmt.Printf("%s\n", out)
			run := exec.Command("npm", "run", "dev", "&")
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
