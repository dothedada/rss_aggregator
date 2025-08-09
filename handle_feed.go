package main

import (
	"context"
	"fmt"
	"time"

	"github.com/dothedada/rss_aggregator/internal/database"
	"github.com/google/uuid"
)

type FeedData struct {
	Name     string
	Url      string
	UserName string
}

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

	fmt.Printf("%s, you successfully added a new feed...\n", userData.Name)
	printFeed(FeedData{Name: rssFeed.Name, Url: rssFeed.Url})
	return nil
}

func printFeed(f FeedData) {
	fmt.Printf("name: 		%s\n", f.Name)
	fmt.Printf("url: 		%s\n", f.Url)
	if f.UserName != "" {
		fmt.Printf("Subscriber:	%s\n", f.UserName)
	}
}

func handlerListFeeds(s *State, _ command) error {
	feeds, err := s.db.ListFeeds(context.Background())
	if err != nil {
		return err
	}

	fmt.Println("Followed feeds:")
	for i, feed := range feeds {
		fmt.Printf("%d) ", i)
		printFeed(FeedData{
			Name:     feed.FeedName,
			Url:      feed.Url,
			UserName: feed.SuscribedUser,
		})
	}
	fmt.Println("--- That's all folks!!! ---")

	return nil
}
