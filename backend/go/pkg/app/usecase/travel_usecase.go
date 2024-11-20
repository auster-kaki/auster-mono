package usecase

import (
	"context"
	"fmt"

	"github.com/auster-kaki/auster-mono/pkg/app/repository"
	"github.com/auster-kaki/auster-mono/pkg/entity"
	"github.com/auster-kaki/auster-mono/pkg/util/austerid"
)

type TravelUseCase struct {
	repository repository.Repository
}

func NewTravelSpotUseCase(r repository.Repository) *TravelUseCase {
	return &TravelUseCase{repository: r}
}

func (u *TravelUseCase) GetTravelSpots(ctx context.Context, userID entity.UserID, hobbyID entity.HobbyID) (entity.TravelSpots, error) {
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

type CrateDiaryOutput struct {
	Title       string
	Description string
	Body        []byte
}

func (u *TravelUseCase) CreateDiary(ctx context.Context, userID entity.UserID, travelSpotID entity.TravelSpotID) (*CrateDiaryOutput, error) {
	return &CrateDiaryOutput{
		Title:       fmt.Sprintf("userID %s, travelSpotIDs %v", userID, travelSpotID),
		Description: austerid.Generate[string](),
	}, nil
}
