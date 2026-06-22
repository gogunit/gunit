package httpassert

import (
	"bytes"
	"io"
	"net/http"
	"regexp"
	"strings"

	"github.com/gogunit/gunit/hammy"
)

func Response(resp *http.Response) *Resp {
	return &Resp{actual: resp}
}

type Resp struct {
	actual *http.Response
}

// HasStatus asserts that the response status equals expected. It is an alias for Status.
func (r *Resp) HasStatus(expected int) hammy.AssertionMessage {
	return r.Status(expected)
}

func (r *Resp) Status(expected int) hammy.AssertionMessage {
	resp := r.response()
	if resp == nil {
		return hammy.Assert(false, "got nil response, wanted status <%d>", expected)
	}
	return hammy.Assert(resp.StatusCode == expected, "got status <%d>, wanted <%d>", resp.StatusCode, expected)
}

// HasStatusInRange asserts that the response status is between min and max, inclusive. It is an alias for StatusInRange.
func (r *Resp) HasStatusInRange(min, max int) hammy.AssertionMessage {
	return r.StatusInRange(min, max)
}

func (r *Resp) StatusInRange(min, max int) hammy.AssertionMessage {
	resp := r.response()
	if min > max {
		return hammy.Assert(false, "got invalid status range <%d..%d>", min, max)
	}
	if resp == nil {
		return hammy.Assert(false, "got nil response, wanted status in range <%d..%d>", min, max)
	}
	return hammy.Assert(resp.StatusCode >= min && resp.StatusCode <= max, "got status <%d>, wanted in range <%d..%d>", resp.StatusCode, min, max)
}

// HeaderEqualTo asserts that the response header equals expected. It is an alias for Header.
func (r *Resp) HeaderEqualTo(key, expected string) hammy.AssertionMessage {
	return r.Header(key, expected)
}

// HasHeader asserts that the response header equals expected. It is an alias for Header.
func (r *Resp) HasHeader(key, expected string) hammy.AssertionMessage {
	return r.Header(key, expected)
}

func (r *Resp) Header(key, expected string) hammy.AssertionMessage {
	resp := r.response()
	if resp == nil {
		return hammy.Assert(false, "got nil response, wanted header <%s> equal to <%s>", key, expected)
	}
	actual := resp.Header.Get(key)
	return hammy.Assert(actual == expected, "got header <%s>=<%s>, wanted <%s>", key, actual, expected)
}

// HasHeaderContaining asserts that the response header contains expected. It is an alias for HeaderContains.
func (r *Resp) HasHeaderContaining(key, expected string) hammy.AssertionMessage {
	return r.HeaderContains(key, expected)
}

func (r *Resp) HeaderContains(key, expected string) hammy.AssertionMessage {
	resp := r.response()
	if resp == nil {
		return hammy.Assert(false, "got nil response, wanted header <%s> containing <%s>", key, expected)
	}
	actual := resp.Header.Get(key)
	return hammy.Assert(strings.Contains(actual, expected), "got header <%s>=<%s>, wanted containing <%s>", key, actual, expected)
}

// BodyEqualTo asserts that the response body equals expected. It is an alias for BodyEqual.
func (r *Resp) BodyEqualTo(expected string) hammy.AssertionMessage {
	return r.BodyEqual(expected)
}

func (r *Resp) BodyEqual(expected string) hammy.AssertionMessage {
	actual, result := readBody(r.response())
	if !result.IsSuccessful {
		return result
	}
	return hammy.Assert(actual == expected, "got body <%s>, wanted equal to <%s>", actual, expected)
}

// HasBodyContaining asserts that the response body contains expected. It is an alias for BodyContains.
func (r *Resp) HasBodyContaining(expected string) hammy.AssertionMessage {
	return r.BodyContains(expected)
}

func (r *Resp) BodyContains(expected string) hammy.AssertionMessage {
	actual, result := readBody(r.response())
	if !result.IsSuccessful {
		return result
	}
	return hammy.Assert(strings.Contains(actual, expected), "got body <%s>, wanted containing <%s>", actual, expected)
}

// BodyMatches asserts that the response body matches pattern. It is an alias for BodyMatchesRegexp.
func (r *Resp) BodyMatches(pattern string) hammy.AssertionMessage {
	return r.BodyMatchesRegexp(pattern)
}

func (r *Resp) BodyMatchesRegexp(pattern string) hammy.AssertionMessage {
	actual, result := readBody(r.response())
	if !result.IsSuccessful {
		return result
	}

	re, err := regexp.Compile(pattern)
	if err != nil {
		return hammy.Assert(false, "invalid regexp <%s>: %v", pattern, err)
	}
	return hammy.Assert(re.MatchString(actual), "got body <%s>, wanted regexp <%s>", actual, pattern)
}

func (r *Resp) response() *http.Response {
	if r == nil {
		return nil
	}
	return r.actual
}

func readBody(resp *http.Response) (string, hammy.AssertionMessage) {
	if resp == nil {
		return "", hammy.Assert(false, "got nil response, wanted response body")
	}
	if resp.Body == nil {
		resp.Body = io.NopCloser(bytes.NewReader(nil))
		return "", hammy.Assert(true, "read body")
	}

	body, err := io.ReadAll(resp.Body)
	_ = resp.Body.Close()
	resp.Body = io.NopCloser(bytes.NewReader(body))
	if err != nil {
		return "", hammy.Assert(false, "got body read error: %v", err)
	}
	return string(body), hammy.Assert(true, "read body")
}
