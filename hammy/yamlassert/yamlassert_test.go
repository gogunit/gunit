package yamlassert_test

import (
	"errors"
	"io"
	"strings"
	"testing"

	"github.com/gogunit/gunit/eye"
	"github.com/gogunit/gunit/hammy"
	"github.com/gogunit/gunit/hammy/yamlassert"
)

func Test_Equal_success_reordered_mapping_keys(t *testing.T) {
	assert := hammy.New(t)

	assert.Is(yamlassert.String("name: Ada\nage: 37\n").EqualTo("age: 37.0\nname: Ada\n"))
}

func Test_Equal_success_anchor_alias(t *testing.T) {
	assert := hammy.New(t)

	assert.Is(yamlassert.String("name: &name Ada\nuser:\n  name: *name\n").EqualTo("name: Ada\nuser:\n  name: Ada\n"))
}

func Test_Equal_failure_array_order_mismatch(t *testing.T) {
	result := yamlassert.String("values: [1, 2, 3]\n").EqualTo("values: [3, 2, 1]\n")

	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(result)
	aSpy.HadErrorContaining(t, "YAML mismatch (-want +got):")
}

func Test_Equal_failure_multiple_documents(t *testing.T) {
	result := yamlassert.String("---\nname: Ada\n---\nname: Grace\n").EqualTo("name: Ada\n")

	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(result)
	aSpy.HadErrorContaining(t, "actual YAML invalid: multiple YAML documents")
}

func Test_EqualBytes_success(t *testing.T) {
	assert := hammy.New(t)

	assert.Is(yamlassert.Bytes([]byte("one: 1\n")).EqualTo([]byte("one: 1.0\n")))
}

func Test_EqualReader_success(t *testing.T) {
	assert := hammy.New(t)

	assert.Is(yamlassert.Reader(strings.NewReader("one: 1\n")).EqualTo(strings.NewReader("one: 1.0\n")))
}

func Test_Valid_success(t *testing.T) {
	assert := hammy.New(t)

	assert.Is(yamlassert.String("name: Ada\n").IsValid())
}

func Test_Valid_failure_invalid_yaml(t *testing.T) {
	result := yamlassert.String("name: [Ada\n").IsValid()

	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(result)
	aSpy.HadErrorContaining(t, "YAML invalid:")
}

func Test_Valid_failure_duplicate_key(t *testing.T) {
	result := yamlassert.String("name: Ada\nname: Grace\n").IsValid()

	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(result)
	aSpy.HadErrorContaining(t, "duplicate YAML key <name>")
}

func Test_ValidBytes_success(t *testing.T) {
	assert := hammy.New(t)

	assert.Is(yamlassert.Bytes([]byte("name: Ada\n")).IsValid())
}

func Test_ValidReader_success(t *testing.T) {
	assert := hammy.New(t)

	assert.Is(yamlassert.Reader(strings.NewReader("name: Ada\n")).IsValid())
}

func Test_ValidReader_failure_read_error(t *testing.T) {
	result := yamlassert.Reader(errorReader{}).IsValid()

	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(result)
	aSpy.HadErrorContaining(t, "actual YAML read error: read failed")
}

func Test_Contains_success_mapping_subset(t *testing.T) {
	assert := hammy.New(t)

	assert.Is(yamlassert.String("status: ok\nmeta:\n  page: 1\n  request_id: abc\n").Contains("meta:\n  page: 1.0\n"))
}

func Test_Contains_failure_missing_field(t *testing.T) {
	result := yamlassert.String("status: ok\n").Contains("meta:\n  page: 1\n")

	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(result)
	aSpy.HadErrorContaining(t, "YAML path <$.meta> missing")
}

func Test_ContainsBytes_success(t *testing.T) {
	assert := hammy.New(t)

	assert.Is(yamlassert.Bytes([]byte("status: ok\nextra: true\n")).Contains([]byte("status: ok\n")))
}

func Test_PathExists_success(t *testing.T) {
	assert := hammy.New(t)

	assert.Is(yamlassert.String("user:\n  name: Ada\n").PathExists("user.name"))
}

func Test_PathExists_success_array_index(t *testing.T) {
	assert := hammy.New(t)

	assert.Is(yamlassert.String("items:\n  - id: 1\n").PathExists("items[0].id"))
}

func Test_PathExists_failure_missing(t *testing.T) {
	result := yamlassert.String("user:\n  name: Ada\n").PathExists("user.email")

	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(result)
	aSpy.HadErrorContaining(t, "YAML path <user.email> missing")
}

func Test_PathMissing_success(t *testing.T) {
	assert := hammy.New(t)

	assert.Is(yamlassert.String("user:\n  name: Ada\n").PathMissing("user.email"))
}

func Test_PathMissing_failure_exists(t *testing.T) {
	result := yamlassert.String("user:\n  name: Ada\n").PathMissing("user.name")

	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(result)
	aSpy.HadErrorContaining(t, "YAML path <user.name> exists, wanted missing")
}

