package actions

import (
	"os"
	"os/user"
	"path/filepath"
)

func CreateDefaultLogFileName(projectName, logFileName string) string {
	dir, err := CreateLogDir(projectName)
	if err != nil {
		return ""
	}
	return filepath.Clean(dir + "/" + logFileName)
}

func CreateConfigDir(projectName string) (string, error) {
	return createHomeDir(projectName, "", "XDG_CONFIG_HOME")
}

func CreateLogDir(projectName string) (string, error) {
	return createHomeDir(projectName, "log", "XDG_DATA_HOME")
}

func createHomeDir(projectNameDir, subDir, envVarName string) (string, error) {
	envDir := os.Getenv(envVarName)
	if envDir == "" {
		u, _ := user.Current()
		envDir = u.HomeDir
		projectNameDir = "." + projectNameDir
	}

	homeDir := filepath.Clean(envDir + "/" + projectNameDir + "/" + subDir)
	if _, err := os.Stat(homeDir); err != nil {
		if err := os.MkdirAll(homeDir, os.FileMode(0755)); err != nil {
			return homeDir, err
		}
	}
	return homeDir, nil
}
