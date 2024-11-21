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
	Reservations                        entity.Reservations
	TravelSpotItinerariesByTravelSpotID map[entity.TravelSpotID]entity.TravelSpotItineraries
	TravelSpotDiaryByID                 map[entity.TravelSpotDiaryID]*entity.TravelSpotDiary
	PhotoByTravelSpotDiaryID            map[entity.TravelSpotDiaryID][]byte
}

func (u *ReservationUseCase) List(ctx context.Context, input ListOutputInput) (*ListOutput, error) {
	var (
		res entity.Reservations
		err error
	)
	switch input.Status {
	case "yet":
		res, err = u.repository.Reservation().GetUpcomingReservations(ctx, input.UserID)
	case "done":
		res, err = u.repository.Reservation().GetEndedReservations(ctx, input.UserID)
	default:
		return nil, errors.New(fmt.Sprintf("invalid status: %s", input.Status))
	}
	if err != nil {
		return nil, err
	}

	travelSpotItineraries, err := u.repository.TravelSpotItinerary().GetByTravelSpotIDs(ctx, res.TravelSpotIDs())
	if err != nil {
		return nil, err
	}

	travelSpotDiaries, err := u.repository.TravelSpotDiary().GetByIDs(ctx, res.TravelSpotDiaryIDs())
	if err != nil {
		return nil, err
	}

	photoByTravelSpotDiaryID := map[entity.TravelSpotDiaryID][]byte{}
	travelSpotDiaryByID := map[entity.TravelSpotDiaryID]*entity.TravelSpotDiary{}
	for _, travelSpotDiary := range travelSpotDiaries {
		photo, err := austerstorage.Get(travelSpotDiary.PhotoPath)
		if err != nil {
			return nil, err
		}
		photoByTravelSpotDiaryID[travelSpotDiary.ID] = photo
		travelSpotDiaryByID[travelSpotDiary.ID] = travelSpotDiary
	}

	travelSpotItinerariesByTravelSpotID := map[entity.TravelSpotID]entity.TravelSpotItineraries{}
	for _, travelSpotItinerary := range travelSpotItineraries {
		travelSpotItinerariesByTravelSpotID[travelSpotItinerary.TravelSpotID] = append(travelSpotItinerariesByTravelSpotID[travelSpotItinerary.TravelSpotID], travelSpotItinerary)
	}

	return &ListOutput{
		Reservations:                        res,
		TravelSpotItinerariesByTravelSpotID: travelSpotItinerariesByTravelSpotID,
		TravelSpotDiaryByID:                 travelSpotDiaryByID,
		PhotoByTravelSpotDiaryID:            photoByTravelSpotDiaryID,
	}, nil
}

type GetReservationOutput struct {
	Reservation           *entity.Reservation
	TravelSpotItineraries entity.TravelSpotItineraries
	TravelSpotDiary       *entity.TravelSpotDiary
	Photo                 []byte
}

func (u *ReservationUseCase) GetReservation(ctx context.Context, id entity.ReservationID) (*GetReservationOutput, error) {
	reservation, err := u.repository.Reservation().FindByID(ctx, id)
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

	photo, err := austerstorage.Get(travelSpotDiary.PhotoPath)
	if err != nil {
		return nil, err
	}

	return &GetReservationOutput{
		Reservation:           reservation,
		TravelSpotItineraries: travelSpotItineraries,
		TravelSpotDiary:       travelSpotDiary,
		Photo:                 photo,
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
