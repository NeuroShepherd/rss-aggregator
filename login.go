package main

import (
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("username is required")
	}

	s.cfg.CurrentUserName = cmd.args[0]
	err := s.cfg.Write()
	if err != nil {
		return fmt.Errorf("write config: %w", err)
	}
	fmt.Printf("Logged in as %s\n", s.cfg.CurrentUserName)

	return nil
}
