package main

import (
	"context"
	"fmt"

	"github.com/dothedada/rss_aggregator/internal/rss"
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
