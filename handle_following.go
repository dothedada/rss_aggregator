package main

import (
	"context"
	"fmt"

	"github.com/dothedada/rss_aggregator/internal/database"
)

func handlerFollowing(s *State, _ command, user database.User) error {
	feeds, err := s.db.GetFeedsFollowedByUser(context.Background(), user.Name)
	if err != nil {
		return err
	}

	fmt.Println()
	fmt.Printf("Feed follows for %s:\n", s.cfg.CurrentUserName)
	for i, feed := range feeds {
		amount := i + 1
		fmt.Printf("%d) %s\n", amount, feed.FeedName)
	}
	fmt.Println()

	return nil
}
