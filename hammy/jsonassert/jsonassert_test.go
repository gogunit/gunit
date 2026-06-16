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

	assert.Is(jsonassert.Equal(`{"name":"Ada","age":37}`, `{
		"name": "Ada",
		"age": 37
	}`))
}

func Test_Equal_success_reordered_object_keys(t *testing.T) {
	assert := hammy.New(t)

	assert.Is(jsonassert.Equal(`{"name":"Ada","age":37}`, `{"age":37,"name":"Ada"}`))
}

func Test_Equal_failure_array_order_mismatch(t *testing.T) {
	result := jsonassert.Equal(`{"values":[1,2,3]}`, `{"values":[3,2,1]}`)

	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(result)
	aSpy.HadErrorContaining(t, "JSON mismatch (-want +got):")
}

func Test_Equal_success_numeric_spelling_equivalence(t *testing.T) {
	assert := hammy.New(t)

	assert.Is(jsonassert.Equal(`{"one":1,"small":0.10}`, `{"one":1.0,"small":1e-1}`))
}

func Test_Equal_success_large_numeric_spelling_equivalence(t *testing.T) {
	assert := hammy.New(t)

	assert.Is(jsonassert.Equal(
		`{"id":123456789012345678901234567890}`,
		`{"id":123456789012345678901234567890.0}`,
	))
}

func Test_Equal_failure_invalid_actual_json(t *testing.T) {
	result := jsonassert.Equal(`{"name":`, `{"name":"Ada"}`)

	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(result)
	aSpy.HadErrorContaining(t, "actual JSON invalid:")
}

func Test_Equal_failure_invalid_expected_json(t *testing.T) {
	result := jsonassert.Equal(`{"name":"Ada"}`, `{"name":`)

	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(result)
	aSpy.HadErrorContaining(t, "expected JSON invalid:")
}

func Test_Equal_failure_multiple_actual_json_values(t *testing.T) {
	result := jsonassert.Equal(`{"name":"Ada"} {"name":"Grace"}`, `{"name":"Ada"}`)

	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(result)
	aSpy.HadErrorContaining(t, "actual JSON invalid: multiple JSON values")
}

func Test_Equal_failure_includes_diff(t *testing.T) {
	result := jsonassert.Equal(`{"name":"Ada"}`, `{"name":"Grace"}`)

	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(result)
	aSpy.HadErrorContaining(t, "JSON mismatch (-want +got):")
	aSpy.HadErrorContaining(t, `Grace`)
	aSpy.HadErrorContaining(t, `Ada`)
}

func Test_EqualBytes_success(t *testing.T) {
	assert := hammy.New(t)

	assert.Is(jsonassert.EqualBytes([]byte(`{"name":"Ada"}`), []byte(`{"name":"Ada"}`)))
}

func Test_EqualReader_success(t *testing.T) {
	assert := hammy.New(t)

	assert.Is(jsonassert.EqualReader(
		strings.NewReader(`{"one":1}`),
		strings.NewReader(`{"one":1.0}`),
	))
}

func Test_EqualReader_failure_nil_actual_reader(t *testing.T) {
	result := jsonassert.EqualReader(nil, strings.NewReader(`{"name":"Ada"}`))

	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(result)
	aSpy.HadErrorContaining(t, "actual JSON reader is nil")
}

func Test_EqualReader_failure_expected_read_error(t *testing.T) {
	result := jsonassert.EqualReader(strings.NewReader(`{"name":"Ada"}`), errorReader{})

	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(result)
	aSpy.HadErrorContaining(t, "expected JSON read error: read failed")
}

func Test_EqualLines_success_multiline(t *testing.T) {
	assert := hammy.New(t)

	assert.Is(jsonassert.EqualLines(
		`{"name":"Ada","age":37}`+"\n"+`{"tags":["go","json"]}`,
		`{"age":37.0,"name":"Ada"}`+"\n"+`{"tags":["go","json"]}`,
	))
}

func Test_EqualLines_success_trailing_newline(t *testing.T) {
	assert := hammy.New(t)

	assert.Is(jsonassert.EqualLines(
		"{\"id\":1}\n{\"id\":2}\n",
		"{\"id\":1.0}\n{\"id\":2.0}\n",
	))
}