func Test_PathEqual_success(t *testing.T) {
	assert := hammy.New(t)

	assert.Is(yamlassert.String("user:\n  age: 37\n").PathEqual("user.age", "37.0\n"))
}

func Test_PathEqual_failure_mismatch(t *testing.T) {
	result := yamlassert.String("user:\n  name: Ada\n").PathEqual("user.name", "Grace\n")

	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(result)
	aSpy.HadErrorContaining(t, "YAML path <user.name> mismatch (-want +got):")
	aSpy.HadErrorContaining(t, "Grace")
	aSpy.HadErrorContaining(t, "Ada")
}

func Test_PathEqualBytes_success(t *testing.T) {
	assert := hammy.New(t)

	assert.Is(yamlassert.Bytes([]byte("user:\n  name: Ada\n")).PathEqual("user.name", []byte("Ada\n")))
}

func Test_ArrayContains_success(t *testing.T) {
	assert := hammy.New(t)

	assert.Is(yamlassert.String("items:\n  - id: 1\n  - id: 2\n").ArrayContains("items", "id: 2.0\n"))
}

func Test_ArrayContains_failure_missing_element(t *testing.T) {
	result := yamlassert.String("items:\n  - id: 1\n").ArrayContains("items", "id: 2\n")

	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(result)
	aSpy.HadErrorContaining(t, "got no matching element at YAML path <items>")
}

func Test_ArrayContainsBytes_success(t *testing.T) {
	assert := hammy.New(t)

	assert.Is(yamlassert.Bytes([]byte("items: [1, 2]\n")).ArrayContains("items", []byte("2.0\n")))
}

func Test_EqualWithOptions_success_ignore_paths(t *testing.T) {
	assert := hammy.New(t)

	assert.Is(yamlassert.String("status: ok\nmeta:\n  request_id: abc\n").EqualTo(
		"status: ok\nmeta:\n  request_id: xyz\n",
		yamlassert.IgnorePaths("meta.request_id"),
	))
}

func Test_EqualWithOptions_success_unordered_array(t *testing.T) {
	assert := hammy.New(t)

	assert.Is(yamlassert.String("tags: [go, test, yaml]\n").EqualTo(
		"tags: [yaml, go, test]\n",
		yamlassert.UnorderedArraysAt("tags"),
	))
}

func Test_EqualBytesWithOptions_success(t *testing.T) {
	assert := hammy.New(t)

	assert.Is(yamlassert.Bytes([]byte("tags: [go, test]\n")).EqualTo(
		[]byte("tags: [test, go]\n"),
		yamlassert.UnorderedArraysAt("tags"),
	))
}

func Test_DocumentCount_success(t *testing.T) {
	assert := hammy.New(t)

	assert.Is(yamlassert.String("---\nname: Ada\n---\nname: Grace\n").DocumentCount(2))
}

func Test_DocumentCount_failure(t *testing.T) {
	result := yamlassert.String("---\nname: Ada\n---\nname: Grace\n").DocumentCount(1)

	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(result)
	aSpy.HadErrorContaining(t, "got YAML document count <2>, wanted <1>")
}

func Test_DocumentCountBytes_success(t *testing.T) {
	assert := hammy.New(t)

	assert.Is(yamlassert.Bytes([]byte("---\nname: Ada\n---\nname: Grace\n")).DocumentCount(2))
}

func Test_DocumentEqual_success(t *testing.T) {
	assert := hammy.New(t)

	assert.Is(yamlassert.String("---\nname: Ada\n---\nname: Grace\n").DocumentEqual(1, "name: Grace\n"))
}

func Test_DocumentEqual_success_with_options(t *testing.T) {
	assert := hammy.New(t)

	assert.Is(yamlassert.String("---\nstatus: ok\nmeta:\n  request_id: abc\n").DocumentEqual(
		0,
		"status: ok\nmeta:\n  request_id: xyz\n",
		yamlassert.IgnorePaths("meta.request_id"),
	))
}

func Test_DocumentEqual_failure_index_out_of_range(t *testing.T) {
	result := yamlassert.String("---\nname: Ada\n").DocumentEqual(1, "name: Ada\n")

	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(result)
	aSpy.HadErrorContaining(t, "got YAML document index <1> out of range for count <1>")
}

func Test_DocumentEqualBytes_success(t *testing.T) {
	assert := hammy.New(t)

	assert.Is(yamlassert.Bytes([]byte("---\nname: Ada\n")).DocumentEqual(0, []byte("name: Ada\n")))
}

func Test_DocumentContains_success(t *testing.T) {
	assert := hammy.New(t)

	assert.Is(yamlassert.String("---\nstatus: ok\nmeta:\n  page: 1\n").DocumentContains(0, "meta:\n  page: 1.0\n"))
}

func Test_DocumentContainsBytes_success(t *testing.T) {
	assert := hammy.New(t)

	assert.Is(yamlassert.Bytes([]byte("---\nstatus: ok\n")).DocumentContains(0, []byte("status: ok\n")))
}

type errorReader struct{}

func (errorReader) Read([]byte) (int, error) {
	return 0, errors.New("read failed")
}

var _ io.Reader = errorReader{}
