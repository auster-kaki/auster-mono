package table

import (
	"context"
	"github.com/auster-kaki/auster-mono/pkg/entity"
	"github.com/uptrace/bun"
)

type DiaryTag struct {
	db *bun.DB
}

func NewDiaryTag(db *bun.DB) *DiaryTag {
	return &DiaryTag{db: db}
}

func (t *DiaryTag) Create(ctx context.Context, ents ...entity.DiaryTag) error {
	if _, err := t.db.NewInsert().Model(&ents).Exec(ctx); err != nil {
		return err
	}
	return nil
}

func (t *DiaryTag) GetByDiaryID(ctx context.Context, diaryID entity.DiaryID) (entity.DiaryTags, error) {
	res := entity.DiaryTags{}
	if err := t.db.NewSelect().Model(&res).Where("diary_id = ?", diaryID).Scan(ctx); err != nil {
		return nil, err
	}
	return res, nil
}
