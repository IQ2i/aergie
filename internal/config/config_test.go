package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReplaceVar(t *testing.T) {
	is := assert.New(t)

	config := Config{
		Variables: map[string]string{"test": "test environment"},
		Commands: []Command{Command{
			Name: "env",
			Steps: []string{
				"${test}",
			},
		}},
	}
	replaceVar(&config)

	expected := Config{
		Variables: map[string]string{"test": "test environment"},
		Commands: []Command{Command{
			Name: "env",
			Steps: []string{
				"test environment",
			},
		}},
	}

	is.Equal(expected, config)
}
