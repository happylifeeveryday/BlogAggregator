package main

import "fmt"

type command struct {
	Name      string
	Arguments []string
}

type commands struct {
	CommandMap map[string]func(*state, command) error
}

func (c *commands) Register(name string, f func(*state, command) error) {
	c.CommandMap[name] = f
}

func (c *commands) Run(s *state, cmd command) error {
	value, exists := c.CommandMap[cmd.Name]
	if !exists {
		return fmt.Errorf("function %v does not exist", cmd.Name)
	}
	functionRun := value
	err := functionRun(s, cmd)
	if err != nil {
		return err
	}
	return nil
}
