package table

import (
	"github.com/uptrace/bun"
)

type TravelSpotItem struct {
	db *bun.DB
}

func NewTravelSpotItem(db *bun.DB) *TravelSpotItem {
	return &TravelSpotItem{db: db}
}
