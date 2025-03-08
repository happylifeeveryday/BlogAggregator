package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	q := s.db
	err := q.DeleteAllUsers(context.Background())
	if err != nil {
		return fmt.Errorf("meet error %v when delete all users", err)
	}
	return nil
}
