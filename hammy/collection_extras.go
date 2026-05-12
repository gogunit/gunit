package hammy

import (
	"fmt"

	"github.com/google/go-cmp/cmp"
)

func OneOf[T any](expected ...T) Matcher[T] {
	return MatchFunc(func(actual T) AssertionMessage {
		for _, item := range expected {
			if cmp.Equal(actual, item) {
				return Assert(true, "got <%v>, wanted one of <%v>", actual, expected)
			}
		}
		return Assert(false, "got <%v>, wanted one of <%v>", actual, expected)
	})
}

func Capacity[T any](expected int) Matcher[[]T] {
	return MatchFunc(func(actual []T) AssertionMessage {
		actualCap := cap(actual)
		return Assert(actualCap == expected, "got cap()=%d, wanted %d", actualCap, expected)
	})
}

func ContainsAny[T any](expected ...T) Matcher[[]T] {
	return MatchFunc(func(actual []T) AssertionMessage {
		for _, actualItem := range actual {
			for _, expectedItem := range expected {
				if cmp.Equal(actualItem, expectedItem) {
					return Assert(true, "got matching item <%v>, wanted any of <%v>", actualItem, expected)
				}
			}
		}
		return Assert(false, "got no matching item, wanted any of <%v>", expected)
	})
}

func SubsetOf[T any](superset ...T) Matcher[[]T] {
	return MatchFunc(func(actual []T) AssertionMessage {
		var outside []string
		for i, item := range actual {
			if !containsEqual(superset, item) {
				outside = append(outside, fmt.Sprintf("index %d: %v", i, item))
			}
		}
		return Assert(len(outside) == 0, "got items outside expected set <%v>, wanted subset of <%v>", outside, superset)
	})
}

func NotSubsetOf[T any](superset ...T) Matcher[[]T] {
	return MatchFunc(func(actual []T) AssertionMessage {
		for _, item := range actual {
			if !containsEqual(superset, item) {
				return Assert(true, "got item <%v> outside expected set <%v>", item, superset)
			}
		}
		return Assert(false, "got subset <%v>, wanted at least one item outside <%v>", actual, superset)
	})
}

func HasEntries[K comparable, V any](expected map[K]V) Matcher[map[K]V] {
	return MatchFunc(func(actual map[K]V) AssertionMessage {
		var missing []K
		var mismatched []K
		for key, expectedValue := range expected {
			actualValue, ok := actual[key]
			if !ok {
				missing = append(missing, key)
				continue
			}
			if !cmp.Equal(actualValue, expectedValue) {
				mismatched = append(mismatched, key)
			}
		}
		return Assert(len(missing) == 0 && len(mismatched) == 0, "got missing keys <%v> and mismatched keys <%v>, wanted entries <%v>", missing, mismatched, expected)
	})
}

func NotHasEntries[K comparable, V any](expected map[K]V) Matcher[map[K]V] {
	return MatchFunc(func(actual map[K]V) AssertionMessage {
		for key, expectedValue := range expected {
			actualValue, ok := actual[key]
			if !ok || !cmp.Equal(actualValue, expectedValue) {
				return Assert(true, "got entry <%v:%v> absent or different", key, expectedValue)
			}
		}
		return Assert(false, "got all entries <%v>, wanted at least one absent or different", expected)
	})
}

func containsEqual[T any](items []T, target T) bool {
	for _, item := range items {
		if cmp.Equal(item, target) {
			return true
		}
	}
	return false
}
