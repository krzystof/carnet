// Package app is the top-level bubbletea app
package app

import (
	"charm.land/bubbles/v2/textinput"
	"github.com/krzystof/carnet/internal/config"
)

type Model struct {
	cfg       *config.Config
	state     state
	textInput textinput.Model
	err       error
}

type state int

const (
	stateStarting = iota
	stateInitConfig
	stateLoadPage
	stateReady
	stateError
)

func InitialModel() Model {
	ti := textinput.New()

	return Model{
		state:     stateStarting,
		textInput: ti,
	}
}
