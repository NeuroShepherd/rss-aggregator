package main

import (
	"fmt"
	"os"
	"time"
)

func handlerAgg(s *state, cmd command) error {

	if len(cmd.args) != 1 {
		return fmt.Errorf("must supply a parseable time duration e.g. 2m15s")
	}

	timeArg, err := time.ParseDuration(cmd.args[0])
	if err != nil {
		return fmt.Errorf("invalid time argument: %w", err)
	}

	ticker := time.NewTicker(timeArg)
	defer ticker.Stop()

	for ; ; <-ticker.C {
		fmt.Printf("scraping feeds every %s", timeArg)
		scrapeFeeds(s)
		if err != nil {
			fmt.Fprintln(os.Stderr, "scrape feeds:", err)
		}
	}

}
