package hammy

import (
	"math"
)

func Number[N Numeric](actual N) *Num[N] {
	return &Num[N]{actual: actual}
}

type Num[N Numeric] struct {
	actual N
}

func (n *Num[N]) EqualTo(expected N) AssertionMessage {
	return Assert(n.actual == expected, "want <%v> equal to <%v>", n.actual, expected)
}

func (n *Num[N]) NotEqual(expected N) AssertionMessage {
	return Assert(n.actual != expected, "want <%v> not equal to <%v>", n.actual, expected)
}

func (n *Num[N]) LessThan(expected N) AssertionMessage {
	return Assert(n.actual < expected, "want <%v> less than <%v>", n.actual, expected)
}

func (n *Num[N]) GreaterThan(expected N) AssertionMessage {
	return Assert(n.actual > expected, "want <%v> greater than <%v>", n.actual, expected)
}

func (n *Num[N]) LessOrEqual(expected N) AssertionMessage {
	return Assert(n.actual <= expected, "want <%v> less or equal to <%v>", n.actual, expected)
}

func (n *Num[N]) GreaterOrEqual(expected N) AssertionMessage {
	return Assert(n.actual >= expected, "want <%v> greater or equal to <%v>", n.actual, expected)
}

func (n *Num[N]) IsZero() AssertionMessage {
	return Assert(n.actual == 0, "want <%v> equal to zero", n.actual)
}

func (n *Num[N]) Within(expected N, error float64) AssertionMessage {
	diff := math.Abs(float64(n.actual - expected))
	return Assert(diff <= error, "want <%v> greater or equal to <%v>", n.actual, expected)
}
