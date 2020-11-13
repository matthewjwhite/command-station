package command

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io"
	"os/exec"
)

// Command represents a command to be run via the command station.
type Command struct {
	Name    string
	Command string
}

// ErrUnknown constant for an unknown command.
var ErrUnknown = fmt.Errorf("command unknown")

// Execute executes the current command, returning the combined output
// (stdout + err), along with an error if the command failed to complete
// execution.
func (c Command) Execute() ([]byte, error) {
	cmd := exec.Command("sh", "-c", c.Command)
	combined, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}

	return combined, nil
}

// Commands returns a slice of Commands, given a Reader with YAML data.
// io.Reader selected for simplified, common live and test use (tests can
// can mock an io.Reader, live can use os.Open).
func Commands(reader io.Reader) ([]Command, error) {
	var commands []Command

	decoder := yaml.NewDecoder(reader)
	if err := decoder.Decode(&commands); err != nil {
		return nil, err
	}

	return commands, nil
}

// Collection is an alias for []Command, allowing us to apply methods to a Command collection.
type Collection []Command

// Get retrieves a Command by name.
func (c Collection) Get(name string) (Command, error) {
	for _, command := range c {
		if command.Name == name {
			return command, nil
		}
	}

	return Command{}, ErrUnknown
}
