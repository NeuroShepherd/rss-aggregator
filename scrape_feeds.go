package main

import (
	"context"
	"fmt"
)

func scrapeFeeds(s *state) error {
	//  get the next feed to fetch from the db:

	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}

	err = s.db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		return err
	}

	rss_feed, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		return err
	}

	for _, item := range rss_feed.Channel.Item {
		fmt.Println(item.Title)
	}

	return nil
}
