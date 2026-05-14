package app

import (
	tea "charm.land/bubbletea/v2"
	"github.com/krzystof/carnet/internal/commands"
)

func (m Model) Init() tea.Cmd {
	return commands.LoadConfig()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {

	case tea.KeyPressMsg:

		switch msg.String() {
		case "enter":
			if m.state == stateInitConfig {
				return m, commands.SetInitConfig(m.textInput.Value())
			}

		case "q":
			// Dont quit when typing "q" on a text input
			if m.state != stateInitConfig {
				return m, tea.Quit
			}

		case "ctrl+c":
			return m, tea.Quit
		}

	case commands.ConfigLoadedMsg:
		if msg.Cfg == nil {
			m.state = stateInitConfig
			m.textInput.Placeholder = "..."
			m.textInput.SetVirtualCursor(false)
			m.textInput.Focus()
		} else {
			m.state = stateLoadPage
			// TODO fire the command
		}

	case commands.ConfigPathErrorMsg:
		m.err = msg.Err
		m.state = stateError
	}

	// Update submodels
	m.textInput, cmd = m.textInput.Update(msg)

	return m, cmd
}
