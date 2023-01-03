package user

import (
	"time"
)

type User struct {
	id        int
	username  string
	hash      string
	createdAt time.Time
	updatedAt time.Time
	deletedAt time.Time
}

func New(username, hash string) *User {
	return &User{
		username:  username,
		hash:      hash,
		createdAt: time.Now(),
	}
}

func (u User) Username() string {
	return u.username
}

func (u User) Hash() string {
	return u.hash
}

func (u User) CreatedAt() time.Time {
	return u.createdAt
}

func (u User) UpdatedAt() time.Time {
	return u.updatedAt
}

func (u User) DeletedAt() time.Time {
	return u.deletedAt
}

func (u *User) SetId(id int) {
	u.id = id
}
