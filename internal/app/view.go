package app

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/krzystof/carnet/internal/layout"
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

		// TODO p2 continue moving the sizing logic inside components

		// Layout dimensions
		sidebarW := 50
		mainW := m.width - 50
		headerH := 5
		mainH := m.height - headerH
		mainColsW := mainW / 2
		mainRightH := mainH * 2 / 3
		mainBottomRightH := mainH - mainRightH

		// Content
		sidebarContent := m.monthlyCalendar.View()

		mainRight := lipgloss.JoinVertical(
			lipgloss.Left,
			styles.Box(mainColsW, mainRightH, m.activeComponent == layout.TaskListComponent).Render(m.tasks.View(m.page)),
			styles.Box(mainColsW, mainBottomRightH, m.activeComponent == layout.EventDetailComponent).Render("event details"),
		)

		mainCols := lipgloss.JoinHorizontal(
			lipgloss.Top,
			styles.
				Box(m.timeline.Width, m.timeline.Height, m.activeComponent == layout.TimelineComponent).
				Render(m.timeline.View()),
			mainRight,
		)

		main := lipgloss.JoinVertical(
			lipgloss.Left,
			styles.Box(mainW, headerH, m.activeComponent == layout.HeaderComponent).Render(m.header.View()),
			mainCols,
		)

		ui := lipgloss.JoinHorizontal(
			lipgloss.Top,
			styles.Box(sidebarW, m.height, m.activeComponent == layout.SidebarComponent).Align(lipgloss.Center).Render(sidebarContent),
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
