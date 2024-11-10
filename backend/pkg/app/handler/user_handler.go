package handler

import (
	"fmt"
	"net/http"

	"github.com/auster-kaki/auster-mono/pkg/app/usecase"
)

type UserHandler struct {
	userUseCase *usecase.UserUseCase
}

func NewUserHandler(u *usecase.UserUseCase) *UserHandler {
	return &UserHandler{userUseCase: u}
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "GetUsers")
}
