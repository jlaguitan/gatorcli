package main

import (
	"errors"
	"fmt"
)

type command struct {
	name string
	args []string
}

type commands struct {
	registeredCommands map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	regCommand, ok := c.registeredCommands[cmd.name]
	if ok == false {
		return errors.New("command is not registered")
	}

	return regCommand(s, cmd)
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.registeredCommands[name] = f
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return errors.New("username is required")
	}

	userName := cmd.args[0]

	err := s.cfg.SetUser(userName)
	if err != nil {
		return errors.New("command failed: user cannot be set.")
	}
	fmt.Printf("User: %v has been set successfully\n", userName)
	return nil
}
