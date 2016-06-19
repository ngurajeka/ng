package ng

import (
	"net/http"
	"strconv"
)

//Error standard structure of error from jsonapi.org
type Error struct {
	ID     int         `json:"id,omitempty"`
	Code   string      `json:"code,omitempty"`
	Detail string      `json:"detail,omitempty"`
	Links  string      `json:"links,omitempty"`
	Status string      `json:"status,omitempty"`
	Field  string      `json:"field,omitempty"`
	Title  string      `json:"title,omitempty"`
	Source Source      `json:"source,omitempty"`
	Meta   interface{} `json:"meta,omitempty"`
}

//Source is a part of Error Struct
type Source struct {
	Pointer   string `json:"pointer,omitempty"`
	Parameter string `json:"parameter,omitempty"`
}

//NewEmptyError create new standard error,
//but empty one
func NewEmptyError() *Error {
	return &Error{}
}

//NewError create new standard error,
//status and detail
func NewError(status, detail string) *Error {
	return &Error{Status: status, Detail: detail}
}

//NewFieldError create new standard error,
//using field and detail value
func NewFieldError(status, field, detail string) *Error {
	return &Error{
		Field:  field,
		Detail: detail,
		Status: status,
	}
}

//NewFromError create new standard error,
//using standard error from Go Package
func NewFromError(status string, err error) *Error {
	return &Error{
		Detail: err.Error(),
		Status: status,
	}
}

//NewBasicError create new standard error,
//http://jsonapi.org/examples/#error-objects-basics
func NewBasicError(status, pointer, title, detail string) *Error {
	return &Error{
		Status: status,
		Source: Source{Pointer: pointer},
		Title:  title,
		Detail: detail,
	}
}

//NewInvalidQueryParameter create new standard error,
//http://jsonapi.org/examples/#error-objects-invalid-query-parameters
func NewInvalidQueryParameter(parameter, title, detail string) *Error {
	return &Error{
		Source: Source{Parameter: parameter},
		Title:  title,
		Detail: detail,
	}
}

//NotFound create new standard error with 404 not found
func NotFound(detail string) *Error {
	return &Error{
		Code:   strconv.Itoa(http.StatusNotFound),
		Title:  http.StatusText(http.StatusNotFound),
		Detail: detail,
	}
}

//InternalServerError create new standard error with 500 internal server error
func InternalServerError(detail string) *Error {
	return &Error{
		Code:   strconv.Itoa(http.StatusInternalServerError),
		Title:  http.StatusText(http.StatusInternalServerError),
		Detail: detail,
	}
}

//BadRequest create new standard error with 400 bad request
func BadRequest(detail string) *Error {
	return &Error{
		Code:   strconv.Itoa(http.StatusBadRequest),
		Title:  http.StatusText(http.StatusBadRequest),
		Detail: detail,
	}
}

//Unauthorized create new standard error with 401 unauthorized
func Unauthorized(detail string) *Error {
	return &Error{
		Code:   strconv.Itoa(http.StatusUnauthorized),
		Title:  http.StatusText(http.StatusUnauthorized),
		Detail: detail,
	}
}

//Forbidden create new standard error with 403 forbidden
func Forbidden(detail string) *Error {
	return &Error{
		Code:   strconv.Itoa(http.StatusForbidden),
		Title:  http.StatusText(http.StatusForbidden),
		Detail: detail,
	}
}

//Conflict create new standard error with 409 conflict
func Conflict(detail string) *Error {
	return &Error{
		Code:   strconv.Itoa(http.StatusConflict),
		Title:  http.StatusText(http.StatusConflict),
		Detail: detail,
	}
}
