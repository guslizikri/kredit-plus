package util

import "net/http"

type HttpError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Err     interface{} `json:"error,omitempty"`
}

func (e *HttpError) Error() string {
	return e.Message
}

func InternalServerError(message string) *HttpError {
	return &HttpError{
		Code:    http.StatusInternalServerError,
		Message: message,
	}
}

func BadRequest(message string) *HttpError {
	return &HttpError{
		Code:    http.StatusBadRequest,
		Message: message,
	}
}

func Unauthorized(message string) *HttpError {
	return &HttpError{
		Code:    http.StatusUnauthorized,
		Message: message,
	}
}

func Forbidden(message string) *HttpError {
	return &HttpError{
		Code:    http.StatusForbidden,
		Message: message,
	}
}

func NotFound(message string) *HttpError {
	return &HttpError{
		Code:    http.StatusNotFound,
		Message: message,
	}
}

func Conflict(message string) *HttpError {
	return &HttpError{
		Code:    http.StatusConflict,
		Message: message,
	}
}

func ToHttpError(err error) *HttpError {
	if httpErr, ok := err.(*HttpError); ok {
		return httpErr
	}

	return InternalServerError(err.Error())
}
