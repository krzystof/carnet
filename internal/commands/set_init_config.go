package commands

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	tea "charm.land/bubbletea/v2"
	"github.com/krzystof/carnet/internal/config"
)

func SetInitConfig(dataPath string) tea.Cmd {
	return func() tea.Msg {
		cfgPath, err := config.Path()

		if err != nil {
			return ConfigPathErrorMsg{Err: err}
		}

		if err := createCfgDirIfNotExists(cfgPath); err != nil {
			return err
		}

		appDir, err := config.Dir()

		if err != nil {
			return ConfigPathErrorMsg{Err: err}
		}

		if dataPath == "" {
			dataPath = appDir + "/diary"
		}

		cfg := config.Config{
			UserDataPath: dataPath,
		}

		data, err := json.MarshalIndent(cfg, "", "  ")

		if err != nil {
			return ConfigLoadedMsg{Cfg: nil, Err: err}
		}

		err = os.WriteFile(cfgPath, data, 0644)

		fmt.Println("[DEBUG] > ", "four", err)

		if err != nil {
			return ConfigLoadedMsg{Cfg: nil, Err: err}
		}

		return ConfigLoadedMsg{Cfg: &cfg, Err: nil}
	}
}

func createCfgDirIfNotExists(cfgPath string) error {
	cfgDir := filepath.Dir(cfgPath)

	if err := os.MkdirAll(cfgDir, 0755); err != nil {
		return err
	}
	return nil
}
