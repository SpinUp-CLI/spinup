package config

import (
	"log"
	"os"
	"path/filepath"
	"fmt"
	"gopkg.in/yaml.v3"
)

func CreateDefaultConfig(init bool) error {
	// Go to home directory
	home, errDir := os.UserHomeDir()
	if errDir != nil {
		log.Fatalf("Error: %s\n", errDir)
		return errDir
	}

	// Generate the path of defaults.yaml
	configPath := filepath.Join(home, ".config", "spinup", "defaults.yaml")

	if init {
		// Check if the file exists
		_, errStat := os.Stat(configPath);
		if errStat == nil {
			fmt.Println("Config file already exists. If you want to reset your configuration, run 'spin reset'.")
		}
		if errStat != nil && os.IsExist(errStat) {
			log.Fatalf("Error: %s\n", errStat)
			return errStat
		}

		// Create the dir if it does not exist.
		errMkdir := os.MkdirAll(filepath.Dir(configPath), 0755)
		if errMkdir != nil {
			log.Fatalf("Error: %s\n", errMkdir)
			return errMkdir
		}
	} else {
		// Check if the file does not exist.
		_, errStat := os.Stat(configPath);
		if os.IsNotExist(errStat) {
			fmt.Println("Config file does not exist. Try doing 'spin init' to create it in your home directory.")
			return errStat
		}
	}

	// Create a new instance of configuration
	cfg := NewDefaultConfig()
	data, errMarshal := yaml.Marshal(cfg)
	if errMarshal != nil {
		log.Fatalf("Error: %s\n", errMarshal)
		return errMarshal
	}

	// Write the YAML data to the configuration file
	errWrite := os.WriteFile(configPath, data, 0644)
	if errWrite != nil {
		log.Fatalf("Error: %s\n", errWrite)
		return errWrite
	}
	return nil
}
