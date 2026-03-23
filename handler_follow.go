package main

import (
	"context"
	"fmt"

	"time"

	"github.com/google/uuid"

	"github.com/neuroshepherd/rss-aggregator/internal/database"
)

func handlerFollow(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: %s <feed_url>", cmd.name)
	}

	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}

	feed, err := s.db.GetFeedByURL(context.Background(), cmd.args[0])
	if err != nil {
		return fmt.Errorf("failed to get feed: %w", err)
	}

	feedFollow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		UserID:    user.ID,
		FeedID:    feed.ID,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})
	if err != nil {
		return fmt.Errorf("failed to create feed follow: %w", err)
	}

	fmt.Printf("Feed: %s\n", feedFollow.FeedName)
	fmt.Printf("User: %s\n", feedFollow.UserName)
	return nil
}
