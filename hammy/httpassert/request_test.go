package httpassert_test

import (
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/gogunit/gunit/eye"
	"github.com/gogunit/gunit/hammy"
	"github.com/gogunit/gunit/hammy/httpassert"
)

func Test_Request_Method_success(t *testing.T) {
	assert := hammy.New(t)
	req := newRequest(http.MethodPost, "https://example.com/widgets", "")

	assert.Is(httpassert.Request(req).Method(http.MethodPost))
}

func Test_Request_Method_failure(t *testing.T) {
	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	req := newRequest(http.MethodPost, "https://example.com/widgets", "")

	assert.Is(httpassert.Request(req).Method(http.MethodGet))

	aSpy.HadErrorContaining(t, "got method <POST>, wanted <GET>")
}

func Test_Request_Method_failure_nil_request(t *testing.T) {
	aSpy := eye.Spy()
	assert := hammy.New(aSpy)

	assert.Is(httpassert.Request(nil).Method(http.MethodGet))

	aSpy.HadErrorContaining(t, "got nil request")
}

func Test_Request_Path_success(t *testing.T) {
	assert := hammy.New(t)
	req := newRequest(http.MethodGet, "https://example.com/widgets/1?expand=true", "")

	assert.Is(httpassert.Request(req).Path("/widgets/1"))
}

func Test_Request_Path_failure(t *testing.T) {
	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	req := newRequest(http.MethodGet, "https://example.com/widgets/1", "")

	assert.Is(httpassert.Request(req).Path("/widgets/2"))

	aSpy.HadErrorContaining(t, "got path </widgets/1>, wanted </widgets/2>")
}

func Test_Request_URL_success(t *testing.T) {
	assert := hammy.New(t)
	req := newRequest(http.MethodGet, "https://example.com/widgets/1?expand=true", "")

	assert.Is(httpassert.Request(req).URL("https://example.com/widgets/1?expand=true"))
	assert.Is(httpassert.Request(req).URLEqual("https://example.com/widgets/1?expand=true"))
}

func Test_Request_URL_failure(t *testing.T) {
	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	req := newRequest(http.MethodGet, "https://example.com/widgets/1", "")

	assert.Is(httpassert.Request(req).URL("https://example.com/widgets/2"))

	aSpy.HadErrorContaining(t, "got URL <https://example.com/widgets/1>, wanted <https://example.com/widgets/2>")
}

func Test_Request_URL_failure_nil_URL(t *testing.T) {
	aSpy := eye.Spy()
	assert := hammy.New(aSpy)

	assert.Is(httpassert.Request(&http.Request{}).URL("https://example.com"))

	aSpy.HadErrorContaining(t, "got nil request URL")
}

func Test_Request_Host_success(t *testing.T) {
	assert := hammy.New(t)
	req := newRequest(http.MethodGet, "https://example.com/widgets", "")

	assert.Is(httpassert.Request(req).Host("example.com"))
}

func Test_Request_Host_failure(t *testing.T) {
	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	req := newRequest(http.MethodGet, "https://example.com/widgets", "")

	assert.Is(httpassert.Request(req).Host("api.example.com"))

	aSpy.HadErrorContaining(t, "got host <example.com>, wanted <api.example.com>")
}

func Test_Request_Header_success(t *testing.T) {
	assert := hammy.New(t)
	req := newRequest(http.MethodGet, "https://example.com/widgets", "")
	req.Header.Set("Accept", "application/json")

	assert.Is(httpassert.Request(req).Header("Accept", "application/json"))
}

func Test_Request_Header_failure(t *testing.T) {
	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	req := newRequest(http.MethodGet, "https://example.com/widgets", "")
	req.Header.Set("Accept", "text/plain")

	assert.Is(httpassert.Request(req).Header("Accept", "application/json"))

	aSpy.HadErrorContaining(t, "got header <Accept>=<text/plain>, wanted <application/json>")
}

