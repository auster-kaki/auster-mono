package table

import (
	"context"

	"github.com/auster-kaki/auster-mono/pkg/entity"

	"github.com/uptrace/bun"
)

type Reservation struct {
	db *bun.DB
}

func NewReservation(db *bun.DB) *Reservation { return &Reservation{db: db} }

func (t *Reservation) Create(ctx context.Context, ents ...entity.Reservation) error {
	if _, err := t.db.NewInsert().Model(&ents).Exec(ctx); err != nil {
		return handleError(err)
	}
	return nil
}

func (t *Reservation) FindByID(ctx context.Context, id entity.ReservationID) (*entity.Reservation, error) {
	res := &entity.Reservation{}
	if err := t.db.NewSelect().Model(res).Where("id = ?", id).Scan(ctx); err != nil {
		return nil, handleError(err)
	}
	return res, nil
}

func (t *Reservation) GetEndedReservations(ctx context.Context, userID entity.UserID) (entity.Reservations, error) {
	res := entity.Reservations{}
	if err := t.db.NewSelect().Model(&res).
		Where("user_id = ?", userID).
		Where("date < NOW()").Scan(ctx); err != nil {
		return nil, handleError(err)
	}
	return res, nil
}

func (t *Reservation) GetUpcomingReservations(ctx context.Context, userID entity.UserID) (entity.Reservations, error) {
	res := entity.Reservations{}
	if err := t.db.NewSelect().Model(&res).
		Where("user_id = ?", userID).
		Where("date >= NOW()").Scan(ctx); err != nil {
		return nil, handleError(err)
	}
	return res, nil
}