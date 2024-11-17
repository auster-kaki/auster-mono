package entity

import "github.com/uptrace/bun"

type DiaryTagID string

type DiaryTag struct {
	ID      DiaryTagID
	DiaryID DiaryID
	Name    string

	// これで設定したtableタグでbunがテーブル名を解決する
	bun.BaseModel `bun:"table:diary_tag,alias:dt"`
}

type DiaryTags []*DiaryTag
