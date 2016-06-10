package ng_test

import (
	"errors"
	"testing"

	"github.com/ngurajeka/ng"
)

func TestError(t *testing.T) {

	var err *ng.Error

	if err != nil {
		t.Errorf("Err should be null")
	}

	err = ng.NewEmptyError()

	if err == new(ng.Error) {
		t.Errorf("Err should not be null")
	}

	err = ng.NewFieldError("404", "field1", "field1 is required")

	if err.Field == "" && err.Detail == "" {
		t.Errorf("Field and Detail should not be \"\"")
	}

	err = ng.NewFromError("404", errors.New("Empty Data"))

	if err.Detail != "Empty Data" {
		t.Errorf("Detail should not be \"Empty Data\"")
	}

}
