package user

type CreateSvc struct {
	userRepo Repo
}

func NewCreateSvc(userRepo Repo) *CreateSvc {
	return &CreateSvc{userRepo: userRepo}
}

func (srv *CreateSvc) Create(username, password string) error {
	if srv.userRepo.Exists(username) {
		return ExistsError{Username: username}
	}
	_, err := srv.userRepo.Create(username, password)
	return err
}
