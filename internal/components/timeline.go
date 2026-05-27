package components

import (
	"strconv"

	"charm.land/lipgloss/v2"
)

type Timeline struct {
	// cursorStart int
	// cursorDuration int
}

// j+k move up and down the cursor
// select an event if overlaps
// otherwise, highlight the relevant squares

func (t Timeline) View(w int) string {
	// use lipgloss layers
	return lipgloss.JoinHorizontal(lipgloss.Top, hourLabels(), " clock ", hoursBlocks(w))
}

func hourLabels() string {
	cells := []string{}

	for h := range 25 {
		cell := []string{
			strconv.Itoa(h),
			"",
		}

		cells = append(cells, cell...)
	}

	return lipgloss.JoinVertical(
		lipgloss.Right,
		cells...,
	)
}

func hoursBlocks(w int) string {
	s := lipgloss.NewStyle().Width(w)

	blocks := []string{}

	for i := 0; i <= 24*60; i += 30 {
		b := ""

		if i%60 == 0 {
			b = "---"
		}

		blocks = append(blocks, b)
	}

	col := lipgloss.JoinVertical(lipgloss.Left, blocks...)

	return s.Render(col)
}
