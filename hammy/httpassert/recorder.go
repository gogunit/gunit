package httpassert

import (
	"net/http"
	"net/http/httptest"

	"github.com/gogunit/gunit/hammy"
)

// Recorder wraps an httptest.ResponseRecorder with HTTP assertion helpers.
func Recorder(rec *httptest.ResponseRecorder) *RecorderAssert {
	return &RecorderAssert{actual: rec}
}

// RecorderAssert provides HTTP assertions for an httptest.ResponseRecorder.
type RecorderAssert struct {
	actual *httptest.ResponseRecorder
}

// Status asserts that the recorded response status equals expected.
func (r *RecorderAssert) Status(expected int) hammy.AssertionMessage {
	resp, result := r.response()
	if !result.IsSuccessful {
		return hammy.Assert(false, "got nil response recorder, wanted status <%d>", expected)
	}
	return Status(resp, expected)
}

// StatusInRange asserts that the recorded response status is between min and max, inclusive.
func (r *RecorderAssert) StatusInRange(min, max int) hammy.AssertionMessage {
	if min > max {
		return hammy.Assert(false, "got invalid status range <%d..%d>", min, max)
	}
	resp, result := r.response()
	if !result.IsSuccessful {
		return hammy.Assert(false, "got nil response recorder, wanted status in range <%d..%d>", min, max)
	}
	return StatusInRange(resp, min, max)
}

// Header asserts that the recorded response header equals expected.
func (r *RecorderAssert) Header(key, expected string) hammy.AssertionMessage {
	resp, result := r.response()
	if !result.IsSuccessful {
		return hammy.Assert(false, "got nil response recorder, wanted header <%s> equal to <%s>", key, expected)
	}
	return Header(resp, key, expected)
}

// HeaderContains asserts that the recorded response header contains expected.
func (r *RecorderAssert) HeaderContains(key, expected string) hammy.AssertionMessage {
	resp, result := r.response()
	if !result.IsSuccessful {
		return hammy.Assert(false, "got nil response recorder, wanted header <%s> containing <%s>", key, expected)
	}
	return HeaderContains(resp, key, expected)
}

// BodyEqual asserts that the recorded response body equals expected.
func (r *RecorderAssert) BodyEqual(expected string) hammy.AssertionMessage {
	resp, result := r.response()
	if !result.IsSuccessful {
		return hammy.Assert(false, "got nil response recorder, wanted response body")
	}
	return BodyEqual(resp, expected)
}

// BodyContains asserts that the recorded response body contains expected.
func (r *RecorderAssert) BodyContains(expected string) hammy.AssertionMessage {
	resp, result := r.response()
	if !result.IsSuccessful {
		return hammy.Assert(false, "got nil response recorder, wanted response body")
	}
	return BodyContains(resp, expected)
}

// BodyMatchesRegexp asserts that the recorded response body matches pattern.
func (r *RecorderAssert) BodyMatchesRegexp(pattern string) hammy.AssertionMessage {
	resp, result := r.response()
	if !result.IsSuccessful {
		return hammy.Assert(false, "got nil response recorder, wanted response body")
	}
	return BodyMatchesRegexp(resp, pattern)
}

func (r *RecorderAssert) response() (*http.Response, hammy.AssertionMessage) {
	if r == nil || r.actual == nil {
		return nil, hammy.Assert(false, "got nil response recorder")
	}
	return r.actual.Result(), hammy.Assert(true, "got response recorder")
}
