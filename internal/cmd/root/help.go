package root

import (
	"fmt"
	"strings"

	"github.com/iq2i/aergie/internal/io"
	"github.com/spf13/cobra"
)

func HelpFunc(cmd *cobra.Command, args []string) {
	commands := []string{}
	for _, c := range cmd.Commands() {
		if c.Short == "" {
			continue
		}
		if c.Hidden {
			continue
		}

		s := io.Rpad(c.Name(), c.NamePadding()+1) + c.Short
		commands = append(commands, s)
	}

	type helpEntry struct {
		Title string
		Body  string
	}

	helpEntries := []helpEntry{}
	if cmd.Long != "" {
		helpEntries = append(helpEntries, helpEntry{"", cmd.Long})
	} else if cmd.Short != "" {
		helpEntries = append(helpEntries, helpEntry{"", cmd.Short})
	}
	helpEntries = append(helpEntries, helpEntry{"USAGE", cmd.UseLine()})
	if len(commands) > 0 {
		helpEntries = append(helpEntries, helpEntry{"COMMANDS", strings.Join(commands, "\n")})
	}

	flagUsages := cmd.LocalFlags().FlagUsages()
	if flagUsages != "" {
		helpEntries = append(helpEntries, helpEntry{"FLAGS", io.Dedent(flagUsages)})
	}
	inheritedFlagUsages := cmd.InheritedFlags().FlagUsages()
	if inheritedFlagUsages != "" {
		helpEntries = append(helpEntries, helpEntry{"INHERITED FLAGS", io.Dedent(inheritedFlagUsages)})
	}
	if _, ok := cmd.Annotations["help:arguments"]; ok {
		helpEntries = append(helpEntries, helpEntry{"ARGUMENTS", cmd.Annotations["help:arguments"]})
	}

	out := cmd.OutOrStdout()
	for _, e := range helpEntries {
		if e.Title != "" {
			fmt.Fprintln(out, io.Bold(e.Title))
			fmt.Fprintln(out, io.Indent(strings.Trim(e.Body, "\r\n"), "  "))
		} else {
			fmt.Fprintln(out, e.Body)
		}
		fmt.Fprintln(out)
	}
}
