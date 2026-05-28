package components

import (
	"strconv"
	"strings"

	"charm.land/lipgloss/v2"
	"github.com/krzystof/carnet/internal/styles"
)

type Timeline struct {
	cursorStart    int
	cursorDuration int
}

func NewTimeline() Timeline {
	return Timeline{
		cursorStart:    8 * 60,
		cursorDuration: 30,
	}
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

	visibleRows := lipgloss.JoinHorizontal(
		lipgloss.Top,
		hourLabels(slots),
		" clock ",
		hoursBlocks(t, slots, width-18)) // w - space for hours labels, clock, and extra stuff on right

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

func hoursBlocks(t Timeline, slots []int, width int) string {
	activeSlotStyle := lipgloss.NewStyle().Background(styles.Theme.ItemActiveBackgroundDim)

	blocks := []string{}

	for _, minutes := range slots {
		b := strings.Repeat(" ", width)

		// If below cursor:
		if minutes >= t.cursorStart && minutes < (t.cursorStart+t.cursorDuration) {
			b = activeSlotStyle.Render(b)
		}

		blocks = append(blocks, b)
	}

	return lipgloss.JoinVertical(lipgloss.Left, blocks...)
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
