package tpl

// AppHelpTemplate is the text template for the Default help topic.
var AppHelpTemplate = `<info>Aergie</> version <comment>{{ .Version }}</> {{ .Compiled.Format "2006-01-02 15:04:05" }}

<comment>Usage:</>
  ae [global options] command [options]

<comment>Options:</>
{{range .VisibleFlags}}  <info>{{range .Aliases}}-{{.}}{{end}}{{"\t"}}--{{.Name}}</>{{"\t"}}{{.Usage}}{{ "\n" }}{{end}}{{if .VisibleCommands}}
<comment>Available commands:</>{{range .VisibleCategories}}{{if .Name}}
  <comment>{{.Name}}</>{{range .VisibleCommands}}
	  <info>{{join .Names ", "}}</>{{ "\t"}}{{.Usage}}{{end}}{{else}}{{range .VisibleCommands}}
  <info>{{join .Names ", "}}</>{{ "\t"}}{{.Usage}}{{end}}{{end}}{{end}}
{{end}}
`

// CommandHelpTemplate is the text template for the command help topic.
var CommandHelpTemplate = `{{if .Usage}}<comment>Description:</>
  {{.Usage}}

{{end}}<comment>Usage:</>
  ae {{.Name}} [options]

<comment>Options:</>
{{range .VisibleFlags}}  <info>{{range .Aliases}}-{{.}}{{end}}{{"\t"}}--{{.Name}}</>{{"\t"}}{{.Usage}}{{ "\n" }}{{end}}
`

// VersionTemplate is the text template for version command.
var VersionTemplate = "<info>Aergie cli</> version <comment>%s</>\n"
