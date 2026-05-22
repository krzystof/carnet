package commands

import (
	"errors"
	"os"
	"path/filepath"
	"time"

	tea "charm.land/bubbletea/v2"
	"github.com/krzystof/carnet/internal/core"
)

type PageLoadedMsg struct {
	Page core.Page
}

func LoadTodaysPage(userDataPath string) tea.Cmd {
	return func() tea.Msg {
		pageTime := time.Now()
		todayFile := pageTime.Format(time.DateOnly)

		pagePath := filepath.Join(userDataPath, "dailies", todayFile+".md")

		content, err := os.ReadFile(pagePath)

		if errors.Is(err, os.ErrNotExist) {
			// TODO: populate empty page based on tasks, recurring events, etc save it on disk

			p := core.EmptyPage(pageTime, todayFile)

			err = os.WriteFile(pagePath, templateContent(todayFile), 0644)

			if err != nil {
				return err
			}

			return PageLoadedMsg{Page: p}
		}

		// TODO: actually parse string into Page
		p := core.ParseStrToPage(pageTime, todayFile, content)

		return PageLoadedMsg{Page: p}
	}
}

func templateContent(todayDate string) []byte {
	c := "# " + todayDate

	return []byte(c)
}
