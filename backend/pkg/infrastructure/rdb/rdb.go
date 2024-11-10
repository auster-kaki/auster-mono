package rdb

import (
	"github.com/auster-kaki/auster-mono/pkg/app/repository"
	"github.com/auster-kaki/auster-mono/pkg/infrastructure/rdb/table"
)

type rdb struct {
	user *table.User
}

func NewDB() (*rdb, error) {
	return &rdb{
		user: table.NewUser(nil),
	}, nil
}

func (r *rdb) User() repository.UserRepository {
	return r.user
}
