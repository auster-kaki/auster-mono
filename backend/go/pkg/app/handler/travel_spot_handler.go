package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/auster-kaki/auster-mono/pkg/app/presenter/request"
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

	var req request.TravelSpot
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.HandleError(ctx, w, err)
		return
	}

	out, err := h.travelSpotUseCase.GetTravelSpots(ctx, entity.UserID(req.UserID), entity.HobbyID(req.HobbyID))
	if err != nil {
		response.HandleError(ctx, w, err)
		return
	}
	response.OK(w, out)
}
