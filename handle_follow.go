package main

import (
	"context"
	"fmt"
	"time"

	"github.com/dothedada/rss_aggregator/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *State, cmd command, user database.User) error {
	ctx := context.Background()
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: %s <rss url>\n", cmd.name)
	}

	feed, err := s.db.GetFeedByUrl(ctx, cmd.args[0])
	if err != nil {
		return fmt.Errorf("Cannot get the feed: %w", err)
	}

	followData := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	}

	feedFollowed, err := s.db.CreateFeedFollow(ctx, followData)
	if err != nil {
		return fmt.Errorf("Cannot create the feed follow: %w", err)
	}

	fmt.Println("Succesfully created a new follow", feedFollowed.FeedName)
	fmt.Printf(
		"Name:	%s\nUser:	%s",
		feedFollowed.FeedName,
		feedFollowed.UserName,
	)

	return nil
}
