package gunit

import (
	"math"
	"reflect"
)

type LengthBounded interface {
	Len() int
}

func Length(t T, actual LengthBounded) *Len {
	if isNilLengthBounded(actual) {
		return invalidLen(t, "length-bounded item")
	}
	return newLen(t, actual.Len())
}

type Len struct {
	T
	actual  int
	valid   bool
	subject string
}

func newLen(t T, actual int) *Len {
	return &Len{T: t, actual: actual, valid: true}
}

func invalidLen(t T, subject string) *Len {
	return &Len{T: t, subject: subject}
}

func isNilLengthBounded(actual LengthBounded) bool {
	if actual == nil {
		return true
	}

	value := reflect.ValueOf(actual)
	switch value.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Pointer, reflect.Slice:
		return value.IsNil()
	default:
		return false
	}
}

func (l *Len) EqualTo(expected int) {
	l.Helper()
	if l.reportInvalid("len()=%d", expected) {
		return
	}
	if l.actual != expected {
		l.Errorf("got len()=%d, wanted %d", l.actual, expected)
	}
}

func (l *Len) NotEqualTo(expected int) {
	l.Helper()
	if l.reportInvalid("len() not equal to %d", expected) {
		return
	}
	if l.actual == expected {
		l.Errorf("got len()=%d, wanted not equal to %d", l.actual, expected)
	}
}

func (l *Len) LessThan(expected int) {
	l.Helper()
	if l.reportInvalid("len() less than %d", expected) {
		return
	}
	if l.actual >= expected {
		l.Errorf("got len()=%d, wanted less than %d", l.actual, expected)
	}
}

func (l *Len) LessOrEqual(expected int) {
	l.Helper()
	if l.reportInvalid("len() less or equal to %d", expected) {
		return
	}
	if l.actual > expected {
		l.Errorf("got len()=%d, wanted less or equal to %d", l.actual, expected)
	}
}

func (l *Len) GreaterThan(expected int) {
	l.Helper()
	if l.reportInvalid("len() greater than %d", expected) {
		return
	}
	if l.actual <= expected {
		l.Errorf("got len()=%d, wanted greater than %d", l.actual, expected)
	}
}

func (l *Len) GreaterOrEqual(expected int) {
	l.Helper()
	if l.reportInvalid("len() greater or equal to %d", expected) {
		return
	}
	if l.actual < expected {
		l.Errorf("got len()=%d, wanted greater or equal to %d", l.actual, expected)
	}
}

func (l *Len) Within(expected int, delta float64) {
	l.Helper()
	if l.reportInvalid("len() within %v of %d", delta, expected) {
		return
	}
	diff := math.Abs(float64(l.actual - expected))
	if diff > delta {
		l.Errorf("got len()=%d, wanted within %v of %d", l.actual, delta, expected)
	}
}

func (l *Len) IsZero() {
	l.Helper()
	if l.reportInvalid("len()=0") {
		return
	}
	if l.actual != 0 {
		l.Errorf("got len()=%d, wanted 0", l.actual)
	}
}

func (l *Len) reportInvalid(format string, args ...any) bool {
	if l.valid {
		return false
	}
	l.Errorf("got nil %s, wanted "+format, append([]any{l.subject}, args...)...)
	return true
}
