package handler

import (
	"encoding/json"
	"io"
	"mime/multipart"
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
		{method: http.MethodGet, path: "/reservations", handler: h.List},
		{method: http.MethodGet, path: "/reservations/{reservation_id}", handler: h.GetReservation},
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

func (h *ReservationHandler) List(w http.ResponseWriter, r *http.Request) {
	var (
		ctx    = r.Context()
		userID = r.URL.Query().Get("user_id")
		filter = r.URL.Query().Get("filter")
	)

	out, err := h.reservationUseCase.List(ctx, usecase.ListOutputInput{
		UserID: entity.UserID(userID),
		Status: filter,
	})
	if err != nil {
		response.HandleError(ctx, w, err)
		return
	}

	res := map[string]any{
		"reservations": []any{},
	}
	for _, reservation := range out.Reservations {
		res["reservations"] = append(res["reservations"].([]any), map[string]any{
			"id":                      reservation.ID,
			"from_date":               reservation.FromDate,
			"to_date":                 reservation.ToDate,
			"travel_spot_title":       out.TravelSpotByID[reservation.TravelSpotID].Title,
			"travel_spot_description": out.TravelSpotByID[reservation.TravelSpotID].Description,
			"diary_photo_path":        out.DiaryByID[reservation.TravelSpotDiaryID].PhotoPath,
		})
	}
	response.OK(w, res)
}

func (h *ReservationHandler) GetReservation(w http.ResponseWriter, r *http.Request) {
	var (
		ctx = r.Context()
		id  = entity.ReservationID(r.PathValue("reservation_id"))
	)

	out, err := h.reservationUseCase.GetReservation(ctx, id)
	if err != nil {
		response.HandleError(ctx, w, err)
		return
	}

	// 画像ファイルとjsonを返す処理
	w.Header().Set("Content-Type", "multipart/mixed; boundary=boundary")

	// マルチパートレスポンスを生成
	multipartWriter := multipart.NewWriter(w)
	if err := multipartWriter.SetBoundary("boundary123"); err != nil {
		response.HandleError(r.Context(), w, err)
		return
	}

	reservationJSON, err := json.Marshal(map[string]any{
		"Reservation":           out.Reservation,
		"TravelSpotItineraries": out.TravelSpotItineraries,
		"TravelSpotDiary":       out.TravelSpotDiary,
	})
	if err != nil {
		response.HandleError(ctx, w, err)
		return
	}

	// ユーザー情報をマルチパートレスポンスに書き込む
	reservationPart, err := multipartWriter.CreatePart(map[string][]string{
		"Content-Type": {"application/json"},
	})
	if err != nil {
		response.HandleError(ctx, w, err)
		return
	}

	if _, err := reservationPart.Write(reservationJSON); err != nil {
		response.HandleError(ctx, w, err)
		return
	}

	// 画像をマルチパートレスポンスに書き込む
	if out.Photo != nil {
		diaryPhotoPart, err := multipartWriter.CreatePart(map[string][]string{
			"Content-Type": {"image/jpeg"},
		})
		if err != nil {
			response.HandleError(ctx, w, err)
			return
		}
		if _, err := diaryPhotoPart.Write(out.Photo); err != nil {
			response.HandleError(ctx, w, err)
			return
		}
	}

	// マルチパートレスポンスを閉じる
	if err := multipartWriter.Close(); err != nil {
		response.HandleError(ctx, w, err)
		return
	}
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
