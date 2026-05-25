// Package components group independent components that
// will be grouped together in the main View() layout.
package components

type ComponentName int

const (
	SidebarComponent = iota
	HeaderComponent
	TimelineComponent
	TaskListComponent
	EventDetailComponent
)
