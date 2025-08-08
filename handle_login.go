package main

import (
	"context"
	"fmt"
	"log"
)

func handlerLogin(s *State, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("Login command expects a single parameter")
	}

	user, err := s.db.GetUser(context.Background(), cmd.args[0])
	if err != nil {
		log.Fatalf("The user '%s' doesn't exists", cmd.args[0])
	}

	if err := s.cfg.SetUser(user.Name); err != nil {
		return err
	}

	fmt.Printf("Hello %s, welcome to your feed\n", cmd.args[0])
	return nil
}
