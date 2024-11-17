package handler

import (
	"github.com/auster-kaki/auster-mono/pkg/app/usecase"
)

type TravelSpotHandler struct {
	travelSpotUseCase usecase.TravelSpotUseCase
}

func NewTravelSpotHandler(u *usecase.TravelSpotUseCase) []Input {
	//h := &TravelSpotHandler{travelSpotUseCase: u}
	return []Input{}
}
