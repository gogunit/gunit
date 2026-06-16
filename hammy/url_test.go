package hammy_test

import (
	"net/url"
	"testing"

	"github.com/gogunit/gunit/eye"
	a "github.com/gogunit/gunit/hammy"
)

func Test_URL_Scheme_success(t *testing.T) {
	assert := a.New(t)
	actual, err := url.Parse("https://example.com/path?name=hammy&empty=")
	assert.Is(a.NilError(err))

	assert.Is(a.URL(actual).Scheme("https"))
}

func Test_URL_Host_success(t *testing.T) {
	assert := a.New(t)
	actual, err := url.Parse("https://example.com/path?name=hammy&empty=")
	assert.Is(a.NilError(err))

	assert.Is(a.URL(actual).Host("example.com"))
}

func Test_URL_Path_success(t *testing.T) {
	assert := a.New(t)
	actual, err := url.Parse("https://example.com/path?name=hammy&empty=")
	assert.Is(a.NilError(err))

	assert.Is(a.URL(actual).Path("/path"))
}

func Test_URL_RawQuery_success(t *testing.T) {
	assert := a.New(t)
	actual, err := url.Parse("https://example.com/path?name=hammy&empty=")
	assert.Is(a.NilError(err))

	assert.Is(a.URL(actual).RawQuery("name=hammy&empty="))
}

func Test_URL_QueryParam_success(t *testing.T) {
	assert := a.New(t)
	actual, err := url.Parse("https://example.com/path?name=hammy&empty=")
	assert.Is(a.NilError(err))

	assert.Is(a.URL(actual).QueryParam("name", "hammy"))
}

func Test_URL_QueryParam_empty_success(t *testing.T) {
	assert := a.New(t)
	actual, err := url.Parse("https://example.com/path?name=hammy&empty=")
	assert.Is(a.NilError(err))

	assert.Is(a.URL(actual).QueryParam("empty", ""))
}

func Test_URL_NoQueryParam_success(t *testing.T) {
	assert := a.New(t)
	actual, err := url.Parse("https://example.com/path?name=hammy&empty=")
	assert.Is(a.NilError(err))

	assert.Is(a.URL(actual).NoQueryParam("missing"))
}

func Test_URL_String_success(t *testing.T) {
	assert := a.New(t)
	actual, err := url.Parse("https://example.com/path?name=hammy&empty=")
	assert.Is(a.NilError(err))

	assert.Is(a.URL(actual).String("https://example.com/path?name=hammy&empty="))
}

func Test_ParseURL_success(t *testing.T) {
	assert := a.New(t)
	assert.Is(a.ParseURL("https://example.com/path?name=hammy").Host("example.com"))
}

func Test_URL_nil_failure(t *testing.T) {
	aSpy := eye.Spy()
	assert := a.New(aSpy)
	assert.Is(a.URL(nil).Scheme("https"))
	aSpy.HadErrorContaining(t, "got nil URL")
}

func Test_ParseURL_invalid_failure(t *testing.T) {
	aSpy := eye.Spy()
	assert := a.New(aSpy)
	assert.Is(a.ParseURL("http://[::1").Host("example.com"))
	aSpy.HadErrorContaining(t, "invalid URL <http://[::1>")
}

func Test_URL_Scheme_failure(t *testing.T) {
	aSpy := eye.Spy()
	assert := a.New(aSpy)
	actual, _ := url.Parse("http://example.com")
	assert.Is(a.URL(actual).Scheme("https"))
	aSpy.HadErrorContaining(t, "got URL scheme <http>, wanted <https>")
}

func Test_URL_Host_failure(t *testing.T) {
	aSpy := eye.Spy()
	assert := a.New(aSpy)
	actual, _ := url.Parse("https://example.com")
	assert.Is(a.URL(actual).Host("example.org"))
	aSpy.HadErrorContaining(t, "got URL host <example.com>, wanted <example.org>")
}

func Test_URL_Path_failure(t *testing.T) {
	aSpy := eye.Spy()
	assert := a.New(aSpy)
	actual, _ := url.Parse("https://example.com/path")
	assert.Is(a.URL(actual).Path("/other"))
	aSpy.HadErrorContaining(t, "got URL path </path>, wanted </other>")
}

func Test_URL_RawQuery_failure(t *testing.T) {
	aSpy := eye.Spy()
	assert := a.New(aSpy)
	actual, _ := url.Parse("https://example.com/path?a=1")
	assert.Is(a.URL(actual).RawQuery("a=2"))
	aSpy.HadErrorContaining(t, "got URL raw query <a=1>, wanted <a=2>")
}

func Test_URL_QueryParam_missing_failure(t *testing.T) {
	aSpy := eye.Spy()
	assert := a.New(aSpy)
	actual, _ := url.Parse("https://example.com/path?a=1")
	assert.Is(a.URL(actual).QueryParam("b", "2"))
	aSpy.HadErrorContaining(t, "missing URL query parameter <b>, wanted <2>")
}

func Test_URL_QueryParam_value_failure(t *testing.T) {
	aSpy := eye.Spy()
	assert := a.New(aSpy)
	actual, _ := url.Parse("https://example.com/path?a=1")
	assert.Is(a.URL(actual).QueryParam("a", "2"))
	aSpy.HadErrorContaining(t, "got URL query parameter <a> value <1>, wanted <2>")
}

func Test_URL_NoQueryParam_failure(t *testing.T) {
	aSpy := eye.Spy()
	assert := a.New(aSpy)
	actual, _ := url.Parse("https://example.com/path?a=1")
	assert.Is(a.URL(actual).NoQueryParam("a"))
	aSpy.HadErrorContaining(t, "got URL query parameter <a>, wanted it absent")
}

func Test_URL_String_failure(t *testing.T) {
	aSpy := eye.Spy()
	assert := a.New(aSpy)
	actual, _ := url.Parse("https://example.com/path")
	assert.Is(a.URL(actual).String("https://example.org/path"))
	aSpy.HadErrorContaining(t, "got URL string <https://example.com/path>, wanted <https://example.org/path>")
}
