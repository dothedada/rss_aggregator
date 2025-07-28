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
	fmt.Println("db:", conf.CurrentUserName)

	if err = config.SetUser("mmejiaaaaa :)"); err != nil {
		fmt.Println("No pudo grabar: %w", err)
	}

	conf, err = config.Read()
	if err != nil {
		fmt.Println("some shit happened: %w", err)
	}

	fmt.Println("db:", conf.DbUrl)
	fmt.Println("db:", conf.CurrentUserName)
}
