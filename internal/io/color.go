package io

import "github.com/mgutz/ansi"

var (
	bold   = ansi.ColorFunc("default+b")
	yellow = ansi.ColorFunc("yellow")
	cyan   = ansi.ColorFunc("cyan")
)

func Bold(t string) string {
	if len(t) == 0 {
		return ""
	}
	return bold(t)
}

func Yellow(t string) string {
	return yellow(t)
}

func Cyan(t string) string {
	return cyan(t)
}
