package usecase

import (
	"context"
	"errors"
	"fmt"
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

type ListOutputInput struct {
	UserID entity.UserID
	Status string
}

type ListOutput struct {
	Reservations   entity.Reservations
	TravelSpotByID map[entity.TravelSpotID]*entity.TravelSpot
	DiaryByID      map[entity.TravelSpotDiaryID]*entity.TravelSpotDiary
}

func (u *ReservationUseCase) List(ctx context.Context, input ListOutputInput) (*ListOutput, error) {
	var (
		reservations entity.Reservations
		err          error
	)
	switch input.Status {
	case "yet":
		reservations, err = u.repository.Reservation().GetUpcomingReservations(ctx, input.UserID)
	case "done":
		reservations, err = u.repository.Reservation().GetEndedReservations(ctx, input.UserID)
	default:
		return nil, errors.New(fmt.Sprintf("invalid status: %s", input.Status))
	}
	if err != nil {
		return nil, err
	}

	travelSpots, err := u.repository.TravelSpot().GetByIDs(ctx, reservations.TravelSpotIDs())
	if err != nil {
		return nil, err
	}

	travelSpotByID := make(map[entity.TravelSpotID]*entity.TravelSpot, len(travelSpots))
	for _, travelSpot := range travelSpots {
		travelSpotByID[travelSpot.ID] = travelSpot
	}

	diaries, err := u.repository.TravelSpotDiary().GetByIDs(ctx, reservations.TravelSpotDiaryIDs())
	if err != nil {
		return nil, err
	}
	diaryByID := make(map[entity.TravelSpotDiaryID]*entity.TravelSpotDiary, len(diaries))
	for _, diary := range diaries {
		diaryByID[diary.ID] = diary
	}

	return &ListOutput{
		Reservations:   reservations,
		TravelSpotByID: travelSpotByID,
		DiaryByID:      diaryByID,
	}, nil
}

type GetReservationOutput struct {
	Reservation           *entity.Reservation
	TravelSpot            *entity.TravelSpot
	TravelSpotItineraries entity.TravelSpotItineraries
	TravelSpotDiary       *entity.TravelSpotDiary
}

func (u *ReservationUseCase) GetReservation(ctx context.Context, id entity.ReservationID) (*GetReservationOutput, error) {
	reservation, err := u.repository.Reservation().FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	travelSpot, err := u.repository.TravelSpot().FindByID(ctx, reservation.TravelSpotID)
	if err != nil {
		return nil, err
	}

	travelSpotItineraries, err := u.repository.TravelSpotItinerary().GetByTravelSpotID(ctx, reservation.TravelSpotID)
	if err != nil {
		return nil, err
	}
	travelSpotDiary, err := u.repository.TravelSpotDiary().FindByID(ctx, reservation.TravelSpotDiaryID)
	if err != nil {
		return nil, err
	}

	return &GetReservationOutput{
		Reservation:           reservation,
		TravelSpot:            travelSpot,
		TravelSpotItineraries: travelSpotItineraries,
		TravelSpotDiary:       travelSpotDiary,
	}, nil
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
