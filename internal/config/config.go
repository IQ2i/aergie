package config

import (
	"fmt"
	"io/ioutil"
	"strings"

	"gopkg.in/yaml.v2"
)

// Config contains data from configuration file
type Config struct {
	Variables map[string]string  `yaml:"variables"`
	Commands  map[string]Command `yaml:"commands"`
}

// Command is a command definition
type Command struct {
	Help  string   `yaml:"help"`
	Steps []string `yaml:"steps"`
}

// AppConfig is the application configuration, accessible from anywhere
var AppConfig Config

// Init the configuration
func Init() error {
	config := Config{}
	defer func() { AppConfig = config }()

	// get config file content
	data := configFile()
	if len(data) == 0 {
		return nil
	}

	// cast to YAML
	err := yaml.Unmarshal(data, &config)
	if err != nil {
		return fmt.Errorf("Configuration file is invalid.")
	}

	// replace variables in steps
	config = replaceVar(config)

	return nil
}

func configFile() []byte {
	data, _ := ioutil.ReadFile(".aergie.yml")
	if len(data) == 0 {
		data, _ = ioutil.ReadFile(".aergie.yaml")
	}
	if len(data) == 0 {
		data, _ = ioutil.ReadFile(".ae.yml")
	}
	if len(data) == 0 {
		data, _ = ioutil.ReadFile(".ae.yaml")
	}

	return data
}

func replaceVar(config Config) Config {
	for name, command := range config.Commands {
		var steps []string
		for _, step := range command.Steps {
			for n, v := range config.Variables {
				step = strings.ReplaceAll(step, "${"+n+"}", v)
				break
			}
			steps = append(steps, step)
		}
		command.Steps = steps
		config.Commands[name] = command
	}

	return config
}
