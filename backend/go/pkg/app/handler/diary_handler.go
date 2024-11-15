package handler

import (
	"net/http"

	"github.com/auster-kaki/auster-mono/pkg/app/usecase"
)

type DiaryHandler struct {
	diaryUseCase usecase.DiaryUseCase
}

func NewDiaryHandler(u *usecase.DiaryUseCase) map[string]http.HandlerFunc {
	//h := &DiaryHandler{diaryUseCase: u}
	return map[string]http.HandlerFunc{}
}
