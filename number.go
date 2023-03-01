package gunit

import "math"

type Numeric interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

func Number[N Numeric](t T, actual N) *Num[N] {
	return &Num[N]{t, actual}
}

type Num[N Numeric] struct {
	T
	actual N
}

func (n *Num[N]) EqualTo(expected N) {
	n.Helper()
	if n.actual != expected {
		n.Errorf("want <%v> equal to <%v>", n.actual, expected)
	}
}

func (n *Num[N]) NotEqualTo(expected N) {
	n.Helper()
	if n.actual == expected {
		n.Errorf("want <%v> not equal to <%v>", n.actual, expected)
	}
}

func (n *Num[N]) LessThan(expected N) {
	n.Helper()
	if n.actual >= expected {
		n.Errorf("want <%v> less than <%v>", n.actual, expected)
	}
}

func (n *Num[N]) LessOrEqual(expected N) {
	n.Helper()
	if n.actual > expected {
		n.Errorf("want <%v> less or equal to <%v>", n.actual, expected)
	}
}

func (n *Num[N]) GreaterThan(expected N) {
	n.Helper()
	if n.actual <= expected {
		n.Errorf("want <%v> greater than <%v>", n.actual, expected)
	}
}

func (n *Num[N]) GreaterOrEqual(expected N) {
	n.Helper()
	if n.actual < expected {
		n.Errorf("want <%v> greater or equal to <%v>", n.actual, expected)
	}
}

func (n *Num[N]) Within(expected N, error float64) {
	n.Helper()
	diff := math.Abs(float64(n.actual - expected))
	if diff > error {
		n.Errorf("want <%v> greater or equal to <%v>", n.actual, expected)
	}
}

func (n *Num[N]) IsZero() {
	n.Helper()
	if n.actual != 0 {
		n.Errorf("want <%v> equal to zero", n.actual)
	}
}
