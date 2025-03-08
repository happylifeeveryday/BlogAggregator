package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/happylifeeveryday/BlogAggregator/internal/config"
	"github.com/happylifeeveryday/BlogAggregator/internal/database"
	_ "github.com/lib/pq"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

func main() {

	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	fmt.Printf("Read config: %+v\n", cfg)

	db, err := sql.Open("postgres", cfg.DbURL)
	if err != nil {
		log.Fatalf("error open database: %v", err)
	}

	dbQueries := database.New(db)

	std := state{
		db:  dbQueries,
		cfg: &cfg,
	}

	commands := commands{
		CommandMap: make(map[string]func(*state, command) error),
	}

	commands.Register("login", handlerLogin)
	commands.Register("register", handlerRegister)

	if len(os.Args) < 2 {
		log.Fatalf("not enough arguments were provided")
	}

	command := command{
		Name:      os.Args[1],
		Arguments: os.Args[2:],
	}
	err = commands.Run(&std, command)
	if err != nil {
		log.Fatalf("%v", err)
	}

	cfg, err = config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	fmt.Printf("Read config again: %+v\n", cfg)
}
