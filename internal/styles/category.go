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

var bright = []string{
	"#DB1D18", // red
	"#22B313", // green
	"#E6DA3F", // yellow
	"#F24700", // orange
	"#2151C0", // blue
	"#24B32E", // green
	"#CD629F", // pink
	"#7D3E93", // purple
}

var neutral = []string{
	"#AA3939", // red
	"#38912F", // green
	"#B6AD37", // yellow
	"#D86800", // orange
	"#0D3AA2", // blue
	"#1D9124", // green
	"#B34D81", // pink
	"#692682", // purple
}

var dim = []string{
	"#801515", // red
	"#084F01", // green
	"#938A19", // yellow
	"#B05400", // orange
	"#092E82", // blue
	"#0B7712", // green
	"#A21F69", // pink
	"#531469", // purple
}

type Palette = []string

var palettes = map[string]Palette{
	"pastel":  pastel,
	"dark":    dark,
	"bright":  bright,
	"neutral": neutral,
	"dim":     dim,
}

var grey = map[string]string{
	"bright":  "#808080",
	"neutral": "#676767",
	"dim":     "#5A5A5A",
}

func GetCategoryColor(categoryName, palette string) color.Color {
	if categoryName == "" {
		return lipgloss.Color(grey[palette])
	}

	h := fnv.New32a()
	h.Write([]byte(categoryName))

	colors := palettes[palette]
	idx := h.Sum32() % uint32(len(colors))

	return lipgloss.Color(colors[int(idx)])
}
