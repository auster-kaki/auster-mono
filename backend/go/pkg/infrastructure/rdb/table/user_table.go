package table

import (
	"context"

	"github.com/auster-kaki/auster-mono/pkg/entity"

	"github.com/uptrace/bun"
)

type User struct {
	db *bun.DB
}

func NewUser(db *bun.DB) *User {
	return &User{db: db}
}

func (t *User) GetAll(ctx context.Context) (entity.Users, error) {
	res := entity.Users{}
	if err := t.db.NewSelect().Model(&entity.User{}).Scan(ctx, &res); err != nil {
		return nil, err
	}
	return res, nil
}
