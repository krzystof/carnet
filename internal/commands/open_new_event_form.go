package commands

import tea "charm.land/bubbletea/v2"

type OpenNewEventFormMsg struct {
	StartTime int
}

func OpenNewEventForm(startTime int) tea.Cmd {
	return func() tea.Msg {
		return OpenNewEventFormMsg{StartTime: startTime}
	}
}
