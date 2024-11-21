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
		handlers    = [][]Input{
			NewUserHandler(usecase.NewUserUseCase(r)),
			NewVendorHandler(usecase.NewVendorUseCase(r)),
			NewTravelSpotHandler(usecase.NewTravelSpotUseCase(r)),
			NewReservationHandler(usecase.NewReservationUseCase(r)),
			NewEncounterHandler(usecase.NewEncounterUseCase(r)),
		}
	)
	for _, hs := range handlers {
		for _, h := range hs {
			path := fmt.Sprintf("%s %s", h.method, h.path)
			if _, ok := alreadyPath[path]; ok {
				panic(fmt.Sprintf("duplicate path: %s", path))
			}
			alreadyPath[path] = struct{}{}
			handlerMap[path] = h.handler
		}
	}
	return handlerMap
}
