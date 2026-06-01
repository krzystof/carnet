// Package layout contains all top-level related dimensions
package layout

import tea "charm.land/bubbletea/v2"

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

type LayoutSizes struct {
	Width                 int
	Height                int
	SidebarWidth          int
	MainWidth             int
	HeaderHeight          int
	MainHeight            int
	MainColumnsWidth      int
	MainRightHeight       int
	MainBottomRightHeight int
}

type LayoutSizesChangedMsg struct {
	LayoutSizes LayoutSizes
}

func CalculateComponentsSizes(width, height int) tea.Cmd {
	SidebarWidth := 50
	MainWidth := width - 50
	HeaderHeight := 5
	MainHeight := height - HeaderHeight
	MainColumnsWidth := MainWidth / 2
	MainRightHeight := MainHeight * 2 / 3
	MainBottomRightHeight := MainHeight - MainRightHeight

	return func() tea.Msg {
		return LayoutSizesChangedMsg{LayoutSizes{
			Width:                 width,
			Height:                height,
			SidebarWidth:          SidebarWidth,
			MainWidth:             MainWidth,
			HeaderHeight:          HeaderHeight,
			MainHeight:            MainHeight,
			MainColumnsWidth:      MainColumnsWidth,
			MainRightHeight:       MainRightHeight,
			MainBottomRightHeight: MainBottomRightHeight,
		}}
	}
}
