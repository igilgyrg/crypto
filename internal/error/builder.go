package error

import "net/http"

func ErrNotFound(err error, message string) *Error {
	return newError(err, message, ItemNotFound, http.StatusNotFound)
}
