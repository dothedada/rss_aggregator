package main

import (
	"context"
	"fmt"
)

func handlerReset(s *State, cmd command) error {
	err := s.db.DropUsers(context.Background())
	if err != nil {
		return error(err)
	}

	fmt.Println("Clean users DB...")
	return nil
}
