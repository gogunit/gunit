package hammy

import "strings"

func String[S Stringy](actual S) *Str[S] {
	return &Str[S]{actual: actual}
}

type Str[S Stringy] struct {
	actual S
}

func (s *Str[S]) EqualTo(expected S) AssertionMessage {
	return Assert(s.actual == expected, "want <%v> equal to <%v>", s.actual, expected)
}

func (s *Str[S]) Contains(expected string) AssertionMessage {
	return Assert(strings.Contains(string(s.actual), expected), "want <%v> to contain <%v>", s.actual, expected)
}

func (s *Str[S]) HasPrefix(expected string) AssertionMessage {
	return Assert(strings.HasPrefix(string(s.actual), expected), "want <%v> with prefix <%v>", s.actual, expected)
}

func (s *Str[S]) HasSuffix(expected string) AssertionMessage {
	return Assert(strings.HasSuffix(string(s.actual), expected), "want <%v> with suffix <%v>", s.actual, expected)
}

func (s *Str[S]) IsEmpty() AssertionMessage {
	return Assert(len(s.actual) == 0, "want <%v> as empty string", len(s.actual))
}

func (s *Str[S]) ToLowerEqualTo(expected string) AssertionMessage {
	lowerActual := strings.ToLower(string(s.actual))
	lowerExpected := strings.ToLower(expected)
	return Assert(lowerActual == lowerExpected, "want <%v>, got <%v>", lowerExpected, lowerActual)
}
