package gunit

import "time"

func Time(t T, actual time.Time) *Tim {
	return &Tim{t: t, actual: actual}
}

type Tim struct {
	t      T
	actual time.Time
}

func (t *Tim) EqualTo(expected time.Time) {
	t.t.Helper()
	if !t.actual.Equal(expected) {
		t.t.Errorf("got <%v>, wanted equal to <%v>", t.actual, expected)
	}
}

func (t *Tim) Before(expected time.Time) {
	t.t.Helper()
	if !t.actual.Before(expected) {
		t.t.Errorf("got <%v>, wanted before <%v>", t.actual, expected)
	}
}

func (t *Tim) BeforeOrEqual(expected time.Time) {
	t.t.Helper()
	if !t.actual.Before(expected) && !t.actual.Equal(expected) {
		t.t.Errorf("got <%v>, wanted before or equal to <%v>", t.actual, expected)
	}
}

func (t *Tim) After(expected time.Time) {
	t.t.Helper()
	if !t.actual.After(expected) {
		t.t.Errorf("got <%v>, wanted after <%v>", t.actual, expected)
	}
}

func (t *Tim) AfterOrEqual(expected time.Time) {
	t.t.Helper()
	if !t.actual.After(expected) && !t.actual.Equal(expected) {
		t.t.Errorf("got <%v>, wanted after or equal to <%v>", t.actual, expected)
	}
}

func (t *Tim) WithinDuration(expected time.Time, delta time.Duration) {
	t.t.Helper()
	diff := t.actual.Sub(expected)
	if diff < 0 {
		diff = -diff
	}
	if diff > delta {
		t.t.Errorf("got <%v>, wanted within <%v> of <%v>", t.actual, delta, expected)
	}
}

func (t *Tim) WithinRange(start, end time.Time) {
	t.t.Helper()
	if (t.actual.Before(start) && !t.actual.Equal(start)) || (t.actual.After(end) && !t.actual.Equal(end)) {
		t.t.Errorf("got <%v>, wanted in range <%v> to <%v>", t.actual, start, end)
	}
}
