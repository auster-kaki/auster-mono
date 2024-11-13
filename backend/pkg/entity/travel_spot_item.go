package entity

import (
	"github.com/uptrace/bun"
)

type TravelSpotItemID string

type TravelSpotItem struct {
	ID           TravelSpotItemID
	TravelSpotID TravelSpotID
	Name         string
	Number       int

	// これで設定したtableタグでbunがテーブル名を解決する
	bun.BaseModel `bun:"table:travel_spot_item,alias:tsi"`
}

type TravelSpotItems []*TravelSpotItem