func Test_Request_HeaderContains_success(t *testing.T) {
	assert := hammy.New(t)
	req := newRequest(http.MethodGet, "https://example.com/widgets", "")
	req.Header.Set("Accept", "application/json; charset=utf-8")

	assert.Is(httpassert.Request(req).HeaderContains("Accept", "application/json"))
}

func Test_Request_HeaderContains_failure(t *testing.T) {
	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	req := newRequest(http.MethodGet, "https://example.com/widgets", "")
	req.Header.Set("Accept", "text/plain")

	assert.Is(httpassert.Request(req).HeaderContains("Accept", "application/json"))

	aSpy.HadErrorContaining(t, "got header <Accept>=<text/plain>, wanted containing <application/json>")
}

func Test_Request_QueryParam_success(t *testing.T) {
	assert := hammy.New(t)
	req := newRequest(http.MethodGet, "https://example.com/widgets?expand=true", "")

	assert.Is(httpassert.Request(req).QueryParam("expand", "true"))
}

func Test_Request_QueryParam_failure(t *testing.T) {
	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	req := newRequest(http.MethodGet, "https://example.com/widgets?expand=false", "")

	assert.Is(httpassert.Request(req).QueryParam("expand", "true"))

	aSpy.HadErrorContaining(t, "got query param <expand>=<false>, wanted <true>")
}

func Test_Request_BodyEqual_success(t *testing.T) {
	assert := hammy.New(t)
	req := newRequest(http.MethodPost, "https://example.com/widgets", "hello world")

	assert.Is(httpassert.Request(req).BodyEqual("hello world"))
}

func Test_Request_BodyEqual_failure(t *testing.T) {
	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	req := newRequest(http.MethodPost, "https://example.com/widgets", "hello world")

	assert.Is(httpassert.Request(req).BodyEqual("goodbye"))

	aSpy.HadErrorContaining(t, "got body <hello world>, wanted equal to <goodbye>")
}

func Test_Request_BodyContains_success(t *testing.T) {
	assert := hammy.New(t)
	req := newRequest(http.MethodPost, "https://example.com/widgets", "hello world")

	assert.Is(httpassert.Request(req).BodyContains("world"))
}

func Test_Request_BodyContains_failure(t *testing.T) {
	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	req := newRequest(http.MethodPost, "https://example.com/widgets", "hello world")

	assert.Is(httpassert.Request(req).BodyContains("goodbye"))

	aSpy.HadErrorContaining(t, "got body <hello world>, wanted containing <goodbye>")
}

func Test_Request_Body_assertions_restore_body(t *testing.T) {
	assert := hammy.New(t)
	req := newRequest(http.MethodPost, "https://example.com/widgets", "hello world")

	assert.Is(httpassert.Request(req).BodyContains("hello"))
	assert.Is(httpassert.Request(req).BodyEqual("hello world"))
}

func Test_Request_Body_assertion_failure_nil_request(t *testing.T) {
	aSpy := eye.Spy()
	assert := hammy.New(aSpy)

	assert.Is(httpassert.Request(nil).BodyEqual("hello"))

	aSpy.HadErrorContaining(t, "got nil request")
}

func Test_Request_Body_assertion_failure_read_error(t *testing.T) {
	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	req := newRequest(http.MethodPost, "https://example.com/widgets", "")
	req.Body = requestErrorReadCloser{}

	assert.Is(httpassert.Request(req).BodyEqual("hello"))

	aSpy.HadErrorContaining(t, "got body read error: read failed")
}

func newRequest(method, rawURL, body string) *http.Request {
	req := &http.Request{
		Method: method,
		Header: make(http.Header),
		Body:   io.NopCloser(strings.NewReader(body)),
	}
	req.URL = mustParseURL(rawURL)
	req.Host = req.URL.Host
	return req
}

func mustParseURL(rawURL string) *url.URL {
	parsed, err := url.Parse(rawURL)
	if err != nil {
		panic(err)
	}
	return parsed
}

type requestErrorReadCloser struct{}

func (requestErrorReadCloser) Read([]byte) (int, error) {
	return 0, errors.New("read failed")
}

func (requestErrorReadCloser) Close() error {
	return nil
}
