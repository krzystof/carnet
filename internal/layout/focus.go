// package layout groups all the section of the app, and the navigation between them
package layout

import (
	tea "charm.land/bubbletea/v2"
)

type ComponentName int

const (
	SidebarComponent = iota
	HeaderComponent
	TimelineComponent
	TaskListComponent
	EventDetailComponent
)

type FocusChangedMsg struct {
	Comp ComponentName
}

type direction int

const (
	up = iota
	down
	left
	right
)

type NavParams struct {
	comp ComponentName
	dir  direction
}

var nav = map[NavParams]ComponentName{
	{comp: SidebarComponent, dir: right}:    TimelineComponent,
	{comp: TimelineComponent, dir: left}:    SidebarComponent,
	{comp: TimelineComponent, dir: up}:      HeaderComponent,
	{comp: TimelineComponent, dir: right}:   TaskListComponent,
	{comp: HeaderComponent, dir: left}:      SidebarComponent,
	{comp: HeaderComponent, dir: down}:      TimelineComponent,
	{comp: TaskListComponent, dir: left}:    TimelineComponent,
	{comp: TaskListComponent, dir: up}:      HeaderComponent,
	{comp: TaskListComponent, dir: down}:    EventDetailComponent,
	{comp: EventDetailComponent, dir: left}: TimelineComponent,
	{comp: EventDetailComponent, dir: up}:   TaskListComponent,
}
var keyToDir = map[string]direction{
	"ctrl+h": left,
	"ctrl+j": down,
	"ctrl+k": up,
	"ctrl+l": right,
}

func ChangeFocus(current ComponentName, keypress string) tea.Cmd {
	dir := keyToDir[keypress]

	params := NavParams{comp: current, dir: dir}

	comp, ok := nav[params]

	return func() tea.Msg {
		if ok {
			return FocusChangedMsg{Comp: comp}
		} else {
			return nil
		}
	}
}
