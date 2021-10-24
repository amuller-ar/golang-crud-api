package errors

import (
	"errors"
	"fmt"
)

type (
	MultiCauseError struct {
		message string
		causes  []string
	}
)

func NewMultiCauseError(message string, causes ...string) MultiCauseError {
	return MultiCauseError{
		message: message,
		causes:  causes,
	}
}

func (e MultiCauseError) Message() string {
	return e.message
}

func (e MultiCauseError) Causes() []string {
	return e.causes
}

func (e MultiCauseError) ErrorCauses() []error {
	var errs []error
	for i := range e.causes {
		errs = append(errs, errors.New(e.causes[i]))
	}
	return errs
}

func (e MultiCauseError) Error() string {
	if len(e.causes) == 0 {
		return e.message
	}

	return fmt.Sprintf("%s because: %s", e.message, e.causes)
}
