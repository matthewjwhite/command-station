package config

import (
	"github.com/matthewjwhite/command-station/command"
	"gopkg.in/yaml.v2"
	"io"
)

// Config contains the configuration parameters for rendering the command station.
type Config struct {
	Title    string
	Commands []command.Command
}

// Parse produces a Config struct from a Reader.
func Parse(reader io.Reader) (Config, error) {
	var config Config

	if err := yaml.NewDecoder(reader).Decode(&config); err != nil {
		return Config{}, nil
	}

	return config, nil
}
