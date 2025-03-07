package main

import (
	"fmt"
	"log"
	"os"

	"github.com/happylifeeveryday/BlogAggregator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	fmt.Printf("Read config: %+v\n", cfg)

	state := config.State{
		ConfigPtr: &cfg,
	}

	commands := config.Commands{
		CommandMap: make(map[string]func(*config.State, config.Command) error),
	}

	commands.Register("login", config.HandlerLogin)

	if len(os.Args) < 2 {
		log.Fatalf("not enough arguments were provided")
	}

	command := config.Command{
		Name:      os.Args[1],
		Arguments: os.Args[2:],
	}
	err = commands.Run(&state, command)
	if err != nil {
		log.Fatalf("error login : %v", err)
	}

	cfg, err = config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	fmt.Printf("Read config again: %+v\n", cfg)
}
