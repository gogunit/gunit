package yamlassert_test

import (
	"errors"
	"io"
	"strings"
	"testing"

	"github.com/gogunit/gunit/hammy"
	"github.com/gogunit/gunit/hammy/yamlassert"
)

func Test_Equal_success_reordered_mapping_keys(t *testing.T) {
	assert := hammy.New(t)

	assert.Is(yamlassert.Equal("name: Ada\nage: 37\n", "age: 37.0\nname: Ada\n"))
}

func Test_Equal_success_anchor_alias(t *testing.T) {
	assert := hammy.New(t)

	assert.Is(yamlassert.Equal("name: &name Ada\nuser:\n  name: *name\n", "name: Ada\nuser:\n  name: Ada\n"))
}

func Test_Equal_failure_array_order_mismatch(t *testing.T) {
	result := yamlassert.Equal("values: [1, 2, 3]\n", "values: [3, 2, 1]\n")

	requireFailure(t, result, "YAML mismatch (-want +got):")
}

func Test_Equal_failure_multiple_documents(t *testing.T) {
	result := yamlassert.Equal("---\nname: Ada\n---\nname: Grace\n", "name: Ada\n")

	requireFailure(t, result, "actual YAML invalid: multiple YAML documents")
}

func Test_EqualBytes_success(t *testing.T) {
	assert := hammy.New(t)

	assert.Is(yamlassert.EqualBytes([]byte("one: 1\n"), []byte("one: 1.0\n")))
}

func Test_EqualReader_success(t *testing.T) {
	assert := hammy.New(t)

	assert.Is(yamlassert.EqualReader(strings.NewReader("one: 1\n"), strings.NewReader("one: 1.0\n")))
}

func Test_Valid_success(t *testing.T) {
	assert := hammy.New(t)

	assert.Is(yamlassert.Valid("name: Ada\n"))
}

func Test_Valid_failure_invalid_yaml(t *testing.T) {
	result := yamlassert.Valid("name: [Ada\n")

	requireFailure(t, result, "YAML invalid:")
}

func Test_Valid_failure_duplicate_key(t *testing.T) {
	result := yamlassert.Valid("name: Ada\nname: Grace\n")

	requireFailure(t, result, "duplicate YAML key <name>")
}

func Test_ValidBytes_success(t *testing.T) {
	assert := hammy.New(t)

	assert.Is(yamlassert.ValidBytes([]byte("name: Ada\n")))
}

func Test_ValidReader_success(t *testing.T) {
	assert := hammy.New(t)

	assert.Is(yamlassert.ValidReader(strings.NewReader("name: Ada\n")))
}

func Test_ValidReader_failure_read_error(t *testing.T) {
	result := yamlassert.ValidReader(errorReader{})

	requireFailure(t, result, "actual YAML read error: read failed")
}

func Test_Contains_success_mapping_subset(t *testing.T) {
	assert := hammy.New(t)

	assert.Is(yamlassert.Contains("status: ok\nmeta:\n  page: 1\n  request_id: abc\n", "meta:\n  page: 1.0\n"))
}

func Test_Contains_failure_missing_field(t *testing.T) {
	result := yamlassert.Contains("status: ok\n", "meta:\n  page: 1\n")

	requireFailure(t, result, "YAML path <$.meta> missing")
}

func Test_ContainsBytes_success(t *testing.T) {
	assert := hammy.New(t)

	assert.Is(yamlassert.ContainsBytes([]byte("status: ok\nextra: true\n"), []byte("status: ok\n")))
}

func Test_PathExists_success(t *testing.T) {
	assert := hammy.New(t)

	assert.Is(yamlassert.PathExists("user:\n  name: Ada\n", "user.name"))
}

func Test_PathExists_success_array_index(t *testing.T) {
	assert := hammy.New(t)

	assert.Is(yamlassert.PathExists("items:\n  - id: 1\n", "items[0].id"))
}

func Test_PathExists_failure_missing(t *testing.T) {
	result := yamlassert.PathExists("user:\n  name: Ada\n", "user.email")

	requireFailure(t, result, "YAML path <user.email> missing")
}

func Test_PathMissing_success(t *testing.T) {
	assert := hammy.New(t)

	assert.Is(yamlassert.PathMissing("user:\n  name: Ada\n", "user.email"))
}

func Test_PathMissing_failure_exists(t *testing.T) {
	result := yamlassert.PathMissing("user:\n  name: Ada\n", "user.name")

	requireFailure(t, result, "YAML path <user.name> exists, wanted missing")
}

func Test_PathEqual_success(t *testing.T) {
	assert := hammy.New(t)

	assert.Is(yamlassert.PathEqual("user:\n  age: 37\n", "user.age", "37.0\n"))
}

