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

func (a Slc[I]) Len(expected int) AssertionMessage {
	sz := len(a.actual)
	return AssertionMessage{
		IsSuccessful: sz == expected,
		Message:      fmt.Sprintf("want len of <%v>, got <%v>", sz, expected),
	}
}

func (a Slc[I]) Contains(expected ...I) AssertionMessage {
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
	return AssertionMessage{
		IsSuccessful: isSuccessful,
		Message:      fmt.Sprintf("want <%v> matched, but no match found for expected items <%v>", len(expected), unmatched),
	}
}
