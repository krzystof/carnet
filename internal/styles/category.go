package styles

import (
	"hash/fnv"
	"image/color"

	"charm.land/lipgloss/v2"
)

var pastel = []string{
	"#8C5E7A", // dusty rose
	"#8B6D5C", // muted peach
	"#8A7A52", // muted gold
	"#5F7F5F", // sage green
	"#4F7A72", // seafoam
	"#4F6F8C", // dusty blue
	"#6B6FA3", // periwinkle
	"#7B5FA8", // lavender
	"#8A5F8F", // mauve
	"#5F6B8A", // slate blue
	"#6F7A5A", // olive sage
	"#7A647A", // muted plum
}

var dark = []string{
	"#5C3D4A", // dusty rose
	"#5A473D", // muted brown
	"#5A5338", // olive gold
	"#3F5A46", // sage
	"#355A55", // teal
	"#3A4F63", // dusty blue
	"#484D73", // periwinkle
	"#58457A", // lavender
	"#5E4560", // mauve
	"#434D63", // slate
	"#4C5A3F", // moss
	"#56465A", // plum
}

type Palette = []string

var palettes = map[string]Palette{
	"pastel": pastel,
	"dark":   dark,
}

func GetCategoryColor(categoryName, palette string) color.Color {
	h := fnv.New32a()
	h.Write([]byte(categoryName))

	colors := palettes[palette]
	idx := h.Sum32() % uint32(len(colors))

	return lipgloss.Color(colors[int(idx)])
}
