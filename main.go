package main

import (
	"fmt"
	"log"

	"github.com/happylifeeveryday/BlogAggregator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	fmt.Printf("Read config: %+v\n", cfg)
	err = cfg.SetUser("yihan")
	if err != nil {
		log.Fatalf("error set user: %v", err)
	}
	cfg, err = config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	fmt.Printf("Read config again: %+v\n", cfg)
}
