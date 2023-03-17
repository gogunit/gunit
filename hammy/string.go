package hammy

import "fmt"

func String[S Stringy](actual S) *Str[S] {
	return &Str[S]{actual: actual}
}

type Str[S Stringy] struct {
	actual S
}

func (s Str[S]) EqualTo(expected S) AssertionMessage {
	return AssertionMessage{
		IsSuccessful: s.actual == expected,
		Message:      fmt.Sprintf("want <%v> equal to <%v>", s.actual, expected),
	}
}
