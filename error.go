package ng

//Error standard structure of error from jsonapi.org
type Error struct {
	ID     int         `json:"id,omitempty"`
	Code   string      `json:"code,omitempty"`
	Detail string      `json:"detail,omitempty"`
	Links  string      `json:"links,omitempty"`
	Status int         `json:"status,omitempty"`
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
func NewFieldError(field, detail string) *Error {
	return &Error{
		Field:  field,
		Detail: detail,
	}
}

//NewFromError create new standard error,
//using standard error from Go Package
func NewFromError(err error) *Error {
	return &Error{
		Detail: err.Error(),
	}
}
