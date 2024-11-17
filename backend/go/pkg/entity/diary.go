package entity

import (
	"github.com/uptrace/bun"
	"time"
)

type DiaryID string

type Diary struct {
	ID        DiaryID
	Title     string
	Date      time.Time
	PhotoPath string
	Body      string

	// これで設定したtableタグでbunがテーブル名を解決する
	bun.BaseModel `bun:"table:diary,alias:d"`
}

type Diaries []*Diary
