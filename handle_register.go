package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/dothedada/rss_aggregator/internal/database"
	"github.com/google/uuid"
)

func handlerRegister(s *State, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("Login command expects a single parameter")
	}

	ctx := context.Background()
	userData := database.User{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.args[0],
	}

	userInDB, err := s.db.GetUser(ctx, userData.Name)
	if err == nil && userInDB.Name != "" {
		log.Fatalf("User '%s' already exists", userInDB.Name)
	}

	user, err := s.db.CreateUser(ctx, database.CreateUserParams(userData))
	if err != nil {
		return error(err)
	}

	err = s.cfg.SetUser(user.Name)
	if err != nil {
		return err
	}

	fmt.Println("I hope you enjoy this rss reader")
	fmt.Printf(
		"UUID: %s\nCreated: %v\nLogged: %v\nName: %s",
		user.ID,
		user.CreatedAt,
		user.UpdatedAt,
		user.Name,
	)

	return nil
}
