package handler

import (
	"net/http"

	"github.com/auster-kaki/auster-mono/pkg/app/presenter/request"
	"github.com/auster-kaki/auster-mono/pkg/app/presenter/response"
	"github.com/auster-kaki/auster-mono/pkg/app/usecase"
	"github.com/auster-kaki/auster-mono/pkg/entity"
)

type EncounterHandler struct {
	encounterUseCase *usecase.EncounterUseCase
}

func NewEncounterHandler(e *usecase.EncounterUseCase) []Input {
	h := &EncounterHandler{encounterUseCase: e}
	return []Input{
		{method: http.MethodGet, path: "/{user_id}/encounters", handler: h.GetEncounters},
		{method: http.MethodPost, path: "/{user_id}/encounters", handler: h.Create},
		{method: http.MethodPatch, path: "/{user_id}/encounters/{id}", handler: h.Update},
	}
}

func (h *EncounterHandler) GetEncounters(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	out, err := h.encounterUseCase.GetEncounters(r.Context(), entity.UserID(userID))
	if err != nil {
		http.Error(w, "failed to get encounters", http.StatusInternalServerError)
		return
	}
	response.OK(w, out)
}

func (h *EncounterHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req request.Encounter
	if err := request.Decode(r, &req); err != nil {
		http.Error(w, "failed to decode request", http.StatusBadRequest)
		return
	}
	if err := h.encounterUseCase.Create(r.Context(), &usecase.CrateEncounterInput{
		UserID:      entity.UserID(r.URL.Query().Get("user_id")),
		Name:        req.Name,
		Place:       req.Place,
		Date:        req.Date,
		Description: req.Description,
	}); err != nil {
		http.Error(w, "failed to create encounter", http.StatusInternalServerError)
		return
	}
	response.Created(w, nil)
}

func (h *EncounterHandler) Update(w http.ResponseWriter, r *http.Request) {
	var req request.Encounter
	if err := request.Decode(r, &req); err != nil {
		http.Error(w, "failed to decode request", http.StatusBadRequest)
		return
	}
	if err := h.encounterUseCase.Update(r.Context(), &usecase.UpdateEncounterInput{
		ID:          entity.EncounterID(r.URL.Query().Get("id")),
		Name:        req.Name,
		Place:       req.Place,
		Date:        req.Date,
		Description: req.Description,
	}); err != nil {
		http.Error(w, "failed to update encounter", http.StatusInternalServerError)
		return
	}
	response.OK(w, nil)
}
