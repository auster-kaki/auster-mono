package table

import (
	"github.com/uptrace/bun"
)

type TravelSpotPhoto struct {
	db *bun.DB
}

func NewTravelSpotPhoto(db *bun.DB) *TravelSpotPhoto {
	return &TravelSpotPhoto{db: db}
}
