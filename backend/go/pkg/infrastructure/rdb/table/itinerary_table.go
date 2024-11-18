package table

import (
	"context"

	"github.com/auster-kaki/auster-mono/pkg/entity"

	"github.com/uptrace/bun"
)

type Itinerary struct {
	db *bun.DB
}

func NewItinerary(db *bun.DB) *Itinerary {
	return &Itinerary{db: db}
}

func (t *Itinerary) Create(ctx context.Context, ents ...entity.Itinerary) error {
	if _, err := t.db.NewInsert().Model(&ents).Exec(ctx); err != nil {
		return handleError(err)
	}
	return nil
}

func (t *Itinerary) GetByUserItineraryHistoryID(ctx context.Context, userItineraryDiaryID entity.UserItineraryHistoryID) (entity.Itineraries, error) {
	res := entity.Itineraries{}
	if err := t.db.NewSelect().Model(&res).Where("user_itinerary_diary_id = ?", userItineraryDiaryID).Scan(ctx); err != nil {
		return nil, handleError(err)
	}
	return res, nil
}
