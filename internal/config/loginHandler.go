package config

import "fmt"

func HandlerLogin(s *State, cmd Command) error {
	if len(cmd.Arguments) < 1 {
		return fmt.Errorf("login func missing parameter")
	}
	err := s.ConfigPtr.SetUser(cmd.Arguments[0])
	if err != nil {
		return fmt.Errorf("error setting name")
	}
	fmt.Printf("User is set to %v \n", cmd.Arguments[0])
	return nil
}
