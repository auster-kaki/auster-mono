package table

import (
	"context"

	"github.com/auster-kaki/auster-mono/pkg/entity"

	"github.com/uptrace/bun"
)

type Hobby struct {
	db *bun.DB
}

func NewHobby(db *bun.DB) *Hobby { return &Hobby{db: db} }

func (t *Hobby) GetAll(ctx context.Context) (entity.Hobbies, error) {
	res := entity.Hobbies{}
	if err := t.db.NewSelect().Model(&res).Scan(ctx); err != nil {
		return nil, handleError(err)
	}
	return res, nil
}

func (t *Hobby) GetByIDs(ctx context.Context, ids []entity.HobbyID) (entity.Hobbies, error) {
	res := entity.Hobbies{}
	if err := t.db.NewSelect().Model(&res).Where("id IN (?)", ids).Scan(ctx); err != nil {
		return nil, handleError(err)
	}
	return res, nil
}
