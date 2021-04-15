package main

import (
	"os"

	"github.com/iq2i/aergie/cmd"
	"github.com/iq2i/aergie/internal/config"
)

var version = "DEV"

func main() {
	os.Setenv("AE_VERSION", version)

	config.InitAppConfig()
	config.InitUserConfig()

	cmd.Execute()
}
