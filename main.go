package main

import (
	"fmt"

	"github.com/happylifeeveryday/BlogAggregator/internal/config"
)

func main() {
	cfg, _ := config.Read()
	fmt.Println(cfg.CurrentUserName, cfg.DbURL)
	_ = cfg.SetUser("yihan")
	cfg, _ = config.Read()
	fmt.Println(cfg.CurrentUserName, cfg.DbURL)
}
