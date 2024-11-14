package entity

type DiaryUser struct {
	DiaryID DiaryID
	UserID  UserID
}

type DiaryUsers []*DiaryUser
