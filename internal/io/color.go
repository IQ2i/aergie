package io

import "github.com/mgutz/ansi"

var (
	bold   = ansi.ColorFunc("default+b")
	green  = ansi.ColorFunc("green")
	yellow = ansi.ColorFunc("yellow")
)

// Bold is a function to render text bold in terminal
func Bold(t string) string {
	if len(t) == 0 {
		return ""
	}
	return bold(t)
}

func Green(t string) string {
	if len(t) == 0 {
		return ""
	}
	return green(t)
}

func Yellow(t string) string {
	if len(t) == 0 {
		return ""
	}
	return yellow(t)
}
