package cmd

import (
	"fmt"
	"os"

	completionCmd "github.com/iq2i/aergie/internal/cmd/completion"
	helpCmd "github.com/iq2i/aergie/internal/cmd/help"
	userCmd "github.com/iq2i/aergie/internal/cmd/user"
	versionCmd "github.com/iq2i/aergie/internal/cmd/version"
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
			for _, c := range cmd.Commands() {
				if !c.IsAvailableCommand() {
					continue
				}

				list = append(list, fmt.Sprintf("%s\t%s", c.Use, c.Short))
			}
			return list, cobra.ShellCompDirectiveNoSpace | cobra.ShellCompDirectiveNoFileComp
		},
	}

	expandedArgs := []string{}
	if len(os.Args) > 0 {
		expandedArgs = os.Args[1:]
	}

	// translate `ae help <command>` to `ae <command> --help
	if len(expandedArgs) == 2 && expandedArgs[0] == "help" {
		expandedArgs = []string{expandedArgs[1], "--help"}
	}

	rootCmd.SetArgs(expandedArgs)

	rootCmd.PersistentFlags().Bool("help", false, "Show help for command")
	rootCmd.SetHelpFunc(helpCmd.Format)
	rootCmd.SetHelpCommand(&cobra.Command{Hidden: true})

	rootCmd.PersistentFlags().Bool("version", false, "Show ae version")
	rootCmd.SetVersionTemplate(versionCmd.Format(version))

	rootCmd.AddCommand(versionCmd.NewVersionCommand(version))
	rootCmd.AddCommand(completionCmd.NewCompletionCommand())
	rootCmd.AddCommand(userCmd.NewUserCommands()...)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
