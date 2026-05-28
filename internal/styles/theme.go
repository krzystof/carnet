package styles

import (
	"image/color"

	"charm.land/lipgloss/v2"
)

// Base background color: #24273a

type ThemeDef struct {
	BorderInactiveColor     color.Color
	BorderActiveColor1      color.Color
	BorderActiveColor2      color.Color
	TextDimColor            color.Color
	TextBrightColor         color.Color
	ItemActiveBackground    color.Color
	ItemActiveBackgroundDim color.Color
}

var Theme = ThemeDef{
	BorderInactiveColor:     lipgloss.Color("#414351"),
	BorderActiveColor1:      lipgloss.Color("#c64c1f"),
	BorderActiveColor2:      lipgloss.Color("#e8bd14"),
	TextDimColor:            lipgloss.Color("#898889"),
	TextBrightColor:         lipgloss.Color("#f7f7f7"),
	ItemActiveBackground:    lipgloss.Color("#c64c1f"),
	ItemActiveBackgroundDim: lipgloss.Color("#7f3a21"),
}
