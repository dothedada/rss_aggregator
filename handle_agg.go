package main

import (
	"context"
	"database/sql"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/dothedada/rss_aggregator/internal/database"
	"github.com/google/uuid"
)

type RSSFeed struct {
	Channel channel `xml:"channel"`
}

type channel struct {
	Title       string    `xml:"title"`
	Link        string    `xml:"link"`
	Description string    `xml:"description"`
	Items       []RSSItem `xml:"item"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func handlerAggregation(s *State, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: %s <time between scraps>\n", cmd.name)
	}

	timeBetweenRequest, err := time.ParseDuration(cmd.args[0])
	if err != nil {
		return fmt.Errorf("Invalid time duration: %w", err)
	}

	ticker := time.NewTicker(timeBetweenRequest)
	for ; ; <-ticker.C {
		scrapeFeed(s)
	}
}

func scrapeFeed(s *State) (*RSSFeed, error) {
	ctx := context.Background()

	feedSource, err := s.db.GetNextFeedToFetch(ctx)
	if err != nil {
		return &RSSFeed{}, fmt.Errorf("Cannot get the next feed to fetch: %w", err)
	}

	err = s.db.MarkFeedFetched(ctx, feedSource.ID)
	if err != nil {
		return &RSSFeed{}, fmt.Errorf("Cannot update feed fetch update: %w", err)
	}

	feeds, err := FetchFeed(ctx, feedSource.Url)
	if err != nil {
		return &RSSFeed{}, fmt.Errorf("Error while scraping feed: %w", err)
	}

	fmt.Printf("Title: %s\n", feeds.Channel.Title)

	for i, feed := range feeds.Channel.Items {
		feedData := database.AddPostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now().UTC(),
			UpdatedAt:   time.Now().UTC(),
			Title:       feed.Title,
			Url:         feed.Link,
			Description: feed.Description,
			PublishedAt: dateParser(feed.PubDate),
			FeedID:      feedSource.ID,
		}
		err := s.db.AddPost(ctx, feedData)
		if err != nil && !strings.Contains(err.Error(), "duplicate") {
			return &RSSFeed{}, fmt.Errorf("Cannot save the posts: %w", err)
		}

		fmt.Printf("%d) %s\n", i, feed.Title)
		fmt.Printf("%s\n", feed.Description)
	}

	return feeds, nil
}

func dateParser(dateString string) sql.NullTime {
	layouts := []string{
		time.RFC1123,
		time.RFC1123Z,
		time.RFC3339,
		"2006-01-02",
	}

	for _, dateLayout := range layouts {
		parsedDate, err := time.Parse(dateLayout, dateString)
		if err == nil {
			return sql.NullTime{
				Time:  parsedDate,
				Valid: true,
			}
		}
	}

	return sql.NullTime{
		Time:  time.Time{},
		Valid: false,
	}
}

func FetchFeed(ctx context.Context, url string) (*RSSFeed, error) {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "RSS-Aggregator")
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode > 299 {
		return nil, fmt.Errorf("Status error: %d", res.StatusCode)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	feed := RSSFeed{}
	err = xml.Unmarshal(data, &feed)
	if err != nil {
		return nil, err
	}

	unescapeRSSFeed(&feed)

	return &feed, nil
}

func unescapeRSSFeed(feed *RSSFeed) {
	feed.Channel.Title = html.UnescapeString(feed.Channel.Title)
	feed.Channel.Description = html.UnescapeString(feed.Channel.Description)

	for i := range feed.Channel.Items {
		item := &feed.Channel.Items[i]
		item.Title = html.UnescapeString(item.Title)
		item.Description = html.UnescapeString(item.Description)
	}
}
