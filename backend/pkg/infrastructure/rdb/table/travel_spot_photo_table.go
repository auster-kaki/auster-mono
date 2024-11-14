package table

import (
	"context"

	"github.com/auster-kaki/auster-mono/pkg/entity"

	"github.com/uptrace/bun"
)

type TravelSpotPhoto struct {
	db *bun.DB
}

func NewTravelSpotPhoto(db *bun.DB) *TravelSpotPhoto {
	return &TravelSpotPhoto{db: db}
}

func (t *TravelSpotPhoto) Create(ctx context.Context, ents ...entity.TravelSpotPhoto) error {
	if _, err := t.db.NewInsert().Model(&ents).Exec(ctx); err != nil {
		return err
	}
	return nil
}

func (t *TravelSpotPhoto) GetByTravelSpotID(ctx context.Context, travelSpotID entity.TravelSpotID) (entity.TravelSpotPhotos, error) {
	res := entity.TravelSpotPhotos{}
	if err := t.db.NewSelect().Model(&res).Where("travel_spot_id = ?", travelSpotID).Scan(ctx); err != nil {
		return nil, err
	}
	return res, nil
}
