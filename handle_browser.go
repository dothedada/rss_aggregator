package main

import (
	"context"
	"fmt"
	"strconv"
)

func handlerBrowse(s *State, cmd command) error {
	postAmount := int32(2)
	if len(cmd.args) >= 1 {
		argInt, err := strconv.ParseInt(cmd.args[0], 10, 32)
		if err != nil {
			return fmt.Errorf("usage: %s <integer>", cmd.name)
		}

		postAmount = int32(argInt)
	}

	posts, err := s.db.GetPostsForUser(context.Background(), postAmount)
	if err != nil {
		return fmt.Errorf("Cannot get the posts data: %w", err)
	}

	fmt.Println("Your posts feed:")
	fmt.Println("================")
	fmt.Println()
	for i, post := range posts {
		fmt.Printf("%d) %s\n", i+1, post.Title)
		fmt.Println(post.Description)
		fmt.Println(post.Url)
		fmt.Println()
	}

	return nil
}
