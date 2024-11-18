package table

import (
	"context"

	"github.com/auster-kaki/auster-mono/pkg/entity"

	"github.com/uptrace/bun"
)

type ItineraryResult struct {
	db *bun.DB
}

func NewItineraryResult(db *bun.DB) *ItineraryResult {
	return &ItineraryResult{db: db}
}

func (t *ItineraryResult) Create(ctx context.Context, ents ...entity.ItineraryResult) error {
	if _, err := t.db.NewInsert().Model(&ents).Exec(ctx); err != nil {
		return handleError(err)
	}
	return nil
}

func (t *ItineraryResult) GetByItineraryID(ctx context.Context, itineraryID entity.ItineraryID) (entity.ItineraryResults, error) {
	res := entity.ItineraryResults{}
	if err := t.db.NewSelect().Model(&res).Where("itinerary_id = ?", itineraryID).Scan(ctx); err != nil {
		return nil, handleError(err)
	}
	return res, nil
}
