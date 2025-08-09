package main

import (
	"context"
	"fmt"
	"time"

	"github.com/dothedada/rss_aggregator/internal/database"
	"github.com/dothedada/rss_aggregator/internal/rss"
	"github.com/google/uuid"
)

func handlerAggregation(s *State, cmd command) error {
	feed, err := rss.FetchFeed(
		context.Background(),
		"https://www.wagslane.dev/index.xml",
	)
	if err != nil {
		return err
	}

	fmt.Printf("%v", feed)

	return nil
}

func handlerAddFeed(s *State, cmd command) error {
	if len(cmd.args) != 2 {
		return fmt.Errorf("usage: %s <rss name> <rss url>", cmd.name)
	}

	ctx := context.Background()

	userData, err := s.db.GetUser(ctx, s.cfg.CurrentUserName)
	if err != nil {
		return err
	}

	params := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      cmd.args[0],
		Url:       cmd.args[1],
		UserID:    userData.ID,
	}

	rssFeed, err := s.db.CreateFeed(ctx, params)
	if err != nil {
		return nil
	}

	fmt.Printf("Successfully added '%s' to your feed.\n", rssFeed.Name)
	fmt.Printf("It would fetch data from '%s'.\n", rssFeed.Url)
	return nil

}
