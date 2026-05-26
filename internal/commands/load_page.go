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

func LoadPage(userDataPath string, date time.Time) tea.Cmd {
	return func() tea.Msg {
		todayFile := date.Format(time.DateOnly)
		pagePath := filepath.Join(userDataPath, "dailies", todayFile+".md")
		content, err := os.ReadFile(pagePath)

		var p core.Page

		if errors.Is(err, os.ErrNotExist) {
			// TODO: populate empty page based on tasks, recurring events, etc save it on disk

			p = core.EmptyPage(date, todayFile)

			// TODO do we want to save it here? if its not dirty? Probably not
			// raw := core.PageToString(p)
			// err = os.WriteFile(pagePath, []byte(raw), 0644)
			//
			// if err != nil {
			// 	return err
			// }
		} else {
			// TODO: actually parse string into Page
			p = core.ParseStrToPage(date, todayFile, content)
		}

		return PageLoadedMsg{Page: p}
	}
}
