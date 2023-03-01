package gunit_test

import (
	"github.com/google/go-cmp/cmp"
	. "github.com/nfisher/gunit"
	"testing"
)

func Test_struct_EqualTo_success(t *testing.T) {
	Struct(t, &s{A: "Hello"}).EqualTo(&s{A: "Hello"})
}

func Test_struct_EqualTo_failure(t *testing.T) {
	aSpy := spy()
	Struct(aSpy, &s{A: "Hello"}).EqualTo(&s{A: "Good-bye"})
	aSpy.HadError(t)
}

type s struct{ A string }

type St[S any] struct {
	T
	actual S
}

func (s St[S]) EqualTo(expected S) {
	s.Helper()
	if diff := cmp.Diff(expected, s.actual); diff != "" {
		s.Errorf("struct mismatch (-want +got):\n%s", diff)
	}
}

func Struct[S any](t T, actual S) *St[S] {
	return &St[S]{t, actual}
}
