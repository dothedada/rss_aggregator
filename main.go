package main

import (
	"log"
	"os"

	"github.com/dothedada/rss_aggregator/internal/config"
)

type State struct {
	cfg *config.Config
}

func main() {
	conf, err := config.Read()
	if err != nil {
		log.Fatalf("some shit happened: %w", err)
	}

	state := &State{
		cfg: &conf,
	}

	cmds := commands{
		handlers: make(map[string]func(*State, command) error),
	}
	cmds.register("login", handlerLogin)

	if len(os.Args) < 2 {
		log.Fatalf("Usage: cli <command> [args...]")
	}

	cmd := command{
		name: os.Args[1],
		args: os.Args[2:],
	}

	if err = cmds.run(state, cmd); err != nil {
		log.Fatal(err)
	}
}
