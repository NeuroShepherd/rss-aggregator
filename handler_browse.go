package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/neuroshepherd/rss-aggregator/internal/database"
)

func handlerBrowse(s *state, cmd command) error {
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}

	var limit int32 = 2
	if len(cmd.args) == 1 {
		parsed, err := strconv.Atoi(cmd.args[0])
		if err != nil {
			return err
		}
		limit = int32(parsed)
	}

	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  limit,
	})
	if err != nil {
		return err
	}

	for _, post := range posts {
		fmt.Printf("%s\n%s\n\n", post.Title, post.Url)
	}

	return nil
}
