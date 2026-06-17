package gunit

import "github.com/google/go-cmp/cmp"

func Slice[I any](t T, actual []I) *Slc[I] {
	return &Slc[I]{T: t, actual: actual}
}

type Slc[I any] struct {
	T
	actual []I
}

func (s *Slc[I]) Contains(expected ...I) {
	s.Helper()
	missing := missingItems(s.actual, expected)
	if len(missing) > 0 {
		s.Errorf("got missing items <%v>, wanted slice containing <%v>", missing, expected)
	}
}

func (s *Slc[I]) ContainsAny(expected ...I) {
	s.Helper()
	for _, actual := range s.actual {
		for _, want := range expected {
			if cmp.Equal(actual, want) {
				return
			}
		}
	}
	s.Errorf("got no matching item, wanted any of <%v>", expected)
}

func (s *Slc[I]) NotContains(expected ...I) {
	s.Helper()
	var present []I
	for _, want := range expected {
		for _, actual := range s.actual {
			if cmp.Equal(actual, want) {
				present = append(present, want)
				break
			}
		}
	}
	if len(present) > 0 {
		s.Errorf("got items <%v> present in slice, wanted absent", present)
	}
}

func (s *Slc[I]) EqualTo(expected ...I) {
	s.Helper()
	if diff := cmp.Diff(expected, s.actual); diff != "" {
		s.Errorf("slice mismatch (-want +got):\n%s", diff)
	}
}

func (s *Slc[I]) Len(expected int) {
	s.Helper()
	if len(s.actual) != expected {
		s.Errorf("got len()=%d, wanted %d", len(s.actual), expected)
	}
}

func (s *Slc[I]) Cap(expected int) {
	s.Helper()
	if cap(s.actual) != expected {
		s.Errorf("got cap()=%d, wanted %d", cap(s.actual), expected)
	}
}

func (s *Slc[I]) IsEmpty() {
	s.Helper()
	if len(s.actual) != 0 {
		s.Errorf("got len()=%d, wanted 0", len(s.actual))
	}
}

func (s *Slc[I]) IsNotEmpty() {
	s.Helper()
	if len(s.actual) == 0 {
		s.Errorf("got len()=0, wanted > 0")
	}
}

func (s *Slc[I]) NotEmpty() { s.IsNotEmpty() }

func (s *Slc[I]) ContainsExactly(expected ...I) {
	s.Helper()
	if len(s.actual) != len(expected) {
		s.Errorf("length mismatch got %d, want %d", len(s.actual), len(expected))
		return
	}
	s.Contains(expected...)
}

func (s *Slc[I]) SubsetOf(expected ...I) {
	s.Helper()
	for _, actual := range s.actual {
		if !containsEqual(expected, actual) {
			s.Errorf("got item <%v> outside expected set <%v>", actual, expected)
			return
		}
	}
}

func (s *Slc[I]) NotSubsetOf(expected ...I) {
	s.Helper()
	for _, actual := range s.actual {
		if !containsEqual(expected, actual) {
			return
		}
	}
	s.Errorf("got subset <%v>, wanted at least one item outside <%v>", s.actual, expected)
}

func missingItems[I any](actual, expected []I) []I {
	var missing []I
	for _, want := range expected {
		if !containsEqual(actual, want) {
			missing = append(missing, want)
		}
	}
	return missing
}

func containsEqual[I any](items []I, target I) bool {
	for _, item := range items {
		if cmp.Equal(item, target) {
			return true
		}
	}
	return false
}
