package hammy

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"github.com/google/go-cmp/cmp"
)

type MatcherFunc[T any] func(actual T) AssertionMessage

func (m MatcherFunc[T]) Match(actual T) AssertionMessage {
	return m(actual)
}

// MatchFunc adapts a closure into a Matcher without requiring a dedicated type.
func MatchFunc[T any](fn func(actual T) AssertionMessage) Matcher[T] {
	return MatcherFunc[T](fn)
}

// Match evaluates a generic matcher against an actual value and returns the
// resulting AssertionMessage.
//
// This is a package-level function rather than a method on Hammy because Go
// does not support generic methods on non-generic types. A receiver form such
// as:
//
//	func (h *Hammy) Matches[T any](actual T, matcher Matcher[T])
//
// is not legal Go. Keeping Match as a generic package function preserves the
// compile-time type safety between actual values and Matcher[T] instances while
// still composing naturally with assertions:
//
//	assert := hammy.New(t)
//	assert.Is(hammy.Match(value, hammy.EqualTo(expected)))
func Match[T any](actual T, matcher Matcher[T]) AssertionMessage {
	return matcher.Match(actual)
}

func Not[T any](matcher Matcher[T]) Matcher[T] {
	return MatchFunc(func(actual T) AssertionMessage {
		result := matcher.Match(actual)
		return AssertionMessage{
			IsSuccessful: !result.IsSuccessful,
			Message:      "not(" + result.Message + ")",
		}
	})
}

func AllOf[T any](matchers ...Matcher[T]) Matcher[T] {
	return MatchFunc(func(actual T) AssertionMessage {
		for _, matcher := range matchers {
			result := matcher.Match(actual)
			if !result.IsSuccessful {
				return result
			}
		}
		return Assert(true, "matched all %d matchers", len(matchers))
	})
}

func AnyOf[T any](matchers ...Matcher[T]) Matcher[T] {
	return MatchFunc(func(actual T) AssertionMessage {
		if len(matchers) == 0 {
			return Assert(false, "got no matchers, wanted at least one")
		}

		failures := make([]string, 0, len(matchers))
		for _, matcher := range matchers {
			result := matcher.Match(actual)
			if result.IsSuccessful {
				return Assert(true, "matched one of %d matchers", len(matchers))
			}
			if result.Message != "" {
				failures = append(failures, result.Message)
			}
		}

		return Assert(false, "got no matching matcher, failures: %s", strings.Join(failures, "; "))
	})
}

func Describe[T any](msg string, matcher Matcher[T]) Matcher[T] {
	return MatchFunc(func(actual T) AssertionMessage {
		result := matcher.Match(actual)
		if result.IsSuccessful {
			return result
		}
		if result.Message == "" {
			return Assert(false, msg)
		}
		return Assert(false, "%s: %s", msg, result.Message)
	})
}

func EqualTo[T any](expected T) Matcher[T] {
	return MatchFunc(func(actual T) AssertionMessage {
		return Assert(cmp.Equal(actual, expected), "got <%v>, wanted equal to <%v>", actual, expected)
	})
}

func GreaterThan[N Numeric](expected N) Matcher[N] {
	return MatchFunc(func(actual N) AssertionMessage {
		return Assert(actual > expected, "got <%v>, wanted greater than <%v>", actual, expected)
	})
}

func GreaterOrEqual[N Numeric](expected N) Matcher[N] {
	return MatchFunc(func(actual N) AssertionMessage {
		return Assert(actual >= expected, "got <%v>, wanted greater or equal to <%v>", actual, expected)
	})
}

func LessThan[N Numeric](expected N) Matcher[N] {
	return MatchFunc(func(actual N) AssertionMessage {
		return Assert(actual < expected, "got <%v>, wanted less than <%v>", actual, expected)
	})
}

func LessOrEqual[N Numeric](expected N) Matcher[N] {
	return MatchFunc(func(actual N) AssertionMessage {
		return Assert(actual <= expected, "got <%v>, wanted less or equal to <%v>", actual, expected)
	})
}

func Zero[N Numeric]() Matcher[N] {
	return MatchFunc(func(actual N) AssertionMessage {
		return Assert(actual == 0, "got <%v>, wanted equal to zero", actual)
	})
}

func Within[N Numeric](expected N, delta float64) Matcher[N] {
	return MatchFunc(func(actual N) AssertionMessage {
		diff := actual - expected
		if diff < 0 {
			diff = -diff
		}
		return Assert(float64(diff) <= delta, "got <%v>, wanted within <%v> of <%v>", actual, delta, expected)
	})
}

func Contains(expected string) Matcher[string] {
	return MatchFunc(func(actual string) AssertionMessage {
		return Assert(strings.Contains(actual, expected), "got <%s>, wanted substring <%s>", actual, expected)
	})
}

