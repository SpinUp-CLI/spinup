package commands

import (
	"errors"
	"os"
	"path/filepath"
	cmdutils "spinup/internal/commands/utils"
	"spinup/internal/config"
	"spinup/pkg/iostream"

	"github.com/spf13/cobra"
)

func CreateCmd(cmd *cobra.Command, args []string) {
	configExists := cmdutils.ConfigExists()
	if !configExists {
		iostream.Warning(
			iostream.OutString{},
			iostream.OutString{String: "Configuration file not found"},
			errors.New("creating one"),
		)

		configPath := cmdutils.WriteConfig()
		iostream.Success("Configuration file created. Configuration file path: %s", configPath)
	}

	cfg, err := config.ParseConfig()
	iostream.Error(
		iostream.OutString{String: "Parsed the configuration file"},
		iostream.OutString{String: "Cannot parse the configuration file"},
		err,
	)
	if len(args) > 0 {
		cfg.Defaults.ProjectName = args[0]
	}

	cwd, err := os.Getwd()
	iostream.Error(
		iostream.OutString{String: "Fetched the current working directory"},
		iostream.OutString{String: "Cannot get the current working directory"},
		err,
	)

	projectPath := filepath.Join(cwd, cfg.Defaults.ProjectName)
	err = os.Mkdir(projectPath, 0755)
	iostream.Error(
		iostream.OutString{String: "Project directory created at %s", Args: []any{projectPath}},
		iostream.OutString{String: "cannot create the project directory"},
		err,
	)

	iostream.Log("Creating project")
	err = cmdutils.CreateProject(projectPath, cfg)
	iostream.Error(
		iostream.OutString{},
		iostream.OutString{String: "cannot create the project directory"},
		err,
	)

	iostream.Success("Project created at %s", projectPath)
}
