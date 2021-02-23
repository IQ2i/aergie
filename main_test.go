package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestVersion(t *testing.T) {
	for name, tt := range map[string]struct {
		version, date string
		out           string
	}{
		"all empty": {
			out: "",
		},
		"only version": {
			version: "1.2.3",
			out:     "1.2.3",
		},
		"version and date": {
			version: "1.2.3",
			date:    "12/12/12",
			out:     "1.2.3\nbuilt at: 12/12/12",
		},
	} {
		tt := tt
		t.Run(name, func(t *testing.T) {
			require.Equal(t, tt.out, buildVersion(tt.version, tt.date))
		})
	}
}
