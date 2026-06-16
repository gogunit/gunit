package httpassert

import (
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"

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
	return hammy.Assert(resp.StatusCode == expected, "got status <%d>, wanted <%d>", resp.StatusCode, expected)
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
	return hammy.Assert(resp.StatusCode >= min && resp.StatusCode <= max, "got status <%d>, wanted in range <%d..%d>", resp.StatusCode, min, max)
}

// Header asserts that the recorded response header equals expected.
func (r *RecorderAssert) Header(key, expected string) hammy.AssertionMessage {
	resp, result := r.response()
	if !result.IsSuccessful {
		return hammy.Assert(false, "got nil response recorder, wanted header <%s> equal to <%s>", key, expected)
	}
	actual := resp.Header.Get(key)
	return hammy.Assert(actual == expected, "got header <%s>=<%s>, wanted <%s>", key, actual, expected)
}

// HeaderContains asserts that the recorded response header contains expected.
func (r *RecorderAssert) HeaderContains(key, expected string) hammy.AssertionMessage {
	resp, result := r.response()
	if !result.IsSuccessful {
		return hammy.Assert(false, "got nil response recorder, wanted header <%s> containing <%s>", key, expected)
	}
	actual := resp.Header.Get(key)
	return hammy.Assert(strings.Contains(actual, expected), "got header <%s>=<%s>, wanted containing <%s>", key, actual, expected)
}

// BodyEqual asserts that the recorded response body equals expected.
func (r *RecorderAssert) BodyEqual(expected string) hammy.AssertionMessage {
	actual, result := r.body()
	if !result.IsSuccessful {
		return hammy.Assert(false, "got nil response recorder, wanted response body")
	}
	return hammy.Assert(actual == expected, "got body <%s>, wanted equal to <%s>", actual, expected)
}

// BodyContains asserts that the recorded response body contains expected.
func (r *RecorderAssert) BodyContains(expected string) hammy.AssertionMessage {
	actual, result := r.body()
	if !result.IsSuccessful {
		return hammy.Assert(false, "got nil response recorder, wanted response body")
	}
	return hammy.Assert(strings.Contains(actual, expected), "got body <%s>, wanted containing <%s>", actual, expected)
}

// BodyMatchesRegexp asserts that the recorded response body matches pattern.
func (r *RecorderAssert) BodyMatchesRegexp(pattern string) hammy.AssertionMessage {
	actual, result := r.body()
	if !result.IsSuccessful {
		return hammy.Assert(false, "got nil response recorder, wanted response body")
	}
	re, err := regexp.Compile(pattern)
	if err != nil {
		return hammy.Assert(false, "invalid regexp <%s>: %v", pattern, err)
	}
	return hammy.Assert(re.MatchString(actual), "got body <%s>, wanted regexp <%s>", actual, pattern)
}

func (r *RecorderAssert) response() (*http.Response, hammy.AssertionMessage) {
	if r == nil || r.actual == nil {
		return nil, hammy.Assert(false, "got nil response recorder")
	}
	return r.actual.Result(), hammy.Assert(true, "got response recorder")
}

func (r *RecorderAssert) body() (string, hammy.AssertionMessage) {
	if r == nil || r.actual == nil {
		return "", hammy.Assert(false, "got nil response recorder")
	}
	if r.actual.Body == nil {
		return "", hammy.Assert(true, "got response recorder body")
	}
	return r.actual.Body.String(), hammy.Assert(true, "got response recorder body")
}
