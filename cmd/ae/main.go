package main

import (
	"fmt"
	"io"
	"os"

	"github.com/gookit/color"
	"github.com/iq2i/aergie/internal/build"
	"github.com/iq2i/aergie/internal/command"
	"github.com/iq2i/aergie/internal/config"
	"github.com/iq2i/aergie/internal/logger"
	"github.com/iq2i/aergie/internal/tpl"
	"github.com/urfave/cli/v2"
)

func init() {
	// init config
	config.Init()

	// init commands
	command.Init()

	// init cli
	cli.HelpFlag = &cli.BoolFlag{
		Name:    "help",
		Aliases: []string{"h"},
		Usage:   "Print this help message",
	}
	cli.VersionFlag = &cli.BoolFlag{
		Name:  "version",
		Usage: "Print this application version",
	}
	cli.AppHelpTemplate = tpl.AppHelpTemplate
	cli.CommandHelpTemplate = tpl.CommandHelpTemplate
	cli.HelpPrinter = func(w io.Writer, templ string, data interface{}) {
		cli.HelpPrinterCustom(w, color.Sprintf(templ), data, nil)
	}
	cli.VersionPrinter = func(c *cli.Context) {
		color.Fprintf(c.App.Writer, tpl.VersionTemplate, c.App.Version)
	}
}

func main() {
	app := cli.App{
		HelpName:             "ae",
		HideHelpCommand:      true,
		EnableBashCompletion: true,
		Version:              build.Version,
		Compiled:             build.Date,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "verbose",
				Aliases: []string{"v"},
				Usage:   "Increase the verbosity of messages",
			},
			&cli.BoolFlag{
				Name:    "quiet",
				Aliases: []string{"q"},
				Usage:   "Do not output message except error messages",
			},
		},
		CommandNotFound: func(c *cli.Context, command string) {
			logger.Error(fmt.Errorf("Command \"%s\" is not defined.", command))
		},
		Commands: command.AppCommands,
	}

	err := app.Run(os.Args)
	if err != nil {
		logger.Error(err)
		os.Exit(1)
	}
}
