package app

import (
	"strings"

	"charm.land/lipgloss/v2"
)

func WizardView(m Model) string {
	str := lipgloss.JoinVertical(lipgloss.Top, m.headerView(), m.textInput.View(), m.footerView())

	return str
}

func (m Model) headerView() string {
	content := []string{
		"Welcome to Carnet! Looks like there is no config setup yet.",
		"",
		"Where should we store your files? (leave empty to use the default path)",
	}

	return strings.Join(content, "\n")
}

func (m Model) footerView() string {
	return "\n(ctrl+c to quit)"
}
