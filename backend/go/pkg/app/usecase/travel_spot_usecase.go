package usecase

import (
	"github.com/auster-kaki/auster-mono/pkg/app/repository"
)

type TravelSpotUseCase struct {
	repository repository.Repository
}

func NewTravelSpotUseCase(r repository.Repository) *TravelSpotUseCase {
	return &TravelSpotUseCase{repository: r}
}
