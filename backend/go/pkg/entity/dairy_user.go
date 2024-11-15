package entity

type DiaryUser struct {
	DiaryID DiaryID
	UserID  UserID
}

type DiaryUsers []*DiaryUser

func (d DiaryUsers) DiaryIDs() []DiaryID {
	ids := make([]DiaryID, len(d))
	for i, v := range d {
		ids[i] = v.DiaryID
	}
	return ids
}
