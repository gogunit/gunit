package hammy

func String[S Stringy](actual S) *Str[S] {
	return &Str[S]{actual: actual}
}

type Str[S Stringy] struct {
	actual S
}

func (s Str[S]) EqualTo(expected S) AssertionMessage {
	return Assert(s.actual == expected, "want <%v> equal to <%v>", s.actual, expected)
}
