package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/happylifeeveryday/BlogAggregator/internal/database"
)

func handlerRegister(s *state, cmd command) error {
	q := s.db
	name := cmd.Arguments[0]
	_, err := q.GetUser(context.Background(), sql.NullString{
		String: name,
		Valid:  true,
	})
	if err == nil {
		return fmt.Errorf("user already exist")
	}

	user := database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
		UpdatedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
		Name: sql.NullString{
			String: name,
			Valid:  true,
		},
	}

	u, err := q.CreateUser(context.Background(), user)
	if err != nil {
		return err
	}
	fmt.Printf("user was created\n")
	fmt.Printf("User info: %+v\n", u)
	_ = handlerLogin(s, cmd)

	return nil
}
