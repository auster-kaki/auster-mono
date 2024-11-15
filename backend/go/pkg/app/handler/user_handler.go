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

func NewUserHandler(u *usecase.UserUseCase) map[string]http.HandlerFunc {
	h := &UserHandler{userUseCase: u}
	return map[string]http.HandlerFunc{
		"/users":      h.GetUsers,
		"/users/{id}": h.GetUser,
	}
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

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.PathValue("id"))
	fmt.Println(r.PathValue("name"))
	response.OK(w, nil)
}
