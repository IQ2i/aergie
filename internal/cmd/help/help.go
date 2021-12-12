package help

import (
	"fmt"
	"strings"

	"github.com/iq2i/aergie/internal/io"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// Format prints help text
func Format(cmd *cobra.Command, args []string) {
	coreCommands := []string{}
	userCommands := make(map[string][]string)
	for _, c := range cmd.Commands() {
		if !c.IsAvailableCommand() {
			continue
		}

		s := io.Green(io.Rpad(c.Name(), c.NamePadding()+1)) + c.Short

		if _, ok := c.Annotations["IsCore"]; ok {
			coreCommands = append(coreCommands, s)
		} else {
			category := strings.Split(c.Name(), ":")[0]
			userCommands[category] = append(userCommands[category], s)
		}
	}

	type helpEntry struct {
		Title   string
		Body    string
		NewLine bool
	}

	helpEntries := []helpEntry{}

	// description
	if !cmd.HasParent() {
		helpEntries = append(helpEntries, helpEntry{"", cmd.Short, true})
	} else {
		helpEntries = append(helpEntries, helpEntry{"Description:", cmd.Short, true})
	}

	// usage
	usage := ""
	if cmd.HasParent() {
		usage += cmd.Parent().CommandPath() + " "
	}
	usage += cmd.Use + " [options]"
	helpEntries = append(helpEntries, helpEntry{"Usage:", usage, true})

	// options
	options := formatFlags(cmd.LocalFlags(), cmd.InheritedFlags())
	if options != "" {
		helpEntries = append(helpEntries, helpEntry{"Options", io.Dedent(options), true})
	}

	// core commands
	if len(coreCommands) > 0 {
		helpEntries = append(helpEntries, helpEntry{"Available commands:", strings.Join(coreCommands, "\n"), false})
	}

	// user commands
	for cat, cmds := range userCommands {
		helpEntries = append(helpEntries, helpEntry{" " + cat, strings.Join(cmds, "\n"), false})
	}

	out := cmd.OutOrStdout()
	for _, e := range helpEntries {
		if e.Title != "" {
			fmt.Fprintln(out, io.Bold(io.Yellow(e.Title)))
			fmt.Fprint(out, io.Indent(strings.Trim(e.Body, "\r\n"), "  "))
		} else {
			fmt.Fprint(out, e.Body)
		}
		if e.NewLine {
			fmt.Fprintln(out)
		}
		fmt.Fprintln(out)
	}
}

func formatFlags(localFlags *pflag.FlagSet, inheritedFlags *pflag.FlagSet) string {

	flags := []*pflag.Flag{}
	localFlags.VisitAll(func(flag *pflag.Flag) {
		flags = append(flags, flag)
	})
	inheritedFlags.VisitAll(func(flag *pflag.Flag) {
		flags = append(flags, flag)
	})

	lines := []string{}
	maxlen := 0
	for _, flag := range flags {
		if flag.Hidden {
			continue
		}

		line := ""
		if flag.Shorthand != "" && flag.ShorthandDeprecated == "" {
			line = fmt.Sprintf("-%s, --%s", flag.Shorthand, flag.Name)
		} else {
			line = fmt.Sprintf("    --%s", flag.Name)
		}

		varname, usage := pflag.UnquoteUsage(flag)
		if varname != "" {
			line += " " + varname
		}
		if flag.NoOptDefVal != "" {
			switch flag.Value.Type() {
			case "string":
				line += fmt.Sprintf("[=\"%s\"]", flag.NoOptDefVal)
			case "bool":
				if flag.NoOptDefVal != "true" {
					line += fmt.Sprintf("[=%s]", flag.NoOptDefVal)
				}
			case "count":
				if flag.NoOptDefVal != "+1" {
					line += fmt.Sprintf("[=%s]", flag.NoOptDefVal)
				}
			default:
				line += fmt.Sprintf("[=%s]", flag.NoOptDefVal)
			}
		}

		// This special character will be replaced with spacing once the
		// correct alignment is calculated
		line += "\x00"
		if len(line) > maxlen {
			maxlen = len(line)
		}

		line += usage
		if len(flag.Deprecated) != 0 {
			line += fmt.Sprintf(" (DEPRECATED: %s)", flag.Deprecated)
		}

		lines = append(lines, line)
	}

	options := []string{}
	for _, line := range lines {
		s := strings.Split(line, "\x00")
		res := io.Green(io.Rpad(s[0], maxlen+2)) + s[1]

		options = append(options, res)
	}

	return strings.Join(options, "\n")
}
