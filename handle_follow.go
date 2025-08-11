package main

import (
	"context"
	"fmt"
	"time"

	"github.com/dothedada/rss_aggregator/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *State, cmd command) error {
	ctx := context.Background()
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: %s <rss url>\n", cmd.name)
	}

	user, err := s.db.GetUser(ctx, s.cfg.CurrentUserName)
	if err != nil {
		return err
	}

	feed, err := s.db.GetFeedByUrl(ctx, cmd.args[0])
	if err != nil {
		return err
	}

	followData := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	}

	feedFollowed, err := s.db.CreateFeedFollow(context.Background(), followData)
	if err != nil {
		return err
	}

	fmt.Println("now you follow", feedFollowed.FeedName)
	fmt.Println("of", feedFollowed.UserName)

	return nil
}
