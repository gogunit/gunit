package httpassert_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gogunit/gunit/hammy"
	"github.com/gogunit/gunit/hammy/httpassert"
)

func Test_Recorder_Status_success(t *testing.T) {
	assert := hammy.New(t)
	rec := newRecorder(http.StatusCreated, nil, "")

	assert.Is(httpassert.Recorder(rec).Status(http.StatusCreated))
}

func Test_Recorder_Status_failure(t *testing.T) {
	rec := newRecorder(http.StatusCreated, nil, "")
	result := httpassert.Recorder(rec).Status(http.StatusOK)

	requireFailure(t, result, "got status <201>, wanted <200>")
}

func Test_Recorder_Status_failure_nil_recorder(t *testing.T) {
	result := httpassert.Recorder(nil).Status(http.StatusOK)

	requireFailure(t, result, "got nil response recorder")
}

func Test_Recorder_StatusInRange_success(t *testing.T) {
	assert := hammy.New(t)
	rec := newRecorder(http.StatusNoContent, nil, "")

	assert.Is(httpassert.Recorder(rec).StatusInRange(200, 299))
}

func Test_Recorder_StatusInRange_failure(t *testing.T) {
	rec := newRecorder(http.StatusBadRequest, nil, "")
	result := httpassert.Recorder(rec).StatusInRange(200, 299)

	requireFailure(t, result, "got status <400>, wanted in range <200..299>")
}

func Test_Recorder_StatusInRange_failure_invalid_range(t *testing.T) {
	rec := newRecorder(http.StatusOK, nil, "")
	result := httpassert.Recorder(rec).StatusInRange(299, 200)

	requireFailure(t, result, "got invalid status range <299..200>")
}

func Test_Recorder_Header_success(t *testing.T) {
	assert := hammy.New(t)
	rec := newRecorder(http.StatusOK, http.Header{"Content-Type": {"application/json"}}, "")

	assert.Is(httpassert.Recorder(rec).Header("Content-Type", "application/json"))
}

func Test_Recorder_Header_failure(t *testing.T) {
	rec := newRecorder(http.StatusOK, http.Header{"Content-Type": {"text/plain"}}, "")
	result := httpassert.Recorder(rec).Header("Content-Type", "application/json")

	requireFailure(t, result, "got header <Content-Type>=<text/plain>, wanted <application/json>")
}

func Test_Recorder_HeaderContains_success(t *testing.T) {
	assert := hammy.New(t)
	rec := newRecorder(http.StatusOK, http.Header{"Content-Type": {"application/json; charset=utf-8"}}, "")

	assert.Is(httpassert.Recorder(rec).HeaderContains("Content-Type", "application/json"))
}

func Test_Recorder_HeaderContains_failure(t *testing.T) {
	rec := newRecorder(http.StatusOK, http.Header{"Content-Type": {"text/plain"}}, "")
	result := httpassert.Recorder(rec).HeaderContains("Content-Type", "application/json")

	requireFailure(t, result, "got header <Content-Type>=<text/plain>, wanted containing <application/json>")
}

func Test_Recorder_BodyEqual_success(t *testing.T) {
	assert := hammy.New(t)
	rec := newRecorder(http.StatusOK, nil, "hello world")

	assert.Is(httpassert.Recorder(rec).BodyEqual("hello world"))
}

func Test_Recorder_BodyEqual_failure(t *testing.T) {
	rec := newRecorder(http.StatusOK, nil, "hello world")
	result := httpassert.Recorder(rec).BodyEqual("goodbye")

	requireFailure(t, result, "got body <hello world>, wanted equal to <goodbye>")
}

func Test_Recorder_BodyContains_success(t *testing.T) {
	assert := hammy.New(t)
	rec := newRecorder(http.StatusOK, nil, "hello world")

	assert.Is(httpassert.Recorder(rec).BodyContains("world"))
}

func Test_Recorder_BodyContains_failure(t *testing.T) {
	rec := newRecorder(http.StatusOK, nil, "hello world")
	result := httpassert.Recorder(rec).BodyContains("goodbye")

	requireFailure(t, result, "got body <hello world>, wanted containing <goodbye>")
}

func Test_Recorder_BodyMatchesRegexp_success(t *testing.T) {
	assert := hammy.New(t)
	rec := newRecorder(http.StatusOK, nil, "status 204")

	assert.Is(httpassert.Recorder(rec).BodyMatchesRegexp(`status \d+`))
}

func Test_Recorder_BodyMatchesRegexp_failure(t *testing.T) {
	rec := newRecorder(http.StatusOK, nil, "status 204")
	result := httpassert.Recorder(rec).BodyMatchesRegexp(`status 5\d\d`)

	requireFailure(t, result, "got body <status 204>, wanted regexp <status 5\\d\\d>")
}

func Test_Recorder_BodyMatchesRegexp_failure_invalid_pattern(t *testing.T) {
	rec := newRecorder(http.StatusOK, nil, "status 204")
	result := httpassert.Recorder(rec).BodyMatchesRegexp(`(`)

	requireFailure(t, result, "invalid regexp <(>")
}

func Test_Recorder_Body_assertions_restore_body(t *testing.T) {
	assert := hammy.New(t)
	recorderAssert := httpassert.Recorder(newRecorder(http.StatusOK, nil, "hello world"))

	assert.Is(recorderAssert.BodyContains("hello"))
	assert.Is(recorderAssert.BodyEqual("hello world"))
	assert.Is(recorderAssert.BodyMatchesRegexp(`world$`))
}

func newRecorder(status int, headers http.Header, body string) *httptest.ResponseRecorder {
	recorder := httptest.NewRecorder()
	for key, values := range headers {
		for _, value := range values {
			recorder.Header().Add(key, value)
		}
	}
	recorder.WriteHeader(status)
	_, _ = recorder.WriteString(body)
	return recorder
}
