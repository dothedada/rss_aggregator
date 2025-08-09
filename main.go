package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/dothedada/rss_aggregator/internal/config"
	"github.com/dothedada/rss_aggregator/internal/database"
	_ "github.com/lib/pq"
)

type State struct {
	cfg *config.Config
	db  *database.Queries
}

func main() {
	conf, err := config.Read()
	if err != nil {
		log.Fatalf("some shit happened %v", err)
	}

	state := &State{
		cfg: &conf,
	}

	db, err := sql.Open("postgres", state.cfg.DbUrl)
	if err != nil {
		log.Fatal(err)
	}

	state.db = database.New(db)

	cmds := commands{
		handlers: make(map[string]func(*State, command) error),
	}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("users", handlerGetUsers)
	cmds.register("agg", handlerAggregation)

	cmds.register("reset", func(s *State, cmd command) error {

		err := s.db.DropUsers(context.Background())
		if err != nil {
			return error(err)
		}

		fmt.Println("Clean users DB...")
		return nil
	})

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
