package ng

import (
	"fmt"
	"log"
)

//Errors is the container of the standard error
//from jsonapi.org
type Errors struct {
	Errors []*Error
}

//NewErrors create new empty container
func NewErrors() *Errors {

	return &Errors{}
}

//Append add new error to the container
func (e *Errors) Append(err *Error) {
	e.Errors = append(e.Errors, err)
}

//AppendSlice add new slice of error to the container
func (e *Errors) AppendSlice(errs []*Error) {
	for _, err := range errs {
		e.Errors = append(e.Errors, err)
	}
}

//AppendError add new standard error from Go Package
//to the container
func (e *Errors) AppendError(err error) {
	e.Errors = append(e.Errors, NewFromError(err))
}

//AppendSliceError add new slice of standard error from
//Go package to the container
func (e *Errors) AppendSliceError(errs []error) {
	for _, err := range errs {
		e.Errors = append(e.Errors, NewFromError(err))
	}
}

//NewFieldError add new error with field, and detail
func (e *Errors) NewFieldError(field, detail string) {
	e.Errors = append(e.Errors, NewFieldError(field, detail))
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
		log.Printf("Errors: %v", err)
	}
}

//FmtPrint just print the error message using fmt from
//Go Package
func (e *Errors) FmtPrint() {
	for _, err := range e.Errors {
		fmt.Printf("%s", err.Detail)
		fmt.Printf("Errors: %v", err)
	}
}
