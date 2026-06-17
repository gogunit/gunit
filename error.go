package gunit

import (
	"errors"
	"regexp"
	"strings"
)

func NilError(t T, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("got <%v>, want nil error", err)
	}
}

func Error(t T, err error) {
	t.Helper()
	if err == nil {
		t.Errorf("got nil, want error")
	}
}

func EqualError(t T, err error, expected string) {
	t.Helper()
	if err == nil {
		t.Errorf("got nil error, want error message <%s>", expected)
		return
	}
	if err.Error() != expected {
		t.Errorf("got error message <%s>, want <%s>", err.Error(), expected)
	}
}

func ErrorContains(t T, err error, expected string) {
	t.Helper()
	if err == nil {
		t.Errorf("got nil error, want error containing <%s>", expected)
		return
	}
	if !strings.Contains(err.Error(), expected) {
		t.Errorf("got error message <%s>, want containing <%s>", err.Error(), expected)
	}
}

func ErrorMatchesRegexp(t T, err error, pattern string) {
	t.Helper()
	if err == nil {
		t.Errorf("got nil error, want error matching regexp <%s>", pattern)
		return
	}
	re, compileErr := regexp.Compile(pattern)
	if compileErr != nil {
		t.Errorf("invalid regexp <%s>: %v", pattern, compileErr)
		return
	}
	if !re.MatchString(err.Error()) {
		t.Errorf("got error message <%s>, want regexp <%s>", err.Error(), pattern)
	}
}

func ErrorIs(t T, err error, target error) {
	t.Helper()
	if !errors.Is(err, target) {
		t.Errorf("got <%v>, want error matching <%v>", err, target)
	}
}

func NotErrorIs(t T, err error, target error) {
	t.Helper()
	if errors.Is(err, target) {
		t.Errorf("got <%v>, want error not matching <%v>", err, target)
	}
}

func ErrorAs[E any](t T, err error, target *E) {
	t.Helper()
	if !errors.As(err, target) {
		t.Errorf("got <%v>, want error assignable to <%T>", err, target)
	}
}

func NotErrorAs[E any](t T, err error, target *E) {
	t.Helper()
	if errors.As(err, target) {
		t.Errorf("got <%v>, want error not assignable to <%T>", err, target)
	}
}

func ErrorType[E error](t T, err error) {
	t.Helper()
	var target E
	if !errors.As(err, &target) {
		t.Errorf("got <%v>, want error assignable to <%T>", err, target)
	}
}