func Test_EqualLines_success_crlf(t *testing.T) {
	assert := hammy.New(t)

	assert.Is(jsonassert.EqualLines(
		"{\"id\":1}\r\n{\"id\":2}\r\n",
		"{\"id\":1.0}\r\n{\"id\":2.0}\r\n",
	))
}

func Test_EqualLines_success_empty_inputs(t *testing.T) {
	assert := hammy.New(t)

	assert.Is(jsonassert.EqualLines("", ""))
}

func Test_EqualLinesWithOptions_success_ignore_paths_per_line(t *testing.T) {
	assert := hammy.New(t)

	assert.Is(jsonassert.EqualLinesWithOptions(
		`{"status":"ok","meta":{"request_id":"abc"}}`+"\n"+`{"status":"ok","meta":{"request_id":"def"}}`,
		`{"status":"ok","meta":{"request_id":"uvw"}}`+"\n"+`{"status":"ok","meta":{"request_id":"xyz"}}`,
		jsonassert.IgnorePaths("meta.request_id"),
	))
}

func Test_EqualLinesWithOptions_success_unordered_arrays_per_line(t *testing.T) {
	assert := hammy.New(t)

	assert.Is(jsonassert.EqualLinesWithOptions(
		`{"tags":["go","test"]}`+"\n"+`{"tags":["json","assert"]}`,
		`{"tags":["test","go"]}`+"\n"+`{"tags":["assert","json"]}`,
		jsonassert.UnorderedArraysAt("tags"),
	))
}

func Test_EqualLinesBytes_success(t *testing.T) {
	assert := hammy.New(t)

	assert.Is(jsonassert.EqualLinesBytes(
		[]byte(`{"id":1}`+"\n"+`{"id":2}`),
		[]byte(`{"id":1.0}`+"\n"+`{"id":2.0}`),
	))
}

func Test_EqualLinesBytesWithOptions_success_ignore_paths_per_line(t *testing.T) {
	assert := hammy.New(t)

	assert.Is(jsonassert.EqualLinesBytesWithOptions(
		[]byte(`{"status":"ok","meta":{"request_id":"abc"}}`+"\n"+`{"status":"ok","meta":{"request_id":"def"}}`),
		[]byte(`{"status":"ok","meta":{"request_id":"uvw"}}`+"\n"+`{"status":"ok","meta":{"request_id":"xyz"}}`),
		jsonassert.IgnorePaths("meta.request_id"),
	))
}

func Test_EqualLines_failure_reports_line_index(t *testing.T) {
	result := jsonassert.EqualLines(
		`{"id":1,"name":"Ada"}`+"\n"+`{"id":2,"name":"Grace"}`,
		`{"id":1,"name":"Ada"}`+"\n"+`{"id":2,"name":"Katherine"}`,
	)

	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(result)
	aSpy.HadErrorContaining(t, "JSONL line <1> mismatch (-want +got):")
	aSpy.HadErrorContaining(t, "Grace")
	aSpy.HadErrorContaining(t, "Katherine")
}

func Test_EqualLinesBytes_failure_reports_line_index(t *testing.T) {
	result := jsonassert.EqualLinesBytes(
		[]byte(`{"id":1}`+"\n"+`{"id":2}`),
		[]byte(`{"id":1}`+"\n"+`{"id":3}`),
	)

	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(result)
	aSpy.HadErrorContaining(t, "JSONL line <1> mismatch (-want +got):")
}

func Test_EqualLines_failure_invalid_actual_json_reports_line_index(t *testing.T) {
	result := jsonassert.EqualLines(
		`{"id":1}`+"\n"+`{"id":`,
		`{"id":1}`+"\n"+`{"id":2}`,
	)

	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(result)
	aSpy.HadErrorContaining(t, "actual JSONL line <1> invalid:")
}

func Test_EqualLines_failure_invalid_expected_json_reports_line_index(t *testing.T) {
	result := jsonassert.EqualLines(
		`{"id":1}`+"\n"+`{"id":2}`,
		`{"id":1}`+"\n"+`{"id":`,
	)

	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(result)
	aSpy.HadErrorContaining(t, "expected JSONL line <1> invalid:")
}

