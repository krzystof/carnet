package components

import (
	"strconv"
	"strings"

	"github.com/krzystof/carnet/internal/core"
)

type Tasks struct {
}

func (t Tasks) View(p *core.Page) string {
	// return "## Tasks\n-------\n" + p.RawContent

	parts := []string{}

	for _, v := range p.Events {
		e := "Start=" + strconv.Itoa(v.StartTime) + " | Duration=" + strconv.Itoa(v.DurationMin) + " | Cat=" + v.Category + " | Title=" + v.Title

		parts = append(parts, e)
	}

	return strings.Join(parts, "\n")
}
