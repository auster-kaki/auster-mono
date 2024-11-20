package handler

import (
	"context"
	"encoding/json"
	"mime/multipart"
	"net/http"

	"github.com/auster-kaki/auster-mono/pkg/app/presenter/request"
	"github.com/auster-kaki/auster-mono/pkg/app/presenter/response"
	"github.com/auster-kaki/auster-mono/pkg/app/usecase"
	"github.com/auster-kaki/auster-mono/pkg/entity"
)

type TravelHandler struct {
	travelUseCase *usecase.TravelUseCase
}

func NewTravelSpotHandler(u *usecase.TravelUseCase) []Input {
	h := &TravelHandler{travelUseCase: u}
	return []Input{
		{method: http.MethodGet, path: "/travel_spots", handler: h.GetTravelSpots},
		{method: http.MethodPost, path: "/travel_spots/diary", handler: h.CreateDiary},
	}
}

func (h *TravelHandler) GetTravelSpots(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	var req request.TravelSpot
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.HandleError(ctx, w, err)
		return
	}

	out, err := h.travelUseCase.GetTravelSpots(ctx, entity.UserID(req.UserID), entity.HobbyID(req.HobbyID))
	if err != nil {
		response.HandleError(ctx, w, err)
		return
	}
	response.OK(w, out)
}

func (h *TravelHandler) CreateDiary(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	var req request.CreateDiary
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.HandleError(ctx, w, err)
		return
	}

	out, err := h.travelUseCase.CreateDiary(ctx, entity.UserID(req.UserID), entity.TravelSpotID(req.TravelSpotID))
	if err != nil {
		response.HandleError(ctx, w, err)
		return
	}

	w.Header().Set("Content-Type", "multipart/mixed; boundary=boundary")

	// マルチパートレスポンスを生成
	multipartWriter := multipart.NewWriter(w)
	if err := multipartWriter.SetBoundary("boundary123"); err != nil {
		response.HandleError(r.Context(), w, err)
		return
	}

	// ユーザー情報をjsonで返す
	diaryJSON, err := json.Marshal(map[string]any{
		"title":       out.Title,
		"description": out.Description,
	})
	if err != nil {
		response.HandleError(ctx, w, err)
		return
	}

	// ユーザー情報をマルチパートレスポンスに書き込む
	part, err := multipartWriter.CreatePart(map[string][]string{
		"Content-Type": {"application/json"},
	})
	if err != nil {
		response.HandleError(r.Context(), w, err)
		return
	}

	if _, err := part.Write(diaryJSON); err != nil {
		response.HandleError(ctx, w, err)
		return
	}

	// 画像をマルチパートレスポンスに書き込む
	imagePart, err := multipartWriter.CreatePart(map[string][]string{
		"Content-Type": {"image/jpeg"},
	})
	if err != nil {
		response.HandleError(r.Context(), w, err)
		return
	}

	if _, err := imagePart.Write(out.Body); err != nil {
		response.HandleError(ctx, w, err)
		return
	}

	// マルチパートレスポンスを閉じる
	if err := multipartWriter.Close(); err != nil {
		response.HandleError(ctx, w, err)
		return
	}
}
