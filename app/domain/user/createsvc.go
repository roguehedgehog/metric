package user

type CreateSvc struct {
	users Repo
}

func NewCreateSvc(userRepo Repo) *CreateSvc {
	return &CreateSvc{users: userRepo}
}

func (srv *CreateSvc) Create(username, password string) error {
	if srv.users.Exists(username) {
		return ExistsError{Username: username}
	}
	_, err := srv.users.Create(username, password)
	return err
}
