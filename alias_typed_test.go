package gunit_test

import (
	"github.com/gogunit/gunit"
	"github.com/gogunit/gunit/eye"
	"testing"
)

func Test_string_NotContains_success(t *testing.T) {
	gunit.String(t, "Hello").NotContains("bye")
}

func Test_string_NotContains_failure(t *testing.T) {
	aSpy := eye.Spy()
	gunit.String(aSpy, "Hello").NotContains("ell")
	aSpy.HadErrorContaining(t, "not to contain")
}

func Test_string_NotEmpty_success(t *testing.T) {
	gunit.String(t, "Hello").NotEmpty()
}

func Test_string_NotEmpty_failure(t *testing.T) {
	aSpy := eye.Spy()
	gunit.String(aSpy, "").NotEmpty()
	aSpy.HadErrorContaining(t, "was not")
}

func Test_number_NotEqual_success(t *testing.T) {
	gunit.Number(t, 1).NotEqual(2)
}

func Test_number_NotEqual_failure(t *testing.T) {
	aSpy := eye.Spy()
	gunit.Number(aSpy, 1).NotEqual(1)
	aSpy.HadErrorContaining(t, "not equal")
}
