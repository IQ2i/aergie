package command

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"syscall"

	"github.com/iq2i/aergie/internal/config"
	"github.com/iq2i/aergie/internal/logger"
	"github.com/kballard/go-shellquote"
	"github.com/urfave/cli/v2"
)

// AppCommands is a slice with all commands
var AppCommands []*cli.Command

// Init is a function to init application commands
func Init() {
	for _, command := range config.AppConfig.Commands {
		steps := command.Steps
		cmd := &cli.Command{
			Name:  command.Name,
			Usage: command.Help,
			Action: func(c *cli.Context) error {
				for _, step := range steps {
					exe(c, step)
				}
				return nil
			},
		}

		AppCommands = append(AppCommands, cmd)
	}
}

func exe(c *cli.Context, step string) {
	if !c.Bool("quiet") && c.Bool("verbose") {
		logger.Step(step)
	}

	args, err := shellquote.Split(step)
	if err != nil {
		logger.Error(fmt.Errorf("Invalid instruction \"%s\"", step))
		os.Exit(1)
	}

	cmd := exec.Command(args[0], args[1:]...)
	if !c.Bool("quiet") {
		cmd.Stdout = os.Stdout
	}
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	go func() {
		<-ch
		os.Exit(0)
	}()

	err = cmd.Run()
	if err != nil {
		// catch interrupt in subprocess
		if exitError, ok := err.(*exec.ExitError); ok {
			if exitError.ExitCode() == 130 {
				os.Exit(0)
			}
		}

		logger.Error(fmt.Errorf("Step \"%s\" failed", step))
		os.Exit(1)
	}
}
