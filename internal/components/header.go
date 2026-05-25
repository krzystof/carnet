package components

import (
	"time"

	tea "charm.land/bubbletea/v2"
)

type Header struct {
	SelectedDate time.Time
}

func (h Header) Update(msg tea.Msg) (Header, tea.Cmd) {
	// switch msg := msg.(type) {
	// case commands.PageLoadedMsg:
	// 	h.dateTime = msg.Page.Time
	// }

	return h, nil
}

func (h Header) View() string {
	return h.SelectedDate.Format("Monday 2 January 2006")
}
