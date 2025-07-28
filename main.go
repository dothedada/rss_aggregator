package main

import (
	"fmt"

	"github.com/dothedada/rss_aggregator/internal/config"
)

func main() {
	conf, err := config.Read()
	if err != nil {
		fmt.Println("some shit happened: %w", err)
	}

	fmt.Println("db:", conf.DbUrl)
	fmt.Println("username:", conf.CurrentUserName)

	conf.SetUser("Carajillo")

	conf, err = config.Read()
	if err != nil {
		fmt.Println("some shit happened: %w", err)
	}

	fmt.Printf("again.... : %+v", conf)
}
