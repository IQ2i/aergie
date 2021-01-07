package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func TestCreateCmd(t *testing.T) {
	is := assert.New(t)

	config := Config{
		Variables: map[string]string{"test": "test command"},
		CommandMap: yaml.Node{
			Content: []*yaml.Node{
				&yaml.Node{
					Kind:  yaml.ScalarNode,
					Value: "cmd1",
				},
				&yaml.Node{
					Kind: yaml.MappingNode,
					Content: []*yaml.Node{
						&yaml.Node{
							Kind:  yaml.ScalarNode,
							Value: "help",
						},
						&yaml.Node{
							Kind:  yaml.ScalarNode,
							Value: "help cmd1",
						},
					},
				},
				&yaml.Node{
					Kind:  yaml.ScalarNode,
					Value: "cmd2",
				},
				&yaml.Node{
					Kind: yaml.MappingNode,
					Content: []*yaml.Node{
						&yaml.Node{
							Kind:  yaml.ScalarNode,
							Value: "help",
						},
						&yaml.Node{
							Kind:  yaml.ScalarNode,
							Value: "help cmd2",
						},
					},
				},
			},
		},
	}
	createCmd(&config)

	expected := Config{
		Variables: map[string]string{"test": "test command"},
		Commands: []Command{
			Command{
				Name: "cmd1",
				Help: "help cmd1",
			},
			Command{
				Name: "cmd2",
				Help: "help cmd2",
			},
		},
	}

	is.Equal(expected.Commands, config.Commands)
}

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
