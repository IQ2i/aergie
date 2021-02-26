package io

import (
	"testing"
)

func Test_Indent(t *testing.T) {
	type args struct {
		s      string
		indent string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "empty",
			args: args{
				s:      "",
				indent: "--",
			},
			want: "",
		},
		{
			name: "blank",
			args: args{
				s:      "\n",
				indent: "--",
			},
			want: "\n",
		},
		{
			name: "indent",
			args: args{
				s:      "one\ntwo\nthree",
				indent: "--",
			},
			want: "--one\n--two\n--three",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Indent(tt.args.s, tt.args.indent); got != tt.want {
				t.Errorf("indent() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestDedent(t *testing.T) {
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
			name: "dedent",
			args: args{
				s: "      test",
			},
			want: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Dedent(tt.args.s); got != tt.want {
				t.Errorf("dedent() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestRpad(t *testing.T) {
	type args struct {
		s       string
		padding int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "empty",
			args: args{
				s:       "",
				padding: 2,
			},
			want: "  ",
		},
		{
			name: "less_long",
			args: args{
				s:       "test",
				padding: 2,
			},
			want: "test",
		},
		{
			name: "longer",
			args: args{
				s:       "test",
				padding: 10,
			},
			want: "test      ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Rpad(tt.args.s, tt.args.padding); got != tt.want {
				t.Errorf("rpad() = %q, want %q", got, tt.want)
			}
		})
	}
}
