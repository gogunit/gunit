package jsonassert_test

import (
	"errors"
	"io"
	"strings"
	"testing"

	"github.com/gogunit/gunit/eye"
	"github.com/gogunit/gunit/hammy"
	"github.com/gogunit/gunit/hammy/jsonassert"
)

func Test_Equal_success_compact_vs_formatted(t *testing.T) {
	assert := hammy.New(t)

	assert.Is(jsonassert.String(`{"name":"Ada","age":37}`).EqualTo(`{
		"name": "Ada",
		"age": 37
	}`))
}

func Test_Equal_success_reordered_object_keys(t *testing.T) {
	assert := hammy.New(t)

	assert.Is(jsonassert.String(`{"name":"Ada","age":37}`).EqualTo(`{"age":37,"name":"Ada"}`))
}

func Test_Equal_failure_array_order_mismatch(t *testing.T) {
	result := jsonassert.String(`{"values":[1,2,3]}`).EqualTo(`{"values":[3,2,1]}`)

	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(result)
	aSpy.HadErrorContaining(t, "JSON mismatch (-want +got):")
}

func Test_Equal_success_numeric_spelling_equivalence(t *testing.T) {
	assert := hammy.New(t)

	assert.Is(jsonassert.String(`{"one":1,"small":0.10}`).EqualTo(`{"one":1.0,"small":1e-1}`))
}

func Test_Equal_success_large_numeric_spelling_equivalence(t *testing.T) {
	assert := hammy.New(t)

	assert.Is(jsonassert.String(`{"id":123456789012345678901234567890}`).EqualTo(`{"id":123456789012345678901234567890.0}`))
}

func Test_Equal_failure_invalid_actual_json(t *testing.T) {
	result := jsonassert.String(`{"name":`).EqualTo(`{"name":"Ada"}`)

	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(result)
	aSpy.HadErrorContaining(t, "actual JSON invalid:")
}

func Test_Equal_failure_invalid_expected_json(t *testing.T) {
	result := jsonassert.String(`{"name":"Ada"}`).EqualTo(`{"name":`)

	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(result)
	aSpy.HadErrorContaining(t, "expected JSON invalid:")
}

func Test_Equal_failure_multiple_actual_json_values(t *testing.T) {
	result := jsonassert.String(`{"name":"Ada"} {"name":"Grace"}`).EqualTo(`{"name":"Ada"}`)

	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(result)
	aSpy.HadErrorContaining(t, "actual JSON invalid: multiple JSON values")
}

func Test_Equal_failure_includes_diff(t *testing.T) {
	result := jsonassert.String(`{"name":"Ada"}`).EqualTo(`{"name":"Grace"}`)

	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(result)
	aSpy.HadErrorContaining(t, "JSON mismatch (-want +got):")
	aSpy.HadErrorContaining(t, `Grace`)
	aSpy.HadErrorContaining(t, `Ada`)
}

func Test_EqualBytes_success(t *testing.T) {
	assert := hammy.New(t)

	assert.Is(jsonassert.Bytes([]byte(`{"name":"Ada"}`)).EqualTo([]byte(`{"name":"Ada"}`)))
}

func Test_EqualReader_success(t *testing.T) {
	assert := hammy.New(t)

	assert.Is(jsonassert.Reader(strings.NewReader(`{"one":1}`)).EqualTo(strings.NewReader(`{"one":1.0}`)))
}

func Test_EqualReader_failure_nil_actual_reader(t *testing.T) {
	result := jsonassert.Reader(nil).EqualTo(strings.NewReader(`{"name":"Ada"}`))

	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(result)
	aSpy.HadErrorContaining(t, "actual JSON reader is nil")
}

func Test_EqualReader_failure_expected_read_error(t *testing.T) {
	result := jsonassert.Reader(strings.NewReader(`{"name":"Ada"}`)).EqualTo(errorReader{})

	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(result)
	aSpy.HadErrorContaining(t, "expected JSON read error: read failed")
}

func Test_EqualLines_success_multiline(t *testing.T) {
	assert := hammy.New(t)

	assert.Is(jsonassert.String(`{"name":"Ada","age":37}` + "\n" + `{"tags":["go","json"]}`).LinesEqualTo(`{"age":37.0,"name":"Ada"}` + "\n" + `{"tags":["go","json"]}`))
}

