package table

import (
	"github.com/uptrace/bun"
)

type UserItineraryDiary struct {
	db *bun.DB
}

func NewUserItineraryDiary(db *bun.DB) *UserItineraryDiary {
	return &UserItineraryDiary{db: db}
}
