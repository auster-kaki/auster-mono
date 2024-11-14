package table

import (
	"context"

	"github.com/auster-kaki/auster-mono/pkg/entity"

	"github.com/uptrace/bun"
)

type UserItineraryDiary struct {
	db *bun.DB
}

func NewUserItineraryDiary(db *bun.DB) *UserItineraryDiary {
	return &UserItineraryDiary{db: db}
}

func (t *UserItineraryDiary) Create(ctx context.Context, ents ...entity.UserItineraryDiary) error {
	if _, err := t.db.NewInsert().Model(&ents).Exec(ctx); err != nil {
		return err
	}
	return nil
}

func (t *UserItineraryDiary) GetByUserID(ctx context.Context, userID entity.UserID) (entity.UserItineraryDiaries, error) {
	res := entity.UserItineraryDiaries{}
	if err := t.db.NewSelect().Model(&res).Where("user_id = ?", userID).Scan(ctx); err != nil {
		return nil, err
	}
	return res, nil
}
