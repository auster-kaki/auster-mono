package table

import (
	"context"
	"database/sql"

	"github.com/auster-kaki/auster-mono/pkg/entity"
)

type User struct {
	db *sql.DB
}

func NewUser(db *sql.DB) *User {
	return &User{db: db}
}

func (t *User) GetAll(ctx context.Context) (entity.Users, error) {
	return nil, nil
}
