package entity

import (
	"github.com/uptrace/bun"
)

type TravelSpotPhotoID string

type TravelSpotPhoto struct {
	ID           TravelSpotPhotoID
	TravelSpotID TravelSpotID
	Name         string
	Path         string

	// これで設定したtableタグでbunがテーブル名を解決する
	bun.BaseModel `bun:"table:travel_spot_photo,alias:tsp"`
}

type TravelSpotPhotos []*TravelSpotPhoto
