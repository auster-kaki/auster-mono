package table

import (
	"context"

	"github.com/auster-kaki/auster-mono/pkg/entity"

	"github.com/uptrace/bun"
)

type TravelSpot struct {
	db *bun.DB
}

func NewTravelSpot(db *bun.DB) *TravelSpot {
	return &TravelSpot{db: db}
}

func (t *TravelSpot) Create(ctx context.Context, ents ...entity.TravelSpot) error {
	if _, err := t.db.NewInsert().Model(&ents).Exec(ctx); err != nil {
		return err
	}
	return nil
}

func (t *TravelSpot) GetByVendorID(ctx context.Context, vendorID entity.VendorID) (entity.TravelSpots, error) {
	res := entity.TravelSpots{}
	if err := t.db.NewSelect().Model(&res).Where("vendor_id = ?", vendorID).Scan(ctx); err != nil {
		return nil, err
	}
	return res, nil
}