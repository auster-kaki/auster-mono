package usecase

import (
	"context"

	"github.com/auster-kaki/auster-mono/pkg/app/repository"
	"github.com/auster-kaki/auster-mono/pkg/entity"
)

type UserUseCase struct {
	repository repository.Repository
}

func NewUserUseCase(r repository.Repository) *UserUseCase {
	return &UserUseCase{repository: r}
}

func (u *UserUseCase) GetUsers(ctx context.Context) (entity.Users, error) {
	users, err := u.repository.User().GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}
