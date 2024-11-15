package usecase

import (
	"context"

	"github.com/auster-kaki/auster-mono/pkg/app/repository"
	"github.com/auster-kaki/auster-mono/pkg/entity"
)

type ItineraryUseCase struct {
	repository repository.Repository
}

func NewItineraryUseCase(r repository.Repository) *ItineraryUseCase {
	return &ItineraryUseCase{repository: r}
}

func (u *ItineraryUseCase) GetUserHistory(ctx context.Context, userID entity.UserID) (entity.UserItineraryHistories, error) {
	histories, err := u.repository.UserItineraryHistory().GetByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return histories, nil
}
