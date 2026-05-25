package app

import "github.com/krzystof/carnet/internal/components"

type direction int

const (
	up = iota
	down
	left
	right
)

type NavParams struct {
	comp components.ComponentName
	dir  direction
}

var keyToDir = map[string]direction{
	"ctrl+h": left,
	"ctrl+j": down,
	"ctrl+k": up,
	"ctrl+l": right,
}

var nav = map[NavParams]components.ComponentName{
	{comp: components.SidebarComponent, dir: right}: components.TimelineComponent,

	{comp: components.TimelineComponent, dir: left}:  components.SidebarComponent,
	{comp: components.TimelineComponent, dir: up}:    components.HeaderComponent,
	{comp: components.TimelineComponent, dir: right}: components.TaskListComponent,

	{comp: components.HeaderComponent, dir: left}: components.SidebarComponent,
	{comp: components.HeaderComponent, dir: down}: components.TimelineComponent,

	{comp: components.TaskListComponent, dir: left}: components.TimelineComponent,
	{comp: components.TaskListComponent, dir: up}:   components.HeaderComponent,
	{comp: components.TaskListComponent, dir: down}: components.EventDetailComponent,

	{comp: components.EventDetailComponent, dir: left}: components.TimelineComponent,
	{comp: components.EventDetailComponent, dir: up}:   components.TaskListComponent,
}

func navigateToComponent(current components.ComponentName, key string) components.ComponentName {
	dir := keyToDir[key]

	params := NavParams{comp: current, dir: dir}

	newComp, ok := nav[params]

	if ok {
		return newComp
	} else {
		return current
	}
}
