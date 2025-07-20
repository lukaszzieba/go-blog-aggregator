package internal

import (
	"github.com/lukaszzieba/go-blog-agregator/internal/database"
)

type State struct {
	db     *database.Queries
	Config *Config
}

func NewState(c *Config, db *database.Queries) *State {
	return &State{Config: c, db: db}
}
