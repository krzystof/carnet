// Package app is the top-level bubbletea app
package app

import (
	"time"

	"charm.land/bubbles/v2/textinput"
	"github.com/krzystof/carnet/internal/components"
	"github.com/krzystof/carnet/internal/config"
	"github.com/krzystof/carnet/internal/core"
	"github.com/krzystof/carnet/internal/layout"
)

type Model struct {

	// app core states
	state           state
	activeComponent layout.ComponentName
	err             error
	cfg             *config.Config
	page            *core.Page
	selectedDate    time.Time

	// global ui stuff
	width  int
	height int

	// reused component whenever we need user input
	textInput textinput.Model

	// ui components
	header          components.Header
	monthlyCalendar components.MonthlyCalendar
	timeline        components.Timeline
	schedule        components.Schedule
	tasks           components.Tasks
	// eventDetail     layout.EventDetail
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
	return Model{
		state:           stateStarting,
		activeComponent: layout.ScheduleComponent,
		textInput:       textinput.New(),
		timeline:        components.NewTimeline(),
		schedule:        components.NewSchedule(),
	}
}
