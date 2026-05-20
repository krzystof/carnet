package app

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

func (m Model) View() tea.View {
	var c *tea.Cursor
	v := tea.NewView("")

	switch m.state {

	case stateStarting:
		v.SetContent("...")

	case stateInitConfig:
		if !m.textInput.VirtualCursor() {
			c = m.textInput.Cursor()
			c.Y += lipgloss.Height(m.headerView())
		}
		v.SetContent(WizardView(m))

	case stateLoadPage:
		v.SetContent("...")

	case stateReady:
		v.SetContent(m.page)

	case stateError:
		v.SetContent("Fatal config path error:\n\n" + m.err.Error() + "\n\nPress q to quit")
	}

	if m.err != nil {
		v.SetContent("/nError: " + m.err.Error() + "/n")
	}

	v.Cursor = c
	v.AltScreen = true

	return v
}
