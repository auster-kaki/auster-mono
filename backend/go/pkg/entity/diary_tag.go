package entity

type DiaryTagID string

type DiaryTag struct {
	ID      DiaryTagID
	DiaryID DiaryID
	Name    string
}

type DiaryTags []*DiaryTag
