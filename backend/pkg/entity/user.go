package entity

import (
	"time"

	"github.com/uptrace/bun"
)

type UserID string

type User struct {
	ID          UserID
	Name        string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedFlag bool

	// これで設定したtableタグでbunがテーブル名を解決する
	bun.BaseModel `bun:"table:user,alias:u"`
}

type Users []*User
