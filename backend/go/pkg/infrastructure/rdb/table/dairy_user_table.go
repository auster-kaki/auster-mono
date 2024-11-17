package table

import (
	"context"
	"github.com/auster-kaki/auster-mono/pkg/entity"
	"github.com/uptrace/bun"
)

type DiaryUser struct {
	db *bun.DB
}

func NewDairyUser(db *bun.DB) *DiaryUser {
	return &DiaryUser{db: db}
}

func (t *DiaryUser) Create(ctx context.Context, ents ...entity.DiaryUser) error {
	if _, err := t.db.NewInsert().Model(&ents).Exec(ctx); err != nil {
		return err
	}
	return nil
}

func (t *DiaryUser) GetByUserID(ctx context.Context, userID entity.UserID) (entity.DiaryUsers, error) {
	res := entity.DiaryUsers{}
	if err := t.db.NewSelect().Model(&res).Where("user_id = ?", userID).Scan(ctx); err != nil {
		return nil, err
	}
	return res, nil
}
