package user

type Repo interface {
	Create(username, password string) (*User, error)
	Exists(username string) bool
	Get(username string) (*User, error)
	CheckPassword(u *User, password string) bool
}
