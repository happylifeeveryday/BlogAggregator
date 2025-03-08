package main

import (
	"context"
	"database/sql"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Arguments) < 1 {
		return fmt.Errorf("login func missing parameter")
	}

	q := s.db
	name := cmd.Arguments[0]
	_, err := q.GetUser(context.Background(), sql.NullString{
		String: name,
		Valid:  true,
	})
	if err != nil {
		return fmt.Errorf("user does not exist")
	}

	err = s.cfg.SetUser(cmd.Arguments[0])
	if err != nil {
		return fmt.Errorf("error setting name")
	}
	fmt.Printf("User is set to %v \n", cmd.Arguments[0])
	return nil
}
