package table

import (
	"context"

	"github.com/auster-kaki/auster-mono/pkg/entity"

	"github.com/uptrace/bun"
)

type UserItineraryHistory struct {
	db *bun.DB
}

func NewUserItineraryDiary(db *bun.DB) *UserItineraryHistory {
	return &UserItineraryHistory{db: db}
}

func (t *UserItineraryHistory) Create(ctx context.Context, ents ...entity.UserItineraryHistory) error {
	//TODO implement me
	panic("implement me")
}

func (t *UserItineraryHistory) GetByUserID(ctx context.Context, userID entity.UserID) (entity.UserItineraryHistories, error) {
	//TODO implement me
	panic("implement me")
}
