package main

import (
	"context"
)

func handlerFeeds(s *state, cmd command) error {

	feed, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return err
	}

	for _, f := range feed {
		println(f.Name, f.Url, f.UserName)
	}

	return nil
}
