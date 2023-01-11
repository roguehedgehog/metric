package user

type LoginSvc struct {
	users Repo
}

func NewLoginSvc(repo Repo) *LoginSvc {
	return &LoginSvc{users: repo}
}

func (svc *LoginSvc) Login(username, password string) error {
	if !svc.users.Exists(username) {
		return NotFoundError{username}
	}

	u, err := svc.users.Get(username)
	if err != nil {
		return err
	}

	if !svc.users.CheckPassword(u, password) {
		return NotAuthorisedError{u.Username, "login"}
	}

	return nil
}
