package main

import (
	"context"
	"fmt"

	"github.com/neuroshepherd/rss-aggregator/internal/database"
)

func handlerFollowing(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 0 {
		return fmt.Errorf("usage: %s", cmd.name)
	}

	following, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("get following: %w", err)
	}

	for _, f := range following {
		fmt.Printf("Feed: %s\n", f.FeedName)
	}

	return nil
}
