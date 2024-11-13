package entity

import (
	"time"

	"github.com/uptrace/bun"
)

type UserItineraryDiaryID string

type UserItineraryDiary struct {
	ID     UserItineraryDiaryID
	UserID UserID
	Name   string
	Date   time.Time

	// これで設定したtableタグでbunがテーブル名を解決する
	bun.BaseModel `bun:"table:user_itinerary_diary,alias:uid"`
}

type UserItineraryDiaries []*UserItineraryDiary
