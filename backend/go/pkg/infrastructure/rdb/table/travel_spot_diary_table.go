package table

import (
	"context"

	"github.com/auster-kaki/auster-mono/pkg/entity"

	"github.com/uptrace/bun"
)

type TravelSpotDiary struct {
	db *bun.DB
}

func NewTravelSpotDiary(db *bun.DB) *TravelSpotDiary {
	return &TravelSpotDiary{db: db}
}

func (t *TravelSpotDiary) Create(ctx context.Context, ents ...entity.TravelSpotDiary) error {
	if _, err := t.db.NewInsert().Model(&ents).Exec(ctx); err != nil {
		return handleError(err)
	}
	return nil
}

func (t *TravelSpotDiary) Update(ctx context.Context, ent *entity.TravelSpotDiary) error {
	if _, err := t.db.NewUpdate().Model(ent).Exec(ctx); err != nil {
		return handleError(err)
	}
	return nil
}

func (t *TravelSpotDiary) FindByID(ctx context.Context, id entity.TravelSpotDiaryID) (*entity.TravelSpotDiary, error) {
	res := &entity.TravelSpotDiary{}
	if err := t.db.NewSelect().Model(&res).Where("id = ?", id).Scan(ctx); err != nil {
		return nil, handleError(err)
	}
	return res, nil
}

func (t *TravelSpotDiary) GetByIDs(ctx context.Context, ids []entity.TravelSpotDiaryID) (entity.TravelSpotDiaries, error) {
	res := entity.TravelSpotDiaries{}
	if err := t.db.NewSelect().Model(&res).Where("id IN (?)", bun.In(ids)).Scan(ctx); err != nil {
		return nil, handleError(err)
	}
	return res, nil
}
