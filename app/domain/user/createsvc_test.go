//go:build integration.test

package user_test

import (
	"github.com/alexedwards/argon2id"
	"github.com/google/go-cmp/cmp"
	"github.com/roguehedgehog/metric/domain/user"
	"github.com/roguehedgehog/metric/infra"
	"testing"
	"time"
)

func TestCreateUser(t *testing.T) {
	// given
	infra.PrimaryDb.Exec("TRUNCATE TABLE user")
	username := "t"
	password := "secret"

	svc := user.NewCreateSvc(infra.NewUserRepo(infra.PrimaryDb))

	// when
	err := svc.Create(username, password)
	if err != nil {
		t.Error(err)
	}

	// then
	var savedUser, hash string
	var created time.Time
	err = infra.PrimaryDb.
		QueryRow("SELECT username, hash, created_at FROM user WHERE user_id = 1").
		Scan(&savedUser, &hash, &created)
	if err != nil {
		t.Error(err)
	}

	if diff := cmp.Diff(username, savedUser); diff != "" {
		t.Error(diff)
	}

	ok, err := argon2id.ComparePasswordAndHash(password, hash)
	if err != nil || !ok {
		t.Error("password hashes do not match")
	}

	if created.IsZero() ||
		created.After(time.Now().Add(time.Duration(1)*time.Second)) ||
		created.Before(time.Now().Add(time.Duration(-2)*time.Minute)) {
		t.Errorf("created time %s does not seem correct, its %s", created, time.Now())
	}
}
