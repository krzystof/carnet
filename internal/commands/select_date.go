package commands

import (
	"time"

	tea "charm.land/bubbletea/v2"
)

type DateSelectedMsg struct {
	Date time.Time
}

func SelectDate(newDate time.Time) tea.Cmd {
	return func() tea.Msg {
		return DateSelectedMsg{Date: newDate}
	}
}
