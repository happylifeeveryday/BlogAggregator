package main

import (
	"fmt"
	"log"
	"os"

	"github.com/happylifeeveryday/BlogAggregator/internal/config"
)

type state struct {
	ConfigPtr *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	fmt.Printf("Read config: %+v\n", cfg)

	std := state{
		ConfigPtr: &cfg,
	}

	commands := commands{
		CommandMap: make(map[string]func(*state, command) error),
	}

	commands.Register("login", handlerLogin)

	if len(os.Args) < 2 {
		log.Fatalf("not enough arguments were provided")
	}

	command := command{
		Name:      os.Args[1],
		Arguments: os.Args[2:],
	}
	err = commands.Run(&std, command)
	if err != nil {
		log.Fatalf("error login : %v", err)
	}

	cfg, err = config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	fmt.Printf("Read config again: %+v\n", cfg)
}
