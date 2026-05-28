package components

import (
	"strconv"

	"charm.land/lipgloss/v2"
	"github.com/krzystof/carnet/internal/styles"
)

type Timeline struct {
	// cursorStart int
	// cursorDuration int
}

// j+k move up and down the cursor
// select an event if overlaps
// otherwise, highlight the relevant squares

func (t Timeline) View(width, height int) string {
	// use lipgloss layers, maybe?

	// 8am by default, will need to check this OR first event
	// TODO this should be init, after its internal state, also do this only if height is not enough
	startFrom := 8 * 60

	slots := visibleSlots(startFrom, height-8) // 2 borders, 2 padding, 2 labels, 2 borders

	visibleRows := lipgloss.JoinHorizontal(lipgloss.Top, hourLabels(slots), " clock ", hoursBlocks(slots))

	s := lipgloss.NewStyle().
		Width(width-4). // (got to remove those borders and padding)
		BorderForeground(styles.Theme.BorderInactiveColor).
		Border(lipgloss.NormalBorder(), true, false)

	return lipgloss.JoinVertical(
		lipgloss.Left,
		"Press k to move up",
		s.Render(visibleRows),
		"Press j to move down",
	)
}

func hourLabels(slots []int) string {
	cells := []string{}

	for _, minutes := range slots {
		cell := ""
		if minutes%60 == 0 {
			hour := minutes / 60
			cell = strconv.Itoa(hour)
		}

		cells = append(cells, cell)
	}

	return lipgloss.JoinVertical(
		lipgloss.Right,
		cells...,
	)
}

// Textures:
// ░  light shade
// ▒  medium shade
// ▓  dark shade
// █  full block
func hoursBlocks(slots []int) string {
	s := lipgloss.NewStyle()

	blocks := []string{}

	for _, minutes := range slots {
		b := ""

		if minutes%60 == 0 {
			b = "---"
		}

		blocks = append(blocks, b)
	}

	col := lipgloss.JoinVertical(lipgloss.Left, blocks...)

	return s.Render(col)
}

// Get a slice of minutes that will be displayed in our component.
// Based on where we start from and how much vertical space we have.
func visibleSlots(startFrom, height int) []int {
	slots := []int{}

	for i := startFrom; i <= 24*60; i += 15 {
		slots = append(slots, i)

		if len(slots) == height {
			break
		}
	}

	return slots
}
