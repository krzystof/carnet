package components

import (
	"fmt"
	"time"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/krzystof/carnet/internal/commands"
	"github.com/krzystof/carnet/internal/layout"
	"github.com/krzystof/carnet/internal/styles"
)

type MonthlyCalendar struct {
	SelectedDate time.Time
	focused      bool
}

func (c MonthlyCalendar) Update(msg tea.Msg) (MonthlyCalendar, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case layout.FocusChangedMsg:
		c.focused = msg.Comp == layout.SidebarComponent

	case tea.KeyPressMsg:
		if c.focused {
			switch msg.String() {
			case "h":
				c.SelectedDate = c.SelectedDate.Add(-1 * 24 * time.Hour)
				cmd = commands.SelectDate(c.SelectedDate)
			case "j":
				c.SelectedDate = c.SelectedDate.Add(7 * 24 * time.Hour)
				cmd = commands.SelectDate(c.SelectedDate)
			case "k":
				c.SelectedDate = c.SelectedDate.Add(-7 * 24 * time.Hour)
				cmd = commands.SelectDate(c.SelectedDate)
			case "l":
				c.SelectedDate = c.SelectedDate.Add(1 * 24 * time.Hour)
				cmd = commands.SelectDate(c.SelectedDate)
			case "t":
				c.SelectedDate = todayDate()
				cmd = commands.SelectDate(c.SelectedDate)
			}

		}
	}

	return c, cmd
}

func (c MonthlyCalendar) View() string {
	startOfRange, endOfRange := getMonthStartAndEnd(c.SelectedDate)

	for startOfRange.Weekday() != time.Monday {
		startOfRange = startOfRange.Add(-1 * 24 * time.Hour)
	}

	for endOfRange.Weekday() != time.Sunday {
		endOfRange = endOfRange.Add(24 * time.Hour)
	}

	days := getCalendar(startOfRange, endOfRange, c.SelectedDate)

	return lipgloss.JoinVertical(
		lipgloss.Left,
		lipgloss.NewStyle().Underline(true).Render(c.SelectedDate.Format("January 2006")),
		"",
		"Mon Tue Wed Thu Fri Sat Sun",
		days,
	)
}

func getMonthStartAndEnd(t time.Time) (time.Time, time.Time) {
	startOfMonth := time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
	endOfMonth := startOfMonth.AddDate(0, 1, -1)
	return startOfMonth, endOfMonth
}

func getCalendar(startOfRange, endOfRange, selectedDate time.Time) string {
	cursor := startOfRange
	end := endOfRange.Add(24 * time.Hour)

	rows := ""
	sel := time.Date(selectedDate.Year(), selectedDate.Month(), selectedDate.Day(), 0, 0, 0, 0, selectedDate.Location())
	today := todayDate()

	for cursor.Before(end) {
		d := fmt.Sprintf("% 3v", cursor.Day()) + " "

		s := lipgloss.NewStyle()

		if cursor.Month() != sel.Month() {
			s = s.Foreground(lipgloss.Color(styles.Theme.TextDimColor))
		}

		if cursor.Equal(today) {
			s = s.
				Foreground(lipgloss.Color(styles.Theme.TextBrightColor)).
				Bold(true).
				Underline(true)
		}

		if cursor.Equal(sel) {
			s = s.Background(lipgloss.Color(styles.Theme.ItemActiveBackground))
		}

		rows = rows + s.Render(d)

		if cursor.Weekday() == time.Sunday {
			rows = rows + " \n"
		}

		cursor = cursor.Add(24 * time.Hour)
	}

	return rows
}

func todayDate() time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
}
