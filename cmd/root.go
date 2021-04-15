package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/iq2i/aergie/internal/cmd/root"
	"github.com/iq2i/aergie/internal/io"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	cobra.EnableCommandSorting = false
}

// Execute is the real main function of Aergie cli
func Execute() {
	var rootCmd = &cobra.Command{
		Use:     "ae <command> [flags]",
		Short:   "Aergie CLI",
		Long:    "An easy alternative to makefile",
		Version: os.Getenv("AE_VERSION"),

		SilenceUsage:  true,
		SilenceErrors: true,

		Args: cobra.NoArgs,

		PersistentPostRun: func(cmd *cobra.Command, args []string) {
			now := time.Now().Add(-24 * time.Hour)
			latestCheck := viper.GetTime("update.latest_check")
			if latestCheck.Before(now) {
				viper.Set("update.latest_version", root.GetLatestVersion())
				viper.Set("update.latest_check", time.Now().Format("2006-01-02 15:04:05"))
				viper.WriteConfig()
			}

			latestVersion := viper.GetString("update.latest_version")
			if os.Getenv("AE_VERSION") != latestVersion {
				fmt.Printf("\n%s %s → %s\n", io.Yellow("A new release of Aergie is available:"), io.Cyan(os.Getenv("AE_VERSION")), io.Cyan(latestVersion))
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
