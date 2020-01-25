package actions

import (
	"github.com/urfave/cli/v2"
	"grename/internal/log"
	"os"
	"os/user"
	"path/filepath"
)

func InitLogging(c *cli.Context) {
	termLevel := log.LevelInfo
	log.AddColorTerminalTarget(termLevel)

	//if fileLevel > log.LevelOff {
	//	f, err := log.AddFileTarget()
	//}
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
		subDir = "."+subDir
	}

	homeDir := filepath.Clean(envDir + "/" + projectNameDir + "/" + subDir)
	if _, err := os.Stat(homeDir); err != nil {
		if err := os.Mkdir(homeDir, os.FileMode(0755)); err != nil {
			return homeDir, err
		}
	}
	return homeDir, nil
}