func Test_EqualLines_failure_blank_middle_line_invalid_json(t *testing.T) {
	result := jsonassert.EqualLines(
		`{"id":1}`+"\n\n"+`{"id":2}`,
		`{"id":1}`+"\n\n"+`{"id":2}`,
	)

	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(result)
	aSpy.HadErrorContaining(t, "actual JSONL line <1> invalid:")
}

func Test_EqualLines_failure_line_count_mismatch_reports_index(t *testing.T) {
	result := jsonassert.EqualLines(
		`{"id":1}`,
		`{"id":1}`+"\n"+`{"id":2}`,
	)

	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(result)
	aSpy.HadErrorContaining(t, "got JSONL line count <1>, wanted <2>; first differing line index <1>")
}

func Test_EqualLinesWithOptions_failure_invalid_option_reports_line_index(t *testing.T) {
	result := jsonassert.EqualLinesWithOptions(
		`{"status":"ok"}`,
		`{"status":"ok"}`,
		jsonassert.IgnorePaths("meta."),
	)

	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(result)
	aSpy.HadErrorContaining(t, "JSONL line <0>: invalid JSON path <meta.>: path ends with dot")
}

func Test_LinesContain_success_full_record(t *testing.T) {
	assert := hammy.New(t)

	assert.Is(jsonassert.LinesContain(
		`{"id":1,"name":"Ada"}`+"\n"+`{"id":2,"score":1.0}`,
		`{"score":1,"id":2}`,
	))
}

func Test_LinesContain_success_ignore_paths(t *testing.T) {
	assert := hammy.New(t)

	assert.Is(jsonassert.LinesContain(
		`{"status":"ok","meta":{"request_id":"abc"}}`+"\n"+`{"status":"done","meta":{"request_id":"def"}}`,
		`{"status":"done","meta":{"request_id":"xyz"}}`,
		jsonassert.IgnorePaths("meta.request_id"),
	))
}

func Test_LinesContain_failure_no_match(t *testing.T) {
	result := jsonassert.LinesContain(
		`{"id":1}`+"\n"+`{"id":2}`,
		`{"id":3}`,
	)

	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(result)
	aSpy.HadErrorContaining(t, "got no matching JSONL line")
}

func Test_LinesContain_failure_invalid_expected_json(t *testing.T) {
	result := jsonassert.LinesContain(`{"id":1}`, `{"id":`)

	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(result)
	aSpy.HadErrorContaining(t, "expected JSON invalid:")
}

func Test_LinesContain_failure_invalid_actual_line(t *testing.T) {
	result := jsonassert.LinesContain(
		`{"id":1}`+"\n"+`{"id":`,
		`{"id":2}`,
	)

	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(result)
	aSpy.HadErrorContaining(t, "actual JSONL line <1> invalid:")
}

func Test_LinesContain_failure_invalid_option_reports_line_index(t *testing.T) {
	result := jsonassert.LinesContain(
		`{"status":"ok"}`,
		`{"status":"ok"}`,
		jsonassert.IgnorePaths("meta."),
	)

	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(result)
	aSpy.HadErrorContaining(t, "JSONL line <0>: invalid JSON path <meta.>: path ends with dot")
}

func Test_LinesContainSubset_success(t *testing.T) {
	assert := hammy.New(t)

	assert.Is(jsonassert.LinesContainSubset(
		`{"status":"ok","meta":{"request_id":"abc","page":1}}`+"\n"+`{"status":"done"}`,
		`{"meta":{"page":1.0}}`,
	))
}

func Test_LinesContainSubset_success_ignore_paths(t *testing.T) {
	assert := hammy.New(t)

	assert.Is(jsonassert.LinesContainSubset(
		`{"status":"ok","meta":{"request_id":"abc","page":1}}`,
		`{"meta":{"request_id":"xyz","page":1.0}}`,
		jsonassert.IgnorePaths("meta.request_id"),
	))
}

func Test_LinesContainSubset_failure_no_match(t *testing.T) {
	result := jsonassert.LinesContainSubset(
		`{"status":"ok"}`+"\n"+`{"status":"done"}`,
		`{"meta":{"page":1}}`,
	)

	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(result)
	aSpy.HadErrorContaining(t, "got no JSONL line containing expected subset")
}

func Test_Valid_success(t *testing.T) {
	assert := hammy.New(t)

	assert.Is(jsonassert.Valid(`{"name":"Ada"}`))
}

