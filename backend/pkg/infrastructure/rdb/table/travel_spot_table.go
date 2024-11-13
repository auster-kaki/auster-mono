package table

import (
	"github.com/uptrace/bun"
)

type TravelSpot struct {
	db *bun.DB
}

func NewTravelSpot(db *bun.DB) *TravelSpot {
	return &TravelSpot{db: db}
}
