package main

import (
	"context"
	"fmt"
)

func handlerFollowing(s *state, cmd command) error {
	if len(cmd.args) != 0 {
		return fmt.Errorf("usage: %s", cmd.name)
	}

	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
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