func Test_EqualLines_success_trailing_newline(t *testing.T) {
	assert := hammy.New(t)

	assert.Is(jsonassert.String("{\"id\":1}\n{\"id\":2}\n").LinesEqualTo("{\"id\":1.0}\n{\"id\":2.0}\n"))
}

func Test_EqualLines_success_crlf(t *testing.T) {
	assert := hammy.New(t)

	assert.Is(jsonassert.String("{\"id\":1}\r\n{\"id\":2}\r\n").LinesEqualTo("{\"id\":1.0}\r\n{\"id\":2.0}\r\n"))
}

func Test_EqualLines_success_empty_inputs(t *testing.T) {
	assert := hammy.New(t)

	assert.Is(jsonassert.String("").LinesEqualTo(""))
}

func Test_EqualLinesWithOptions_success_ignore_paths_per_line(t *testing.T) {
	assert := hammy.New(t)

	assert.Is(jsonassert.String(`{"status":"ok","meta":{"request_id":"abc"}}`+"\n"+`{"status":"ok","meta":{"request_id":"def"}}`).LinesEqualToWithOptions(`{"status":"ok","meta":{"request_id":"uvw"}}`+"\n"+`{"status":"ok","meta":{"request_id":"xyz"}}`, jsonassert.IgnorePaths("meta.request_id")))
}

func Test_EqualLinesWithOptions_success_unordered_arrays_per_line(t *testing.T) {
	assert := hammy.New(t)

	assert.Is(jsonassert.String(`{"tags":["go","test"]}`+"\n"+`{"tags":["json","assert"]}`).LinesEqualToWithOptions(`{"tags":["test","go"]}`+"\n"+`{"tags":["assert","json"]}`, jsonassert.UnorderedArraysAt("tags")))
}

func Test_EqualLinesBytes_success(t *testing.T) {
	assert := hammy.New(t)

	assert.Is(jsonassert.Bytes([]byte(`{"id":1}` + "\n" + `{"id":2}`)).LinesEqualTo([]byte(`{"id":1.0}` + "\n" + `{"id":2.0}`)))
}

func Test_EqualLinesBytesWithOptions_success_ignore_paths_per_line(t *testing.T) {
	assert := hammy.New(t)

	assert.Is(jsonassert.Bytes([]byte(`{"status":"ok","meta":{"request_id":"abc"}}`+"\n"+`{"status":"ok","meta":{"request_id":"def"}}`)).LinesEqualToWithOptions([]byte(`{"status":"ok","meta":{"request_id":"uvw"}}`+"\n"+`{"status":"ok","meta":{"request_id":"xyz"}}`), jsonassert.IgnorePaths("meta.request_id")))
}

func Test_EqualLines_failure_reports_line_index(t *testing.T) {
	result := jsonassert.String(`{"id":1,"name":"Ada"}` + "\n" + `{"id":2,"name":"Grace"}`).LinesEqualTo(`{"id":1,"name":"Ada"}` + "\n" + `{"id":2,"name":"Katherine"}`)

	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(result)
	aSpy.HadErrorContaining(t, "JSONL line <1> mismatch (-want +got):")
	aSpy.HadErrorContaining(t, "Grace")
	aSpy.HadErrorContaining(t, "Katherine")
}

func Test_EqualLinesBytes_failure_reports_line_index(t *testing.T) {
	result := jsonassert.Bytes([]byte(`{"id":1}` + "\n" + `{"id":2}`)).LinesEqualTo([]byte(`{"id":1}` + "\n" + `{"id":3}`))

	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(result)
	aSpy.HadErrorContaining(t, "JSONL line <1> mismatch (-want +got):")
}

func Test_EqualLines_failure_invalid_actual_json_reports_line_index(t *testing.T) {
	result := jsonassert.String(`{"id":1}` + "\n" + `{"id":`).LinesEqualTo(`{"id":1}` + "\n" + `{"id":2}`)

	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(result)
	aSpy.HadErrorContaining(t, "actual JSONL line <1> invalid:")
}

func Test_EqualLines_failure_invalid_expected_json_reports_line_index(t *testing.T) {
	result := jsonassert.String(`{"id":1}` + "\n" + `{"id":2}`).LinesEqualTo(`{"id":1}` + "\n" + `{"id":`)

	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(result)
	aSpy.HadErrorContaining(t, "expected JSONL line <1> invalid:")
}

