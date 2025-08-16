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

func handlerAddFeed(s *State, cmd command, user database.User) error {
	ctx := context.Background()
	if len(cmd.args) != 2 {
		return fmt.Errorf("usage: %s <rss name> <rss url>\n", cmd.name)
	}

	fetchFromData := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      cmd.args[0],
		Url:       cmd.args[1],
		UserID:    user.ID,
	}
	rssFeed, err := s.db.CreateFeed(ctx, fetchFromData)
	if err != nil {
		return fmt.Errorf("Cannot create the feed: %w", err)
	}

	feedFollowData := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    rssFeed.ID,
	}
	feedSubscription, err := s.db.CreateFeedFollow(ctx, feedFollowData)
	if err != nil {
		return fmt.Errorf("Cannot create the feed follow: %w", err)
	}

	fmt.Println("Feed successfully created")
	printFeed(FeedData{Name: rssFeed.Name, Url: rssFeed.Url})

	fmt.Println("Feed successfully followed")
	fmt.Println("Name:", feedSubscription.FeedName)
	fmt.Println("User:", feedSubscription.UserName)

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

func handlerUnfollowFeed(s *State, cmd command, user database.User) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: %s <rss url>\n", cmd.name)
	}

	deleteParams := database.DeleteFeedByURLParams{
		UserID: user.ID,
		Url:    cmd.args[0],
	}

	err := s.db.DeleteFeedByURL(context.Background(), deleteParams)
	if err != nil {
		return fmt.Errorf("Cannot delet folling feed record: %w", err)
	}

	return nil
}
