package main

import "fmt"

func handlerLogin(s *state, cmd command) error {
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
