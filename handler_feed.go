package main

import (
	"context"
	"fmt"

	"time"

	"github.com/google/uuid"

	"github.com/neuroshepherd/rss-aggregator/internal/database"
)

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.args) < 2 {
		return fmt.Errorf("name and url must be provided")
	}
	name := cmd.args[0]
	url := cmd.args[1]

	if name == "" || url == "" {
		return fmt.Errorf("name and url must be provided")
	}

	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("get user: %w", err)
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

	fmt.Printf("Feed created: %+v\n", feed)
	return nil
}
