package entity

import "time"

type DiaryID string

type Diary struct {
	ID        DiaryID
	Title     string
	Date      time.Time
	PhotoPath string
	Body      string
}

type Diaries []*Diary
