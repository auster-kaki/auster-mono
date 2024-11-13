package entity

import "github.com/uptrace/bun"

type ItineraryID string

type ItineraryKind string

const (
	// ItineraryKindMove は「移動」
	ItineraryKindMove ItineraryKind = "move"
	// ItineraryKindSpot は「観光地」
	ItineraryKindSpot ItineraryKind = "spot"
)

type Itinerary struct {
	ID                   ItineraryID
	UserItineraryDiaryID UserItineraryDiaryID
	Kind                 ItineraryKind
	TravelSpotID         TravelSpotID
	From                 string
	To                   string
	TakeTime             int
	Order                int

	// これで設定したtableタグでbunがテーブル名を解決する
	bun.BaseModel `bun:"table:itinerary,alias:i"`
}

type Itineraries []*Itinerary
