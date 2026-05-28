package styles

import (
	"charm.land/lipgloss/v2"
)

func Box(w, h int, active bool) lipgloss.Style {
	s := lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		Width(w).
		Height(h).
		Padding(1)

	if active {
		s = s.Border(lipgloss.ThickBorder()).BorderForegroundBlend(
			Theme.BorderActiveColor1,
			Theme.BorderActiveColor2,
		)
	} else {
		s = s.BorderForeground(
			Theme.BorderInactiveColor,
		)
	}

	return s
}
