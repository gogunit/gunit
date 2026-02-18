package hammy_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/gogunit/gunit/eye"
	a "github.com/gogunit/gunit/hammy"
)

type customError struct {
	message string
}

func (e *customError) Error() string {
	return e.message
}

func Test_is_error_success(t *testing.T) {
	var err = errors.New("hello world")
	a.New(t).Is(a.Error(err))
}

func Test_is_error_failure(t *testing.T) {
	aSpy := eye.Spy()
	a.New(aSpy).Is(a.Error(nil))
	aSpy.HadErrorContaining(t, "got <nil>, want error")
}

func Test_nil_error_success(t *testing.T) {
	a.New(t).Is(a.NilError(nil))
}

func Test_nil_error_failure(t *testing.T) {
	aSpy := eye.Spy()
	var err = errors.New("hello world")
	a.New(aSpy).Is(a.NilError(err))
	aSpy.HadErrorContaining(t, "got <hello world>, want nil error")
}

func Test_error_is_success(t *testing.T) {
	target := errors.New("boom")
	err := fmt.Errorf("wrapped: %w", target)

	a.New(t).Is(a.ErrorIs(err, target))
}

func Test_error_is_failure(t *testing.T) {
	aSpy := eye.Spy()
	target := errors.New("boom")
	err := errors.New("different")

	a.New(aSpy).Is(a.ErrorIs(err, target))
	aSpy.HadErrorContaining(t, "got <different>, want error matching <boom>")
}

func Test_error_as_success(t *testing.T) {
	err := fmt.Errorf("wrapped: %w", &customError{message: "boom"})
	var target *customError

	a.New(t).Is(a.ErrorAs(err, &target))
}

func Test_error_as_failure(t *testing.T) {
	aSpy := eye.Spy()
	err := errors.New("boom")
	var target *customError

	a.New(aSpy).Is(a.ErrorAs(err, &target))
	aSpy.HadErrorContaining(t, "want error assignable")
}
