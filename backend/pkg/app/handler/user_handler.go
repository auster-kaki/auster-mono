package handler

import (
	"github.com/auster-kaki/auster-mono/pkg/app/repository"
	"github.com/auster-kaki/auster-mono/pkg/app/usecase"
)

type UserHandler struct {
	userUseCase *usecase.UserUseCase
}

func NewUserHandler(r repository.Repository) *UserHandler {
	return &UserHandler{userUseCase: usecase.NewUserUseCase(r)}
}
