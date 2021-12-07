package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/iq2i/aergie/internal/io"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Variables map[string]string  `yaml:"variables"`
	Commands  map[string]Command `yaml:"commands"`
}
type Command struct {
	Name  string   `yaml:"-"`
	Help  string   `yaml:"help"`
	Steps []string `yaml:"steps"`
}

var Variables map[string]string
var Commands map[string]Command

func init() {
	Variables = make(map[string]string)
	Commands = make(map[string]Command)
}

func LoadEnv(path string, env string) {
	populate(path)
	populate(fmt.Sprintf("%s.local", path))

	if env != "" {
		populate(fmt.Sprintf("%s.%s", path, env))
		populate(fmt.Sprintf("%s.%s.local", path, env))
	}
}

func populate(filename string) {
	var data = Config{}

	if filepath := fmt.Sprintf("%s.yml", filename); io.FileExists(filepath) {
		data = parseFile(filepath)
	} else if filepath := fmt.Sprintf("%s.yaml", filename); io.FileExists(filepath) {
		data = parseFile(filepath)
	}

	for key, variable := range data.Variables {
		Variables[key] = variable
	}
	for key, cmd := range data.Commands {
		Commands[key] = cmd
	}

	for _, cmd := range Commands {
		for index, step := range cmd.Steps {
			for name, value := range Variables {
				step = strings.ReplaceAll(step, "${"+name+"}", value)
			}

			cmd.Steps[index] = step
		}
	}
}

func parseFile(filepath string) Config {
	var config = Config{}

	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatalf("Can not read %s configuration file", filepath)
	}

	if err := yaml.Unmarshal(data, &config); err != nil {
		log.Fatalf("Invalid %s configuration file", filepath)
	}

	return config
}
