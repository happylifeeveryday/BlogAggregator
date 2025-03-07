package config

import "fmt"

type State struct {
	ConfigPtr *Config
}

type Command struct {
	Name      string
	Arguments []string
}

type Commands struct {
	CommandMap map[string]func(*State, Command) error
}

func (c *Commands) Register(name string, f func(*State, Command) error) {
	c.CommandMap[name] = f
}

func (c *Commands) Run(s *State, cmd Command) error {
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
