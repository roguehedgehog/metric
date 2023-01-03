package user

import "fmt"

type CreationError struct {
	action   string
	username string
	parent   error
}

func NewCreationError(action string, username string, parent error) *CreationError {
	return &CreationError{action: action, username: username, parent: parent}
}

func (e CreationError) Error() string {
	return fmt.Sprintf("error %s user %s: %v", e.action, e.username, e.parent)
}

type ExistsError struct {
	Username string
}

func (e ExistsError) Error() string {
	return fmt.Sprintf("user with username %s exists", e.Username)
}
