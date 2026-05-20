// Package app is the top-level bubbletea app
package app

import (
	"charm.land/bubbles/v2/textinput"
	"github.com/krzystof/carnet/internal/config"
)

type Model struct {
	state     state
	textInput textinput.Model
	err       error
	cfg       *config.Config
	page      string // TODO that's going to be a Page
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
