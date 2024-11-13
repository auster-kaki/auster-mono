package entity

import "github.com/uptrace/bun"

type ItineraryResultID string

type ItineraryResult struct {
	ID          ItineraryResultID
	ItineraryID ItineraryID
	Done        bool
	PhotoPath   string
	Description string

	// これで設定したtableタグでbunがテーブル名を解決する
	bun.BaseModel `bun:"table:itinerary_result,alias:ir"`
}

type ItineraryResults []*ItineraryResult
