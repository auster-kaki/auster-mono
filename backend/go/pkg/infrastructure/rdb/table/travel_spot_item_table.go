package table

import (
	"context"

	"github.com/auster-kaki/auster-mono/pkg/entity"

	"github.com/uptrace/bun"
)

type TravelSpotItem struct {
	db *bun.DB
}

func NewTravelSpotItem(db *bun.DB) *TravelSpotItem {
	return &TravelSpotItem{db: db}
}

func (t *TravelSpotItem) Create(ctx context.Context, ents ...entity.TravelSpotItem) error {
	if _, err := t.db.NewInsert().Model(&ents).Exec(ctx); err != nil {
		return err
	}
	return nil
}

func (t *TravelSpotItem) GetByTravelSpotID(ctx context.Context, travelSpotID entity.TravelSpotID) (entity.TravelSpotItems, error) {
	res := entity.TravelSpotItems{}
	if err := t.db.NewSelect().Model(&res).Where("travel_spot_id = ?", travelSpotID).Scan(ctx); err != nil {
		return nil, err
	}
	return res, nil
}
