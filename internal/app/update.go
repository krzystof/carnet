package app

import (
	"time"

	tea "charm.land/bubbletea/v2"
	"github.com/krzystof/carnet/internal/commands"
	"github.com/krzystof/carnet/internal/layout"
)

func (m Model) Init() tea.Cmd {
	return commands.LoadConfig()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		cmd = layout.CalculateComponentsSizes(msg.Width, msg.Height)
		cmds = append(cmds, cmd)

	case layout.FocusChangedMsg:
		m.activeComponent = msg.Comp

	case commands.ConfigLoadedMsg:
		m.state = stateLoadPage
		m.cfg = &msg.Cfg
		cmd = commands.SelectDate(time.Now())
		cmds = append(cmds, cmd)

	case commands.ConfigNotExistsMsg:
		m.state = stateInitConfig
		m.textInput.Placeholder = "..."
		m.textInput.SetVirtualCursor(false)
		m.textInput.Focus()

	case commands.ConfigLoadFailedMsg:
		m.state = stateError
		m.err = msg.Err

	case commands.PageLoadedMsg:
		m.state = stateReady
		m.page = &msg.Page
		m.selectedDate = msg.Page.Time

	case commands.DateSelectedMsg:
		m.selectedDate = msg.Date

		cmd = commands.LoadPage(m.cfg.UserDataPath, m.selectedDate)
		cmds = append(cmds, cmd)

	case tea.KeyPressMsg:
		switch msg.String() {
		case "enter":
			if m.state == stateInitConfig {
				return m, commands.SetInitConfig(m.textInput.Value())
			}

		case "q":
			// Dont quit when typing "q" on a text input
			if m.state != stateInitConfig {
				return m, tea.Quit
			}

		case "ctrl+c":
			return m, tea.Quit

		case "ctrl+h", "ctrl+j", "ctrl+k", "ctrl+l":
			cmd = layout.ChangeFocus(m.activeComponent, msg.String())
			cmds = append(cmds, cmd)
		}

	}

	// Child updates
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch m.activeComponent {
		case layout.HeaderComponent:
			m.header, cmd = m.header.Update(msg)
			cmds = append(cmds, cmd)
		case layout.SidebarComponent:
			m.monthlyCalendar, cmd = m.monthlyCalendar.Update(msg)
			cmds = append(cmds, cmd)
		case layout.TimelineComponent:
			m.timeline, cmd = m.timeline.Update(msg)
			cmds = append(cmds, cmd)
		case layout.ScheduleComponent:
			m.schedule, cmd = m.schedule.Update(msg)
			cmds = append(cmds, cmd)
		}

	default:
		m.header, cmd = m.header.Update(msg)
		cmds = append(cmds, cmd)
		m.monthlyCalendar, cmd = m.monthlyCalendar.Update(msg)
		cmds = append(cmds, cmd)
		m.timeline, cmd = m.timeline.Update(msg)
		cmds = append(cmds, cmd)
		m.schedule, cmd = m.schedule.Update(msg)
		cmds = append(cmds, cmd)
	}

	// Update submodels
	m.textInput, cmd = m.textInput.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}
