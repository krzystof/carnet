package styles

// Base background color: #24273a

type ThemeDef struct {
	BorderInactiveColor  string
	BorderActiveColor1   string
	BorderActiveColor2   string
	TextDimColor         string
	TextBrightColor      string
	ItemActiveBackground string
}

var Theme = ThemeDef{
	BorderInactiveColor:  "#414351",
	BorderActiveColor1:   "#c64c1f",
	BorderActiveColor2:   "#e8bd14",
	TextDimColor:         "#898889",
	TextBrightColor:      "#f7f7f7",
	ItemActiveBackground: "#c64c1f",
}
