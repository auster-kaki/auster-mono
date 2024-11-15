package usecase

import (
	"context"

	"github.com/auster-kaki/auster-mono/pkg/app/repository"
	"github.com/auster-kaki/auster-mono/pkg/entity"
)

type DiaryUseCase struct {
	repository repository.Repository
}

func NewDiaryUseCase(r repository.Repository) *DiaryUseCase {
	return &DiaryUseCase{repository: r}
}

func (u *DiaryUseCase) GetDiaries(ctx context.Context, userID entity.UserID) (entity.Diaries, error) {
	diaryUsers, err := u.repository.DiaryUser().GetByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	diaries, err := u.repository.Diary().GetByIDs(ctx, diaryUsers.DiaryIDs())
	if err != nil {
		return nil, err
	}
	return diaries, nil
}
