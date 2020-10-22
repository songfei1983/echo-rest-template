package errors

import "fmt"

func NewHTTPError(status int, err error) HTTPError {
	return HTTPError{Code: status, Message: err.Error()}
}

type HTTPError struct {
	Code    int
	Message string
}

func (e HTTPError) Error() string {
	return fmt.Sprintf("%d: %s", e.Code, e.Message)
}

var _ error = (*HTTPError)(nil)
