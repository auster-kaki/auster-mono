package table

import (
	"context"

	"github.com/auster-kaki/auster-mono/pkg/entity"

	"github.com/uptrace/bun"
)

type UserItineraryHistory struct {
	db *bun.DB
}

func NewUserItineraryHistory(db *bun.DB) *UserItineraryHistory {
	return &UserItineraryHistory{db: db}
}

func (t *UserItineraryHistory) Create(ctx context.Context, ents ...entity.UserItineraryHistory) error {
	if _, err := t.db.NewInsert().Model(&ents).Exec(ctx); err != nil {
		return handleError(err)
	}
	return nil
}

func (t *UserItineraryHistory) GetByUserID(ctx context.Context, userID entity.UserID) (entity.UserItineraryHistories, error) {
	res := entity.UserItineraryHistories{}
	if err := t.db.NewSelect().Model(&res).Where("user_id = ?", userID).Scan(ctx); err != nil {
		return nil, handleError(err)
	}
	return res, nil
}
