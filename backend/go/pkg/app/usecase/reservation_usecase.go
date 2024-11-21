package usecase

import (
	"context"
	"time"

	"github.com/auster-kaki/auster-mono/pkg/app/repository"
	"github.com/auster-kaki/auster-mono/pkg/entity"
	"github.com/auster-kaki/auster-mono/pkg/util/austerid"
	"github.com/auster-kaki/auster-mono/pkg/util/austerstorage"
)

type ReservationUseCase struct {
	repository repository.Repository
}

func NewReservationUseCase(r repository.Repository) *ReservationUseCase {
	return &ReservationUseCase{repository: r}
}

type CreateReservationInput struct {
	UserID            entity.UserID
	TravelSpotID      entity.TravelSpotID
	TravelSpotDiaryID entity.TravelSpotDiaryID
	From              time.Time
	To                time.Time
}

func (u *ReservationUseCase) Create(ctx context.Context, input *CreateReservationInput) (entity.ReservationID, error) {
	id := austerid.Generate[entity.ReservationID]()
	if err := u.repository.Reservation().Create(ctx, entity.Reservation{
		ID:                id,
		UserID:            input.UserID,
		TravelSpotID:      input.TravelSpotID,
		TravelSpotDiaryID: input.TravelSpotDiaryID,
		FromDate:          input.From,
		ToDate:            input.To,
	}); err != nil {
		return "", err
	}
	return id, nil
}

type GetReservationInput struct {
	UserID entity.UserID
	Status string
}

func (u *ReservationUseCase) GetEndedReservations(ctx context.Context, input GetReservationInput) (entity.Reservations, error) {
	var (
		res entity.Reservations
		err error
	)
	switch input.Status {
	case "yet":
		res, err = u.repository.Reservation().GetUpcomingReservations(ctx, input.UserID)
	case "done":
		res, err = u.repository.Reservation().GetEndedReservations(ctx, input.UserID)
	}
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (u *ReservationUseCase) UpdateDiaryPhoto(ctx context.Context, id entity.ReservationID, photo []byte) error {
	reservation, err := u.repository.Reservation().FindByID(ctx, id)
	if err != nil {
		return err
	}

	travelSpotDiary, err := u.repository.TravelSpotDiary().FindByID(ctx, reservation.TravelSpotDiaryID)
	if err != nil {
		return err
	}

	if _, err := austerstorage.Save(austerstorage.JPEG, travelSpotDiary.PhotoPath, photo); err != nil {
		return err
	}
	return nil
}

func (u *ReservationUseCase) UpdateDiaryDescription(ctx context.Context, id entity.ReservationID, description string) error {
	reservation, err := u.repository.Reservation().FindByID(ctx, id)
	if err != nil {
		return err
	}

	travelSpotDiary, err := u.repository.TravelSpotDiary().FindByID(ctx, reservation.TravelSpotDiaryID)
	if err != nil {
		return err
	}

	travelSpotDiary.Description = description
	if err := u.repository.TravelSpotDiary().Update(ctx, travelSpotDiary); err != nil {
		return err
	}
	return nil
}
