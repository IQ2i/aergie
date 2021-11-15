package completion

import (
	"os"

	"github.com/spf13/cobra"
)

func NewCompletionCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:    "completion [bash|zsh]",
		Short:  "Generate completion script",
		Hidden: true,

		SilenceUsage:  true,
		SilenceErrors: true,

		DisableFlagsInUseLine: true,
		ValidArgs:             []string{"bash", "zsh"},
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
			}

			return nil
		},
	}

	return cmd
}
