package main

import (
	"fmt"

	"github.com/iq2i/aergie/cmd"
	"github.com/iq2i/aergie/internal/config"
)

var (
	version = "dev"
	date    = ""
)

func main() {
	config.Init()
	cmd.Execute(
		buildVersion(version, date),
	)
}

func buildVersion(version string, date string) string {
	var result = version
	if date != "" {
		result = fmt.Sprintf("%s\nbuilt at: %s", result, date)
	}
	return result
}
