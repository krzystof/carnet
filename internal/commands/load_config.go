// Package commands groups all the commands emitted to run through the main bubbletea Update()
package commands

import (
	"errors"
	"os"

	tea "charm.land/bubbletea/v2"
	"github.com/krzystof/carnet/internal/config"
)

type ConfigLoadedMsg struct {
	Cfg *config.Config
	Err error
}

type ConfigPathErrorMsg struct {
	Err error
}

func LoadConfig() tea.Cmd {
	return func() tea.Msg {
		path, err := config.Path()

		if err != nil {
			return ConfigPathErrorMsg{Err: err}
		}

		data, err := os.ReadFile(path)

		if err != nil {
			// If file not exist, trigger the first time setup
			if errors.Is(err, os.ErrNotExist) {
				return ConfigLoadedMsg{
					Cfg: nil,
					Err: err,
				}
			}

			// Other IO error
			return ConfigLoadedMsg{
				Cfg: nil,
				Err: err,
			}
		}

		cfg, err := config.FromJSON(data)

		if err != nil {
			return ConfigLoadedMsg{Cfg: nil, Err: err}
		}

		return ConfigLoadedMsg{Cfg: &cfg, Err: nil}
	}
}
