package repository

import (
	"context"

	"github.com/auster-kaki/auster-mono/pkg/entity"
)

type Repository interface {
	User() UserRepository
}

type UserRepository interface {
	GetAll(ctx context.Context) (entity.Users, error)
}
