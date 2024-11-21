package entity

import (
	"github.com/uptrace/bun"
)

type TravelSpotID string

type TravelSpot struct {
	ID          TravelSpotID
	VendorID    VendorID
	Title       string
	TakeTime    int
	Description string
	Address     string
	PhotoPath   string

	// これで設定したtableタグでbunがテーブル名を解決する
	bun.BaseModel `bun:"table:travel_spot,alias:ts"`
}

type TravelSpots []*TravelSpot
