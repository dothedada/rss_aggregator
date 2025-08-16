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
		return fmt.Errorf("Usage: %s <name>", cmd.name)
	}

	ctx := context.Background()
	userData := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      cmd.args[0],
	}

	user, err := s.db.CreateUser(ctx, userData)
	if err != nil {
		return error(err)
	}

	err = s.cfg.SetUser(user.Name)
	if err != nil {
		return err
	}

	fmt.Println("New user created:")
	printUser(user)

	return nil
}

func handlerLogin(s *State, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("Usage: %s <name>", cmd.name)
	}

	user, err := s.db.GetUser(context.Background(), cmd.args[0])
	if err != nil {
		log.Fatalf("Couldn't find the user '%s'", cmd.args[0])
	}

	if err := s.cfg.SetUser(user.Name); err != nil {
		return err
	}

	fmt.Printf("Hello %s, welcome to your feed\n", cmd.args[0])
	return nil
}

func handlerGetUsers(s *State, _ command) error {
	users, err := s.db.ListUsers(context.Background())
	if err != nil {
		return fmt.Errorf("Cannot get users list: %w", err)
	}

	for _, user := range users {
		var current string
		if s.cfg.CurrentUserName == user {
			current = "(current)"
		}
		fmt.Println(user, current)
	}

	return nil
}

func printUser(user database.User) {
	fmt.Printf("Name:   %s\n", user.Name)
	fmt.Printf("ID: 	%s\n", user.ID)
	fmt.Printf("at: 	%s\n", user.CreatedAt)
}
