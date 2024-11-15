package handler

import (
	"net/http"

	"github.com/auster-kaki/auster-mono/pkg/app/usecase"
)

type TravelSpotHandler struct {
	travelSpotUseCase usecase.TravelSpotUseCase
}

func NewTravelSpotHandler(u *usecase.TravelSpotUseCase) map[string]http.HandlerFunc {
	//h := &TravelSpotHandler{travelSpotUseCase: u}
	return map[string]http.HandlerFunc{}
}
