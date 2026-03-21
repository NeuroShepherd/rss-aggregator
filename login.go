package main

import (
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("username is required")
	}

	s.config.CurrentUserName = cmd.args[0]
	err := s.config.Write()
	if err != nil {
		return fmt.Errorf("write config: %w", err)
	}
	fmt.Printf("Logged in as %s\n", s.config.CurrentUserName)

	return nil
}
