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
	//TODO implement me
	panic("implement me")
}

func (t *DiaryTag) GetByDiaryID(ctx context.Context, diaryID entity.DiaryID) (entity.DiaryTags, error) {
	//TODO implement me
	panic("implement me")
}
