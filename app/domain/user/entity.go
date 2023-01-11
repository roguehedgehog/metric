package user

import (
	"time"
)

type User struct {
	Id        int
	Username  string
	Hash      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func New(username, hash string) *User {
	return &User{
		Username:  username,
		Hash:      hash,
		CreatedAt: time.Now(),
	}
}
