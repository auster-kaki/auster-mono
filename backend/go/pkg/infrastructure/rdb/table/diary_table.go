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
	if _, err := t.db.NewInsert().Model(&ents).Exec(ctx); err != nil {
		return handleError(err)
	}
	return nil
}

func (t *Diary) FindByID(ctx context.Context, id entity.DiaryID) (*entity.Diary, error) {
	diary := &entity.Diary{}
	if err := t.db.NewSelect().Model(diary).Where("id = ?", id).Scan(ctx); err != nil {
		return nil, handleError(err)
	}
	return diary, nil
}

func (t *Diary) GetByIDs(ctx context.Context, ids []entity.DiaryID) (entity.Diaries, error) {
	res := entity.Diaries{}
	if err := t.db.NewSelect().Model(&res).Where("id IN (?)", ids).Scan(ctx); err != nil {
		return nil, handleError(err)
	}
	return res, nil
}
