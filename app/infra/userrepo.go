package infra

import (
	"database/sql"
	"github.com/alexedwards/argon2id"
	"github.com/roguehedgehog/metric/domain/user"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (repo *UserRepo) Create(username, password string) (*user.User, error) {
	hash, err := argon2id.CreateHash(password, argon2id.DefaultParams)
	if err != nil {
		return &user.User{}, user.NewCreationError("hashing password", username, err)
	}

	u := user.New(username, hash)
	if err = repo.save(u); err != nil {
		return u, user.NewCreationError("saving", username, err)
	}

	return u, nil
}

func (repo *UserRepo) Exists(username string) bool {
	var exists bool
	q := "SELECT 1 FROM user WHERE username = ?"
	err := repo.db.QueryRow(q, username).Scan(&exists)
	return err != sql.ErrNoRows && exists
}

func (repo *UserRepo) save(u *user.User) error {
	var id int
	q := "INSERT INTO user (username, hash, created_at) VALUES (?, ?, ?)"
	s, err := repo.db.Prepare(q)
	if err != nil {
		return user.NewCreationError("preparing query", u.Username(), err)
	}

	defer s.Close()
	_, err = s.Exec(u.Username(), u.Hash(), u.CreatedAt())
	if err != nil {
		return user.NewCreationError("saving", u.Username(), err)
	}

	err = repo.db.QueryRow("SELECT LAST_INSERT_ID()").Scan(&id)
	if err != nil {
		return user.NewCreationError("getting id for", u.Username(), err)
	}
	u.SetId(id)
	return nil
}
