package ng

import "fmt"

const (
	IsRequired      = "Is Required"
	IsInvalidValue  = "Has Invalid Value"
	ShouldLengthMin = "Should have minimum length of"
	ShouldMin       = "Should have minimum value of"
	ShouldLengthMax = "Should have maximum length of"
	ShouldMax       = "Should have maximum value of"
	ShouldBetween   = "Should between"
)

func Required(name string, i interface{}) *Errors {
	errs := NewErrors()

	switch v := i.(type) {
	case string:
		if v == "" {
			errs.Append(NewFieldError("404", name, isRequired(name)))
		}
	case int, int8, int16, int32, int64:
		if v == 0 {
			errs.Append(NewFieldError("400", name, isInvalidValue(name)))
		}
	}

	return errs
}

func LengthMin(name string, i string, min int) *Errors {
	errs := NewErrors()

	if len(i) < min {
		errs.Append(NewFieldError("400", name, shouldLengthMin(name, min)))
	}

	return errs
}

func LengthMax(name string, i string, max int) *Errors {
	errs := NewErrors()

	if len(i) > max {
		errs.Append(NewFieldError("400", name, shouldLengthMax(name, max)))
	}

	return errs
}

func Range(name string, i, min, max int) *Errors {
	errs := NewErrors()

	if min > i && i < max {
		errs.Append(NewFieldError("400", name, shouldBetween(name, min, max)))
	}
	return errs
}

func isRequired(name string) string {
	return fmt.Sprintf("%s %s", name, IsRequired)
}

func isInvalidValue(name string) string {
	return fmt.Sprintf("%s %s", name, IsInvalidValue)
}

func shouldLengthMin(name string, min int) string {
	return fmt.Sprintf("%s %s %d", name, ShouldLengthMin, min)
}

func shouldMin(name string, min int) string {
	return fmt.Sprintf("%s %s %d", name, ShouldMin, min)
}

func shouldLengthMax(name string, max int) string {
	return fmt.Sprintf("%s %s %d", name, ShouldLengthMax, max)
}

func shouldMax(name string, max int) string {
	return fmt.Sprintf("%s %s %d", name, ShouldMax, max)
}

func shouldBetween(name string, min, max int) string {
	return fmt.Sprintf("%s %s %d and %d", name, ShouldBetween, min, max)
}
