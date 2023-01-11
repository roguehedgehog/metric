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
	return fmt.Sprintf("user with Username %s exists", e.Username)
}

type NotFoundError struct {
	Username string
}

func (e NotFoundError) Error() string {
	return fmt.Sprintf("User %s not found", e.Username)
}

type NotAuthorisedError struct {
	Username string
	Action   string
}

func (e NotAuthorisedError) Error() string {
	return fmt.Sprintf("User %s is not authorised to %s", e.Username, e.Action)
}