func Test_Valid_failure(t *testing.T) {
	result := jsonassert.Valid(`{"name":`)

	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(result)
	aSpy.HadErrorContaining(t, "JSON invalid:")
}

func Test_ValidBytes_success(t *testing.T) {
	assert := hammy.New(t)

	assert.Is(jsonassert.ValidBytes([]byte(`{"name":"Ada"}`)))
}

func Test_ValidReader_success(t *testing.T) {
	assert := hammy.New(t)

	assert.Is(jsonassert.ValidReader(strings.NewReader(`{"name":"Ada"}`)))
}

func Test_ValidReader_failure_read_error(t *testing.T) {
	result := jsonassert.ValidReader(errorReader{})

	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(result)
	aSpy.HadErrorContaining(t, "actual JSON read error: read failed")
}

func Test_Contains_success_object_subset(t *testing.T) {
	assert := hammy.New(t)

	assert.Is(jsonassert.Contains(
		`{"status":"ok","meta":{"request_id":"abc","page":1}}`,
		`{"status":"ok","meta":{"page":1.0}}`,
	))
}

func Test_Contains_failure_missing_field(t *testing.T) {
	result := jsonassert.Contains(`{"status":"ok"}`, `{"meta":{"page":1}}`)

	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(result)
	aSpy.HadErrorContaining(t, "JSON path <$.meta> missing")
}

func Test_Contains_failure_mismatched_value(t *testing.T) {
	result := jsonassert.Contains(`{"status":"ok"}`, `{"status":"failed"}`)

	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(result)
	aSpy.HadErrorContaining(t, "JSON path <$.status> mismatch")
}

func Test_Contains_failure_array_order_sensitive(t *testing.T) {
	result := jsonassert.Contains(`{"values":[1,2,3]}`, `{"values":[3,2,1]}`)

	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(result)
	aSpy.HadErrorContaining(t, "JSON path <$.values> mismatch")
}

func Test_ContainsBytes_success(t *testing.T) {
	assert := hammy.New(t)

	assert.Is(jsonassert.ContainsBytes([]byte(`{"status":"ok","extra":true}`), []byte(`{"status":"ok"}`)))
}

func Test_PathExists_success(t *testing.T) {
	assert := hammy.New(t)

	assert.Is(jsonassert.PathExists(`{"user":{"name":"Ada"}}`, "user.name"))
}

func Test_PathExists_success_array_index(t *testing.T) {
	assert := hammy.New(t)

	assert.Is(jsonassert.PathExists(`{"items":[{"id":1}]}`, "items[0].id"))
}

func Test_PathExists_failure_missing(t *testing.T) {
	result := jsonassert.PathExists(`{"user":{"name":"Ada"}}`, "user.email")

	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(result)
	aSpy.HadErrorContaining(t, "JSON path <user.email> missing")
}

func Test_PathExists_failure_invalid_path(t *testing.T) {
	result := jsonassert.PathExists(`{"user":{"name":"Ada"}}`, "user.")

	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(result)
	aSpy.HadErrorContaining(t, "invalid JSON path <user.>: path ends with dot")
}

func Test_PathMissing_success(t *testing.T) {
	assert := hammy.New(t)

	assert.Is(jsonassert.PathMissing(`{"user":{"name":"Ada"}}`, "user.email"))
}

func Test_PathMissing_failure_exists(t *testing.T) {
	result := jsonassert.PathMissing(`{"user":{"name":"Ada"}}`, "user.name")

	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(result)
	aSpy.HadErrorContaining(t, "JSON path <user.name> exists, wanted missing")
}

func Test_PathEqual_success(t *testing.T) {
	assert := hammy.New(t)

	assert.Is(jsonassert.PathEqual(`{"user":{"name":"Ada","age":37}}`, "user.age", `37.0`))
}

func Test_PathEqual_failure_mismatch(t *testing.T) {
	result := jsonassert.PathEqual(`{"user":{"name":"Ada"}}`, "user.name", `"Grace"`)

	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(result)
	aSpy.HadErrorContaining(t, "JSON path <user.name> mismatch (-want +got):")
	aSpy.HadErrorContaining(t, "Grace")
	aSpy.HadErrorContaining(t, "Ada")
}

func Test_PathEqualBytes_success(t *testing.T) {
	assert := hammy.New(t)

	assert.Is(jsonassert.PathEqualBytes([]byte(`{"user":{"name":"Ada"}}`), "user.name", []byte(`"Ada"`)))
}

