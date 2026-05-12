package hammy_test

import (
	"testing"
	"time"

	"github.com/gogunit/gunit/eye"
	a "github.com/gogunit/gunit/hammy"
)

func Test_Time_EqualTo_success(t *testing.T) {
	assert := a.New(t)
	actual := time.Date(2026, 5, 12, 10, 30, 0, 0, time.UTC)

	assert.Is(a.Time(actual).EqualTo(actual))
}

func Test_Time_Before_success(t *testing.T) {
	assert := a.New(t)
	actual := time.Date(2026, 5, 12, 10, 30, 0, 0, time.UTC)

	assert.Is(a.Time(actual).Before(actual.Add(time.Second)))
}

func Test_Time_Before_failure(t *testing.T) {
	aSpy := eye.Spy()
	assert := a.New(aSpy)
	actual := time.Date(2026, 5, 12, 10, 30, 0, 0, time.UTC)

	assert.Is(a.Time(actual).Before(actual))

	aSpy.HadErrorContaining(t, "wanted before")
}

func Test_Time_BeforeOrEqual_success(t *testing.T) {
	assert := a.New(t)
	actual := time.Date(2026, 5, 12, 10, 30, 0, 0, time.UTC)

	assert.Is(a.Time(actual).BeforeOrEqual(actual))
}

func Test_Time_After_success(t *testing.T) {
	assert := a.New(t)
	actual := time.Date(2026, 5, 12, 10, 30, 0, 0, time.UTC)

	assert.Is(a.Time(actual).After(actual.Add(-time.Second)))
}

func Test_Time_After_failure(t *testing.T) {
	aSpy := eye.Spy()
	assert := a.New(aSpy)
	actual := time.Date(2026, 5, 12, 10, 30, 0, 0, time.UTC)

	assert.Is(a.Time(actual).After(actual))

	aSpy.HadErrorContaining(t, "wanted after")
}

func Test_Time_AfterOrEqual_success(t *testing.T) {
	assert := a.New(t)
	actual := time.Date(2026, 5, 12, 10, 30, 0, 0, time.UTC)

	assert.Is(a.Time(actual).AfterOrEqual(actual))
}

func Test_Time_WithinDuration_success(t *testing.T) {
	assert := a.New(t)
	expected := time.Date(2026, 5, 12, 10, 30, 0, 0, time.UTC)

	assert.Is(a.Time(expected.Add(500*time.Millisecond)).WithinDuration(expected, time.Second))
}

func Test_Time_WithinDuration_failure(t *testing.T) {
	aSpy := eye.Spy()
	assert := a.New(aSpy)
	expected := time.Date(2026, 5, 12, 10, 30, 0, 0, time.UTC)

	assert.Is(a.Time(expected.Add(2*time.Second)).WithinDuration(expected, time.Second))

	aSpy.HadErrorContaining(t, "wanted within <1s>")
}

func Test_Time_WithinRange_success_inclusive(t *testing.T) {
	assert := a.New(t)
	start := time.Date(2026, 5, 12, 10, 30, 0, 0, time.UTC)
	end := start.Add(time.Hour)

	assert.Is(a.Time(start).WithinRange(start, end))
	assert.Is(a.Time(end).WithinRange(start, end))
}

func Test_Time_WithinRange_failure(t *testing.T) {
	aSpy := eye.Spy()
	assert := a.New(aSpy)
	start := time.Date(2026, 5, 12, 10, 30, 0, 0, time.UTC)
	end := start.Add(time.Hour)

	assert.Is(a.Time(start.Add(-time.Second)).WithinRange(start, end))

	aSpy.HadErrorContaining(t, "wanted in range")
}

func Test_Time_Matches_success(t *testing.T) {
	assert := a.New(t)
	actual := time.Date(2026, 5, 12, 10, 30, 0, 0, time.UTC)

	assert.Is(a.Time(actual).Matches(a.AllOf(
		a.After(actual.Add(-time.Second)),
		a.Before(actual.Add(time.Second)),
	)))
}
