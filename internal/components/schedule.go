package components

import (
	"fmt"
	"strings"

	"github.com/krzystof/carnet/internal/core"
	"github.com/krzystof/carnet/internal/layout"
	"github.com/krzystof/carnet/internal/styles"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

// Schedule is a list of events for a given day

type Schedule struct {
	Width  int
	Height int
}

func NewSchedule() Schedule {
	return Schedule{}
}

func (t Schedule) Update(msg tea.Msg) (Schedule, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	// case tea.KeyPressMsg:
	// 	switch msg.String() {
	// 	case "j":
	// 		// go down
	// 		t.cursorStart = t.cursorStart + t.cursorDuration
	// 		t.cursorDuration = defaultCursorDuration
	//
	// 		if t.cursorStart+t.cursorDuration >= 24*60 {
	// 			t.cursorStart = 24*60 - defaultCursorDuration
	// 		}
	//
	// 		cursorEnd := t.cursorStart + t.cursorDuration
	// 		if cursorEnd >= t.maxVisibleSlot() {
	// 			t.displayFrom += t.cursorDuration
	// 		}
	//
	// 		// TODO 2 - what if an event?
	//
	// 	case "k":
	// 		// go up
	// 		t.cursorStart = max(t.cursorStart-defaultCursorDuration, 0)
	// 		t.cursorDuration = defaultCursorDuration
	//
	// 		// If you go up and the new start is not display from, update
	// 		if t.cursorStart < t.displayFrom {
	// 			t.displayFrom = t.cursorStart
	// 		}
	//
	// 		// TODO 2 - what if an event?
	// 	}

	case layout.LayoutSizesChangedMsg:
		t.Width = msg.LayoutSizes.MainColumnsWidth
		t.Height = msg.LayoutSizes.MainHeight
	}

	return t, cmd
}

func (s Schedule) View(p *core.Page) string {
	if len(p.Events) == 0 {
		st := lipgloss.NewStyle().
			Width(s.Width).
			Height(s.Height - 4).
			Align(lipgloss.Center).
			AlignVertical(lipgloss.Center)

		return st.Render("Nothing planned")
	}

	boxes := []string{}

	for _, e := range p.Events {
		// b := renderEvent(s.Width, e)
		b := renderEventV2(s.Width, e)
		boxes = append(boxes, b)
	}

	return lipgloss.JoinVertical(
		lipgloss.Left,
		boxes...,
	)
}

func formatClock(minutes int) string {
	return fmt.Sprintf("%2d:%02d", minutes/60, minutes%60)
}

func renderEvent(width int, e *core.Event) string {
	s := lipgloss.NewStyle().
		Width(width-8).
		BorderForeground(styles.Theme.BorderInactiveColor).
		Border(lipgloss.NormalBorder()).
		Padding(0, 2)

	b := fmt.Sprintf("%s - %s %s %s",
		formatClock(e.StartTime),
		formatClock(e.StartTime+e.DurationMin),
		(e.Category),
		(e.Title),
	)

	return s.Render(b)
}

func renderEventV2(width int, e *core.Event) string {
	// height := (e.DurationMin / 15) + 2 // add the borders

	slotsCount := min(e.DurationMin/15, 2)
	fgColor := styles.GetCategoryColor(e.Category, "dark")

	s := lipgloss.NewStyle().
		Width(width-8).
		// Height(height).
		BorderForeground(fgColor).
		Border(lipgloss.ThickBorder()).
		Padding(0, 2)

	lines := fmt.Sprintf("%s    %s %s",
		formatClock(e.StartTime),
		formatCategoryTag(e.Category),
		(e.Title),
	)

	// add empty lines
	lines += strings.Repeat("\n", slotsCount-2)

	lines += "\n" + formatClock(e.StartTime+e.DurationMin)

	return s.Render(lines)
}

func formatCategoryTag(c string) string {
	if c == "" {
		return c
	}

	bgColor := styles.GetCategoryColor(c, "dark")

	s := lipgloss.NewStyle().Background(bgColor).Padding(0, 1)

	return s.Render(c)
}
