package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

func newCompletionCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:    "completion [bash|zsh|fish|powershell]",
		Short:  "Generate completion script",
		Hidden: true,

		SilenceUsage:  true,
		SilenceErrors: true,

		DisableFlagsInUseLine: true,
		ValidArgs:             []string{"bash", "zsh", "fish", "powershell"},
		Args:                  cobra.ExactValidArgs(1),

		RunE: func(cmd *cobra.Command, args []string) error {
			switch args[0] {
			case "bash":
				err := cmd.Root().GenBashCompletion(os.Stdout)
				if err != nil {
					return err
				}
			case "zsh":
				err := cmd.Root().GenZshCompletion(os.Stdout)
				if err != nil {
					return err
				}
			case "fish":
				err := cmd.Root().GenFishCompletion(os.Stdout, true)
				if err != nil {
					return err
				}
			case "powershell":
				err := cmd.Root().GenPowerShellCompletionWithDesc(os.Stdout)
				if err != nil {
					return err
				}
			}

			return nil
		},
	}

	return cmd
}
