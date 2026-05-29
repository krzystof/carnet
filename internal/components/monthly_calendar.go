package components

import (
	"fmt"
	"time"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/krzystof/carnet/internal/commands"
	"github.com/krzystof/carnet/internal/styles"
)

type MonthlyCalendar struct {
	SelectedDate time.Time
}

func (c MonthlyCalendar) Update(msg tea.Msg) (MonthlyCalendar, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case commands.DateSelectedMsg:
		c.SelectedDate = msg.Date

	case tea.KeyPressMsg:
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
		var d string

		// Here we could use different styles for today, if we wanted to
		if cursor.Equal(today) {
			d = fmt.Sprintf(" %2v ", cursor.Day())
		} else {
			d = fmt.Sprintf(" %2v ", cursor.Day())
		}

		s := lipgloss.NewStyle()

		if cursor.Month() != sel.Month() {
			s = s.Foreground(styles.Theme.TextDimColor)
		}

		if cursor.Equal(today) {
			s = s.
				Foreground(styles.Theme.TextBrightColor).
				Bold(true).
				Underline(true)
		}

		if cursor.Equal(sel) {
			s = s.Background(styles.Theme.ItemActiveBackground)
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