func Test_EqualLines_failure_blank_middle_line_invalid_json(t *testing.T) {
	result := jsonassert.String(`{"id":1}` + "\n\n" + `{"id":2}`).LinesEqualTo(`{"id":1}` + "\n\n" + `{"id":2}`)

	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(result)
	aSpy.HadErrorContaining(t, "actual JSONL line <1> invalid:")
}

func Test_EqualLines_failure_line_count_mismatch_reports_index(t *testing.T) {
	result := jsonassert.String(`{"id":1}`).LinesEqualTo(`{"id":1}` + "\n" + `{"id":2}`)

	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(result)
	aSpy.HadErrorContaining(t, "got JSONL line count <1>, wanted <2>; first differing line index <1>")
}

func Test_EqualLinesWithOptions_failure_invalid_option_reports_line_index(t *testing.T) {
	result := jsonassert.String(`{"status":"ok"}`).LinesEqualToWithOptions(`{"status":"ok"}`, jsonassert.IgnorePaths("meta."))

	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(result)
	aSpy.HadErrorContaining(t, "JSONL line <0>: invalid JSON path <meta.>: path ends with dot")
}

func Test_LinesContain_success_full_record(t *testing.T) {
	assert := hammy.New(t)

	assert.Is(jsonassert.String(`{"id":1,"name":"Ada"}` + "\n" + `{"id":2,"score":1.0}`).LinesContain(`{"score":1,"id":2}`))
}

func Test_LinesContain_success_ignore_paths(t *testing.T) {
	assert := hammy.New(t)

	assert.Is(jsonassert.String(`{"status":"ok","meta":{"request_id":"abc"}}`+"\n"+`{"status":"done","meta":{"request_id":"def"}}`).LinesContain(`{"status":"done","meta":{"request_id":"xyz"}}`, jsonassert.IgnorePaths("meta.request_id")))
}

func Test_LinesContain_failure_no_match(t *testing.T) {
	result := jsonassert.String(`{"id":1}` + "\n" + `{"id":2}`).LinesContain(`{"id":3}`)

	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(result)
	aSpy.HadErrorContaining(t, "got no matching JSONL line")
}

func Test_LinesContain_failure_invalid_expected_json(t *testing.T) {
	result := jsonassert.String(`{"id":1}`).LinesContain(`{"id":`)

	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(result)
	aSpy.HadErrorContaining(t, "expected JSON invalid:")
}

func Test_LinesContain_failure_invalid_actual_line(t *testing.T) {
	result := jsonassert.String(`{"id":1}` + "\n" + `{"id":`).LinesContain(`{"id":2}`)

	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(result)
	aSpy.HadErrorContaining(t, "actual JSONL line <1> invalid:")
}

func Test_LinesContain_failure_invalid_option_reports_line_index(t *testing.T) {
	result := jsonassert.String(`{"status":"ok"}`).LinesContain(`{"status":"ok"}`, jsonassert.IgnorePaths("meta."))

	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(result)
	aSpy.HadErrorContaining(t, "JSONL line <0>: invalid JSON path <meta.>: path ends with dot")
}

func Test_LinesContainSubset_success(t *testing.T) {
	assert := hammy.New(t)

	assert.Is(jsonassert.String(`{"status":"ok","meta":{"request_id":"abc","page":1}}` + "\n" + `{"status":"done"}`).LinesContainSubset(`{"meta":{"page":1.0}}`))
}

func Test_LinesContainSubset_success_ignore_paths(t *testing.T) {
	assert := hammy.New(t)

	assert.Is(jsonassert.String(`{"status":"ok","meta":{"request_id":"abc","page":1}}`).LinesContainSubset(`{"meta":{"request_id":"xyz","page":1.0}}`, jsonassert.IgnorePaths("meta.request_id")))
}

func Test_LinesContainSubset_failure_no_match(t *testing.T) {
	result := jsonassert.String(`{"status":"ok"}` + "\n" + `{"status":"done"}`).LinesContainSubset(`{"meta":{"page":1}}`)

	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(result)
	aSpy.HadErrorContaining(t, "got no JSONL line containing expected subset")
}

func Test_Valid_success(t *testing.T) {
	assert := hammy.New(t)

	assert.Is(jsonassert.String(`{"name":"Ada"}`).IsValid())
}

