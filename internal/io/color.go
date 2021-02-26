package io

import "github.com/mgutz/ansi"

var (
	bold = ansi.ColorFunc("default+b")
)

// Bold is a function to render text bold in terminal
func Bold(t string) string {
	if len(t) == 0 {
		return ""
	}
	return bold(t)
}
