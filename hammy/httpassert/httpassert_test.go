package httpassert_test

import (
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gogunit/gunit/eye"
	"github.com/gogunit/gunit/hammy"
	"github.com/gogunit/gunit/hammy/httpassert"
)

func Test_Status_success(t *testing.T) {
	assert := hammy.New(t)
	resp := newResponse(http.StatusCreated, nil, "")

	assert.Is(httpassert.Response(resp).Status(http.StatusCreated))
}

func Test_Status_failure(t *testing.T) {
	result := httpassert.Response(newResponse(http.StatusCreated, nil, "")).Status(http.StatusOK)

	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(result)
	aSpy.HadErrorContaining(t, "got status <201>, wanted <200>")
}

func Test_Status_failure_nil_response(t *testing.T) {
	result := httpassert.Response(nil).Status(http.StatusOK)

	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(result)
	aSpy.HadErrorContaining(t, "got nil response")
}

func Test_StatusInRange_success(t *testing.T) {
	assert := hammy.New(t)
	resp := newResponse(http.StatusNoContent, nil, "")

	assert.Is(httpassert.Response(resp).StatusInRange(200, 299))
}

func Test_StatusInRange_failure(t *testing.T) {
	result := httpassert.Response(newResponse(http.StatusBadRequest, nil, "")).StatusInRange(200, 299)

	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(result)
	aSpy.HadErrorContaining(t, "got status <400>, wanted in range <200..299>")
}

func Test_StatusInRange_failure_invalid_range(t *testing.T) {
	result := httpassert.Response(newResponse(http.StatusOK, nil, "")).StatusInRange(299, 200)

	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(result)
	aSpy.HadErrorContaining(t, "got invalid status range <299..200>")
}

func Test_Header_success(t *testing.T) {
	assert := hammy.New(t)
	resp := newResponse(http.StatusOK, http.Header{"Content-Type": {"application/json"}}, "")

	assert.Is(httpassert.Response(resp).Header("Content-Type", "application/json"))
}

func Test_Header_failure(t *testing.T) {
	resp := newResponse(http.StatusOK, http.Header{"Content-Type": {"text/plain"}}, "")
	result := httpassert.Response(resp).Header("Content-Type", "application/json")

	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(result)
	aSpy.HadErrorContaining(t, "got header <Content-Type>=<text/plain>, wanted <application/json>")
}

func Test_HeaderContains_success(t *testing.T) {
	assert := hammy.New(t)
	resp := newResponse(http.StatusOK, http.Header{"Content-Type": {"application/json; charset=utf-8"}}, "")

	assert.Is(httpassert.Response(resp).HeaderContains("Content-Type", "application/json"))
}

func Test_HeaderContains_failure(t *testing.T) {
	resp := newResponse(http.StatusOK, http.Header{"Content-Type": {"text/plain"}}, "")
	result := httpassert.Response(resp).HeaderContains("Content-Type", "application/json")

	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(result)
	aSpy.HadErrorContaining(t, "got header <Content-Type>=<text/plain>, wanted containing <application/json>")
}

func Test_BodyEqual_success(t *testing.T) {
	assert := hammy.New(t)
	resp := newResponse(http.StatusOK, nil, "hello world")

	assert.Is(httpassert.Response(resp).BodyEqual("hello world"))
}

func Test_BodyEqual_failure(t *testing.T) {
	resp := newResponse(http.StatusOK, nil, "hello world")
	result := httpassert.Response(resp).BodyEqual("goodbye")

	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(result)
	aSpy.HadErrorContaining(t, "got body <hello world>, wanted equal to <goodbye>")
}

func Test_BodyContains_success(t *testing.T) {
	assert := hammy.New(t)
	resp := newResponse(http.StatusOK, nil, "hello world")

	assert.Is(httpassert.Response(resp).BodyContains("world"))
}

func Test_BodyContains_failure(t *testing.T) {
	resp := newResponse(http.StatusOK, nil, "hello world")
	result := httpassert.Response(resp).BodyContains("goodbye")

	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(result)
	aSpy.HadErrorContaining(t, "got body <hello world>, wanted containing <goodbye>")
}

func Test_BodyMatchesRegexp_success(t *testing.T) {
	assert := hammy.New(t)
	resp := newResponse(http.StatusOK, nil, "status 204")

	assert.Is(httpassert.Response(resp).BodyMatchesRegexp(`status \d+`))
}

