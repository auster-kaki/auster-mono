package table

import (
	"context"
	"github.com/auster-kaki/auster-mono/pkg/entity"
	"github.com/uptrace/bun"
)

type DairyUser struct {
	db *bun.DB
}

func NewDairyUserTable(db *bun.DB) *DairyUser {
	return &DairyUser{db: db}
}

func (t *DairyUser) Create(ctx context.Context, ents ...entity.DiaryUser) error {
	//TODO implement me
	panic("implement me")
}

func (t *DairyUser) GetByUserID(ctx context.Context, userID entity.UserID) (entity.DiaryUsers, error) {
	//TODO implement me
	panic("implement me")
}
