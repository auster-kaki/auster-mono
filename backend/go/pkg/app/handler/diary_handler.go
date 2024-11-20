package handler

import (
	"encoding/json"
	"net/http"

	"github.com/auster-kaki/auster-mono/pkg/app/presenter/request"
	"github.com/auster-kaki/auster-mono/pkg/app/presenter/response"
	"github.com/auster-kaki/auster-mono/pkg/app/usecase"
)

type DiaryHandler struct {
	diaryUseCase *usecase.DiaryUseCase
}

func NewDiaryHandler(u *usecase.DiaryUseCase) []Input {
	h := &DiaryHandler{diaryUseCase: u}
	return []Input{
		{method: http.MethodPost, path: "/diaries", handler: h.CreateDiary},
	}
}

func (h *DiaryHandler) CreateDiary(w http.ResponseWriter, r *http.Request) {
	var req request.Diary
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.HandleError(r.Context(), w, err)
		return
	}

	out, err := h.diaryUseCase.CreateDiary(r.Context(), &req)
	if err != nil {
		response.HandleError(r.Context(), w, err)
		return
	}
	response.OK(w, out)
}
