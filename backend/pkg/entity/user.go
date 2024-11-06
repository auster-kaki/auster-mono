package entity

type UserID string

type User struct {
	ID   UserID
	Name string
}

type Users []*User
