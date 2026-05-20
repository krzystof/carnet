// Package commands groups all the commands emitted to run through the main bubbletea Update()
package commands

import (
	"errors"
	"os"

	tea "charm.land/bubbletea/v2"
	"github.com/krzystof/carnet/internal/config"
)

type ConfigLoadedMsg struct {
	Cfg config.Config
}

type ConfigNotExistsMsg struct{}

type ConfigLoadFailedMsg struct {
	Err error
}

func LoadConfig() tea.Cmd {
	return func() tea.Msg {
		path, err := config.Path()

		if err != nil {
			return ConfigLoadFailedMsg{Err: err}
		}

		data, err := os.ReadFile(path)

		if err != nil {
			// If file not exist, trigger the first time setup
			if errors.Is(err, os.ErrNotExist) {
				return ConfigNotExistsMsg{}
			}

			// Other IO error
			return ConfigLoadFailedMsg{Err: err}
		}

		cfg, err := config.FromJSON(data)

		if err != nil {
			return ConfigLoadFailedMsg{Err: err}
		}

		return ConfigLoadedMsg{Cfg: cfg}
	}
}
