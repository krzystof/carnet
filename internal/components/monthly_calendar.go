package components

import (
	"fmt"
	"time"

	"charm.land/lipgloss/v2"
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

	days := getCalendar(startOfRange, endOfRange)

	return lipgloss.JoinVertical(
		lipgloss.Left,
		c.SelectedDate.Format("January 2006"),
		"",
		"Mon Tue Wed Thu Fri Sat Sun",
		startOfRange.Format("2")+" -> "+endOfRange.Format("2"),
		days,
	)
}

func getMonthStartAndEnd(t time.Time) (time.Time, time.Time) {
	startOfMonth := time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
	endOfMonth := startOfMonth.AddDate(0, 1, -1)
	return startOfMonth, endOfMonth
}

func getCalendar(startOfRange, endOfRange time.Time) string {
	cursor := startOfRange
	rows := ""

	end := endOfRange.Add(24 * time.Hour)
	for cursor.Before(end) {
		d := fmt.Sprintf("% 3v", cursor.Day())

		// TODO <p1> format:
		// selected date
		// today
		// day not in the selected month

		rows = rows + d

		if cursor.Weekday() == time.Sunday {
			rows = rows + "\n"
		} else {
			rows = rows + " "
		}

		cursor = cursor.Add(24 * time.Hour)
	}

	return rows
}
