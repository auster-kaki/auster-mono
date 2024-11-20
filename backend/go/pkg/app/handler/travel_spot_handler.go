package handler

import (
	"context"
	"net/http"

	"github.com/auster-kaki/auster-mono/pkg/app/presenter/response"
	"github.com/auster-kaki/auster-mono/pkg/app/usecase"
	"github.com/auster-kaki/auster-mono/pkg/entity"
)

type TravelSpotHandler struct {
	travelSpotUseCase *usecase.TravelSpotUseCase
}

func NewTravelSpotHandler(u *usecase.TravelSpotUseCase) []Input {
	h := &TravelSpotHandler{travelSpotUseCase: u}
	return []Input{
		{method: http.MethodGet, path: "/travel_spots", handler: h.GetTravelSpots},
	}
}

func (h *TravelSpotHandler) GetTravelSpots(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	userID := r.URL.Query().Get("user_id")
	hobbyID := r.URL.Query().Get("hobby_id")

	out, err := h.travelSpotUseCase.GetTravelSpots(ctx, entity.UserID(userID), entity.HobbyID(hobbyID))
	if err != nil {
		response.HandleError(ctx, w, err)
		return
	}
	response.OK(w, out)
}
