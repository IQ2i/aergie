package cmd

import (
	"fmt"
	"regexp"

	"github.com/spf13/cobra"
)

func newVersionCommand(version string) *cobra.Command {
	cmd := &cobra.Command{
		Use:    "version",
		Hidden: true,

		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Fprint(cmd.OutOrStdout(), versionFormat(version))
			return nil
		},
	}

	return cmd
}

func versionFormat(version string) string {
	return fmt.Sprintf("ae version %s\n%s\n", version, changelogURL(version))
}

func changelogURL(version string) string {
	path := "https://github.com/IQ2i/aergie"

	r := regexp.MustCompile(`^\d+\.\d+\.\d+$`)
	if !r.MatchString(version) {
		return fmt.Sprintf("%s/releases/latest", path)
	}

	return fmt.Sprintf("https://github.com/IQ2i/aergie/releases/tag/v%s", version)
}
