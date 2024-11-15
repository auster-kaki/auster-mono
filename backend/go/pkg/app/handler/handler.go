package handler

import (
	"fmt"
	"net/http"

	"github.com/auster-kaki/auster-mono/pkg/app/repository"
	"github.com/auster-kaki/auster-mono/pkg/app/usecase"
)

func NewHandlerMap(r repository.Repository) map[string]http.HandlerFunc {
	handlerMap := make(map[string]http.HandlerFunc)
	for path, handlerFunc := range NewUserHandler(usecase.NewUserUseCase(r)) {
		if _, ok := handlerMap[path]; ok {
			panic(fmt.Sprintf("duplicate path: %s", path))
		}
		handlerMap[path] = handlerFunc
	}
	for path, handlerFunc := range NewVendorHandler(usecase.NewVendorUseCase(r)) {
		if _, ok := handlerMap[path]; ok {
			panic(fmt.Sprintf("duplicate path: %s", path))
		}
		handlerMap[path] = handlerFunc
	}
	for path, handlerFunc := range NewDiaryHandler(usecase.NewDiaryUseCase(r)) {
		if _, ok := handlerMap[path]; ok {
			panic(fmt.Sprintf("duplicate path: %s", path))
		}
		handlerMap[path] = handlerFunc
	}
	for path, handlerFunc := range NewTravelSpotHandler(usecase.NewTravelSpotUseCase(r)) {
		if _, ok := handlerMap[path]; ok {
			panic(fmt.Sprintf("duplicate path: %s", path))
		}
		handlerMap[path] = handlerFunc
	}
	return handlerMap
}
