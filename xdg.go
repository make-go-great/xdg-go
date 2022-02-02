// https://specifications.freedesktop.org/basedir-spec/basedir-spec-latest.html

package xdg

import (
	"log"
	"os"
	"path/filepath"
)

const (
	EnvDataHome   = "XDG_DATA_HOME"
	EnvConfigHome = "XDG_CONFIG_HOME"
)

// GetHome return home from environment
func GetHome() string {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatalln("Failed to get user home dir", err)
	}

	return home
}

// GetDataHome return data home from environment first
// falback to default later
func GetDataHome() string {
	dataHome := os.Getenv(EnvDataHome)
	if dataHome != "" {
		return dataHome
	}

	home := GetHome()
	dataHome = filepath.Join(home, ".local/share")
	return dataHome
}

// GetConfigHome return config home from environment first
// fallback to default later
func GetConfigHome() string {
	configHome := os.Getenv(EnvConfigHome)
	if configHome != "" {
		return configHome
	}

	home := GetHome()
	configHome = filepath.Join(home, ".config")
	return configHome
}
