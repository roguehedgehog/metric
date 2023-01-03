package user

type Repo interface {
	Create(username, password string) (*User, error)
	Exists(username string) bool
}
