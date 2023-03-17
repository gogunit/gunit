package gunit_test

import (
	"github.com/gogunit/gunit"
	"github.com/gogunit/gunit/eye"
	"testing"
)

func Test_struct_EqualTo_success(t *testing.T) {
	gunit.Struct(t, &s{A: "Hello"}).EqualTo(&s{A: "Hello"})
}

func Test_struct_EqualTo_failure(t *testing.T) {
	aSpy := eye.Spy()
	gunit.Struct(aSpy, &s{A: "Hello"}).EqualTo(&s{A: "Good-bye"})
	aSpy.HadError(t)
}

type s struct{ A string }
