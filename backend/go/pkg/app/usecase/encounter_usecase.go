package usecase

import (
	"context"
	"time"

	"github.com/auster-kaki/auster-mono/pkg/app/repository"
	"github.com/auster-kaki/auster-mono/pkg/entity"
	"github.com/auster-kaki/auster-mono/pkg/util/austerid"
)

type EncounterUseCase struct {
	repository repository.Repository
}

func NewEncounterUseCase(r repository.Repository) *EncounterUseCase {
	return &EncounterUseCase{repository: r}
}

func (u *EncounterUseCase) GetEncounters(ctx context.Context, userID entity.UserID) (entity.Encounters, error) {
	return u.repository.Encounter().GetByUserID(ctx, userID)
}

type CrateEncounterInput struct {
	UserID      entity.UserID
	Name        string
	Place       string
	Date        time.Time
	Description string
}

func (u *EncounterUseCase) CreateEncounter(ctx context.Context, input *CrateEncounterInput) error {
	return u.repository.Encounter().Create(ctx, entity.Encounter{
		ID:          austerid.Generate[entity.EncounterID](),
		Name:        input.Name,
		Place:       input.Place,
		UserID:      input.UserID,
		Date:        input.Date,
		Description: input.Description,
	})
}
