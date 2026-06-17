package gunit_test

import (
	"github.com/gogunit/gunit"
	"github.com/gogunit/gunit/eye"
	"testing"
)

func Test_url_Scheme_success(t *testing.T) { testURL(t).Scheme("https") }
func Test_url_Scheme_failure(t *testing.T) {
	aSpy := eye.Spy()
	testURL(aSpy).Scheme("http")
	aSpy.HadErrorContaining(t, "wanted <http>")
}
func Test_url_Scheme_failure_for_nil(t *testing.T) {
	aSpy := eye.Spy()
	gunit.URL(aSpy, nil).Scheme("https")
	aSpy.HadErrorContaining(t, "got nil URL")
}

func Test_url_Host_success(t *testing.T) { testURL(t).Host("example.com") }
func Test_url_Host_failure(t *testing.T) {
	aSpy := eye.Spy()
	testURL(aSpy).Host("other.example")
	aSpy.HadErrorContaining(t, "wanted <other.example>")
}

func Test_url_Path_success(t *testing.T) { testURL(t).Path("/path") }
func Test_url_Path_failure(t *testing.T) {
	aSpy := eye.Spy()
	testURL(aSpy).Path("/other")
	aSpy.HadErrorContaining(t, "wanted </other>")
}

func Test_url_RawQuery_success(t *testing.T) { testURL(t).RawQuery("q=go") }
func Test_url_RawQuery_failure(t *testing.T) {
	aSpy := eye.Spy()
	testURL(aSpy).RawQuery("q=rust")
	aSpy.HadErrorContaining(t, "wanted <q=rust>")
}

func Test_url_QueryParam_success(t *testing.T) { testURL(t).QueryParam("q", "go") }
func Test_url_QueryParam_failure_for_missing(t *testing.T) {
	aSpy := eye.Spy()
	testURL(aSpy).QueryParam("missing", "go")
	aSpy.HadErrorContaining(t, "missing URL query parameter")
}
func Test_url_QueryParam_failure_for_mismatch(t *testing.T) {
	aSpy := eye.Spy()
	testURL(aSpy).QueryParam("q", "rust")
	aSpy.HadErrorContaining(t, "wanted <rust>")
}

func Test_url_NoQueryParam_success(t *testing.T) { testURL(t).NoQueryParam("missing") }
func Test_url_NoQueryParam_failure(t *testing.T) {
	aSpy := eye.Spy()
	testURL(aSpy).NoQueryParam("q")
	aSpy.HadErrorContaining(t, "wanted it absent")
}

func Test_url_String_success(t *testing.T) { testURL(t).String("https://example.com/path?q=go") }
func Test_url_String_failure(t *testing.T) {
	aSpy := eye.Spy()
	testURL(aSpy).String("https://example.com/other")
	aSpy.HadErrorContaining(t, "wanted <https://example.com/other>")
}
func Test_url_ParseURL_String_success(t *testing.T) {
	gunit.ParseURL(t, "https://example.com/path?q=go").String("https://example.com/path?q=go")
}
func Test_url_ParseURL_String_failure_for_invalid_url(t *testing.T) {
	aSpy := eye.Spy()
	gunit.ParseURL(aSpy, "%zz").String("x")
	aSpy.HadErrorContaining(t, "invalid URL")
}
