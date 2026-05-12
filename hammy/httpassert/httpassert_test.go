package httpassert_test

import (
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gogunit/gunit/hammy"
	"github.com/gogunit/gunit/hammy/httpassert"
)

func Test_Status_success(t *testing.T) {
	assert := hammy.New(t)
	resp := newResponse(http.StatusCreated, nil, "")

	assert.Is(httpassert.Status(resp, http.StatusCreated))
}

func Test_Status_failure(t *testing.T) {
	result := httpassert.Status(newResponse(http.StatusCreated, nil, ""), http.StatusOK)

	requireFailure(t, result, "got status <201>, wanted <200>")
}

func Test_Status_failure_nil_response(t *testing.T) {
	result := httpassert.Status(nil, http.StatusOK)

	requireFailure(t, result, "got nil response")
}

func Test_StatusInRange_success(t *testing.T) {
	assert := hammy.New(t)
	resp := newResponse(http.StatusNoContent, nil, "")

	assert.Is(httpassert.StatusInRange(resp, 200, 299))
}

func Test_StatusInRange_failure(t *testing.T) {
	result := httpassert.StatusInRange(newResponse(http.StatusBadRequest, nil, ""), 200, 299)

	requireFailure(t, result, "got status <400>, wanted in range <200..299>")
}

func Test_StatusInRange_failure_invalid_range(t *testing.T) {
	result := httpassert.StatusInRange(newResponse(http.StatusOK, nil, ""), 299, 200)

	requireFailure(t, result, "got invalid status range <299..200>")
}

func Test_Header_success(t *testing.T) {
	assert := hammy.New(t)
	resp := newResponse(http.StatusOK, http.Header{"Content-Type": {"application/json"}}, "")

	assert.Is(httpassert.Header(resp, "Content-Type", "application/json"))
}

func Test_Header_failure(t *testing.T) {
	resp := newResponse(http.StatusOK, http.Header{"Content-Type": {"text/plain"}}, "")
	result := httpassert.Header(resp, "Content-Type", "application/json")

	requireFailure(t, result, "got header <Content-Type>=<text/plain>, wanted <application/json>")
}

func Test_HeaderContains_success(t *testing.T) {
	assert := hammy.New(t)
	resp := newResponse(http.StatusOK, http.Header{"Content-Type": {"application/json; charset=utf-8"}}, "")

	assert.Is(httpassert.HeaderContains(resp, "Content-Type", "application/json"))
}

func Test_HeaderContains_failure(t *testing.T) {
	resp := newResponse(http.StatusOK, http.Header{"Content-Type": {"text/plain"}}, "")
	result := httpassert.HeaderContains(resp, "Content-Type", "application/json")

	requireFailure(t, result, "got header <Content-Type>=<text/plain>, wanted containing <application/json>")
}

func Test_BodyEqual_success(t *testing.T) {
	assert := hammy.New(t)
	resp := newResponse(http.StatusOK, nil, "hello world")

	assert.Is(httpassert.BodyEqual(resp, "hello world"))
}

func Test_BodyEqual_failure(t *testing.T) {
	resp := newResponse(http.StatusOK, nil, "hello world")
	result := httpassert.BodyEqual(resp, "goodbye")

	requireFailure(t, result, "got body <hello world>, wanted equal to <goodbye>")
}

func Test_BodyContains_success(t *testing.T) {
	assert := hammy.New(t)
	resp := newResponse(http.StatusOK, nil, "hello world")

	assert.Is(httpassert.BodyContains(resp, "world"))
}

func Test_BodyContains_failure(t *testing.T) {
	resp := newResponse(http.StatusOK, nil, "hello world")
	result := httpassert.BodyContains(resp, "goodbye")

	requireFailure(t, result, "got body <hello world>, wanted containing <goodbye>")
}

func Test_BodyMatchesRegexp_success(t *testing.T) {
	assert := hammy.New(t)
	resp := newResponse(http.StatusOK, nil, "status 204")

	assert.Is(httpassert.BodyMatchesRegexp(resp, `status \d+`))
}

func Test_BodyMatchesRegexp_failure(t *testing.T) {
	resp := newResponse(http.StatusOK, nil, "status 204")
	result := httpassert.BodyMatchesRegexp(resp, `status 5\d\d`)

	requireFailure(t, result, "got body <status 204>, wanted regexp <status 5\\d\\d>")
}

func Test_BodyMatchesRegexp_failure_invalid_pattern(t *testing.T) {
	resp := newResponse(http.StatusOK, nil, "status 204")
	result := httpassert.BodyMatchesRegexp(resp, `(`)

	requireFailure(t, result, "invalid regexp <(>")
}

func Test_Body_assertions_restore_body(t *testing.T) {
	assert := hammy.New(t)
	resp := newResponse(http.StatusOK, nil, "hello world")

	assert.Is(httpassert.BodyContains(resp, "hello"))
	assert.Is(httpassert.BodyEqual(resp, "hello world"))
	assert.Is(httpassert.BodyMatchesRegexp(resp, `world$`))
}

func Test_Body_assertion_failure_nil_response(t *testing.T) {
	result := httpassert.BodyEqual(nil, "hello")

	requireFailure(t, result, "got nil response")
}

func Test_Body_assertion_failure_read_error(t *testing.T) {
	resp := &http.Response{Body: errorReadCloser{}}
	result := httpassert.BodyEqual(resp, "hello")

	requireFailure(t, result, "got body read error: read failed")
}

func newResponse(status int, headers http.Header, body string) *http.Response {
	recorder := httptest.NewRecorder()
	for key, values := range headers {
		for _, value := range values {
			recorder.Header().Add(key, value)
		}
	}
	recorder.WriteHeader(status)
	_, _ = recorder.WriteString(body)
	return recorder.Result()
}

func requireFailure(t *testing.T, result hammy.AssertionMessage, contains string) {
	t.Helper()
	if result.IsSuccessful {
		t.Fatalf("got success, wanted failure containing %q", contains)
	}
	if !strings.Contains(result.Message, contains) {
		t.Fatalf("got message %q, wanted containing %q", result.Message, contains)
	}
}

type errorReadCloser struct{}

func (errorReadCloser) Read([]byte) (int, error) {
	return 0, errors.New("read failed")
}

func (errorReadCloser) Close() error {
	return nil
}

var _ io.ReadCloser = errorReadCloser{}
