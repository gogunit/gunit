package gunit_test

import (
	"testing"

	"github.com/gogunit/gunit"
	"github.com/gogunit/gunit/eye"
)

func Test_bool_true_success(t *testing.T) {
	gunit.True(t, true)
}

func Test_bool_false_success(t *testing.T) {
	gunit.False(t, false)
}

func Test_bool_true_failure(t *testing.T) {
	aSpy := eye.Spy()
	gunit.True(aSpy, false)
	aSpy.HadErrorContaining(t, "got false, wanted true")
}

func Test_bool_false_failure(t *testing.T) {
	aSpy := eye.Spy()
	gunit.False(aSpy, true)
	aSpy.HadErrorContaining(t, "got true, wanted false")
}