func Test_PathEqual_failure_mismatch(t *testing.T) {
	result := yamlassert.PathEqual("user:\n  name: Ada\n", "user.name", "Grace\n")

	requireFailure(t, result, "YAML path <user.name> mismatch (-want +got):")
	requireFailure(t, result, "Grace")
	requireFailure(t, result, "Ada")
}

func Test_PathEqualBytes_success(t *testing.T) {
	assert := hammy.New(t)

	assert.Is(yamlassert.PathEqualBytes([]byte("user:\n  name: Ada\n"), "user.name", []byte("Ada\n")))
}

func Test_ArrayContains_success(t *testing.T) {
	assert := hammy.New(t)

	assert.Is(yamlassert.ArrayContains("items:\n  - id: 1\n  - id: 2\n", "items", "id: 2.0\n"))
}

func Test_ArrayContains_failure_missing_element(t *testing.T) {
	result := yamlassert.ArrayContains("items:\n  - id: 1\n", "items", "id: 2\n")

	requireFailure(t, result, "got no matching element at YAML path <items>")
}

func Test_ArrayContainsBytes_success(t *testing.T) {
	assert := hammy.New(t)

	assert.Is(yamlassert.ArrayContainsBytes([]byte("items: [1, 2]\n"), "items", []byte("2.0\n")))
}

func Test_EqualWithOptions_success_ignore_paths(t *testing.T) {
	assert := hammy.New(t)

	assert.Is(yamlassert.EqualWithOptions(
		"status: ok\nmeta:\n  request_id: abc\n",
		"status: ok\nmeta:\n  request_id: xyz\n",
		yamlassert.IgnorePaths("meta.request_id"),
	))
}

func Test_EqualWithOptions_success_unordered_array(t *testing.T) {
	assert := hammy.New(t)

	assert.Is(yamlassert.EqualWithOptions(
		"tags: [go, test, yaml]\n",
		"tags: [yaml, go, test]\n",
		yamlassert.UnorderedArraysAt("tags"),
	))
}

func Test_EqualBytesWithOptions_success(t *testing.T) {
	assert := hammy.New(t)

	assert.Is(yamlassert.EqualBytesWithOptions(
		[]byte("tags: [go, test]\n"),
		[]byte("tags: [test, go]\n"),
		yamlassert.UnorderedArraysAt("tags"),
	))
}

func Test_DocumentCount_success(t *testing.T) {
	assert := hammy.New(t)

	assert.Is(yamlassert.DocumentCount("---\nname: Ada\n---\nname: Grace\n", 2))
}

func Test_DocumentCount_failure(t *testing.T) {
	result := yamlassert.DocumentCount("---\nname: Ada\n---\nname: Grace\n", 1)

	requireFailure(t, result, "got YAML document count <2>, wanted <1>")
}

func Test_DocumentCountBytes_success(t *testing.T) {
	assert := hammy.New(t)

	assert.Is(yamlassert.DocumentCountBytes([]byte("---\nname: Ada\n---\nname: Grace\n"), 2))
}

func Test_DocumentEqual_success(t *testing.T) {
	assert := hammy.New(t)

	assert.Is(yamlassert.DocumentEqual("---\nname: Ada\n---\nname: Grace\n", 1, "name: Grace\n"))
}

func Test_DocumentEqual_success_with_options(t *testing.T) {
	assert := hammy.New(t)

	assert.Is(yamlassert.DocumentEqual(
		"---\nstatus: ok\nmeta:\n  request_id: abc\n",
		0,
		"status: ok\nmeta:\n  request_id: xyz\n",
		yamlassert.IgnorePaths("meta.request_id"),
	))
}

func Test_DocumentEqual_failure_index_out_of_range(t *testing.T) {
	result := yamlassert.DocumentEqual("---\nname: Ada\n", 1, "name: Ada\n")

	requireFailure(t, result, "got YAML document index <1> out of range for count <1>")
}

func Test_DocumentEqualBytes_success(t *testing.T) {
	assert := hammy.New(t)

	assert.Is(yamlassert.DocumentEqualBytes([]byte("---\nname: Ada\n"), 0, []byte("name: Ada\n")))
}

func Test_DocumentContains_success(t *testing.T) {
	assert := hammy.New(t)

	assert.Is(yamlassert.DocumentContains("---\nstatus: ok\nmeta:\n  page: 1\n", 0, "meta:\n  page: 1.0\n"))
}

func Test_DocumentContainsBytes_success(t *testing.T) {
	assert := hammy.New(t)

	assert.Is(yamlassert.DocumentContainsBytes([]byte("---\nstatus: ok\n"), 0, []byte("status: ok\n")))
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
