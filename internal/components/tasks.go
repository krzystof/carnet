package components

import "github.com/krzystof/carnet/internal/core"

type Tasks struct {
}

func (t Tasks) View(p *core.Page) string {
	return "## Tasks\n-------\n" + p.RawContent
}
