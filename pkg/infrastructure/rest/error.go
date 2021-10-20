package rest

import (
	"fmt"
	"strings"
)

type Error struct {
	Status  int
	Message string
	Causes  []string
}

func NewError(status int, message string, causes ...error) Error {
	var causeText []string
	for _, err := range causes {
		causeText = append(causeText, err.Error())
	}

	return Error{
		Status:  status,
		Message: message,
		Causes:  causeText,
	}
}

func (e Error) Error() string {
	return fmt.Sprintf("%s: %v", e.Message, strings.Join(e.Causes, ", "))
}
