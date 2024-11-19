package entity

import "github.com/uptrace/bun"

type DiaryUser struct {
	DiaryID DiaryID
	UserID  UserID

	// これで設定したtableタグでbunがテーブル名を解決する
	bun.BaseModel `bun:"table:diary_user,alias:du"`
}

type DiaryUsers []*DiaryUser

func (d DiaryUsers) DiaryIDs() []DiaryID {
	ids := make([]DiaryID, len(d))
	for i, v := range d {
		ids[i] = v.DiaryID
	}
	return ids
}
