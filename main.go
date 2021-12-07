package main

import (
	"os"

	"github.com/iq2i/aergie/cmd"
	"github.com/iq2i/aergie/internal/config"
)

var version = "DEV"

func main() {
	config.LoadEnv(".aergie", os.Getenv("AERGIE_ENV"))
	cmd.Execute(version)
}
