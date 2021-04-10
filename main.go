package main

import (
	"github.com/iq2i/aergie/cmd"
	"github.com/iq2i/aergie/internal/config"
)

var version = "DEV"

func main() {
	config.InitAppConfig()
	config.InitUserConfig()

	cmd.Execute(version)
}
