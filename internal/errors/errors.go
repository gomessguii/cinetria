package errors

import "fmt"

type ErrorWithCode struct {
	code int
	err error
}

func NewErrorWithCode(code int, err error) error {
	return &ErrorWithCode{code, err}
}

func (e ErrorWithCode) Error() string {
	return fmt.Sprintf("%d: %s", e.code, e.err.Error())
}