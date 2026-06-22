package httpassert

import (
	"bytes"
	"io"
	"net/http"
	"strings"

	"github.com/gogunit/gunit/hammy"
)

func Request(req *http.Request) *Req {
	return &Req{actual: req}
}

type Req struct {
	actual *http.Request
}

// HasMethod asserts that the request method equals expected. It is an alias for Method.
func (r *Req) HasMethod(expected string) hammy.AssertionMessage { return r.Method(expected) }

func (r *Req) Method(expected string) hammy.AssertionMessage {
	if r == nil || r.actual == nil {
		return hammy.Assert(false, "got nil request, wanted method <%s>", expected)
	}
	return hammy.Assert(r.actual.Method == expected, "got method <%s>, wanted <%s>", r.actual.Method, expected)
}

// HasPath asserts that the request path equals expected. It is an alias for Path.
func (r *Req) HasPath(expected string) hammy.AssertionMessage { return r.Path(expected) }

func (r *Req) Path(expected string) hammy.AssertionMessage {
	if r == nil || r.actual == nil {
		return hammy.Assert(false, "got nil request, wanted path <%s>", expected)
	}
	if r.actual.URL == nil {
		return hammy.Assert(false, "got nil request URL, wanted path <%s>", expected)
	}
	return hammy.Assert(r.actual.URL.Path == expected, "got path <%s>, wanted <%s>", r.actual.URL.Path, expected)
}

// URLEqualTo asserts that the request URL equals expected. It is an alias for URL.
func (r *Req) URLEqualTo(expected string) hammy.AssertionMessage { return r.URL(expected) }

func (r *Req) URL(expected string) hammy.AssertionMessage {
	if r == nil || r.actual == nil {
		return hammy.Assert(false, "got nil request, wanted URL <%s>", expected)
	}
	if r.actual.URL == nil {
		return hammy.Assert(false, "got nil request URL, wanted URL <%s>", expected)
	}
	actual := r.actual.URL.String()
	return hammy.Assert(actual == expected, "got URL <%s>, wanted <%s>", actual, expected)
}

func (r *Req) URLEqual(expected string) hammy.AssertionMessage {
	return r.URL(expected)
}

// HasHost asserts that the request host equals expected. It is an alias for Host.
func (r *Req) HasHost(expected string) hammy.AssertionMessage { return r.Host(expected) }

func (r *Req) Host(expected string) hammy.AssertionMessage {
	if r == nil || r.actual == nil {
		return hammy.Assert(false, "got nil request, wanted host <%s>", expected)
	}
	return hammy.Assert(r.actual.Host == expected, "got host <%s>, wanted <%s>", r.actual.Host, expected)
}

// HeaderEqualTo asserts that the request header equals expected. It is an alias for Header.
func (r *Req) HeaderEqualTo(key, expected string) hammy.AssertionMessage {
	return r.Header(key, expected)
}

// HasHeader asserts that the request header equals expected. It is an alias for Header.
func (r *Req) HasHeader(key, expected string) hammy.AssertionMessage { return r.Header(key, expected) }

func (r *Req) Header(key, expected string) hammy.AssertionMessage {
	if r == nil || r.actual == nil {
		return hammy.Assert(false, "got nil request, wanted header <%s> equal to <%s>", key, expected)
	}
	actual := r.actual.Header.Get(key)
	return hammy.Assert(actual == expected, "got header <%s>=<%s>, wanted <%s>", key, actual, expected)
}

// HasHeaderContaining asserts that the request header contains expected. It is an alias for HeaderContains.
func (r *Req) HasHeaderContaining(key, expected string) hammy.AssertionMessage {
	return r.HeaderContains(key, expected)
}

func (r *Req) HeaderContains(key, expected string) hammy.AssertionMessage {
	if r == nil || r.actual == nil {
		return hammy.Assert(false, "got nil request, wanted header <%s> containing <%s>", key, expected)
	}
	actual := r.actual.Header.Get(key)
	return hammy.Assert(strings.Contains(actual, expected), "got header <%s>=<%s>, wanted containing <%s>", key, actual, expected)
}

// HasQueryParam asserts that the request query parameter equals expected. It is an alias for QueryParam.
func (r *Req) HasQueryParam(key, expected string) hammy.AssertionMessage {
	return r.QueryParam(key, expected)
}

func (r *Req) QueryParam(key, expected string) hammy.AssertionMessage {
	if r == nil || r.actual == nil {
		return hammy.Assert(false, "got nil request, wanted query param <%s> equal to <%s>", key, expected)
	}
	if r.actual.URL == nil {
		return hammy.Assert(false, "got nil request URL, wanted query param <%s> equal to <%s>", key, expected)
	}
	actual := r.actual.URL.Query().Get(key)
	return hammy.Assert(actual == expected, "got query param <%s>=<%s>, wanted <%s>", key, actual, expected)
}

// BodyEqualTo asserts that the request body equals expected. It is an alias for BodyEqual.
func (r *Req) BodyEqualTo(expected string) hammy.AssertionMessage { return r.BodyEqual(expected) }

func (r *Req) BodyEqual(expected string) hammy.AssertionMessage {
	actual, result := readRequestBody(r)
	if !result.IsSuccessful {
		return result
	}
	return hammy.Assert(actual == expected, "got body <%s>, wanted equal to <%s>", actual, expected)
}

// HasBodyContaining asserts that the request body contains expected. It is an alias for BodyContains.
func (r *Req) HasBodyContaining(expected string) hammy.AssertionMessage {
	return r.BodyContains(expected)
}

func (r *Req) BodyContains(expected string) hammy.AssertionMessage {
	actual, result := readRequestBody(r)
	if !result.IsSuccessful {
		return result
	}
	return hammy.Assert(strings.Contains(actual, expected), "got body <%s>, wanted containing <%s>", actual, expected)
}

func readRequestBody(r *Req) (string, hammy.AssertionMessage) {
	if r == nil || r.actual == nil {
		return "", hammy.Assert(false, "got nil request, wanted request body")
	}
	if r.actual.Body == nil {
		r.actual.Body = io.NopCloser(bytes.NewReader(nil))
		return "", hammy.Assert(true, "read body")
	}

	body, err := io.ReadAll(r.actual.Body)
	_ = r.actual.Body.Close()
	r.actual.Body = io.NopCloser(bytes.NewReader(body))
	if err != nil {
		return "", hammy.Assert(false, "got body read error: %v", err)
	}
	return string(body), hammy.Assert(true, "read body")
}
