package handler

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/auster-kaki/auster-mono/pkg/app/presenter/request"
	"github.com/auster-kaki/auster-mono/pkg/app/presenter/response"
	"github.com/auster-kaki/auster-mono/pkg/app/usecase"
	"github.com/auster-kaki/auster-mono/pkg/entity"
)

type ReservationHandler struct {
	reservationUseCase *usecase.ReservationUseCase
}

func NewReservationHandler(r *usecase.ReservationUseCase) []Input {
	h := &ReservationHandler{reservationUseCase: r}
	return []Input{
		{method: http.MethodPost, path: "/reservations", handler: h.Create},
		{method: http.MethodGet, path: "/reservations", handler: h.Get},
		{method: http.MethodPatch, path: "/reservations/{reservation_id}/diary_photo", handler: h.UpdateDiaryPhoto},
		{method: http.MethodPatch, path: "/reservations/{reservation_id}/diary_description", handler: h.UpdateDescription},
	}
}

func (h *ReservationHandler) Create(w http.ResponseWriter, r *http.Request) {
	var (
		req request.Reservation
		ctx = r.Context()
	)
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.HandleError(ctx, w, err)
		return
	}

	from, err := time.ParseInLocation(time.DateOnly, req.FromDate, time.Local)
	if err != nil {
		response.HandleError(ctx, w, err)
		return
	}
	to, err := time.ParseInLocation(time.DateOnly, req.ToDate, time.Local)
	if err != nil {
		response.HandleError(ctx, w, err)
	}

	out, err := h.reservationUseCase.Create(ctx, &usecase.CreateReservationInput{
		UserID:            entity.UserID(req.UserID),
		TravelSpotID:      entity.TravelSpotID(req.TravelSpotID),
		TravelSpotDiaryID: entity.TravelSpotDiaryID(req.TravelSpotDiaryID),
		From:              from,
		To:                to,
	})
	if err != nil {
		response.HandleError(ctx, w, err)
		return
	}
	response.Created(w, out)
}

func (h *ReservationHandler) Get(w http.ResponseWriter, r *http.Request) {
	var (
		ctx    = r.Context()
		userID = r.URL.Query().Get("user_id")
		filter = r.URL.Query().Get("filter")
	)

	out, err := h.reservationUseCase.GetEndedReservations(ctx, usecase.GetReservationInput{
		UserID: entity.UserID(userID),
		Status: filter,
	})
	if err != nil {
		response.HandleError(ctx, w, err)
		return
	}
	response.OK(w, out)
}

func (h *ReservationHandler) UpdateDiaryPhoto(w http.ResponseWriter, r *http.Request) {
	var (
		ctx = r.Context()
		id  = entity.ReservationID(r.PathValue("reservation_id"))
	)

	file, _, err := r.FormFile("photo")
	if err != nil {
		response.HandleError(ctx, w, err)
		return
	}
	defer file.Close()

	photo, err := io.ReadAll(file)
	if err != nil {
		response.HandleError(ctx, w, err)
		return
	}

	if err := h.reservationUseCase.UpdateDiaryPhoto(ctx, id, photo); err != nil {
		response.HandleError(ctx, w, err)
		return
	}
	response.NoContent(w)
}

func (h *ReservationHandler) UpdateDescription(w http.ResponseWriter, r *http.Request) {
	var (
		req request.DiaryDescription
		id  = entity.ReservationID(r.PathValue("reservation_id"))
		ctx = r.Context()
	)
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.HandleError(ctx, w, err)
		return
	}

	if err := h.reservationUseCase.UpdateDiaryDescription(ctx, id, req.Description); err != nil {
		response.HandleError(ctx, w, err)
		return
	}
	response.NoContent(w)
}
