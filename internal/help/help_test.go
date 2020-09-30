package help

import (
	"testing"

	"github.com/gookit/color"
	"github.com/stretchr/testify/assert"
)

func TestApp(t *testing.T) {
	color.ForceOpenColor()
	is := assert.New(t)

	expected := "\x1b[0;32mAergie\x1b[0m version \x1b[0;33m{{ .Version }}\x1b[0m\n\n\x1b[0;33mUsage:\x1b[0m\n  ae [global options] command [options]\n\n\x1b[0;33mOptions:\x1b[0m\n{{range .VisibleFlags}}  \x1b[0;32m{{range .Aliases}}-{{.}}{{end}}{{\"\\t\"}}--{{.Name}}\x1b[0m{{\"\\t\"}}{{.Usage}}{{ \"\\n\" }}{{end}}{{if .VisibleCommands}}\n\x1b[0;33mAvailable commands:\x1b[0m{{range .VisibleCategories}}{{if .Name}}\n  \x1b[0;33m{{.Name}}\x1b[0m{{range .VisibleCommands}}\n\t  \x1b[0;32m{{join .Names \", \"}}\x1b[0m{{ \"\\t\"}}{{.Usage}}{{end}}{{else}}{{range .VisibleCommands}}\n  \x1b[0;32m{{join .Names \", \"}}\x1b[0m{{ \"\\t\"}}{{.Usage}}{{end}}{{end}}{{end}}\n{{end}}\n"
	is.Equal(expected, App())
}

func TestStep(t *testing.T) {
	color.ForceOpenColor()
	is := assert.New(t)

	expected := "{{if .Usage}}\x1b[0;33mDescription:\x1b[0m\n  {{.Usage}}\n\n{{end}}\x1b[0;33mUsage:\x1b[0m\nae {{.Name}} [options]\n\n\x1b[0;33mOptions:\x1b[0m\n{{range .VisibleFlags}}  \x1b[0;32m{{range .Aliases}}-{{.}}{{end}}{{\"\\t\"}}--{{.Name}}\x1b[0m{{\"\\t\"}}{{.Usage}}{{ \"\\n\" }}{{end}}\n"
	is.Equal(expected, Command())
}
