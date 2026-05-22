package styles

import "charm.land/lipgloss/v2"

func Box(content string, w, h int) string {
	s := lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		Width(w).
		Height(h).
		Padding(1)

	return s.Render(content)
}
