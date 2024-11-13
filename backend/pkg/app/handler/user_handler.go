package handler

import (
	"fmt"
	"net/http"

	"github.com/auster-kaki/auster-mono/pkg/app/presenter/response"
	"github.com/auster-kaki/auster-mono/pkg/app/usecase"
)

type UserHandler struct {
	userUseCase *usecase.UserUseCase
}

func NewUserHandler(u *usecase.UserUseCase) *UserHandler {
	return &UserHandler{userUseCase: u}
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	out, err := h.userUseCase.GetUsers(r.Context())
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to get users: %v", err), http.StatusInternalServerError)
		return
	}
	for i, user := range out {
		fmt.Println(i, user)
	}
	response.OK(w, out)
}
