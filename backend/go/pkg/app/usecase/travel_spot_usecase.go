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

func (u *TravelSpotUseCase) GetTravelSpots(ctx context.Context, userID entity.UserID, hobbyID entity.HobbyID) (entity.TravelSpots, error) {
	travelSpotHobbies, err := u.repository.TravelSpotHobby().GetByHobbyID(ctx, hobbyID)
	if err != nil {
		return nil, err
	}

	travelSpots, err := u.repository.TravelSpot().GetByIDs(ctx, travelSpotHobbies.TravelSpotIDs())
	if err != nil {
		return nil, err
	}
	return travelSpots, nil
}
