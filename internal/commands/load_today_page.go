package commands

import (
	"errors"
	"os"
	"path/filepath"
	"time"

	tea "charm.land/bubbletea/v2"
)

type PageLoadedMsg struct {
	Page string // TODO should be core.Page, once I implement the parser
}

func LoadTodaysPage(userDataPath string) tea.Cmd {
	return func() tea.Msg {
		todayFile := time.Now().Format(time.DateOnly)

		pagePath := filepath.Join(userDataPath, "dailies", todayFile+".md")

		content, err := os.ReadFile(pagePath)

		if errors.Is(err, os.ErrNotExist) {
			// TODO:
			// initializeEmptyPage
			// populate empty page based on tasks, recurring events, etc
			// save it on disk
			content := "New one"
			return PageLoadedMsg{Page: "Page loaded:" + pagePath + " content:" + content}

		} else {
			// TODO:
			// parse the content into Page
		}

		return PageLoadedMsg{Page: string(content)}
	}
}
