package entity

import (
	"github.com/uptrace/bun"
)

type TravelSpotItineraryID string

type TravelSpotItineraryKind string

const (
	// TravelSpotItineraryKindKindMove は「移動」
	TravelSpotItineraryKindKindMove TravelSpotItineraryKind = "move"
	// TravelSpotItineraryKindKindSpot は「観光地」
	TravelSpotItineraryKindKindSpot TravelSpotItineraryKind = "spot"
)

type TravelSpotItinerary struct {
	ID           TravelSpotItineraryID
	Kind         TravelSpotItineraryKind
	TravelSpotID TravelSpotID
	TakeTime     int
	Price        int
	Order        int

	// これで設定したtableタグでbunがテーブル名を解決する
	bun.BaseModel `bun:"table:travel_spot_itinerary,alias:tsi"`
}

type TravelSpotItineraries []*TravelSpotItinerary

func (e *TravelSpotItineraries) IDs() []TravelSpotItineraryID {
	ids := make([]TravelSpotItineraryID, len(*e))
	for i, v := range *e {
		ids[i] = v.ID
	}
	return ids
}
