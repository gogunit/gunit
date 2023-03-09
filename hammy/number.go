package hammy

import (
	"fmt"
	"github.com/nfisher/gunit"
	"math"
)

func New(t gunit.T) *Hammy {
	return &Hammy{t}
}

type Hammy struct {
	gunit.T
}

func (h *Hammy) That(a AssertionMessage) {
	h.Helper()
	if !a.IsSuccessful {
		h.Errorf(a.Message)
	}
}

func Number[N Numeric](actual N) *Num[N] {
	return &Num[N]{actual: actual}
}

type Num[N Numeric] struct {
	actual N
}

func (n *Num[N]) EqualTo(expected N) AssertionMessage {
	return AssertionMessage{
		IsSuccessful: n.actual == expected,
		Message:      fmt.Sprintf("want <%v> equal to <%v>", n.actual, expected),
	}
}

func (n *Num[N]) NotEqual(expected N) AssertionMessage {
	return AssertionMessage{
		IsSuccessful: n.actual != expected,
		Message:      fmt.Sprintf("want <%v> not equal to <%v>", n.actual, expected),
	}
}

func (n *Num[N]) LessThan(expected N) AssertionMessage {
	return AssertionMessage{
		IsSuccessful: n.actual < expected,
		Message:      fmt.Sprintf("want <%v> less than <%v>", n.actual, expected),
	}
}

func (n *Num[N]) GreaterThan(expected N) AssertionMessage {
	return AssertionMessage{
		IsSuccessful: n.actual > expected,
		Message:      fmt.Sprintf("want <%v> greater than <%v>", n.actual, expected),
	}
}

func (n *Num[N]) LessOrEqual(expected N) AssertionMessage {
	return AssertionMessage{
		IsSuccessful: n.actual <= expected,
		Message:      fmt.Sprintf("want <%v> less or equal to <%v>", n.actual, expected),
	}
}

func (n *Num[N]) GreaterOrEqual(expected N) AssertionMessage {
	return AssertionMessage{
		IsSuccessful: n.actual >= expected,
		Message:      fmt.Sprintf("want <%v> greater or equal to <%v>", n.actual, expected),
	}
}

func (n *Num[N]) IsZero() AssertionMessage {
	return AssertionMessage{
		IsSuccessful: n.actual == 0,
		Message:      fmt.Sprintf("want <%v> equal to zero", n.actual),
	}
}

func (n *Num[N]) Within(expected N, error float64) AssertionMessage {
	diff := math.Abs(float64(n.actual - expected))
	return AssertionMessage{
		IsSuccessful: diff <= error,
		Message:      fmt.Sprintf("want <%v> greater or equal to <%v>", n.actual, expected),
	}
}

type AssertionMessage struct {
	Message      string
	IsSuccessful bool
}
