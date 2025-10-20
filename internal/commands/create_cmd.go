package commands

import (
	"fmt"
	"os"
	"path/filepath"
	cmdutils "spinup/internal/commands/utils"
	"spinup/internal/config"
	"spinup/internal/project"
	"spinup/pkg/iostream"

	"github.com/spf13/cobra"
)

func CreateCmd(cmd *cobra.Command, args []string) {
	// Check if configuration exists. If not, a default one is created and then the project goes to the creation
	// pipeline.
	iostream.Log("Checking current configuration.")
	configExists, err := cmdutils.ConfigExists()
	if err != nil {
		iostream.Error(err)
		os.Exit(1)
	}
	if !configExists {
		iostream.Info("Configuration does not exist yet. Creating one.")

		_, err := cmdutils.WriteConfig()
		if err != nil {
			iostream.Error(err)
			os.Exit(1)
		}

		iostream.Success()
	} else {
		iostream.Success()
	}

	// Get the configuration by parsing the yaml file.
	iostream.Log("Parsing configuration.")
	cfg, err := config.ParseConfig()
	if err != nil {
		iostream.Error(err)
		os.Exit(1)
	}
	if len(args) > 0 {
		iostream.Info(fmt.Sprintf("Using %s (not default) as project name.", args[0]))
		cfg.Defaults.ProjectName = args[0]
	} else {
		iostream.Info(fmt.Sprintf("Using %s (default) as project name.", cfg.Defaults.ProjectName))
	}
	iostream.Success()

	// Setup the directory for the project.
	iostream.Log("Setting up the project directory.")
	cwd, err := os.Getwd()
	if err != nil {
		iostream.Error(err)
		os.Exit(1)
	}
	projectPath := filepath.Join(cwd, cfg.Defaults.ProjectName)
	err = os.Mkdir(projectPath, 0755)
	if err != nil {
		iostream.Error(err)
		os.Exit(1)
	}
	iostream.Success()

	// Create the project stacks/services.
	iostream.Log("Creating project stacks/services.")
	err = project.CreateProject(projectPath, cfg)
	if err != nil {
		iostream.Error(err)
		os.Exit(1)
	}
}
