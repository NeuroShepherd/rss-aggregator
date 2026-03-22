package main

import (
	"context"
	"fmt"

	"time"

	"os"

	"github.com/google/uuid"
	"github.com/neuroshepherd/rss-aggregator/internal/database"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("username is required")
	}

	name := cmd.args[0]

	user, err := s.db.CreateUser(
		context.Background(),
		database.CreateUserParams{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Name:      name,
		},
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, "create user: %w", err)
		os.Exit(1)
	}

	err = s.cfg.SetUser(name)
	if err != nil {
		fmt.Fprintln(os.Stderr, "set user: %w", err)
		os.Exit(1)
	}

	fmt.Printf("Registered user: %s (ID: %s)\n", user.Name, user.ID)
	return nil
}

func handlerUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("get users: %w", err)
	}

	fmt.Printf("Users:\n")
	for _, user := range users {
		if user.Name == s.cfg.CurrentUserName {
			user.Name += " (current)"
		}
		fmt.Printf("* %s\n", user.Name)
	}
	return nil

}
