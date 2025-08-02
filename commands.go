package main

import "errors"

type command struct {
	name string
	args []string
}

type commands struct {
	handlers map[string]func(*State, command) error
}

func (c *commands) run(s *State, cmd command) error {
	handler, ok := c.handlers[cmd.name]
	if !ok {
		return errors.New("Command not found")
	}

	return handler(s, cmd)
}

func (c *commands) register(name string, f func(*State, command) error) {
	c.handlers[name] = f
}
