package main

import (
	"github.com/JtorresBear/gator_project/internal/config"
	"github.com/JtorresBear/gator_project/internal/database"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}
