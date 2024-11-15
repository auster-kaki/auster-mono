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
	//TODO implement me
	panic("implement me")
}

func (t *DiaryUser) GetByUserID(ctx context.Context, userID entity.UserID) (entity.DiaryUsers, error) {
	//TODO implement me
	panic("implement me")
}
