package main

import (
	"context"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("username is required")
	}

	user, err := s.db.GetUser(context.Background(), cmd.args[0])
	if err != nil {
		return fmt.Errorf("get user: %w", err)
	}

	s.cfg.CurrentUserName = user.Name
	err = s.cfg.Write()
	if err != nil {
		return fmt.Errorf("write config: %w", err)
	}
	fmt.Printf("Logged in as %s\n", s.cfg.CurrentUserName)

	return nil
}
