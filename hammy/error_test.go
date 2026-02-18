package hammy_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/gogunit/gunit/eye"
	hammy "github.com/gogunit/gunit/hammy"
)

type customError struct {
	message string
}

func (e *customError) Error() string {
	return e.message
}

func Test_is_error_success(t *testing.T) {
	var err = errors.New("hello world")
	hammy.New(t).Is(hammy.Error(err))
}

func Test_is_error_failure(t *testing.T) {
	aSpy := eye.Spy()
	hammy.New(aSpy).Is(hammy.Error(nil))
	aSpy.HadErrorContaining(t, "got <nil>, want error")
}

func Test_nil_error_success(t *testing.T) {
	hammy.New(t).Is(hammy.NilError(nil))
}

func Test_nil_error_failure(t *testing.T) {
	aSpy := eye.Spy()
	var err = errors.New("hello world")
	hammy.New(aSpy).Is(hammy.NilError(err))
	aSpy.HadErrorContaining(t, "got <hello world>, want nil error")
}

func Test_error_is_success(t *testing.T) {
	target := errors.New("boom")
	err := fmt.Errorf("wrapped: %w", target)

	hammy.New(t).Is(hammy.ErrorIs(err, target))
}

func Test_error_is_failure(t *testing.T) {
	aSpy := eye.Spy()
	target := errors.New("boom")
	err := errors.New("different")

	hammy.New(aSpy).Is(hammy.ErrorIs(err, target))
	aSpy.HadErrorContaining(t, "got <different>, want error matching <boom>")
}

func Test_error_as_success(t *testing.T) {
	err := fmt.Errorf("wrapped: %w", &customError{message: "boom"})
	var target *customError

	hammy.New(t).Is(hammy.ErrorAs(err, &target))
}

func Test_error_as_failure(t *testing.T) {
	aSpy := eye.Spy()
	err := errors.New("boom")
	var target *customError

	hammy.New(aSpy).Is(hammy.ErrorAs(err, &target))
	aSpy.HadErrorContaining(t, "want error assignable")
}
