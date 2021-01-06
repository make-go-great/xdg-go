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

func GetHome() string {
	home := os.Getenv(EnvHome)
	return home
}

func GetDataHome() string {
	dataHome := os.Getenv(EnvDataHome)
	if dataHome != "" {
		return dataHome
	}

	home := GetHome()
	dataHome = filepath.Join(home, ".local/share")
	return dataHome
}

func GetConfigHome() string {
	configHome := os.Getenv(EnvConfigHome)
	if configHome != "" {
		return configHome
	}

	home := GetHome()
	configHome = filepath.Join(home, ".config")
	return configHome
}
