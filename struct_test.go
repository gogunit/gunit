package gunit_test

import (
	. "github.com/nfisher/gunit"
	. "github.com/nfisher/gunit/testing"
	"testing"
)

func Test_struct_EqualTo_success(t *testing.T) {
	Struct(t, &s{A: "Hello"}).EqualTo(&s{A: "Hello"})
}

func Test_struct_EqualTo_failure(t *testing.T) {
	aSpy := Spy()
	Struct(aSpy, &s{A: "Hello"}).EqualTo(&s{A: "Good-bye"})
	aSpy.HadError(t)
}

type s struct{ A string }
