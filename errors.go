package ng

import (
	"fmt"
	"log"
)

//Errors is the container of the standard error
//from jsonapi.org
type Errors struct {
	Errors []*Error `json:"errors"`
}

//NewErrors create new empty container
func NewErrors() *Errors {

	return &Errors{}
}

//Append add new error to the container
func (e *Errors) Append(errs ...*Error) {
	e.Errors = append(e.Errors, errs...)
}

//AppendSlice add new slice of error to the container
func (e *Errors) AppendSlice(errs []*Error) {
	for _, err := range errs {
		e.Errors = append(e.Errors, err)
	}
}

//AppendErrors add new errors from the same errors structure
func (e *Errors) AppendErrors(errs *Errors) {
	if len(errs.Errors) > 0 {
		e.Errors = append(e.Errors, errs.Errors...)
	}
}

//AppendError add new standard error from Go Package
//to the container
func (e *Errors) AppendError(status string, err error) {
	e.Errors = append(e.Errors, NewFromError(status, err))
}

//AppendSliceError add new slice of standard error from
//Go package to the container
func (e *Errors) AppendSliceError(status string, errs []error) {
	for _, err := range errs {
		e.Errors = append(e.Errors, NewFromError(status, err))
	}
}

//NewError add new error with status and detail
func (e *Errors) NewError(status, detail string) {
	e.Errors = append(e.Errors, NewError(status, detail))
}

//NewFieldError add new error with field, and detail
func (e *Errors) NewFieldError(status, field, detail string) {
	e.Errors = append(e.Errors, NewFieldError(status, field, detail))
}

//NotFound add new error with 404 not found
func (e *Errors) NotFound(detail string) {
	e.Errors = append(e.Errors, NotFound(detail))
}

//InternalServerError add new error with 500 internal server error
func (e *Errors) InternalServerError(detail string) {
	e.Errors = append(e.Errors, InternalServerError(detail))
}

//BadRequest add new error with 400 bad request
func (e *Errors) BadRequest(detail string) {
	e.Errors = append(e.Errors, BadRequest(detail))
}

//Unauthorized add new error with 401 unauthorized
func (e *Errors) Unauthorized(detail string) {
	e.Errors = append(e.Errors, Unauthorized(detail))
}

//Forbidden add new error with 403 forbidden
func (e *Errors) Forbidden(detail string) {
	e.Errors = append(e.Errors, Forbidden(detail))
}

//Conflict add new error with 401 conflict
func (e *Errors) Conflict(detail string) {
	e.Errors = append(e.Errors, Conflict(detail))
}

//IsError check if it has more than zero error(s)
func (e *Errors) IsError() bool {
	return len(e.Errors) > 0
}

//LogPrint just print the error message using log from
//Go package
func (e *Errors) LogPrint() {
	for _, err := range e.Errors {
		log.Printf("%s", err.Detail)
		log.Println()
	}
}

//FmtPrint just print the error message using fmt from
//Go Package
func (e *Errors) FmtPrint() {
	for _, err := range e.Errors {
		fmt.Printf("%s", err.Detail)
		fmt.Println()
	}
}
