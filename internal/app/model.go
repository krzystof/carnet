// Package app is the top-level bubbletea app
package app

import (
	"charm.land/bubbles/v2/textinput"
	"github.com/krzystof/carnet/internal/components"
	"github.com/krzystof/carnet/internal/config"
	"github.com/krzystof/carnet/internal/core"
)

type Model struct {

	// app core states
	state           state
	activeComponent components.ComponentName
	err             error
	cfg             *config.Config
	page            *core.Page

	// global ui stuff
	width  int
	height int

	// reused component whenever we need user input
	textInput textinput.Model

	// ui components
	header components.Header
	// monthlyCalendar components.MonthlyCalendar
	// timeline        components.Timeline
	// tasks           components.Tasks
	// eventDetail     components.EventDetail
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
		state:           stateStarting,
		activeComponent: components.TimelineComponent,
		textInput:       ti,
	}
}
