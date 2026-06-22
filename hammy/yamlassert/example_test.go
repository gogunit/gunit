package yamlassert_test

import (
	"fmt"
	"strings"

	"github.com/gogunit/gunit/hammy"
	"github.com/gogunit/gunit/hammy/yamlassert"
)

func ExampleStringAssert_EqualTo() {
	actual := "name: Ada\nage: 37\n"
	expected := "age: 37.0\nname: Ada\n"

	printExample(yamlassert.String(actual).EqualTo(expected))
	// Output:
	// message="YAML mismatch (-want +got):\n"
	// success=true
}

func ExampleStringAssert_EqualToWithOptions() {
	actual := "status: ok\nmeta:\n  request_id: abc\n"
	expected := "status: ok\nmeta:\n  request_id: xyz\n"

	printExample(yamlassert.String(actual).EqualTo(expected, yamlassert.IgnorePaths("meta.request_id")))
	// Output:
	// message="YAML mismatch (-want +got):\n"
	// success=true
}

func ExampleBytesAssert_EqualTo() {
	printExample(yamlassert.Bytes([]byte("one: 1\n")).EqualTo([]byte("one: 1.0\n")))
	// Output:
	// message="YAML mismatch (-want +got):\n"
	// success=true
}

func ExampleBytesAssert_EqualToWithOptions() {
	actual := []byte("tags: [go, test]\n")
	expected := []byte("tags: [test, go]\n")

	printExample(yamlassert.Bytes(actual).EqualTo(expected, yamlassert.UnorderedArraysAt("tags")))
	// Output:
	// message="YAML mismatch (-want +got):\n"
	// success=true
}

func ExampleReaderAssert_EqualTo() {
	actual := strings.NewReader("one: 1\n")
	expected := strings.NewReader("one: 1.0\n")

	printExample(yamlassert.Reader(actual).EqualTo(expected))
	// Output:
	// message="YAML mismatch (-want +got):\n"
	// success=true
}

func ExampleStringAssert_IsValid() {
	printExample(yamlassert.String("name: Ada\n").IsValid())
	// Output:
	// message="got valid YAML"
	// success=true
}

func ExampleBytesAssert_IsValid() {
	printExample(yamlassert.Bytes([]byte("name: Ada\n")).IsValid())
	// Output:
	// message="got valid YAML"
	// success=true
}

func ExampleReaderAssert_IsValid() {
	printExample(yamlassert.Reader(strings.NewReader("name: Ada\n")).IsValid())
	// Output:
	// message="got valid YAML"
	// success=true
}

func ExampleStringAssert_Contains() {
	actual := "status: ok\nmeta:\n  page: 1\n  request_id: abc\n"
	expected := "meta:\n  page: 1.0\n"

	printExample(yamlassert.String(actual).Contains(expected))
	// Output:
	// message="YAML contained expected subset"
	// success=true
}

func ExampleBytesAssert_Contains() {
	actual := []byte("status: ok\nextra: true\n")
	expected := []byte("status: ok\n")

	printExample(yamlassert.Bytes(actual).Contains(expected))
	// Output:
	// message="YAML contained expected subset"
	// success=true
}

func ExampleStringAssert_PathExists() {
	printExample(yamlassert.String("user:\n  name: Ada\n").PathExists("user.name"))
	// Output:
	// message="YAML path <user.name> exists"
	// success=true
}

func ExampleStringAssert_PathMissing() {
	printExample(yamlassert.String("user:\n  name: Ada\n").PathMissing("user.email"))
	// Output:
	// message="YAML path <user.email> missing"
	// success=true
}

func ExampleStringAssert_PathEqual() {
	printExample(yamlassert.String("user:\n  age: 37\n").PathEqual("user.age", "37.0\n"))
	// Output:
	// message="YAML path <user.age> mismatch (-want +got):\n"
	// success=true
}

func ExampleBytesAssert_PathEqual() {
	actual := []byte("user:\n  name: Ada\n")
	expected := []byte("Ada\n")

	printExample(yamlassert.Bytes(actual).PathEqual("user.name", expected))
	// Output:
	// message="YAML path <user.name> mismatch (-want +got):\n"
	// success=true
}

func ExampleStringAssert_ArrayContains() {
	actual := "items:\n  - id: 1\n  - id: 2\n"

	printExample(yamlassert.String(actual).ArrayContains("items", "id: 2.0\n"))
	// Output:
	// message="found matching element at YAML path <items> index <1>"
	// success=true
}

func ExampleBytesAssert_ArrayContains() {
	actual := []byte("items: [1, 2]\n")
	expected := []byte("2.0\n")

	printExample(yamlassert.Bytes(actual).ArrayContains("items", expected))
	// Output:
	// message="found matching element at YAML path <items> index <1>"
	// success=true
}

func ExampleIgnorePaths() {
	actual := "status: ok\nmeta:\n  request_id: abc\n"
	expected := "status: ok\nmeta:\n  request_id: xyz\n"

	printExample(yamlassert.String(actual).EqualTo(expected, yamlassert.IgnorePaths("meta.request_id")))
	// Output:
	// message="YAML mismatch (-want +got):\n"
	// success=true
}

func ExampleUnorderedArraysAt() {
	actual := "tags: [go, test, yaml]\n"
	expected := "tags: [yaml, go, test]\n"

	printExample(yamlassert.String(actual).EqualTo(expected, yamlassert.UnorderedArraysAt("tags")))
	// Output:
	// message="YAML mismatch (-want +got):\n"
	// success=true
}

func ExampleStringAssert_DocumentCount() {
	printExample(yamlassert.String("---\nname: Ada\n---\nname: Grace\n").DocumentCount(2))
	// Output:
	// message="got YAML document count <2>, wanted <2>"
	// success=true
}

func ExampleBytesAssert_DocumentCount() {
	printExample(yamlassert.Bytes([]byte("---\nname: Ada\n---\nname: Grace\n")).DocumentCount(2))
	// Output:
	// message="got YAML document count <2>, wanted <2>"
	// success=true
}

func ExampleStringAssert_DocumentEqual() {
	printExample(yamlassert.String("---\nname: Ada\n---\nname: Grace\n").DocumentEqual(1, "name: Grace\n"))
	// Output:
	// message="YAML document <1> mismatch (-want +got):\n"
	// success=true
}

func ExampleBytesAssert_DocumentEqual() {
	printExample(yamlassert.Bytes([]byte("---\nname: Ada\n")).DocumentEqual(0, []byte("name: Ada\n")))
	// Output:
	// message="YAML document <0> mismatch (-want +got):\n"
	// success=true
}

func ExampleStringAssert_DocumentContains() {
	actual := "---\nstatus: ok\nmeta:\n  page: 1\n"
	expected := "meta:\n  page: 1.0\n"

	printExample(yamlassert.String(actual).DocumentContains(0, expected))
	// Output:
	// message="YAML document <0> contained expected subset"
	// success=true
}

func ExampleBytesAssert_DocumentContains() {
	printExample(yamlassert.Bytes([]byte("---\nstatus: ok\n")).DocumentContains(0, []byte("status: ok\n")))
	// Output:
	// message="YAML document <0> contained expected subset"
	// success=true
}

func printExample(result hammy.AssertionMessage) {
	fmt.Printf("message=%q\nsuccess=%t\n", result.Message, result.IsSuccessful)
}
