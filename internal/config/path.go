package config

import (
	"os"
	"path/filepath"
)

const AppName = "carnet"
const FileName = "carnet.json"

func Dir() (string, error) {
	dir, err := os.UserConfigDir()

	if err != nil {
		return "", err
	}

	return filepath.Join(dir, AppName), nil
}

func Path() (string, error) {
	appDir, err := Dir()

	if err != nil {
		return "", err
	}

	return filepath.Join(appDir, FileName), nil
}