func Test_Valid_failure(t *testing.T) {
	result := jsonassert.String(`{"name":`).IsValid()

	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(result)
	aSpy.HadErrorContaining(t, "JSON invalid:")
}

func Test_ValidBytes_success(t *testing.T) {
	assert := hammy.New(t)

	assert.Is(jsonassert.Bytes([]byte(`{"name":"Ada"}`)).IsValid())
}

func Test_ValidReader_success(t *testing.T) {
	assert := hammy.New(t)

	assert.Is(jsonassert.Reader(strings.NewReader(`{"name":"Ada"}`)).IsValid())
}

func Test_ValidReader_failure_read_error(t *testing.T) {
	result := jsonassert.Reader(errorReader{}).IsValid()

	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(result)
	aSpy.HadErrorContaining(t, "actual JSON read error: read failed")
}

func Test_Contains_success_object_subset(t *testing.T) {
	assert := hammy.New(t)

	assert.Is(jsonassert.String(`{"status":"ok","meta":{"request_id":"abc","page":1}}`).Contains(`{"status":"ok","meta":{"page":1.0}}`))
}

func Test_Contains_failure_missing_field(t *testing.T) {
	result := jsonassert.String(`{"status":"ok"}`).Contains(`{"meta":{"page":1}}`)

	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(result)
	aSpy.HadErrorContaining(t, "JSON path <$.meta> missing")
}

func Test_Contains_failure_mismatched_value(t *testing.T) {
	result := jsonassert.String(`{"status":"ok"}`).Contains(`{"status":"failed"}`)

	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(result)
	aSpy.HadErrorContaining(t, "JSON path <$.status> mismatch")
}

func Test_Contains_failure_array_order_sensitive(t *testing.T) {
	result := jsonassert.String(`{"values":[1,2,3]}`).Contains(`{"values":[3,2,1]}`)

	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(result)
	aSpy.HadErrorContaining(t, "JSON path <$.values> mismatch")
}

func Test_ContainsBytes_success(t *testing.T) {
	assert := hammy.New(t)

	assert.Is(jsonassert.Bytes([]byte(`{"status":"ok","extra":true}`)).Contains([]byte(`{"status":"ok"}`)))
}

func Test_PathExists_success(t *testing.T) {
	assert := hammy.New(t)

	assert.Is(jsonassert.String(`{"user":{"name":"Ada"}}`).PathExists("user.name"))
}

func Test_PathExists_success_array_index(t *testing.T) {
	assert := hammy.New(t)

	assert.Is(jsonassert.String(`{"items":[{"id":1}]}`).PathExists("items[0].id"))
}

func Test_PathExists_failure_missing(t *testing.T) {
	result := jsonassert.String(`{"user":{"name":"Ada"}}`).PathExists("user.email")

	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(result)
	aSpy.HadErrorContaining(t, "JSON path <user.email> missing")
}

func Test_PathExists_failure_invalid_path(t *testing.T) {
	result := jsonassert.String(`{"user":{"name":"Ada"}}`).PathExists("user.")

	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(result)
	aSpy.HadErrorContaining(t, "invalid JSON path <user.>: path ends with dot")
}

func Test_PathMissing_success(t *testing.T) {
	assert := hammy.New(t)

	assert.Is(jsonassert.String(`{"user":{"name":"Ada"}}`).PathMissing("user.email"))
}

func Test_PathMissing_failure_exists(t *testing.T) {
	result := jsonassert.String(`{"user":{"name":"Ada"}}`).PathMissing("user.name")

	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(result)
	aSpy.HadErrorContaining(t, "JSON path <user.name> exists, wanted missing")
}

func Test_PathEqual_success(t *testing.T) {
	assert := hammy.New(t)

	assert.Is(jsonassert.String(`{"user":{"name":"Ada","age":37}}`).PathEqual("user.age", `37.0`))
}

func Test_PathEqual_failure_mismatch(t *testing.T) {
	result := jsonassert.String(`{"user":{"name":"Ada"}}`).PathEqual("user.name", `"Grace"`)

	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(result)
	aSpy.HadErrorContaining(t, "JSON path <user.name> mismatch (-want +got):")
	aSpy.HadErrorContaining(t, "Grace")
	aSpy.HadErrorContaining(t, "Ada")
}

