package ng

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
