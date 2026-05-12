package jsonassert_test

import (
	"errors"
	"io"
	"strings"
	"testing"

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

	requireFailure(t, result, "JSON mismatch (-want +got):")
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

	requireFailure(t, result, "actual JSON invalid:")
}

func Test_Equal_failure_invalid_expected_json(t *testing.T) {
	result := jsonassert.Equal(`{"name":"Ada"}`, `{"name":`)

	requireFailure(t, result, "expected JSON invalid:")
}

func Test_Equal_failure_multiple_actual_json_values(t *testing.T) {
	result := jsonassert.Equal(`{"name":"Ada"} {"name":"Grace"}`, `{"name":"Ada"}`)

	requireFailure(t, result, "actual JSON invalid: multiple JSON values")
}

func Test_Equal_failure_includes_diff(t *testing.T) {
	result := jsonassert.Equal(`{"name":"Ada"}`, `{"name":"Grace"}`)

	requireFailure(t, result, "JSON mismatch (-want +got):")
	requireFailure(t, result, `Grace`)
	requireFailure(t, result, `Ada`)
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

	requireFailure(t, result, "actual JSON reader is nil")
}

func Test_EqualReader_failure_expected_read_error(t *testing.T) {
	result := jsonassert.EqualReader(strings.NewReader(`{"name":"Ada"}`), errorReader{})

	requireFailure(t, result, "expected JSON read error: read failed")
}

func Test_Valid_success(t *testing.T) {
	assert := hammy.New(t)

	assert.Is(jsonassert.Valid(`{"name":"Ada"}`))
}

func Test_Valid_failure(t *testing.T) {
	result := jsonassert.Valid(`{"name":`)

	requireFailure(t, result, "JSON invalid:")
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

	requireFailure(t, result, "actual JSON read error: read failed")
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

type errorReader struct{}

func (errorReader) Read([]byte) (int, error) {
	return 0, errors.New("read failed")
}

var _ io.Reader = errorReader{}
