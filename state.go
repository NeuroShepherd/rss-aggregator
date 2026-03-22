package main

import (
	"github.com/neuroshepherd/rss-aggregator/internal/config"
	"github.com/neuroshepherd/rss-aggregator/internal/database"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}
