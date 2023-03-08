package hammy

import (
	"fmt"
	"github.com/nfisher/gunit"
	. "github.com/nfisher/gunit/testing"
	"math"
	"testing"
)

func Test_int_equalTo_success(t *testing.T) {
	assert := New(t)
	assert.That(Number(123).EqualTo(123))
}

func Test_int_is_equalTo_success(t *testing.T) {
	assert := New(t)
	assert.That(Number(123).Is(EqualTo(123)))
}

func Test_int_equalTo_failure(t *testing.T) {
	aSpy := Spy()
	assert := New(aSpy)
	assert.That(Number(123).EqualTo(345))
	aSpy.HadError(t)
}

func Test_int_is_equalTo_failure(t *testing.T) {
	aSpy := Spy()
	assert := New(aSpy)
	assert.That(Number(123).Is(EqualTo(345)))
	aSpy.HadError(t)
}

func Test_int_is_not_equalTo_failure(t *testing.T) {
	aSpy := Spy()
	assert := New(aSpy)
	assert.That(Number(123).Is(Not(EqualTo(123))))
	aSpy.HadError(t)
}

func Test_int_is_within_failure(t *testing.T) {
	aSpy := Spy()
	assert := New(aSpy)
	assert.That(Number(123).Is(Within(345, 1.0)))
	aSpy.HadError(t)
}

type Numeric interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

type Stringy interface {
	~string
}

type Primitives interface {
	Stringy | Numeric
}

func New(t gunit.T) *Hammy {
	return &Hammy{t}
}

type Hammy struct {
	t gunit.T
}

type Asserter func(gunit.T)

func (h *Hammy) That(a Asserter) {
	a(h.t)
}

func Number[N Numeric](actual N) *Num[N] {
	return &Num[N]{actual: actual}
}

type Num[N Numeric] struct {
	actual N
}

func (n *Num[N]) EqualTo(expected N) func(gunit.T) {
	return func(t gunit.T) {
		t.Helper()
		if n.actual != expected {
			t.Errorf("want <%v> equal to <%v>", n.actual, expected)
		}
	}
}

func (n *Num[N]) Is(fn func(actual N, t gunit.T) AssertionMessage) Asserter {
	return func(t gunit.T) {
		t.Helper()
		msg := fn(n.actual, t)
		if !msg.Result {
			t.Errorf(msg.Message)
		}
	}
}

func Not[N Primitives](fn func(actual N, t gunit.T) AssertionMessage) func(actual N, t gunit.T) AssertionMessage {
	return func(actual N, t gunit.T) AssertionMessage {
		msg := fn(actual, t)
		msg.Result = !msg.Result
		return msg
	}
}

func EqualTo[N Primitives](expected N) func(actual N, t gunit.T) AssertionMessage {
	return func(actual N, t gunit.T) AssertionMessage {
		return AssertionMessage{
			Message: fmt.Sprintf("want <%v> equal to <%v>", actual, expected),
			Result:  actual == expected,
		}
	}
}

func LessThan[N Primitives](expected N) func(actual N, t gunit.T) AssertionMessage {
	return func(actual N, t gunit.T) AssertionMessage {
		return AssertionMessage{
			Message: fmt.Sprintf("want <%v> equal to <%v>", actual, expected),
			Result:  actual < expected,
		}
	}
}

func GreaterThan[N Primitives](expected N) func(actual N, t gunit.T) AssertionMessage {
	return func(actual N, t gunit.T) AssertionMessage {
		return AssertionMessage{
			Message: fmt.Sprintf("want <%v> equal to <%v>", actual, expected),
			Result:  actual > expected,
		}
	}
}

func Within[N Numeric](expected N, delta float64) func(actual N, t gunit.T) AssertionMessage {
	return func(actual N, t gunit.T) AssertionMessage {
		diff := math.Abs(float64(actual) - float64(expected))
		return AssertionMessage{
			Message: fmt.Sprintf("want <%v> equal to <%v>", actual, expected),
			Result:  diff < delta,
		}
	}
}

type AssertionMessage struct {
	Message string
	Result  bool
}
