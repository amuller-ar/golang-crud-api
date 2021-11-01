package domain

import "fmt"

type (
	PropertyNotFoundError struct {
		ID uint
	}
	UserNotFoundError struct {
		Email string
	}
)

func (e PropertyNotFoundError) Error() string {
	return fmt.Sprintf("property id %d not found", e.ID)
}

func (u UserNotFoundError) Error() string {
	return fmt.Sprintf("user with email: %s not found", u.Email)
}