func HasPrefix(expected string) Matcher[string] {
	return MatchFunc(func(actual string) AssertionMessage {
		return Assert(strings.HasPrefix(actual, expected), "got <%s>, wanted prefix <%s>", actual, expected)
	})
}

func HasSuffix(expected string) Matcher[string] {
	return MatchFunc(func(actual string) AssertionMessage {
		return Assert(strings.HasSuffix(actual, expected), "got <%s>, wanted suffix <%s>", actual, expected)
	})
}

func EmptyString() Matcher[string] {
	return MatchFunc(func(actual string) AssertionMessage {
		return Assert(actual == "", "got <%s>, wanted an empty string", actual)
	})
}

func NotEmptyString() Matcher[string] {
	return MatchFunc(func(actual string) AssertionMessage {
		return Assert(actual != "", "got an empty string, wanted non-empty string")
	})
}

func EqualIgnoringCase(expected string) Matcher[string] {
	return MatchFunc(func(actual string) AssertionMessage {
		return Assert(strings.EqualFold(actual, expected), "got <%s>, wanted equal to <%s> ignoring case", actual, expected)
	})
}

func HasPrefixIgnoringCase(expected string) Matcher[string] {
	return MatchFunc(func(actual string) AssertionMessage {
		return Assert(strings.HasPrefix(strings.ToLower(actual), strings.ToLower(expected)), "got <%s>, wanted prefix <%s> ignoring case", actual, expected)
	})
}

func HasSuffixIgnoringCase(expected string) Matcher[string] {
	return MatchFunc(func(actual string) AssertionMessage {
		return Assert(strings.HasSuffix(strings.ToLower(actual), strings.ToLower(expected)), "got <%s>, wanted suffix <%s> ignoring case", actual, expected)
	})
}

func MatchesRegexp(pattern string) Matcher[string] {
	return MatchFunc(func(actual string) AssertionMessage {
		re, err := regexp.Compile(pattern)
		if err != nil {
			return Assert(false, "invalid regexp <%s>: %v", pattern, err)
		}
		return Assert(re.MatchString(actual), "got <%s>, wanted regexp <%s>", actual, pattern)
	})
}

func EqualIgnoringWhitespace(expected string) Matcher[string] {
	return MatchFunc(func(actual string) AssertionMessage {
		actualNormalized := strings.Join(strings.Fields(actual), " ")
		expectedNormalized := strings.Join(strings.Fields(expected), " ")
		return Assert(actualNormalized == expectedNormalized, "got <%s>, wanted equal to <%s> ignoring whitespace", actual, expected)
	})
}

func EqualNormalizedWhitespace(expected string) Matcher[string] {
	return EqualIgnoringWhitespace(expected)
}

func SamePointer[T any](expected *T) Matcher[*T] {
	return MatchFunc(func(actual *T) AssertionMessage {
		return Assert(actual == expected, "got pointer <%p>, wanted same pointer as <%p>", actual, expected)
	})
}

func TypeOf[T any]() Matcher[any] {
	expectedType := reflect.TypeOf((*T)(nil)).Elem()
	return MatchFunc(func(actual any) AssertionMessage {
		if actual == nil {
			return Assert(false, "got <nil>, wanted dynamic type <%s>", expectedType)
		}

		actualType := reflect.TypeOf(actual)
		return Assert(actualType == expectedType, "got dynamic type <%s>, wanted <%s>", actualType, expectedType)
	})
}

func AssignableTo[T any]() Matcher[any] {
	expectedType := reflect.TypeOf((*T)(nil)).Elem()
	return MatchFunc(func(actual any) AssertionMessage {
		if actual == nil {
			return Assert(false, "got <nil>, wanted assignable to <%s>", expectedType)
		}

		actualType := reflect.TypeOf(actual)
		return Assert(actualType.AssignableTo(expectedType), "got dynamic type <%s>, wanted assignable to <%s>", actualType, expectedType)
	})
}

func Having[T, U any](selector func(T) U, matcher Matcher[U]) Matcher[T] {
	return HavingField("", selector, matcher)
}

func HavingField[T, U any](name string, selector func(T) U, matcher Matcher[U]) Matcher[T] {
	return MatchFunc(func(actual T) AssertionMessage {
		result := matcher.Match(selector(actual))
		if result.IsSuccessful {
			return result
		}
		if name == "" {
			return Assert(false, "selected value: %s", result.Message)
		}
		return Assert(false, "field %s: %s", name, result.Message)
	})
}

func formatMatcherFailure(prefix string, result AssertionMessage) string {
	if result.Message == "" {
		return prefix
	}
	return fmt.Sprintf("%s: %s", prefix, result.Message)
}
