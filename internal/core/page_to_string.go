package core

import (
	"strings"
)

const template = `# <date>

## Timeline
`

func PageToString(p Page) string {
	content := strings.Replace(template, "<date>", p.Time.Format("Mon 02.01.2006"), 1)

	return content
}
