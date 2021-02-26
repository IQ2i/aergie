package io

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"
)

var lineRE = regexp.MustCompile(`(?m)^`)

// Indent adds characters before each line of the given string
func Indent(s, indent string) string {
	if len(strings.TrimSpace(s)) == 0 {
		return s
	}
	return lineRE.ReplaceAllLiteralString(s, indent)
}

// Dedent removes space before given string
func Dedent(s string) string {
	lines := strings.Split(s, "\n")
	minIndent := -1

	for _, l := range lines {
		if len(l) == 0 {
			continue
		}

		indent := len(l) - len(strings.TrimLeft(l, " "))
		if minIndent == -1 || indent < minIndent {
			minIndent = indent
		}
	}

	if minIndent <= 0 {
		return s
	}

	var buf bytes.Buffer
	for _, l := range lines {
		fmt.Fprintln(&buf, strings.TrimPrefix(l, strings.Repeat(" ", minIndent)))
	}
	return strings.TrimSuffix(buf.String(), "\n")
}

// Rpad adds space after given string if it's shorter than padding
func Rpad(s string, padding int) string {
	template := fmt.Sprintf("%%-%dv", padding)
	return fmt.Sprintf(template, s)
}
