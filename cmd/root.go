package cmd

import (
	"fmt"
	"os"

	"github.com/iq2i/aergie/internal/cmd/root"
	"github.com/spf13/cobra"
)

var quiet bool
var verbose bool

func init() {
	cobra.EnableCommandSorting = false
}

func Execute(version string) {
	var rootCmd = &cobra.Command{
		Use:     "ae <command> [flags]",
		Short:   "Aergie CLI",
		Long:    "An easy alternative to makefile",
		Version: version,

		SilenceUsage:  true,
		SilenceErrors: true,

		Args: cobra.NoArgs,
	}

	rootCmd.SetHelpFunc(root.HelpFunc)

	rootCmd.AddCommand(newUserCommands()...)

	rootCmd.PersistentFlags().Bool("help", false, "Show help for command")
	rootCmd.PersistentFlags().Bool("version", false, "Show ae version")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
