package table

import (
	"github.com/uptrace/bun"
)

type ItineraryResult struct {
	db *bun.DB
}

func NewItineraryResult(db *bun.DB) *ItineraryResult {
	return &ItineraryResult{db: db}
}
