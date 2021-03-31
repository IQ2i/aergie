package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"gopkg.in/yaml.v3"
)

// Config contains data from configuration file
type Config struct {
	Variables  map[string]string `yaml:"variables"`
	CommandMap yaml.Node         `yaml:"commands,flow"`
	Commands   []Command         `yaml:"-"`
}

// Command is a command definition
type Command struct {
	Name  string   `yaml:"-"`
	Help  string   `yaml:"help"`
	Steps []string `yaml:"steps"`
}

// AppConfig is the application configuration, accessible from anywhere
var AppConfig Config

// Init the configuration
func Init() {
	config := Config{}
	defer func() { AppConfig = config }()

	// get config file content
	data := configFile()
	if len(data) == 0 {
		return
	}

	// cast to YAML
	err := yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatal("Configuration file is invalid")
	}

	// create commands from YAML parsing
	createCmd(&config)

	// replace variables in steps
	replaceVar(&config)
}

func configFile() []byte {
	data, _ := ioutil.ReadFile(".aergie.yml")
	if len(data) == 0 {
		data, _ = ioutil.ReadFile(".aergie.yaml")
	}

	return data
}

func createCmd(config *Config) {
	for key, item := range config.CommandMap.Content {
		if yaml.ScalarNode == item.Kind {
			// get command name
			name := item.Value

			for _, cmd := range config.Commands {
				if cmd.Name == name {
					log.Fatal(fmt.Sprintf("Configuration file is invalid, you have defined the same command \"%s\" twice", name))
				}
			}

			// get next loop's item which is the command
			item = config.CommandMap.Content[key+1]
			command := &Command{}
			_ = item.Decode(command)

			// update command name
			command.Name = name

			config.Commands = append(config.Commands, *command)
		}
	}
}

func replaceVar(config *Config) {
	for key, command := range config.Commands {
		var steps []string
		for _, step := range command.Steps {
			for n, v := range config.Variables {
				step = strings.ReplaceAll(step, "${"+n+"}", v)
			}
			steps = append(steps, step)
		}
		command.Steps = steps
		config.Commands[key] = command
	}
}
