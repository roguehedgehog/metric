package infra

import (
	"database/sql"
	"fmt"
	"github.com/alexedwards/argon2id"
	"github.com/roguehedgehog/metric/domain/user"
	"github.com/rs/zerolog/log"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (r *UserRepo) Create(username, password string) (*user.User, error) {
	hash, err := argon2id.CreateHash(password, argon2id.DefaultParams)
	if err != nil {
		return &user.User{}, user.NewCreationError("hashing password", username, err)
	}

	u := user.New(username, hash)
	if err = r.save(u); err != nil {
		return u, user.NewCreationError("saving", username, err)
	}

	return u, nil
}

func (r *UserRepo) Exists(username string) bool {
	var exists bool
	q := "SELECT 1 FROM user WHERE username = ?"
	err := r.db.QueryRow(q, username).Scan(&exists)
	return err != sql.ErrNoRows && exists
}

func (r *UserRepo) Get(username string) (u *user.User, err error) {
	err = r.db.QueryRow("SELECT id, username, hash, created_at, updated_at, deleted_at "+
		"FROM user "+
		"WHERE username = ?", username).
		Scan(
			&u.Id,
			&u.Username,
			&u.Hash,
			&u.CreatedAt,
			&u.UpdatedAt,
			&u.DeletedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			err = user.NotFoundError{Username: username}
			return
		}
		err = fmt.Errorf("error finding user %s, %v", username, err)
		return
	}

	return
}

func (r *UserRepo) CheckPassword(u *user.User, password string) bool {
	ok, err := argon2id.ComparePasswordAndHash(password, u.Hash)
	if err != nil {
		log.Err(fmt.Errorf("error trying to compare password hash for %s: %v", u.Username, err))
		return false
	}

	return ok
}

func (r *UserRepo) save(u *user.User) error {
	q := "INSERT INTO user (username, hash, created_at) VALUES (?, ?, ?)"
	s, err := r.db.Prepare(q)
	if err != nil {
		return user.NewCreationError("preparing query", u.Username, err)
	}

	defer s.Close()
	_, err = s.Exec(u.Username, u.Hash, u.CreatedAt)
	if err != nil {
		return user.NewCreationError("saving", u.Username, err)
	}

	err = r.db.QueryRow("SELECT LAST_INSERT_ID()").Scan(&u.Id)
	if err != nil {
		return user.NewCreationError("getting id for", u.Username, err)
	}

	return nil
}
