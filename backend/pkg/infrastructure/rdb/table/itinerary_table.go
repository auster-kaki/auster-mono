package table

import (
	"github.com/uptrace/bun"
)

type Itinerary struct {
	db *bun.DB
}

func NewItinerary(db *bun.DB) *Itinerary {
	return &Itinerary{db: db}
}
