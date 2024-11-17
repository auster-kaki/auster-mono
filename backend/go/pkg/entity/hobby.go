package entity

import "github.com/uptrace/bun"

type HobbyID string

type Hobby struct {
	ID   HobbyID
	Name string

	// これで設定したtableタグでbunがテーブル名を解決する
	bun.BaseModel `bun:"table:hobby,alias:h"`
}

type Hobbies []*Hobby
