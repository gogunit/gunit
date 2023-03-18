package gunit_test

import (
	"testing"

	"github.com/gogunit/gunit"
	"github.com/gogunit/gunit/eye"
)

func Test_Struct_EqualTo_success(t *testing.T) {
	gunit.Struct(t, &s{A: "Hello"}).EqualTo(&s{A: "Hello"})
}

func Test_Struct_EqualTo_failure(t *testing.T) {
	aSpy := eye.Spy()
	gunit.Struct(aSpy, &s{A: "Hello"}).EqualTo(&s{A: "Good-bye"})
	aSpy.HadError(t)
}

type s struct{ A string }
