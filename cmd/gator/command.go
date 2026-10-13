package main

import "errors"

type command struct {
	name      string
	arguments []string
}

type commands struct {
	cmdMap map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	value, ok := c.cmdMap[cmd.name]
	if !ok {
		return errors.New("The command doesn't exist")
	}
	err := value(s, cmd)
	if err != nil {
		return err
	}
	return nil
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.cmdMap[name] = f
}
