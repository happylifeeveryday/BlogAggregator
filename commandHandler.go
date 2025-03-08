package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/happylifeeveryday/BlogAggregator/internal/database"
)

// login
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

// register
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

// reset: delete all users
func handlerReset(s *state, cmd command) error {
	q := s.db
	err := q.DeleteAllUsers(context.Background())
	if err != nil {
		return fmt.Errorf("meet error %v when delete all users", err)
	}
	return nil
}

// users: get all users
func handlerUsers(s *state, cmd command) error {
	q := s.db
	users, err := q.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("meet error %v when get all users", err)
	}
	for _, user := range users {
		if user.Name.String == s.cfg.CurrentUserName {
			fmt.Printf("* %v (current)\n", user.Name.String)
		} else {
			fmt.Printf("* %v \n", user.Name.String)
		}
	}
	return nil
}

// agg: fetch feed
func handlerAgg(s *state, cmd command) error {
	RSS, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return err
	}
	fmt.Printf("%+v", *RSS)
	return nil
}

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
