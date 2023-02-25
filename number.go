package gunit

import "math"

type Numeric interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 |
		float32 | float64
}

func Number[T Numeric](t test, actual T) *Num[T] {
	return &Num[T]{t, actual}
}

type Num[T Numeric] struct {
	test
	actual T
}

func (n *Num[T]) EqualTo(expected T) {
	n.Helper()
	if n.actual != expected {
		n.Errorf("want <%v> equal to <%v>", n.actual, expected)
	}
}

func (n *Num[T]) NotEqualTo(expected T) {
	n.Helper()
	if n.actual == expected {
		n.Errorf("want <%v> not equal to <%v>", n.actual, expected)
	}
}

func (n *Num[T]) LessThan(expected T) {
	n.Helper()
	if n.actual >= expected {
		n.Errorf("want <%v> less than <%v>", n.actual, expected)
	}
}

func (n *Num[T]) LessOrEqual(expected T) {
	n.Helper()
	if n.actual > expected {
		n.Errorf("want <%v> less or equal to <%v>", n.actual, expected)
	}
}

func (n *Num[T]) GreaterThan(expected T) {
	n.Helper()
	if n.actual <= expected {
		n.Errorf("want <%v> greater than <%v>", n.actual, expected)
	}
}

func (n *Num[T]) GreaterOrEqual(expected T) {
	n.Helper()
	if n.actual < expected {
		n.Errorf("want <%v> greater or equal to <%v>", n.actual, expected)
	}
}

func (n *Num[T]) Within(expected T, error float64) {
	n.Helper()
	diff := math.Abs(float64(n.actual - expected))
	if diff > error {
		n.Errorf("want <%v> greater or equal to <%v>", n.actual, expected)
	}
}
