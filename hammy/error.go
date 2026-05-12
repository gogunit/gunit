package hammy

import (
	"errors"
	"regexp"
	"strings"
)

func NilError(err error) AssertionMessage {
	return Assert(err == nil, "got <%v>, want nil error", err)
}

func Error(err error) AssertionMessage {
	return Assert(err != nil, "got %v, want error", err)
}

func EqualError(err error, expected string) AssertionMessage {
	if err == nil {
		return Assert(false, "got nil error, want error message <%s>", expected)
	}
	return Assert(err.Error() == expected, "got error message <%s>, want <%s>", err.Error(), expected)
}

func ErrorContains(err error, expected string) AssertionMessage {
	if err == nil {
		return Assert(false, "got nil error, want error containing <%s>", expected)
	}
	return Assert(strings.Contains(err.Error(), expected), "got error message <%s>, want containing <%s>", err.Error(), expected)
}

func ErrorMatchesRegexp(err error, pattern string) AssertionMessage {
	if err == nil {
		return Assert(false, "got nil error, want error matching regexp <%s>", pattern)
	}

	re, compileErr := regexp.Compile(pattern)
	if compileErr != nil {
		return Assert(false, "invalid regexp <%s>: %v", pattern, compileErr)
	}
	return Assert(re.MatchString(err.Error()), "got error message <%s>, want regexp <%s>", err.Error(), pattern)
}

func ErrorIs(err error, target error) AssertionMessage {
	return Assert(errors.Is(err, target), "got <%v>, want error matching <%v>", err, target)
}

func NotErrorIs(err error, target error) AssertionMessage {
	return Assert(!errors.Is(err, target), "got <%v>, want error not matching <%v>", err, target)
}

func ErrorAs[T any](err error, target *T) AssertionMessage {
	return Assert(errors.As(err, target), "got <%v>, want error assignable to <%T>", err, target)
}

func NotErrorAs[T any](err error, target *T) AssertionMessage {
	return Assert(!errors.As(err, target), "got <%v>, want error not assignable to <%T>", err, target)
}

func ErrorType[T error](err error) AssertionMessage {
	var target T
	return Assert(errors.As(err, &target), "got <%v>, want error assignable to <%T>", err, target)
}
