package interceptor

import "github.com/google/go-safeweb/safehttp"

type Error struct {
	code safehttp.StatusCode
}

func NewError(code int) *Error {
	return &Error{
		code: safehttp.StatusCode(code),
	}
}

func (e Error) Code() safehttp.StatusCode {
	return e.code
}