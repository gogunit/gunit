package gunit

import "github.com/google/go-cmp/cmp"

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
