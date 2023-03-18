package hammy

import "github.com/google/go-cmp/cmp"

func Struct[S any](actual S) *St[S] {
	return &St[S]{actual}
}

type St[S any] struct {
	actual S
}

func (s *St[S]) EqualTo(expected S) AssertionMessage {
	diff := cmp.Diff(expected, s.actual)
	return Assert(diff == "", "Struct mismatch (-want +got):\n%s", diff)
}
