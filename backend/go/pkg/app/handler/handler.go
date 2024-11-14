package handler

import (
	"net/http"

	"github.com/auster-kaki/auster-mono/pkg/app/repository"
	"github.com/auster-kaki/auster-mono/pkg/app/usecase"
)

func NewHandlerMap(r repository.Repository) map[string]http.HandlerFunc {
	handlerMap := make(map[string]http.HandlerFunc)
	{
		u := usecase.NewUserUseCase(r)
		h := NewUserHandler(u)
		handlerMap["/users"] = h.GetUsers
		handlerMap["/users/{id}/{name}"] = h.GetUser
	}
	return handlerMap
}
