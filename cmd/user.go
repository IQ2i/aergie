package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"syscall"

	"github.com/iq2i/aergie/internal/config"
	"github.com/kballard/go-shellquote"
	"github.com/spf13/cobra"
)

func newUserCommands() []*cobra.Command {
	var cmds = make([]*cobra.Command, 0)

	for _, configCmd := range config.AppConfig.Commands {
		steps := configCmd.Steps
		userCmd := &cobra.Command{
			Use:   configCmd.Name,
			Short: configCmd.Help,

			SilenceUsage:  true,
			SilenceErrors: true,

			Args: cobra.NoArgs,
			RunE: func(c *cobra.Command, args []string) error {
				for _, step := range steps {
					args, err := shellquote.Split(step)
					if err != nil {
						return fmt.Errorf("Invalid instruction \"%s\"", step)
					}

					cmd := exec.Command(args[0], args[1:]...)
					cmd.Stdout = os.Stdout
					cmd.Stdin = os.Stdin
					cmd.Stderr = os.Stderr

					// In case of interrupt signal, stop all command, not only current step
					signals := make(chan os.Signal, 1)
					signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
					go func() {
						<-signals
						os.Exit(0)
					}()

					if err := cmd.Run(); err != nil {
						// catch interrupt in subprocess
						if exitError, ok := err.(*exec.ExitError); ok {
							if exitError.ExitCode() == 130 {
								os.Exit(0)
							}
						}

						return fmt.Errorf("Step \"%s\" failed", step)
					}
				}

				return nil
			},
		}
		cmds = append(cmds, userCmd)
	}

	return cmds
}
