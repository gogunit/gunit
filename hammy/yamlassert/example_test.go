package yamlassert_test

import (
	"fmt"
	"strings"

	"github.com/gogunit/gunit/hammy"
	"github.com/gogunit/gunit/hammy/yamlassert"
)

func ExampleEqual() {
	actual := "name: Ada\nage: 37\n"
	expected := "age: 37.0\nname: Ada\n"

	printExample(yamlassert.Equal(actual, expected))
	// Output:
	// message="YAML mismatch (-want +got):\n"
	// success=true
}

func ExampleEqualWithOptions() {
	actual := "status: ok\nmeta:\n  request_id: abc\n"
	expected := "status: ok\nmeta:\n  request_id: xyz\n"

	printExample(yamlassert.EqualWithOptions(actual, expected, yamlassert.IgnorePaths("meta.request_id")))
	// Output:
	// message="YAML mismatch (-want +got):\n"
	// success=true
}

func ExampleEqualBytes() {
	printExample(yamlassert.EqualBytes([]byte("one: 1\n"), []byte("one: 1.0\n")))
	// Output:
	// message="YAML mismatch (-want +got):\n"
	// success=true
}

func ExampleEqualBytesWithOptions() {
	actual := []byte("tags: [go, test]\n")
	expected := []byte("tags: [test, go]\n")

	printExample(yamlassert.EqualBytesWithOptions(actual, expected, yamlassert.UnorderedArraysAt("tags")))
	// Output:
	// message="YAML mismatch (-want +got):\n"
	// success=true
}

func ExampleEqualReader() {
	actual := strings.NewReader("one: 1\n")
	expected := strings.NewReader("one: 1.0\n")

	printExample(yamlassert.EqualReader(actual, expected))
	// Output:
	// message="YAML mismatch (-want +got):\n"
	// success=true
}

func ExampleValid() {
	printExample(yamlassert.Valid("name: Ada\n"))
	// Output:
	// message="got valid YAML"
	// success=true
}

func ExampleValidBytes() {
	printExample(yamlassert.ValidBytes([]byte("name: Ada\n")))
	// Output:
	// message="got valid YAML"
	// success=true
}

func ExampleValidReader() {
	printExample(yamlassert.ValidReader(strings.NewReader("name: Ada\n")))
	// Output:
	// message="got valid YAML"
	// success=true
}

func ExampleContains() {
	actual := "status: ok\nmeta:\n  page: 1\n  request_id: abc\n"
	expected := "meta:\n  page: 1.0\n"

	printExample(yamlassert.Contains(actual, expected))
	// Output:
	// message="YAML contained expected subset"
	// success=true
}

func ExampleContainsBytes() {
	actual := []byte("status: ok\nextra: true\n")
	expected := []byte("status: ok\n")

	printExample(yamlassert.ContainsBytes(actual, expected))
	// Output:
	// message="YAML contained expected subset"
	// success=true
}

func ExamplePathExists() {
	printExample(yamlassert.PathExists("user:\n  name: Ada\n", "user.name"))
	// Output:
	// message="YAML path <user.name> exists"
	// success=true
}

func ExamplePathMissing() {
	printExample(yamlassert.PathMissing("user:\n  name: Ada\n", "user.email"))
	// Output:
	// message="YAML path <user.email> missing"
	// success=true
}

func ExamplePathEqual() {
	printExample(yamlassert.PathEqual("user:\n  age: 37\n", "user.age", "37.0\n"))
	// Output:
	// message="YAML path <user.age> mismatch (-want +got):\n"
	// success=true
}

func ExamplePathEqualBytes() {
	actual := []byte("user:\n  name: Ada\n")
	expected := []byte("Ada\n")

	printExample(yamlassert.PathEqualBytes(actual, "user.name", expected))
	// Output:
	// message="YAML path <user.name> mismatch (-want +got):\n"
	// success=true
}

func ExampleArrayContains() {
	actual := "items:\n  - id: 1\n  - id: 2\n"

	printExample(yamlassert.ArrayContains(actual, "items", "id: 2.0\n"))
	// Output:
	// message="found matching element at YAML path <items> index <1>"
	// success=true
}

func ExampleArrayContainsBytes() {
	actual := []byte("items: [1, 2]\n")
	expected := []byte("2.0\n")

	printExample(yamlassert.ArrayContainsBytes(actual, "items", expected))
	// Output:
	// message="found matching element at YAML path <items> index <1>"
	// success=true
}

func ExampleIgnorePaths() {
	actual := "status: ok\nmeta:\n  request_id: abc\n"
	expected := "status: ok\nmeta:\n  request_id: xyz\n"

	printExample(yamlassert.EqualWithOptions(actual, expected, yamlassert.IgnorePaths("meta.request_id")))
	// Output:
	// message="YAML mismatch (-want +got):\n"
	// success=true
}

func ExampleUnorderedArraysAt() {
	actual := "tags: [go, test, yaml]\n"
	expected := "tags: [yaml, go, test]\n"

	printExample(yamlassert.EqualWithOptions(actual, expected, yamlassert.UnorderedArraysAt("tags")))
	// Output:
	// message="YAML mismatch (-want +got):\n"
	// success=true
}

func ExampleDocumentCount() {
	printExample(yamlassert.DocumentCount("---\nname: Ada\n---\nname: Grace\n", 2))
	// Output:
	// message="got YAML document count <2>, wanted <2>"
	// success=true
}

func ExampleDocumentCountBytes() {
	printExample(yamlassert.DocumentCountBytes([]byte("---\nname: Ada\n---\nname: Grace\n"), 2))
	// Output:
	// message="got YAML document count <2>, wanted <2>"
	// success=true
}

func ExampleDocumentEqual() {
	printExample(yamlassert.DocumentEqual("---\nname: Ada\n---\nname: Grace\n", 1, "name: Grace\n"))
	// Output:
	// message="YAML document <1> mismatch (-want +got):\n"
	// success=true
}

func ExampleDocumentEqualBytes() {
	printExample(yamlassert.DocumentEqualBytes([]byte("---\nname: Ada\n"), 0, []byte("name: Ada\n")))
	// Output:
	// message="YAML document <0> mismatch (-want +got):\n"
	// success=true
}

func ExampleDocumentContains() {
	actual := "---\nstatus: ok\nmeta:\n  page: 1\n"
	expected := "meta:\n  page: 1.0\n"

	printExample(yamlassert.DocumentContains(actual, 0, expected))
	// Output:
	// message="YAML document <0> contained expected subset"
	// success=true
}

func ExampleDocumentContainsBytes() {
	printExample(yamlassert.DocumentContainsBytes([]byte("---\nstatus: ok\n"), 0, []byte("status: ok\n")))
	// Output:
	// message="YAML document <0> contained expected subset"
	// success=true
}

func printExample(result hammy.AssertionMessage) {
	fmt.Printf("message=%q\nsuccess=%t\n", result.Message, result.IsSuccessful)
}