func Test_PathEqualBytes_success(t *testing.T) {
	assert := hammy.New(t)

	assert.Is(jsonassert.Bytes([]byte(`{"user":{"name":"Ada"}}`)).PathEqual("user.name", []byte(`"Ada"`)))
}

func Test_ArrayContains_success(t *testing.T) {
	assert := hammy.New(t)

	assert.Is(jsonassert.String(`{"items":[{"id":1},{"id":2}]}`).ArrayContains("items", `{"id":2.0}`))
}

func Test_ArrayContains_failure_missing_element(t *testing.T) {
	result := jsonassert.String(`{"items":[{"id":1}]}`).ArrayContains("items", `{"id":2}`)

	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(result)
	aSpy.HadErrorContaining(t, "got no matching element at JSON path <items>")
}

func Test_ArrayContains_failure_non_array_path(t *testing.T) {
	result := jsonassert.String(`{"items":{"id":1}}`).ArrayContains("items", `{"id":1}`)

	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(result)
	aSpy.HadErrorContaining(t, "wanted array")
}

func Test_ArrayContainsBytes_success(t *testing.T) {
	assert := hammy.New(t)

	assert.Is(jsonassert.Bytes([]byte(`{"items":[1,2]}`)).ArrayContains("items", []byte(`2.0`)))
}

func Test_EqualWithOptions_success_ignore_paths(t *testing.T) {
	assert := hammy.New(t)

	assert.Is(jsonassert.String(`{"status":"ok","meta":{"request_id":"abc"}}`).EqualToWithOptions(`{"status":"ok","meta":{"request_id":"xyz"}}`, jsonassert.IgnorePaths("meta.request_id")))
}

func Test_EqualWithOptions_success_ignore_array_item_path(t *testing.T) {
	assert := hammy.New(t)

	assert.Is(jsonassert.String(`{"items":[{"id":1,"etag":"a"}]}`).EqualToWithOptions(`{"items":[{"id":1,"etag":"b"}]}`, jsonassert.IgnorePaths("items[0].etag")))
}

func Test_EqualWithOptions_failure_invalid_ignore_path(t *testing.T) {
	result := jsonassert.String(`{"status":"ok"}`).EqualToWithOptions(`{"status":"ok"}`, jsonassert.IgnorePaths("meta."))

	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(result)
	aSpy.HadErrorContaining(t, "invalid JSON path <meta.>: path ends with dot")
}

func Test_EqualWithOptions_success_unordered_array(t *testing.T) {
	assert := hammy.New(t)

	assert.Is(jsonassert.String(`{"tags":["go","test","json"]}`).EqualToWithOptions(`{"tags":["json","go","test"]}`, jsonassert.UnorderedArraysAt("tags")))
}

func Test_EqualWithOptions_success_unordered_array_of_objects(t *testing.T) {
	assert := hammy.New(t)

	assert.Is(jsonassert.String(`{"items":[{"id":2},{"id":1}]}`).EqualToWithOptions(`{"items":[{"id":1.0},{"id":2.0}]}`, jsonassert.UnorderedArraysAt("items")))
}

func Test_EqualWithOptions_failure_unordered_array_non_array_path(t *testing.T) {
	result := jsonassert.String(`{"items":{"id":1}}`).EqualToWithOptions(`{"items":{"id":1}}`, jsonassert.UnorderedArraysAt("items"))

	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(result)
	aSpy.HadErrorContaining(t, "got JSON path <items> type <map[string]interface {}>, wanted array")
}

func Test_EqualWithOptions_failure_without_ignore_paths(t *testing.T) {
	result := jsonassert.String(`{"status":"ok","meta":{"request_id":"abc"}}`).EqualToWithOptions(`{"status":"ok","meta":{"request_id":"xyz"}}`)

	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(result)
	aSpy.HadErrorContaining(t, "JSON mismatch (-want +got):")
}

func Test_EqualBytesWithOptions_success(t *testing.T) {
	assert := hammy.New(t)

	assert.Is(jsonassert.Bytes([]byte(`{"tags":["go","test"]}`)).EqualToWithOptions([]byte(`{"tags":["test","go"]}`), jsonassert.UnorderedArraysAt("tags")))
}

type errorReader struct{}

func (errorReader) Read([]byte) (int, error) {
	return 0, errors.New("read failed")
}

var _ io.Reader = errorReader{}
