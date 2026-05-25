package app

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/krzystof/carnet/internal/components"
	"github.com/krzystof/carnet/internal/styles"
)

func (m Model) View() tea.View {
	var c *tea.Cursor
	v := tea.NewView("")

	switch m.state {

	case stateStarting:
		v.SetContent("...")

	case stateInitConfig:
		if !m.textInput.VirtualCursor() {
			c = m.textInput.Cursor()
			c.Y += lipgloss.Height(m.headerView())
		}
		v.SetContent(WizardView(m))

	case stateLoadPage:
		v.SetContent("...")

	case stateError:
		v.SetContent("Error:\n\n" + m.err.Error() + "\n\nPress q to quit")

	case stateReady:
		//	Layout:
		//
		//	+---------+-----------------------------------------+
		//	|	sidebar	|	main.header	 														|
		//	|					+-----------------+-----------------------+
		//	|					|	main.left				|	mainRight							|
		//	|					|									|												|
		//	|					|									|												|
		//	|					|									+-----------------------+
		//	|					|									|	main.bottomRight			|
		//	|					|									|												|
		//	+---------+-----------------+-----------------------+
		//

		// Layout dimensions
		sidebarW := 50
		mainW := m.width - 50
		headerH := 20
		mainH := m.height - headerH
		mainColsW := mainW / 2
		mainRightH := mainH * 2 / 3
		mainBottomRightH := mainH - mainRightH

		// Content
		sidebarContent := "TODO: monthlyCalendar.View()"

		mainRight := lipgloss.JoinVertical(
			lipgloss.Left,
			styles.Box("tasks.View", mainColsW, mainRightH, m.activeComponent == components.TaskListComponent),
			styles.Box("event details", mainColsW, mainBottomRightH, m.activeComponent == components.EventDetailComponent),
		)

		mainCols := lipgloss.JoinHorizontal(
			lipgloss.Top,
			styles.Box("timeline", mainColsW, mainH, m.activeComponent == components.TimelineComponent),
			mainRight,
		)

		main := lipgloss.JoinVertical(
			lipgloss.Left,
			styles.Box(m.header.View(), mainW, headerH, m.activeComponent == components.HeaderComponent),
			mainCols,
		)

		ui := lipgloss.JoinHorizontal(
			lipgloss.Top,
			styles.Box(sidebarContent, sidebarW, m.height, m.activeComponent == components.SidebarComponent),
			main,
		)

		v.SetContent(ui)
	}

	if m.err != nil {
		// TODO: render an error overlay or toast or something
		v.SetContent("/nError: " + m.err.Error() + "/n")
	}

	v.Cursor = c
	v.AltScreen = true

	return v
}
