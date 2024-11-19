package usecase

import (
	"context"

	"github.com/auster-kaki/auster-mono/pkg/app/repository"
	"github.com/auster-kaki/auster-mono/pkg/entity"
)

type TravelSpotUseCase struct {
	repository repository.Repository
}

func NewTravelSpotUseCase(r repository.Repository) *TravelSpotUseCase {
	return &TravelSpotUseCase{repository: r}
}

type TravelSpotsGetOutput struct {
	TravelSpots entity.TravelSpots
	Photos      entity.TravelSpotPhotos
}

func (u *TravelSpotUseCase) GetTravelSpots(ctx context.Context, userID entity.UserID, hobbyID entity.HobbyID) (*TravelSpotsGetOutput, error) {
	travelSpotHobbies, err := u.repository.TravelSpotHobby().GetByHobbyID(ctx, hobbyID)
	if err != nil {
		return nil, err
	}
	travelSpotIDs := travelSpotHobbies.TravelSpotIDs()
	travelSpots, err := u.repository.TravelSpot().GetByIDs(ctx, travelSpotIDs)
	if err != nil {
		return nil, err
	}
	travelSpotPhotos, err := u.repository.TravelSpotPhoto().GetByTravelSpotIDs(ctx, travelSpotIDs)
	if err != nil {
		return nil, err
	}
	return &TravelSpotsGetOutput{
		TravelSpots: travelSpots,
		Photos:      travelSpotPhotos,
	}, nil
}
