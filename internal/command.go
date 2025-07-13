package internal

import "fmt"

type Command struct {
	Name string
	Args []string
}

type commands struct {
	available map[string]func(*state, Command) error
}

func NewCommands() *commands {
	available := make(map[string]func(*state, Command) error)
	return &commands{available: available}
}

func (c *commands) Run(s *state, cmd Command) error {
	command, ok := c.available[cmd.Name]
	if !ok {
		return fmt.Errorf("unknown command: %s", cmd.Name)
	}

	err := command(s, cmd)
	if err != nil {
		return err
	}

	return nil
}

func (c *commands) Register(name string, f func(*state, Command) error) {
	c.available[name] = f
}
