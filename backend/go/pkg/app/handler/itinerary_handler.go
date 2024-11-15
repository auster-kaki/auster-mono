package handler

import (
	"net/http"

	"github.com/auster-kaki/auster-mono/pkg/app/usecase"
)

type ItineraryHandler struct {
	itineraryUseCase usecase.ItineraryUseCase
}

func NewItineraryHandler(u *usecase.ItineraryUseCase) map[string]http.HandlerFunc {
	//h := &ItineraryHandler{itineraryUseCase: u}
	return map[string]http.HandlerFunc{}
}
