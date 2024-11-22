package entity

import (
	"github.com/uptrace/bun"
)

type TravelSpotID string

type TravelSpotLevel int

const (
	TravelSpotLevel1 TravelSpotLevel = iota + 1
	TravelSpotLevel2
	TravelSpotLevel3
)

type TravelSpot struct {
	ID          TravelSpotID
	VendorID    VendorID
	Title       string
	TakeTime    int
	Description string
	Address     string
	PhotoPath   string
	Level       TravelSpotLevel

	// これで設定したtableタグでbunがテーブル名を解決する
	bun.BaseModel `bun:"table:travel_spot,alias:ts"`
}

type TravelSpots []*TravelSpot
