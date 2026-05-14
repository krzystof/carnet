package main

import (
	tea "charm.land/bubbletea/v2"
	app "github.com/krzystof/carnet/internal/app"
)

func main() {
	model := app.InitialModel()

	program := tea.NewProgram(model)

	program.Run()
}
