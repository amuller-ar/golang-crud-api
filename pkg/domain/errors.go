package domain

import "fmt"

type (
	PropertyNotFoundError struct {
		ID uint
	}
)

func (e PropertyNotFoundError) Error() string {
	return fmt.Sprintf("property id %d not found", e.ID)
}
