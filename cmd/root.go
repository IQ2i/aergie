package cmd

import (
	"fmt"
	"os"

	"github.com/iq2i/aergie/internal/cmd/root"
	"github.com/spf13/cobra"
)

func init() {
	cobra.EnableCommandSorting = false
}

// Execute is the real main function of Aergie cli
func Execute(version string) {
	var rootCmd = &cobra.Command{
		Use:     "ae <command> [flags]",
		Short:   "Aergie CLI",
		Long:    "An easy alternative to makefile",
		Version: version,

		SilenceUsage:  true,
		SilenceErrors: true,

		Args: cobra.NoArgs,
		ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			list := make([]string, 0)
			for _, cmd := range cmd.Commands() {
				if !cmd.IsAvailableCommand() {
					continue
				}

				list = append(list, fmt.Sprintf("%s\t%s", cmd.Use, cmd.Short))
			}
			return list, cobra.ShellCompDirectiveNoSpace
		},
	}

	rootCmd.SetHelpFunc(root.HelpFunc)
	rootCmd.SetHelpCommand(&cobra.Command{Hidden: true})

	rootCmd.AddCommand(newCompletionCommand())
	rootCmd.AddCommand(newUserCommands()...)

	rootCmd.PersistentFlags().Bool("help", false, "Show help for command")
	rootCmd.PersistentFlags().Bool("version", false, "Show ae version")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
