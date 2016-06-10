package ng_test

import (
	"errors"
	"testing"

	"github.com/ngurajeka/ng"
)

func TestErrors(t *testing.T) {

	var errs *ng.Errors

	if errs != nil {
		t.Errorf("Errs should be null")
	}

	errs = ng.NewErrors()

	if len(errs.Errors) > 0 {
		t.Errorf("Null Errors should have 0 slice of errors")
	}

	errs.Append(ng.NewEmptyError())

	if len(errs.Errors) != 1 {
		t.Errorf("Err should have 1 slice of errors now")
	}

	errs.AppendSlice([]*ng.Error{ng.NewEmptyError()})

	if len(errs.Errors) != 2 {
		t.Errorf("Err should have 2 slice of errors now")
	}

	errs.AppendError("404", errors.New(""))

	if len(errs.Errors) != 3 {
		t.Errorf("Err should have 3 slice of errors now")
	}

	errs.AppendSliceError("404", []error{errors.New("")})

	if len(errs.Errors) != 4 {
		t.Errorf("Err should have 4 slice of errors now")
	}

	errs.NewFieldError("404", "field1", "field1 is required")

	if len(errs.Errors) != 5 {
		t.Errorf("Err should have 5 slice of errors now")
	}

	if errs.IsError() == false {
		t.Errorf("Err should have return true")
	}

}
