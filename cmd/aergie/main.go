package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gookit/color"
	"github.com/iq2i/aergie/internal/command"
	"github.com/iq2i/aergie/internal/config"
	"github.com/iq2i/aergie/internal/help"
	"github.com/iq2i/aergie/internal/logger"
	"github.com/urfave/cli/v2"
)

// init with -ldflags option during run or build
var version string

func init() {
	// init config
	err := config.Init()
	if err != nil {
		logger.Error(err)
		os.Exit(1)
	}

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

	cli.AppHelpTemplate = help.App()
	cli.CommandHelpTemplate = help.Command()
	cli.VersionPrinter = func(c *cli.Context) {
		color.Fprintf(c.App.Writer, "<info>Aergie cli</> version <comment>%s</>\n", c.App.Version)
	}
}

func main() {
	app := cli.App{
		HelpName:             "ae",
		HideHelpCommand:      true,
		EnableBashCompletion: true,
		Version:              version,
		Compiled:             time.Now(),
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
