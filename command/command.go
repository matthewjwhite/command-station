package command

import (
	"fmt"
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
