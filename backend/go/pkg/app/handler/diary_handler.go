package handler

import (
	"github.com/auster-kaki/auster-mono/pkg/app/usecase"
)

type DiaryHandler struct {
	diaryUseCase usecase.DiaryUseCase
}

func NewDiaryHandler(u *usecase.DiaryUseCase) []Input {
	//h := &DiaryHandler{diaryUseCase: u}
	return []Input{}
}
