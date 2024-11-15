package table

import (
	"context"
	"github.com/auster-kaki/auster-mono/pkg/entity"
	"github.com/uptrace/bun"
)

type Diary struct {
	db *bun.DB
}

func NewDairy(db *bun.DB) *Diary {
	return &Diary{db: db}
}

func (t *Diary) Create(ctx context.Context, ents ...entity.Diary) error {
	//TODO implement me
	panic("implement me")
}

func (t *Diary) FindByID(ctx context.Context, id entity.DiaryID) (*entity.Diary, error) {
	//TODO implement me
	panic("implement me")
}

func (t *Diary) GetByIDs(ctx context.Context, ids []entity.DiaryID) (entity.Diaries, error) {
	//TODO implement me
	panic("implement me")
}
