package hammy

import "time"

func Time(actual time.Time) *Tim {
	return &Tim{actual: actual}
}

type Tim struct {
	actual time.Time
}

func (t *Tim) Matches(matcher Matcher[time.Time]) AssertionMessage {
	return matcher.Match(t.actual)
}

func (t *Tim) EqualTo(expected time.Time) AssertionMessage {
	return EqualTo(expected).Match(t.actual)
}

func (t *Tim) Before(expected time.Time) AssertionMessage {
	return Before(expected).Match(t.actual)
}

func (t *Tim) BeforeOrEqual(expected time.Time) AssertionMessage {
	return BeforeOrEqual(expected).Match(t.actual)
}

func (t *Tim) After(expected time.Time) AssertionMessage {
	return After(expected).Match(t.actual)
}

func (t *Tim) AfterOrEqual(expected time.Time) AssertionMessage {
	return AfterOrEqual(expected).Match(t.actual)
}

func (t *Tim) WithinDuration(expected time.Time, delta time.Duration) AssertionMessage {
	return WithinDuration(expected, delta).Match(t.actual)
}

func (t *Tim) WithinRange(start, end time.Time) AssertionMessage {
	return WithinRange(start, end).Match(t.actual)
}

func Before(expected time.Time) Matcher[time.Time] {
	return MatchFunc(func(actual time.Time) AssertionMessage {
		return Assert(actual.Before(expected), "got <%v>, wanted before <%v>", actual, expected)
	})
}

func BeforeOrEqual(expected time.Time) Matcher[time.Time] {
	return MatchFunc(func(actual time.Time) AssertionMessage {
		return Assert(actual.Before(expected) || actual.Equal(expected), "got <%v>, wanted before or equal to <%v>", actual, expected)
	})
}

func After(expected time.Time) Matcher[time.Time] {
	return MatchFunc(func(actual time.Time) AssertionMessage {
		return Assert(actual.After(expected), "got <%v>, wanted after <%v>", actual, expected)
	})
}

func AfterOrEqual(expected time.Time) Matcher[time.Time] {
	return MatchFunc(func(actual time.Time) AssertionMessage {
		return Assert(actual.After(expected) || actual.Equal(expected), "got <%v>, wanted after or equal to <%v>", actual, expected)
	})
}

func WithinDuration(expected time.Time, delta time.Duration) Matcher[time.Time] {
	return MatchFunc(func(actual time.Time) AssertionMessage {
		diff := actual.Sub(expected)
		if diff < 0 {
			diff = -diff
		}
		return Assert(diff <= delta, "got <%v>, wanted within <%v> of <%v>", actual, delta, expected)
	})
}

func WithinRange(start, end time.Time) Matcher[time.Time] {
	return MatchFunc(func(actual time.Time) AssertionMessage {
		return Assert((actual.Equal(start) || actual.After(start)) && (actual.Equal(end) || actual.Before(end)), "got <%v>, wanted in range <%v> to <%v>", actual, start, end)
	})
}
