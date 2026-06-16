package gunit_test

import (
	"testing"
	"time"

	"github.com/gogunit/gunit"
	"github.com/gogunit/gunit/eye"
)

func Test_Time_EqualTo_success(t *testing.T) {
	actual := time.Date(2026, 5, 12, 10, 30, 0, 0, time.UTC)

	gunit.Time(t, actual).EqualTo(actual)
}

func Test_Time_EqualTo_failure(t *testing.T) {
	aSpy := eye.Spy()
	actual := time.Date(2026, 5, 12, 10, 30, 0, 0, time.UTC)

	gunit.Time(aSpy, actual).EqualTo(actual.Add(time.Second))

	aSpy.HadErrorContaining(t, "wanted equal to")
}

func Test_Time_Before_success(t *testing.T) {
	actual := time.Date(2026, 5, 12, 10, 30, 0, 0, time.UTC)

	gunit.Time(t, actual).Before(actual.Add(time.Second))
}

func Test_Time_Before_failure(t *testing.T) {
	aSpy := eye.Spy()
	actual := time.Date(2026, 5, 12, 10, 30, 0, 0, time.UTC)

	gunit.Time(aSpy, actual).Before(actual)

	aSpy.HadErrorContaining(t, "wanted before")
}

func Test_Time_BeforeOrEqual_success(t *testing.T) {
	actual := time.Date(2026, 5, 12, 10, 30, 0, 0, time.UTC)

	gunit.Time(t, actual).BeforeOrEqual(actual)
	gunit.Time(t, actual.Add(-time.Second)).BeforeOrEqual(actual)
}

func Test_Time_BeforeOrEqual_failure(t *testing.T) {
	aSpy := eye.Spy()
	actual := time.Date(2026, 5, 12, 10, 30, 0, 0, time.UTC)

	gunit.Time(aSpy, actual.Add(time.Second)).BeforeOrEqual(actual)

	aSpy.HadErrorContaining(t, "wanted before or equal")
}

func Test_Time_After_success(t *testing.T) {
	actual := time.Date(2026, 5, 12, 10, 30, 0, 0, time.UTC)

	gunit.Time(t, actual).After(actual.Add(-time.Second))
}

func Test_Time_After_failure(t *testing.T) {
	aSpy := eye.Spy()
	actual := time.Date(2026, 5, 12, 10, 30, 0, 0, time.UTC)

	gunit.Time(aSpy, actual).After(actual)

	aSpy.HadErrorContaining(t, "wanted after")
}

func Test_Time_AfterOrEqual_success(t *testing.T) {
	actual := time.Date(2026, 5, 12, 10, 30, 0, 0, time.UTC)

	gunit.Time(t, actual).AfterOrEqual(actual)
	gunit.Time(t, actual.Add(time.Second)).AfterOrEqual(actual)
}

func Test_Time_AfterOrEqual_failure(t *testing.T) {
	aSpy := eye.Spy()
	actual := time.Date(2026, 5, 12, 10, 30, 0, 0, time.UTC)

	gunit.Time(aSpy, actual.Add(-time.Second)).AfterOrEqual(actual)

	aSpy.HadErrorContaining(t, "wanted after or equal")
}

func Test_Time_WithinDuration_success(t *testing.T) {
	expected := time.Date(2026, 5, 12, 10, 30, 0, 0, time.UTC)

	gunit.Time(t, expected.Add(500*time.Millisecond)).WithinDuration(expected, time.Second)
}

func Test_Time_WithinDuration_failure(t *testing.T) {
	aSpy := eye.Spy()
	expected := time.Date(2026, 5, 12, 10, 30, 0, 0, time.UTC)

	gunit.Time(aSpy, expected.Add(2*time.Second)).WithinDuration(expected, time.Second)

	aSpy.HadErrorContaining(t, "wanted within <1s>")
}

func Test_Time_WithinRange_success_inclusive(t *testing.T) {
	start := time.Date(2026, 5, 12, 10, 30, 0, 0, time.UTC)
	end := start.Add(time.Hour)

	gunit.Time(t, start).WithinRange(start, end)
	gunit.Time(t, end).WithinRange(start, end)
}

func Test_Time_WithinRange_failure(t *testing.T) {
	aSpy := eye.Spy()
	start := time.Date(2026, 5, 12, 10, 30, 0, 0, time.UTC)
	end := start.Add(time.Hour)

	gunit.Time(aSpy, start.Add(-time.Second)).WithinRange(start, end)

	aSpy.HadErrorContaining(t, "wanted in range")
}
