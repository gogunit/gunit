package jsonassert_test

import (
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

func requireFailure(t *testing.T, result hammy.AssertionMessage, contains string) {
	t.Helper()
	if result.IsSuccessful {
		t.Fatalf("got success, wanted failure containing %q", contains)
	}
	if !strings.Contains(result.Message, contains) {
		t.Fatalf("got message %q, wanted containing %q", result.Message, contains)
	}
}
