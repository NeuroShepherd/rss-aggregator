package main

import (
	"context"
	"fmt"

	"time"

	"github.com/google/uuid"

	"github.com/neuroshepherd/rss-aggregator/internal/database"
)

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.args) < 2 {
		return fmt.Errorf("name and url must be provided")
	}
	name := cmd.args[0]
	url := cmd.args[1]

	if name == "" || url == "" {
		return fmt.Errorf("name and url must be provided")
	}

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      name,
		Url:       url,
		UserID:    user.ID,
	})
	if err != nil {
		return fmt.Errorf("create feed: %w", err)
	}

	// add a feed follow record for the current user
	_, err = s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		UserID:    user.ID,
		FeedID:    feed.ID,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})
	if err != nil {
		return fmt.Errorf("create feed follow: %w", err)
	}

	fmt.Printf("Feed created: %+v\n", feed)

	return nil
}

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		if s.cfg.CurrentUserName == "" {
			return fmt.Errorf("not logged in")
		}

		user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
		if err != nil {
			return fmt.Errorf("get user: %w", err)
		}

		return handler(s, cmd, user)
	}
}
