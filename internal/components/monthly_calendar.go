package components

import (
	"fmt"
	"time"

	"charm.land/lipgloss/v2"
	"github.com/krzystof/carnet/internal/styles"
)

type MonthlyCalendar struct {
	SelectedDate time.Time
}

// update function: nav hjkl -> change the selected date!

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
	now := time.Now()
	sel := time.Date(selectedDate.Year(), selectedDate.Month(), selectedDate.Day(), 0, 0, 0, 0, selectedDate.Location())
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

	for cursor.Before(end) {
		d := fmt.Sprintf("% 3v", cursor.Day())

		if cursor.Weekday() == time.Sunday {
			d = d + "\n"
		} else {
			d = d + " "
		}

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
		cursor = cursor.Add(24 * time.Hour)
	}

	return rows
}
