package table

import (
	"context"
	"github.com/auster-kaki/auster-mono/pkg/entity"
	"github.com/uptrace/bun"
)

type Dairy struct {
	db *bun.DB
}

func NewDairy(db *bun.DB) *Dairy {
	return &Dairy{db: db}
}

func (t *Dairy) Create(ctx context.Context, ents ...entity.Diary) error {
	//TODO implement me
	panic("implement me")
}

func (t *Dairy) FindByID(ctx context.Context, id entity.DiaryID) (*entity.Diary, error) {
	//TODO implement me
	panic("implement me")
}

func (t *Dairy) GetByID(ctx context.Context, id entity.DiaryID) (entity.Diaries, error) {
	//TODO implement me
	panic("implement me")
}
