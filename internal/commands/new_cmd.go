package commands

import (
	"log"
	"spinup/internal/config"
	"spinup/internal/service_stack"
	"spinup/pkg/fileutils"

	"github.com/spf13/cobra"
)

func NewCmd(cmd *cobra.Command, args []string) {
	path, err := config.GetConfigPath()
	if err != nil {
		log.Fatalf("Oops! An error occurred: %+v\n", err)
		return
	}

	exists, err := fileutils.FileExists(path)
	if err != nil {
		log.Fatalf("Oops! An error occurred: %+v\n", err)
		return
	}

	if !exists {
		return
	}

	config, err := config.ParseConfig()
	if err != nil {
		log.Fatalf("Oops! An error occurred: %+v\n", err)
		return
	}

	if err = service_stack.CreateProject(config); err != nil {
		log.Fatalf("Oops! An error occurred: %+v\n", err)
		return
	}
}
