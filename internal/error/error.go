package error

import (
	"encoding/json"
	"net/http"
)

type Error struct {
	Err          error        `json:"-"`
	Message      string       `json:"message"`
	InternalCode InternalCode `json:"internal_code"`
	HttpCode     int          `json:"code"`
}

func newError(err error, message string, internalCode InternalCode, code int) *Error {
	return &Error{
		Err:          err,
		Message:      message,
		InternalCode: internalCode,
		HttpCode:     code,
	}
}

func (e *Error) Error() string {
	return e.Err.Error()
}

func (e *Error) Unwrap() error {
	return e.Err
}

func (e *Error) Marshal() []byte {
	marshal, err := json.Marshal(e)
	if err != nil {
		return nil
	}

	return marshal
}

func systemError(err error) *Error {
	return newError(err, "system error", SystemError, http.StatusInternalServerError)
}
