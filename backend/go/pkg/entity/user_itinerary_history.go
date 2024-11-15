package entity

import (
	"time"

	"github.com/uptrace/bun"
)

type UserItineraryHistoryID string

type UserItineraryHistory struct {
	ID     UserItineraryHistoryID
	UserID UserID
	Name   string
	Date   time.Time

	// これで設定したtableタグでbunがテーブル名を解決する
	bun.BaseModel `bun:"table:user_itinerary_diary,alias:uid"`
}

type UserItineraryHistories []*UserItineraryHistory
