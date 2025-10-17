package config

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"gopkg.in/yaml.v3"
)

func ExtractEnvPath(yamlContent []byte) (string, error) {
	var minimal MinimalConfig
	if err := yaml.Unmarshal(yamlContent, &minimal); err != nil {
		return "", nil
	}
	return minimal.EnvFile, nil
}

func ResolveEnvPath(envPath, configDir string) string {
	if filepath.IsAbs(envPath) {
		return envPath
	}
	return filepath.Join(configDir, envPath)
}

func ReplaceEnvVars(content []byte) []byte {
	pattern := regexp.MustCompile(`\$\{([^}]+)\}`)
	result := pattern.ReplaceAllStringFunc(string(content), func(match string) string {
		inner := match[2 : len(match)-1]

		parts := strings.SplitN(inner, ":", 2)
		varName := parts[0]

		defaultValue := ""
		if len(parts) > 1 {
			defaultValue = parts[1]
		}

		if value := os.Getenv(varName); value != "" {
			return value
		}
		return defaultValue
	})

	return []byte(result)
}
