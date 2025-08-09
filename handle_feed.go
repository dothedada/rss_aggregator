package main

import (
	"context"
	"fmt"
	"time"

	"github.com/dothedada/rss_aggregator/internal/database"
	"github.com/google/uuid"
)

func handlerAggregation(s *State, cmd command) error {
	feed, err := FetchFeed(
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
	ctx := context.Background()
	userData, err := s.db.GetUser(ctx, s.cfg.CurrentUserName)
	if err != nil {
		return err
	}

	if len(cmd.args) != 2 {
		return fmt.Errorf("usage: %s <rss name> <rss url>", cmd.name)
	}

	fetchFromData := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      cmd.args[0],
		Url:       cmd.args[1],
		UserID:    userData.ID,
	}
	rssFeed, err := s.db.CreateFeed(ctx, fetchFromData)
	if err != nil {
		return nil
	}

	fmt.Printf("Successfully added a new feed")
	printFeed(rssFeed)
	return nil
}

func printFeed(f database.Feed) {
	fmt.Println("Feed data...")
	fmt.Printf("name: 		%s\n", f.Name)
	fmt.Printf("url: 		%s\n", f.Url)
	fmt.Printf("since: 		%s\n", f.CreatedAt)
	fmt.Printf("Updated at:	%s\n", f.UpdatedAt)
}
