package main

import (
	"simple-erp/internal/database"
)

type apiConfig struct {
	db *database.Queries
}
type errorResponse struct {
	Error string `json:"error"`
}

var grainTypes = map[string]struct{}{
	"wheat":     {},
	"barley":    {},
	"corn":      {},
	"sunflower": {},
	"canola":    {},
}
