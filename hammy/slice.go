package hammy

import (
	"fmt"
	"strings"

	"github.com/google/go-cmp/cmp"
)

func Slice[I any](actual []I) *Slc[I] {
	return &Slc[I]{actual}
}

type Slc[I any] struct {
	actual []I
}

func (a *Slc[I]) Matches(matcher Matcher[[]I]) AssertionMessage {
	return matcher.Match(a.actual)
}

// Contains asserts whether the slice contains the expected elements in any order.
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

// ContainsAny asserts that at least one expected element is present.
func (a *Slc[I]) ContainsAny(expected ...I) AssertionMessage {
	return ContainsAny(expected...).Match(a.actual)
}

// NotContains asserts that none of the expected elements are present.
func (a *Slc[I]) NotContains(expected ...I) AssertionMessage {
	var matched []int
	for i, e := range expected {
		for _, item := range a.actual {
			if cmp.Equal(item, e) {
				matched = append(matched, i)
				break
			}
		}
	}
	return Assert(len(matched) == 0, "got items at expected index %v present in slice, wanted all absent", Join(matched, ", "))
}

func Join[T any](a []T, sep string) string {
	if len(a) < 1 {
		return ""
	}

	parts := make([]string, 0, len(a))
	for _, item := range a {
		parts = append(parts, fmt.Sprintf("%v", item))
	}
	return strings.Join(parts, sep)
}

// EqualTo asserts whether the slice is equal to the expected items in both order and values.
func (a *Slc[I]) EqualTo(expected ...I) AssertionMessage {
	diff := cmp.Diff(expected, a.actual)
	return Assert(diff == "", "slice mismatch (-want +got):\\n%s", diff)
}

// Len asserts that the slice contains exactly the number of elements specified.
func (a *Slc[I]) Len(expected int) AssertionMessage {
	sz := len(a.actual)
	return Assert(sz == expected, "got len()=%d, wanted %d", sz, expected)
}

// Cap asserts that the slice has exactly the specified capacity.
func (a *Slc[I]) Cap(expected int) AssertionMessage {
	return Capacity[I](expected).Match(a.actual)
}

// IsEmpty asserts that the slice contains no elements.
func (a *Slc[I]) IsEmpty() AssertionMessage {
	sz := len(a.actual)
	return Assert(sz == 0, "got len()=%d, wanted 0", sz)
}

// NotEmpty asserts that the slice contains at least one element.
func (a *Slc[I]) NotEmpty() AssertionMessage {
	sz := len(a.actual)
	return Assert(sz > 0, "got len()=%d, wanted > 0", sz)
}

// ContainsExactly asserts that the slice contains the exact number of elements in any order.
func (a *Slc[I]) ContainsExactly(expected ...I) AssertionMessage {
	szActual := len(a.actual)
	szExpected := len(expected)
	if szActual != szExpected {
		return Assert(false, "length mismatch got %d, want %d", szActual, szExpected)
	}

	return a.Contains(expected...)
}

// SubsetOf asserts that every actual item is present in the expected superset.
func (a *Slc[I]) SubsetOf(expected ...I) AssertionMessage {
	return SubsetOf(expected...).Match(a.actual)
}

// NotSubsetOf asserts that at least one actual item is absent from the expected superset.
func (a *Slc[I]) NotSubsetOf(expected ...I) AssertionMessage {
	return NotSubsetOf(expected...).Match(a.actual)
}

func (a *Slc[I]) Every(matcher Matcher[I]) AssertionMessage {
	return Every(matcher).Match(a.actual)
}

func (a *Slc[I]) HasItem(matcher Matcher[I]) AssertionMessage {
	return HasItem(matcher).Match(a.actual)
}

func (a *Slc[I]) ContainsInOrder(matchers ...Matcher[I]) AssertionMessage {
	return ContainsInOrder(matchers...).Match(a.actual)
}

func (a *Slc[I]) ContainsInAnyOrder(matchers ...Matcher[I]) AssertionMessage {
	return ContainsInAnyOrder(matchers...).Match(a.actual)
}
