package io

import (
	"testing"
)

func Test_Bold(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "empty",
			args: args{
				s: "",
			},
			want: "",
		},
		{
			name: "bold",
			args: args{
				s: "test",
			},
			want: "\x1b[0;1;39mtest\x1b[0m",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Bold(tt.args.s); got != tt.want {
				t.Errorf("bold() = %q, want %q", got, tt.want)
			}
		})
	}
}
