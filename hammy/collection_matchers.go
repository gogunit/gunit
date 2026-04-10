package hammy

import (
	"fmt"
	"strings"
)

func Every[T any](matcher Matcher[T]) Matcher[[]T] {
	return MatchFunc(func(actual []T) AssertionMessage {
		for i, item := range actual {
			result := matcher.Match(item)
			if !result.IsSuccessful {
				return Assert(false, "%s", formatMatcherFailure(fmt.Sprintf("item at index %d", i), result))
			}
		}
		return Assert(true, "all %d items matched", len(actual))
	})
}

func HasItem[T any](matcher Matcher[T]) Matcher[[]T] {
	return MatchFunc(func(actual []T) AssertionMessage {
		failures := make([]string, 0, min(3, len(actual)))
		for i, item := range actual {
			result := matcher.Match(item)
			if result.IsSuccessful {
				return Assert(true, "found matching item at index %d", i)
			}
			if len(failures) < 3 {
				failures = append(failures, formatMatcherFailure(fmt.Sprintf("item at index %d", i), result))
			}
		}

		if len(failures) == 0 {
			return Assert(false, "got no matching item in empty slice")
		}
		return Assert(false, "got no matching item in slice. Sample failures: %s", strings.Join(failures, "; "))
	})
}

func ContainsInOrder[T any](matchers ...Matcher[T]) Matcher[[]T] {
	return MatchFunc(func(actual []T) AssertionMessage {
		if len(actual) != len(matchers) {
			return Assert(false, "got len()=%d, wanted %d", len(actual), len(matchers))
		}

		for i, matcher := range matchers {
			result := matcher.Match(actual[i])
			if !result.IsSuccessful {
				return Assert(false, "%s", formatMatcherFailure(fmt.Sprintf("item at index %d", i), result))
			}
		}

		return Assert(true, "all %d items matched in order", len(actual))
	})
}

func ContainsInAnyOrder[T any](matchers ...Matcher[T]) Matcher[[]T] {
	return MatchFunc(func(actual []T) AssertionMessage {
		if len(actual) != len(matchers) {
			return Assert(false, "got len()=%d, wanted %d", len(actual), len(matchers))
		}

		edges := make([][]int, len(matchers))
		for i, matcher := range matchers {
			for j, item := range actual {
				if matcher.Match(item).IsSuccessful {
					edges[i] = append(edges[i], j)
				}
			}
			if len(edges[i]) == 0 {
				return Assert(false, "matcher at index %d matched no items", i)
			}
		}

		matchedActual := make([]int, len(actual))
		for i := range matchedActual {
			matchedActual[i] = -1
		}

		for i := range matchers {
			seen := make([]bool, len(actual))
			if !augmentMatch(i, edges, seen, matchedActual) {
				return Assert(false, "matcher at index %d could not be assigned to a unique item", i)
			}
		}

		return Assert(true, "all %d items matched in any order", len(actual))
	})
}

func HasEntry[K comparable, V any](keyMatcher Matcher[K], valueMatcher Matcher[V]) Matcher[map[K]V] {
	return MatchFunc(func(actual map[K]V) AssertionMessage {
		failures := make([]string, 0, min(3, len(actual)))
		for key, value := range actual {
			keyResult := keyMatcher.Match(key)
			if !keyResult.IsSuccessful {
				if len(failures) < 3 {
					failures = append(failures, formatMatcherFailure(fmt.Sprintf("key <%v>", key), keyResult))
				}
				continue
			}

			valueResult := valueMatcher.Match(value)
			if valueResult.IsSuccessful {
				return Assert(true, "found matching entry for key <%v>", key)
			}
			if len(failures) < 3 {
				failures = append(failures, formatMatcherFailure(fmt.Sprintf("value for key <%v>", key), valueResult))
			}
		}

		if len(failures) == 0 {
			return Assert(false, "got no matching entry in empty map")
		}
		return Assert(false, "got no matching map entry. Sample failures: %s", strings.Join(failures, "; "))
	})
}

func HasKeyMatching[K comparable, V any](matcher Matcher[K]) Matcher[map[K]V] {
	return MatchFunc(func(actual map[K]V) AssertionMessage {
		failures := make([]string, 0, min(3, len(actual)))
		for key := range actual {
			result := matcher.Match(key)
			if result.IsSuccessful {
				return Assert(true, "found matching key <%v>", key)
			}
			if len(failures) < 3 {
				failures = append(failures, formatMatcherFailure(fmt.Sprintf("key <%v>", key), result))
			}
		}

		if len(failures) == 0 {
			return Assert(false, "got no matching key in empty map")
		}
		return Assert(false, "got no matching key. Sample failures: %s", strings.Join(failures, "; "))
	})
}

func HasValueMatching[K comparable, V any](matcher Matcher[V]) Matcher[map[K]V] {
	return MatchFunc(func(actual map[K]V) AssertionMessage {
		failures := make([]string, 0, min(3, len(actual)))
		for key, value := range actual {
			result := matcher.Match(value)
			if result.IsSuccessful {
				return Assert(true, "found matching value for key <%v>", key)
			}
			if len(failures) < 3 {
				failures = append(failures, formatMatcherFailure(fmt.Sprintf("value for key <%v>", key), result))
			}
		}

		if len(failures) == 0 {
			return Assert(false, "got no matching value in empty map")
		}
		return Assert(false, "got no matching value. Sample failures: %s", strings.Join(failures, "; "))
	})
}

func augmentMatch(expectedIndex int, edges [][]int, seen []bool, matchedActual []int) bool {
	for _, actualIndex := range edges[expectedIndex] {
		if seen[actualIndex] {
			continue
		}
		seen[actualIndex] = true
		if matchedActual[actualIndex] == -1 || augmentMatch(matchedActual[actualIndex], edges, seen, matchedActual) {
			matchedActual[actualIndex] = expectedIndex
			return true
		}
	}
	return false
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