func Test_BodyMatchesRegexp_failure(t *testing.T) {
	resp := newResponse(http.StatusOK, nil, "status 204")
	result := httpassert.Response(resp).BodyMatchesRegexp(`status 5\d\d`)

	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(result)
	aSpy.HadErrorContaining(t, "got body <status 204>, wanted regexp <status 5\\d\\d>")
}

func Test_BodyMatchesRegexp_failure_invalid_pattern(t *testing.T) {
	resp := newResponse(http.StatusOK, nil, "status 204")
	result := httpassert.Response(resp).BodyMatchesRegexp(`(`)

	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(result)
	aSpy.HadErrorContaining(t, "invalid regexp <(>")
}

func Test_Body_assertion_restores_body(t *testing.T) {
	assert := hammy.New(t)
	resp := newResponse(http.StatusOK, nil, "hello world")

	assert.Is(httpassert.Response(resp).BodyContains("hello"))

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("read restored body: %v", err)
	}
	if string(body) != "hello world" {
		t.Fatalf("got restored body %q, wanted %q", string(body), "hello world")
	}
}

func Test_Body_assertion_failure_nil_response(t *testing.T) {
	result := httpassert.Response(nil).BodyEqual("hello")

	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(result)
	aSpy.HadErrorContaining(t, "got nil response")
}

func Test_Body_assertion_failure_read_error(t *testing.T) {
	resp := &http.Response{Body: errorReadCloser{}}
	result := httpassert.Response(resp).BodyEqual("hello")

	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(result)
	aSpy.HadErrorContaining(t, "got body read error: read failed")
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

type errorReadCloser struct{}

func (errorReadCloser) Read([]byte) (int, error) {
	return 0, errors.New("read failed")
}

func (errorReadCloser) Close() error {
	return nil
}

var _ io.ReadCloser = errorReadCloser{}

func Test_Response_methods_failure_nil_receiver(t *testing.T) {
	var resp *httpassert.Resp
	result := resp.Status(http.StatusOK)

	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(result)
	aSpy.HadErrorContaining(t, "got nil response")
}

func Test_Response_HasStatus_alias_success(t *testing.T) {
	assert := hammy.New(t)
	resp := newResponse(http.StatusOK, nil, "")

	assert.Is(httpassert.Response(resp).HasStatus(http.StatusOK))
}

func Test_Response_HasStatusInRange_alias_success(t *testing.T) {
	assert := hammy.New(t)
	resp := newResponse(http.StatusOK, nil, "")

	assert.Is(httpassert.Response(resp).HasStatusInRange(200, 299))
}

func Test_Response_HeaderEqualTo_alias_success(t *testing.T) {
	assert := hammy.New(t)
	resp := newResponse(http.StatusOK, http.Header{"Content-Type": {"application/json; charset=utf-8"}}, "")

	assert.Is(httpassert.Response(resp).HeaderEqualTo("Content-Type", "application/json; charset=utf-8"))
}

func Test_Response_HasHeader_alias_success(t *testing.T) {
	assert := hammy.New(t)
	resp := newResponse(http.StatusOK, http.Header{"Content-Type": {"application/json; charset=utf-8"}}, "")

	assert.Is(httpassert.Response(resp).HasHeader("Content-Type", "application/json; charset=utf-8"))
}

func Test_Response_HasHeaderContaining_alias_success(t *testing.T) {
	assert := hammy.New(t)
	resp := newResponse(http.StatusOK, http.Header{"Content-Type": {"application/json; charset=utf-8"}}, "")

	assert.Is(httpassert.Response(resp).HasHeaderContaining("Content-Type", "application/json"))
}

func Test_Response_BodyEqualTo_alias_success(t *testing.T) {
	assert := hammy.New(t)
	resp := newResponse(http.StatusOK, nil, "status 204")

	assert.Is(httpassert.Response(resp).BodyEqualTo("status 204"))
}

func Test_Response_HasBodyContaining_alias_success(t *testing.T) {
	assert := hammy.New(t)
	resp := newResponse(http.StatusOK, nil, "status 204")

	assert.Is(httpassert.Response(resp).HasBodyContaining("204"))
}

func Test_Response_BodyMatches_alias_success(t *testing.T) {
	assert := hammy.New(t)
	resp := newResponse(http.StatusOK, nil, "status 204")

	assert.Is(httpassert.Response(resp).BodyMatches(`status \d+`))
}