func Test_ArrayContains_success(t *testing.T) {
	assert := hammy.New(t)

	assert.Is(jsonassert.ArrayContains(`{"items":[{"id":1},{"id":2}]}`, "items", `{"id":2.0}`))
}

func Test_ArrayContains_failure_missing_element(t *testing.T) {
	result := jsonassert.ArrayContains(`{"items":[{"id":1}]}`, "items", `{"id":2}`)

	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(result)
	aSpy.HadErrorContaining(t, "got no matching element at JSON path <items>")
}

func Test_ArrayContains_failure_non_array_path(t *testing.T) {
	result := jsonassert.ArrayContains(`{"items":{"id":1}}`, "items", `{"id":1}`)

	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(result)
	aSpy.HadErrorContaining(t, "wanted array")
}

func Test_ArrayContainsBytes_success(t *testing.T) {
	assert := hammy.New(t)

	assert.Is(jsonassert.ArrayContainsBytes([]byte(`{"items":[1,2]}`), "items", []byte(`2.0`)))
}

func Test_EqualWithOptions_success_ignore_paths(t *testing.T) {
	assert := hammy.New(t)

	assert.Is(jsonassert.EqualWithOptions(
		`{"status":"ok","meta":{"request_id":"abc"}}`,
		`{"status":"ok","meta":{"request_id":"xyz"}}`,
		jsonassert.IgnorePaths("meta.request_id"),
	))
}

func Test_EqualWithOptions_success_ignore_array_item_path(t *testing.T) {
	assert := hammy.New(t)

	assert.Is(jsonassert.EqualWithOptions(
		`{"items":[{"id":1,"etag":"a"}]}`,
		`{"items":[{"id":1,"etag":"b"}]}`,
		jsonassert.IgnorePaths("items[0].etag"),
	))
}

func Test_EqualWithOptions_failure_invalid_ignore_path(t *testing.T) {
	result := jsonassert.EqualWithOptions(`{"status":"ok"}`, `{"status":"ok"}`, jsonassert.IgnorePaths("meta."))

	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(result)
	aSpy.HadErrorContaining(t, "invalid JSON path <meta.>: path ends with dot")
}

func Test_EqualWithOptions_success_unordered_array(t *testing.T) {
	assert := hammy.New(t)

	assert.Is(jsonassert.EqualWithOptions(
		`{"tags":["go","test","json"]}`,
		`{"tags":["json","go","test"]}`,
		jsonassert.UnorderedArraysAt("tags"),
	))
}

func Test_EqualWithOptions_success_unordered_array_of_objects(t *testing.T) {
	assert := hammy.New(t)

	assert.Is(jsonassert.EqualWithOptions(
		`{"items":[{"id":2},{"id":1}]}`,
		`{"items":[{"id":1.0},{"id":2.0}]}`,
		jsonassert.UnorderedArraysAt("items"),
	))
}

func Test_EqualWithOptions_failure_unordered_array_non_array_path(t *testing.T) {
	result := jsonassert.EqualWithOptions(
		`{"items":{"id":1}}`,
		`{"items":{"id":1}}`,
		jsonassert.UnorderedArraysAt("items"),
	)

	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(result)
	aSpy.HadErrorContaining(t, "got JSON path <items> type <map[string]interface {}>, wanted array")
}

func Test_EqualWithOptions_failure_without_ignore_paths(t *testing.T) {
	result := jsonassert.EqualWithOptions(
		`{"status":"ok","meta":{"request_id":"abc"}}`,
		`{"status":"ok","meta":{"request_id":"xyz"}}`,
	)

	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(result)
	aSpy.HadErrorContaining(t, "JSON mismatch (-want +got):")
}

func Test_EqualBytesWithOptions_success(t *testing.T) {
	assert := hammy.New(t)

	assert.Is(jsonassert.EqualBytesWithOptions(
		[]byte(`{"tags":["go","test"]}`),
		[]byte(`{"tags":["test","go"]}`),
		jsonassert.UnorderedArraysAt("tags"),
	))
}

type errorReader struct{}

func (errorReader) Read([]byte) (int, error) {
	return 0, errors.New("read failed")
}

var _ io.Reader = errorReader{}
