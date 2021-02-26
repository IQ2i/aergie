package io

import "github.com/mgutz/ansi"

var (
	bold = ansi.ColorFunc("default+b")
)

func Bold(t string) string {
	return bold(t)
}
