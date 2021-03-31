package cmd

import (
	"fmt"
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

		Run: func(cmd *cobra.Command, args []string) {
			switch args[0] {
			case "bash":
				err := cmd.Root().GenBashCompletion(os.Stdout)
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
			case "zsh":
				err := cmd.Root().GenZshCompletion(os.Stdout)
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
			case "fish":
				err := cmd.Root().GenFishCompletion(os.Stdout, true)
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
			case "powershell":
				err := cmd.Root().GenPowerShellCompletionWithDesc(os.Stdout)
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
			}
		},
	}

	return cmd
}
