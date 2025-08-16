package main

import (
	"context"
	"fmt"

	"github.com/dothedada/rss_aggregator/internal/database"
)

func middlewareLoggedIn(
	handler func(s *State, cmd command, user database.User) error,
) func(*State, command) error {

	return func(s *State, cmd command) error {
		user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
		if err != nil {
			return fmt.Errorf("Couldn't fetch the user data: %w", err)
		}

		return handler(s, cmd, user)
	}
}
