package hammy_test

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
	"github.com/nfisher/gunit/eye"
	"github.com/nfisher/gunit/hammy"
	"testing"
)

func Test_Slice_Len_failure(t *testing.T) {
	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(Slice([]int{1, 2, 3}).Len(2))
	aSpy.HadError(t)
}

func Test_Slice_Len_success(t *testing.T) {
	assert := hammy.New(t)
	assert.Is(Slice([]int{1, 2, 3}).Len(3))
}

func Test_Slice_Contains_failure(t *testing.T) {
	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(Slice([]int{1, 2, 3}).Contains(2, 3, 4))
	aSpy.HadError(t)
}

func Test_Slice_Contains_success(t *testing.T) {
	assert := hammy.New(t)
	assert.Is(Slice([]int{1, 2, 3}).Contains(2, 3))
}

func Slice[I any](actual []I) *Slc[I] {
	return &Slc[I]{actual}
}

type Slc[I any] struct {
	actual []I
}

func (a Slc[I]) Len(expected int) hammy.AssertionMessage {
	sz := len(a.actual)
	return hammy.AssertionMessage{
		IsSuccessful: sz == expected,
		Message:      fmt.Sprintf("want len of <%v>, got <%v>", sz, expected),
	}
}

func (a Slc[I]) Contains(expected ...I) hammy.AssertionMessage {
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
	return hammy.AssertionMessage{
		IsSuccessful: isSuccessful,
		Message:      fmt.Sprintf("want <%v> matched, but no match found for expected items <%v>", len(expected), unmatched),
	}
}
