package entity

import (
	"time"

	"github.com/uptrace/bun"
)

type ReservationID string

type Reservation struct {
	ID                ReservationID
	UserID            UserID
	TravelSpotID      TravelSpotID
	TravelSpotDiaryID TravelSpotDiaryID
	FromDate          time.Time
	ToDate            time.Time
	IsOffer           bool

	// これで設定したtableタグでbunがテーブル名を解決する
	bun.BaseModel `bun:"table:reservation,alias:r"`
}

type Reservations []*Reservation
