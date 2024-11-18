package handler

import (
	"fmt"
	"net/http"

	"github.com/auster-kaki/auster-mono/pkg/app/repository"
	"github.com/auster-kaki/auster-mono/pkg/app/usecase"
)

type Input struct {
	method  string
	path    string
	handler http.HandlerFunc
}

func NewHandlerMap(r repository.Repository) map[string]http.HandlerFunc {
	var (
		alreadyPath = make(map[string]struct{})
		handlerMap  = make(map[string]http.HandlerFunc)
	)
	for _, h := range NewUserHandler(usecase.NewUserUseCase(r)) {
		path := fmt.Sprintf("%s %s", h.method, h.path)
		if _, ok := alreadyPath[path]; ok {
			panic(fmt.Sprintf("duplicate path: %s", path))
		}
		alreadyPath[path] = struct{}{}
		handlerMap[path] = h.handler
	}
	for _, h := range NewVendorHandler(usecase.NewVendorUseCase(r)) {
		path := fmt.Sprintf("%s %s", h.method, h.path)
		if _, ok := alreadyPath[path]; ok {
			panic(fmt.Sprintf("duplicate path: %s", path))
		}
		alreadyPath[path] = struct{}{}
		handlerMap[path] = h.handler
	}
	for _, h := range NewDiaryHandler(usecase.NewDiaryUseCase(r)) {
		path := fmt.Sprintf("%s %s", h.method, h.path)
		if _, ok := alreadyPath[path]; ok {
			panic(fmt.Sprintf("duplicate path: %s", path))
		}
		alreadyPath[path] = struct{}{}
		handlerMap[path] = h.handler
	}
	for _, h := range NewTravelSpotHandler(usecase.NewTravelSpotUseCase(r)) {
		path := fmt.Sprintf("%s %s", h.method, h.path)
		if _, ok := alreadyPath[path]; ok {
			panic(fmt.Sprintf("duplicate path: %s", path))
		}
		alreadyPath[path] = struct{}{}
		handlerMap[path] = h.handler
	}
	for _, h := range NewEncounterHandler(usecase.NewEncounterUseCase(r)) {
		path := fmt.Sprintf("%s %s", h.method, h.path)
		if _, ok := alreadyPath[path]; ok {
			panic(fmt.Sprintf("duplicate path: %s", path))
		}
		alreadyPath[path] = struct{}{}
		handlerMap[path] = h.handler
	}
	return handlerMap
}
