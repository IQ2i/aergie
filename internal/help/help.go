package help

import "github.com/gookit/color"

// App function return help text template for app
func App() string {
	return color.Sprintf(`<info>Aergie</> version <comment>{{ .Version }}</>

<comment>Usage:</>
  ae [global options] command [options]

<comment>Options:</>
{{range .VisibleFlags}}  <info>{{range .Aliases}}-{{.}}{{end}}{{"\t"}}--{{.Name}}</>{{"\t"}}{{.Usage}}{{ "\n" }}{{end}}{{if .VisibleCommands}}
<comment>Available commands:</>{{range .VisibleCategories}}{{if .Name}}
  <comment>{{.Name}}</>{{range .VisibleCommands}}
	  <info>{{join .Names ", "}}</>{{ "\t"}}{{.Usage}}{{end}}{{else}}{{range .VisibleCommands}}
  <info>{{join .Names ", "}}</>{{ "\t"}}{{.Usage}}{{end}}{{end}}{{end}}
{{end}}
`)
}

// Command function return help text template for a specific command
func Command() string {
	return color.Sprintf(`{{if .Usage}}<comment>Description:</>
  {{.Usage}}

{{end}}<comment>Usage:</>
ae {{.Name}} [options]

<comment>Options:</>
{{range .VisibleFlags}}  <info>{{range .Aliases}}-{{.}}{{end}}{{"\t"}}--{{.Name}}</>{{"\t"}}{{.Usage}}{{ "\n" }}{{end}}
`)
}
