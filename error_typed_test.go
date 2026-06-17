package gunit_test

import (
	"errors"
	"fmt"
	"github.com/gogunit/gunit"
	"github.com/gogunit/gunit/eye"
	"testing"
)

func Test_error_NilError_success(t *testing.T) {
	gunit.NilError(t, nil)
}
func Test_error_NilError_failure(t *testing.T) {
	aSpy := eye.Spy()
	gunit.NilError(aSpy, errors.New("sentinel"))
	aSpy.HadErrorContaining(t, "want nil error")
}

func Test_error_Error_success(t *testing.T) {
	gunit.Error(t, errors.New("sentinel"))
}
func Test_error_Error_failure(t *testing.T) {
	aSpy := eye.Spy()
	gunit.Error(aSpy, nil)
	aSpy.HadErrorContaining(t, "want error")
}

func Test_error_EqualError_success(t *testing.T) {
	gunit.EqualError(t, errors.New("sentinel"), "sentinel")
}
func Test_error_EqualError_failure_for_nil(t *testing.T) {
	aSpy := eye.Spy()
	gunit.EqualError(aSpy, nil, "sentinel")
	aSpy.HadErrorContaining(t, "got nil error")
}
func Test_error_EqualError_failure_for_mismatch(t *testing.T) {
	aSpy := eye.Spy()
	gunit.EqualError(aSpy, errors.New("wrapped sentinel"), "sentinel")
	aSpy.HadErrorContaining(t, "want <sentinel>")
}

func Test_error_ErrorContains_success(t *testing.T) {
	gunit.ErrorContains(t, errors.New("wrapped sentinel"), "wrapped")
}
func Test_error_ErrorContains_failure_for_nil(t *testing.T) {
	aSpy := eye.Spy()
	gunit.ErrorContains(aSpy, nil, "wrapped")
	aSpy.HadErrorContaining(t, "got nil error")
}
func Test_error_ErrorContains_failure_for_mismatch(t *testing.T) {
	aSpy := eye.Spy()
	gunit.ErrorContains(aSpy, errors.New("wrapped sentinel"), "missing")
	aSpy.HadErrorContaining(t, "want containing")
}

func Test_error_ErrorMatchesRegexp_success(t *testing.T) {
	gunit.ErrorMatchesRegexp(t, errors.New("wrapped sentinel"), "sentinel")
}
func Test_error_ErrorMatchesRegexp_failure_for_nil(t *testing.T) {
	aSpy := eye.Spy()
	gunit.ErrorMatchesRegexp(aSpy, nil, "wrapped")
	aSpy.HadErrorContaining(t, "got nil error")
}
func Test_error_ErrorMatchesRegexp_failure_for_invalid_regexp(t *testing.T) {
	aSpy := eye.Spy()
	gunit.ErrorMatchesRegexp(aSpy, errors.New("wrapped sentinel"), "[")
	aSpy.HadErrorContaining(t, "invalid regexp")
}
func Test_error_ErrorMatchesRegexp_failure_for_mismatch(t *testing.T) {
	aSpy := eye.Spy()
	gunit.ErrorMatchesRegexp(aSpy, errors.New("wrapped sentinel"), "missing")
	aSpy.HadErrorContaining(t, "want regexp")
}

func Test_error_ErrorIs_success(t *testing.T) {
	sentinel := errors.New("sentinel")
	gunit.ErrorIs(t, fmt.Errorf("wrapped: %w", sentinel), sentinel)
}
func Test_error_ErrorIs_failure(t *testing.T) {
	aSpy := eye.Spy()
	gunit.ErrorIs(aSpy, errors.New("wrapped"), errors.New("sentinel"))
	aSpy.HadErrorContaining(t, "want error matching")
}

func Test_error_NotErrorIs_success(t *testing.T) {
	gunit.NotErrorIs(t, errors.New("wrapped"), errors.New("sentinel"))
}
func Test_error_NotErrorIs_failure(t *testing.T) {
	sentinel := errors.New("sentinel")
	aSpy := eye.Spy()
	gunit.NotErrorIs(aSpy, sentinel, sentinel)
	aSpy.HadErrorContaining(t, "want error not matching")
}

func Test_error_ErrorAs_success(t *testing.T) {
	var target *typedTestError
	gunit.ErrorAs(t, &typedTestError{msg: "wrapped sentinel"}, &target)
}
func Test_error_ErrorAs_failure(t *testing.T) {
	var target *typedTestError
	aSpy := eye.Spy()
	gunit.ErrorAs(aSpy, errors.New("sentinel"), &target)
	aSpy.HadErrorContaining(t, "want error assignable")
}

func Test_error_NotErrorAs_success(t *testing.T) {
	var target *otherTypedTestError
	gunit.NotErrorAs(t, &typedTestError{msg: "wrapped sentinel"}, &target)
}
func Test_error_NotErrorAs_failure(t *testing.T) {
	var target *typedTestError
	aSpy := eye.Spy()
	gunit.NotErrorAs(aSpy, &typedTestError{msg: "wrapped sentinel"}, &target)
	aSpy.HadErrorContaining(t, "want error not assignable")
}

func Test_error_ErrorType_success(t *testing.T) {
	gunit.ErrorType[*typedTestError](t, &typedTestError{msg: "wrapped sentinel"})
}
func Test_error_ErrorType_failure(t *testing.T) {
	aSpy := eye.Spy()
	gunit.ErrorType[*typedTestError](aSpy, errors.New("sentinel"))
	aSpy.HadErrorContaining(t, "want error assignable")
}

type typedTestError struct{ msg string }

func (e *typedTestError) Error() string { return e.msg }

type otherTypedTestError struct{ msg string }

func (e *otherTypedTestError) Error() string { return e.msg }
