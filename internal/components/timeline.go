package components

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/krzystof/carnet/internal/core"
	"github.com/krzystof/carnet/internal/layout"
	"github.com/krzystof/carnet/internal/styles"
)

const defaultCursorDuration = 30
const slotDuration = 15

type Timeline struct {
	Width  int
	Height int
	// In minutes since midnight
	displayFrom int
	// In minutes since midnight
	cursorStart    int
	cursorDuration int
}

func NewTimeline() Timeline {
	return Timeline{
		displayFrom:    -1,
		cursorStart:    time.Now().Hour() * 60,
		cursorDuration: defaultCursorDuration,
	}
}

func (t Timeline) Update(msg tea.Msg) (Timeline, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "j":
			// go down
			t.cursorStart = t.cursorStart + t.cursorDuration
			t.cursorDuration = defaultCursorDuration

			if t.cursorStart+t.cursorDuration >= 24*60 {
				t.cursorStart = 24*60 - defaultCursorDuration
			}

			cursorEnd := t.cursorStart + t.cursorDuration
			if cursorEnd >= t.maxVisibleSlot() {
				t.displayFrom += t.cursorDuration
			}

			// TODO 2 - what if an event?

		case "k":
			// go up
			t.cursorStart = max(t.cursorStart-defaultCursorDuration, 0)
			t.cursorDuration = defaultCursorDuration

			// If you go up and the new start is not display from, update
			if t.cursorStart < t.displayFrom {
				t.displayFrom = t.cursorStart
			}

			// TODO 2 - what if an event?
		}

	case layout.LayoutSizesChangedMsg:
		t.Width = msg.LayoutSizes.MainColumnsWidth
		t.Height = msg.LayoutSizes.MainHeight

		visibleRowsCount := t.Height - 8
		t.displayFrom = calcStartFrom(t.displayFrom, visibleRowsCount)
	}

	return t, cmd
}

func (t Timeline) View(page *core.Page) string {
	slots := visibleSlots(t.displayFrom, t.Height-8) // 2 borders, 2 padding, 2 labels, 2 borders

	visibleRows := lipgloss.JoinHorizontal(
		lipgloss.Top,
		hourLabels(slots),
		" clock ",
		eventsSlots(t, slots, t.Width-18, page)) // w - space for hours labels, clock, and extra stuff on right

	s := lipgloss.NewStyle().
		Width(t.Width-4). // (got to remove those borders and padding)
		BorderForeground(styles.Theme.BorderInactiveColor).
		Border(lipgloss.NormalBorder(), true, false)

	return lipgloss.JoinVertical(
		lipgloss.Left,
		"Press k to move up",
		s.Render(visibleRows),
		("Press j to move down" + " | max visible slot: " + strconv.Itoa(t.maxVisibleSlot())),
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

func eventsSlots(t Timeline, slots []int, width int, page *core.Page) string {
	blocks := []string{}
	events := page.GetEventPerSlots(slotDuration)

	for _, minutes := range slots {
		s := lipgloss.NewStyle()
		b := strings.Repeat(" ", width)

		e, ok := events[minutes]

		fgColor := styles.Theme.BorderActiveColor1
		bgColor := styles.Theme.ItemActiveBackgroundDim

		if ok {
			fgColor = styles.GetCategoryColor(e.Category, "dark")
			bgColor = styles.GetCategoryColor(e.Category, "pastel")

			if e.StartTime == minutes {
				b = fmt.Sprintf("%s %s", e.Category, e.Title)
				b += strings.Repeat(" ", width-len(b))
			}

			s = s.BorderLeft(true).
				BorderStyle(lipgloss.NormalBorder()).
				BorderForeground(fgColor).
				BorderBackground(bgColor).
				Background(bgColor).
				PaddingLeft(1)
		}

		// TODO 1
		// if cursor is first of an event, render the event times, category and title
		// sepcial styles too
		// if still within an event, just border left

		// TODO 3
		// if event selected

		// If below cursor:
		if minutes >= t.cursorStart && minutes < (t.cursorStart+t.cursorDuration) {
			s = s.
				Background(bgColor).
				BorderBackground(bgColor)
		}

		blocks = append(blocks, s.Render(b))
	}

	return lipgloss.JoinVertical(lipgloss.Left, blocks...)
}

// Get a slice of minutes that will be displayed in our component.
// Based on where we start from and how much vertical space we have.
func visibleSlots(startFrom, height int) []int {
	slots := []int{}

	for i := startFrom; i <= 24*60; i += slotDuration {
		slots = append(slots, i)

		if len(slots) == height {
			break
		}
	}

	return slots
}

// Returns a value in minutes
func calcStartFrom(storedDisplayFrom, rowCount int) int {
	if storedDisplayFrom != -1 {
		return storedDisplayFrom
	}

	hourCount := rowCount / 4

	// We have enough real estate to display everything
	if hourCount >= 24 {
		return 0
	}

	// else calculate the most centered window we can by leaving same space before / after
	hoursOffset := ((24 - hourCount) / 2)

	// Convert to blocks
	return 60 * hoursOffset
}

func (t Timeline) maxVisibleSlot() int {
	visibleRowsCount := t.Height - 8
	return t.displayFrom + visibleRowsCount*slotDuration
}
