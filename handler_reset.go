package main

import (
	"database/sql"
	"fmt"
	"net/url"
	"strings"

	"github.com/pressly/goose/v3"
)

func handlerReset(s *state, cmd command) error {
	parsedURL, err := url.Parse(s.cfg.DBURL)
	if err != nil {
		return fmt.Errorf("parse database url: %w", err)
	}

	dbName := strings.TrimPrefix(parsedURL.Path, "/")
	parsedURL.Path = "/postgres"

	db, err := sql.Open("postgres", parsedURL.String())
	if err != nil {
		return fmt.Errorf("open database: %w", err)
	}
	defer db.Close()

	_, err = db.Exec(fmt.Sprintf("DROP DATABASE IF EXISTS %s;", dbName))
	if err != nil {
		return fmt.Errorf("drop database: %w", err)
	}

	_, err = db.Exec(fmt.Sprintf("CREATE DATABASE %s;", dbName))
	if err != nil {
		return fmt.Errorf("create database: %w", err)
	}

	// open admin connection to the new database to run migrations
	db, err = sql.Open("postgres", s.cfg.DBURL)
	if err != nil {
		return fmt.Errorf("open database: %w", err)
	}
	defer db.Close()

	err = goose.Up(db, "./sql/schema")
	if err != nil {
		return fmt.Errorf("run migrations: %w", err)
	}

	fmt.Printf("Database %s have been reset.\n", dbName)

	return nil
}
