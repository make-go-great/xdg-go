// https://specifications.freedesktop.org/basedir-spec/basedir-spec-latest.html

package xdg

import (
	"os"
	"path/filepath"
)

const (
	EnvHome       = "HOME"
	EnvDataHome   = "XDG_DATA_HOME"
	EnvConfigHome = "XDG_CONFIG_HOME"
)

// GetHome return home from environment
func GetHome() string {
	home := os.Getenv(EnvHome)
	return home
}

// GetDataHome return data gome from environment first
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
