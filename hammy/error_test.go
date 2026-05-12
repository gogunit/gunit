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

func Test_equal_error_success(t *testing.T) {
	err := errors.New("hello world")

	a.New(t).Is(a.EqualError(err, "hello world"))
}

func Test_equal_error_failure(t *testing.T) {
	aSpy := eye.Spy()
	err := errors.New("hello world")

	a.New(aSpy).Is(a.EqualError(err, "hello"))
	aSpy.HadErrorContaining(t, "got error message <hello world>, want <hello>")
}

func Test_equal_error_failure_on_nil(t *testing.T) {
	aSpy := eye.Spy()

	a.New(aSpy).Is(a.EqualError(nil, "hello"))
	aSpy.HadErrorContaining(t, "got nil error, want error message <hello>")
}

func Test_error_contains_success(t *testing.T) {
	err := errors.New("hello world")

	a.New(t).Is(a.ErrorContains(err, "world"))
}

func Test_error_contains_failure(t *testing.T) {
	aSpy := eye.Spy()
	err := errors.New("hello world")

	a.New(aSpy).Is(a.ErrorContains(err, "goodbye"))
	aSpy.HadErrorContaining(t, "got error message <hello world>, want containing <goodbye>")
}

func Test_error_matches_regexp_success(t *testing.T) {
	err := errors.New("request failed with status 503")

	a.New(t).Is(a.ErrorMatchesRegexp(err, `status \d+`))
}

func Test_error_matches_regexp_failure(t *testing.T) {
	aSpy := eye.Spy()
	err := errors.New("request failed with status 503")

	a.New(aSpy).Is(a.ErrorMatchesRegexp(err, `status [12]\d\d`))
	aSpy.HadErrorContaining(t, "want regexp <status [12]\\d\\d>")
}

func Test_error_matches_regexp_invalid_pattern_failure(t *testing.T) {
	aSpy := eye.Spy()
	err := errors.New("hello world")

	a.New(aSpy).Is(a.ErrorMatchesRegexp(err, `(`))
	aSpy.HadErrorContaining(t, "invalid regexp <(>")
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

func Test_not_error_is_success(t *testing.T) {
	target := errors.New("boom")
	err := errors.New("different")

	a.New(t).Is(a.NotErrorIs(err, target))
}

func Test_not_error_is_failure(t *testing.T) {
	aSpy := eye.Spy()
	target := errors.New("boom")
	err := fmt.Errorf("wrapped: %w", target)

	a.New(aSpy).Is(a.NotErrorIs(err, target))
	aSpy.HadErrorContaining(t, "want error not matching <boom>")
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

func Test_not_error_as_success(t *testing.T) {
	err := errors.New("boom")
	var target *customError

	a.New(t).Is(a.NotErrorAs(err, &target))
}

func Test_not_error_as_failure(t *testing.T) {
	aSpy := eye.Spy()
	err := fmt.Errorf("wrapped: %w", &customError{message: "boom"})
	var target *customError

	a.New(aSpy).Is(a.NotErrorAs(err, &target))
	aSpy.HadErrorContaining(t, "want error not assignable")
}

func Test_error_type_success(t *testing.T) {
	err := fmt.Errorf("wrapped: %w", &customError{message: "boom"})

	a.New(t).Is(a.ErrorType[*customError](err))
}

func Test_error_type_failure(t *testing.T) {
	aSpy := eye.Spy()
	err := errors.New("boom")

	a.New(aSpy).Is(a.ErrorType[*customError](err))
	aSpy.HadErrorContaining(t, "want error assignable")
}
