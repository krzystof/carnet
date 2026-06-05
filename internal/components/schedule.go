package components

import (
	"fmt"
	"strings"
	"time"

	"github.com/krzystof/carnet/internal/commands"
	"github.com/krzystof/carnet/internal/core"
	"github.com/krzystof/carnet/internal/layout"
	"github.com/krzystof/carnet/internal/styles"
	"github.com/krzystof/carnet/internal/utils"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

// Schedule is a list of events for a given day

type Schedule struct {
	Width         int
	Height        int
	page          *core.Page
	selectedEvent *core.Event
}

func NewSchedule() Schedule {
	return Schedule{}
}

func (s Schedule) Update(msg tea.Msg) (Schedule, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "j":
			// go down
			if s.selectedEvent == nil && len(s.page.Events) > 0 {
				s.selectedEvent = s.page.Events[0]
				break
			}

			idx := utils.IndexOf(s.page.Events, func(e *core.Event) bool {
				return e.Equal(s.selectedEvent)
			})

			if idx >= 0 && (idx+1) < len(s.page.Events) {
				s.selectedEvent = s.page.Events[idx+1]
			}

		case "k":
			if s.selectedEvent == nil && len(s.page.Events) > 0 {
				s.selectedEvent = s.page.Events[0]
				break
			}

			idx := utils.IndexOf(s.page.Events, func(e *core.Event) bool {
				return e.Equal(s.selectedEvent)
			})

			if (idx - 1) >= 0 {
				s.selectedEvent = s.page.Events[idx-1]
			}
		}

	case commands.PageLoadedMsg:
		s.page = &msg.Page

	case layout.LayoutSizesChangedMsg:
		s.Width = msg.LayoutSizes.MainColumnsWidth
		s.Height = msg.LayoutSizes.MainHeight
	}

	return s, cmd
}

func (s Schedule) View() string {
	if len(s.page.Events) == 0 {
		st := lipgloss.NewStyle().
			Width(s.Width).
			Height(s.Height - 4).
			Align(lipgloss.Center).
			AlignVertical(lipgloss.Center)

		return st.Render("Nothing planned")
	}

	boxes := []string{}
	lastEndTime := 0
	eventWidth := s.Width - 8

	t := time.Now()
	clock := ""
	if s.page.IsToday() {
		clock = renderClock(t)
	}
	clockLength := lipgloss.Width(clock)

	for _, e := range s.page.Events {
		endTime := e.StartTime + e.DurationMin

		if lastEndTime != 0 && e.StartTime != lastEndTime {
			boxes = append(boxes, " ")
		}

		isSelected := false
		if s.selectedEvent != nil {
			isSelected = e.Equal(s.selectedEvent)
		}

		b := renderEventV2(e, eventWidth-clockLength, isSelected)
		boxes = append(boxes, b)

		lastEndTime = endTime
	}

	boxesCol := lipgloss.JoinVertical(
		lipgloss.Left,
		boxes...,
	)

	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		lipgloss.PlaceVertical(calcOffset(s.page, t), lipgloss.Bottom, clock),
		lipgloss.NewStyle().
			Width(s.Width-4-clockLength). // remove border and padding
			Align(lipgloss.Center).
			Render(boxesCol),
	)
}

func formatClock(minutes int) string {
	return fmt.Sprintf("%2d:%02d", minutes/60, minutes%60)
}

const debugClock = false

func renderEventV2(e *core.Event, width int, isSelected bool) string {
	slotsCount := max(e.DurationMin/15, 2)
	fgColor := styles.GetCategoryColor(e.Category, "bright")
	bgColor := styles.GetCategoryColor(e.Category, "dim")

	s := lipgloss.NewStyle().
		Width(width).
		BorderForeground(bgColor).
		Border(lipgloss.ThickBorder())

	blockStyle := lipgloss.NewStyle().Foreground(bgColor)

	if isSelected {
		s = s.BorderForeground(fgColor)
		blockStyle = blockStyle.Foreground(fgColor)
	}

	coloredBlock := blockStyle.Render("▒")

	lines := fmt.Sprintf("%s %s    %s %s",
		coloredBlock,
		formatClock(e.StartTime),
		formatCategoryTag(e.Category),
		(e.Title),
	)

	// add empty lines
	if debugClock {
		for i := e.StartTime + 15; i < e.StartTime+e.DurationMin; i += 15 {
			lines += fmt.Sprintf("\n%s %s", coloredBlock, formatClock(i))
		}
	} else {
		lines += strings.Repeat("\n"+coloredBlock, slotsCount-1)
		lines += fmt.Sprintf(" %s", formatClock(e.StartTime+e.DurationMin))
	}

	return s.Render(lines)
}

func formatCategoryTag(c string) string {
	if c == "" {
		return c
	}

	bgColor := styles.GetCategoryColor(c, "dim")

	s := lipgloss.NewStyle().Background(bgColor).Padding(0, 1)

	return s.Render(c)
}

// Returns by how much the given time need to be offset on the schedule
func calcOffset(p *core.Page, t time.Time) int {
	offset := 1
	h, m, _ := t.Clock()
	timeInMinutes := h*60 + m

	lastEndTime := 0

	for _, e := range p.Events {
		if e.StartTime > timeInMinutes {
			break
		}

		// if there is a gap between events, add +1
		if lastEndTime != 0 && lastEndTime < e.StartTime {
			offset += 1
		}

		offset += 1 // border before
		endTime := e.StartTime + e.DurationMin
		addTime := (timeInMinutes - e.StartTime)

		if timeInMinutes > endTime {
			addTime = e.DurationMin
			offset += 1 // border after
		}

		offset += addTime / 15

		// if the event contains the clock, then calc relative offset
		lastEndTime = endTime
	}

	return offset
}

func renderClock(t time.Time) string {
	s := lipgloss.NewStyle().Background(styles.Theme.ItemActiveBackground)
	clock := " " + t.Format("15:04") + " ▶ "
	return s.Render(clock)
}
