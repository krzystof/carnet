package components

import (
	"time"

	tea "charm.land/bubbletea/v2"
	"github.com/krzystof/carnet/internal/commands"
)

type Header struct {
	SelectedDate time.Time
}

func (h Header) Update(msg tea.Msg) (Header, tea.Cmd) {
	switch msg := msg.(type) {
	case commands.DateSelectedMsg:
		h.SelectedDate = msg.Date
	}

	return h, nil
}

func (h Header) View() string {
	return h.SelectedDate.Format("Monday 2 January 2006")
}
