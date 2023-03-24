package hammy

import (
	"fmt"

	"github.com/google/go-cmp/cmp"
)

func Slice[I any](actual []I) *Slc[I] {
	return &Slc[I]{actual}
}

type Slc[I any] struct {
	actual []I
}

func (a *Slc[I]) Contains(expected ...I) AssertionMessage {
	hasMatch := make([]bool, len(expected))
	for _, item := range a.actual {
		for i, e := range expected {
			if cmp.Equal(item, e) {
				hasMatch[i] = true
				break
			}
		}
	}
	isSuccessful := true
	var unmatched []int
	for i, hasMatch := range hasMatch {
		if !hasMatch {
			isSuccessful = false
			unmatched = append(unmatched, i)
		}
	}
	return Assert(isSuccessful, "got %v unmatched items, wanted array containing the %v items. Items at index %v were missing", len(unmatched), len(expected), Join(unmatched, ", "))
}

func Join[T any](a []T, sep string) string {
	var s = ""
	if len(a) < 1 {
		return s
	}
	var i = 0
	for ; i < len(a)-2; i++ {
		s += fmt.Sprintf("%v%s", a[i], sep)
	}
	s += fmt.Sprintf("%v", a[i])

	return s
}

func (a *Slc[I]) EqualTo(expected ...I) AssertionMessage {
	diff := cmp.Diff(expected, a.actual)
	return Assert(diff == "", "slice mismatch (-want +got):\\n%s", diff)
}

func (a *Slc[I]) Len(expected int) AssertionMessage {
	sz := len(a.actual)
	return Assert(sz == expected, "got len()=%d, wanted %d", sz, expected)
}

func (a *Slc[I]) IsEmpty() AssertionMessage {
	sz := len(a.actual)
	return Assert(sz == 0, "got len()=%d, wanted 0", sz)
}
