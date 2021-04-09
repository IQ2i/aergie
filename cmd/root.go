package cmd

import (
	"fmt"
	"os"

	"github.com/iq2i/aergie/internal/io"
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

		PersistentPostRun: func(cmd *cobra.Command, args []string) {
			latestVersion := root.GetLatestVersion()

			if version != latestVersion {
				fmt.Printf("\n%s %s → %s\n", io.Yellow("A new release of Aergie is available:"), io.Cyan(version), io.Cyan(latestVersion))
				fmt.Printf("%s\n\n", io.Yellow(fmt.Sprintf("https://github.com/IQ2i/aergie/releases/tag/%s", latestVersion)))
			}
		},
	}

	rootCmd.SetHelpFunc(root.HelpFunc)

	rootCmd.AddCommand(newCompletionCommand())
	rootCmd.AddCommand(newUserCommands()...)

	rootCmd.PersistentFlags().Bool("help", false, "Show help for command")
	rootCmd.PersistentFlags().Bool("version", false, "Show ae version")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
