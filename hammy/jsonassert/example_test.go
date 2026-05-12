package jsonassert_test

import (
	"fmt"
	"strings"

	"github.com/gogunit/gunit/hammy"
	"github.com/gogunit/gunit/hammy/jsonassert"
)

func ExampleEqual() {
	actual := `{"name":"Ada","age":37}`
	expected := `{
		"name": "Ada",
		"age": 37
	}`

	printExample(jsonassert.Equal(actual, expected))
	// Output:
	// message="JSON mismatch (-want +got):\n"
	// success=true
}

func ExampleEqual_reorderedObjectKeys() {
	actual := `{"name":"Ada","age":37}`
	expected := `{"age":37,"name":"Ada"}`

	printExample(jsonassert.Equal(actual, expected))
	// Output:
	// message="JSON mismatch (-want +got):\n"
	// success=true
}

func ExampleEqualWithOptions() {
	actual := `{"status":"ok","meta":{"request_id":"abc"}}`
	expected := `{"status":"ok","meta":{"request_id":"xyz"}}`

	printExample(jsonassert.EqualWithOptions(actual, expected, jsonassert.IgnorePaths("meta.request_id")))
	// Output:
	// message="JSON mismatch (-want +got):\n"
	// success=true
}

func ExampleEqualBytes() {
	actual := []byte(`{"one":1}`)
	expected := []byte(`{"one":1.0}`)

	printExample(jsonassert.EqualBytes(actual, expected))
	// Output:
	// message="JSON mismatch (-want +got):\n"
	// success=true
}

func ExampleEqualBytesWithOptions() {
	actual := []byte(`{"tags":["go","test"]}`)
	expected := []byte(`{"tags":["test","go"]}`)

	printExample(jsonassert.EqualBytesWithOptions(actual, expected, jsonassert.UnorderedArraysAt("tags")))
	// Output:
	// message="JSON mismatch (-want +got):\n"
	// success=true
}

func ExampleEqualReader() {
	actual := strings.NewReader(`{"one":1}`)
	expected := strings.NewReader(`{"one":1.0}`)

	printExample(jsonassert.EqualReader(actual, expected))
	// Output:
	// message="JSON mismatch (-want +got):\n"
	// success=true
}

func ExampleValid() {
	printExample(jsonassert.Valid(`{"name":"Ada"}`))
	// Output:
	// message="got valid JSON"
	// success=true
}

func ExampleValidBytes() {
	printExample(jsonassert.ValidBytes([]byte(`{"name":"Ada"}`)))
	// Output:
	// message="got valid JSON"
	// success=true
}

func ExampleValidReader() {
	printExample(jsonassert.ValidReader(strings.NewReader(`{"name":"Ada"}`)))
	// Output:
	// message="got valid JSON"
	// success=true
}

func ExampleContains() {
	actual := `{"status":"ok","meta":{"page":1,"request_id":"abc"}}`
	expected := `{"meta":{"page":1.0}}`

	printExample(jsonassert.Contains(actual, expected))
	// Output:
	// message="JSON contained expected subset"
	// success=true
}

func ExampleContainsBytes() {
	actual := []byte(`{"status":"ok","extra":true}`)
	expected := []byte(`{"status":"ok"}`)

	printExample(jsonassert.ContainsBytes(actual, expected))
	// Output:
	// message="JSON contained expected subset"
	// success=true
}

func ExamplePathExists() {
	printExample(jsonassert.PathExists(`{"user":{"name":"Ada"}}`, "user.name"))
	// Output:
	// message="JSON path <user.name> exists"
	// success=true
}

func ExamplePathMissing() {
	printExample(jsonassert.PathMissing(`{"user":{"name":"Ada"}}`, "user.email"))
	// Output:
	// message="JSON path <user.email> missing"
	// success=true
}

func ExamplePathEqual() {
	printExample(jsonassert.PathEqual(`{"user":{"age":37}}`, "user.age", `37.0`))
	// Output:
	// message="JSON path <user.age> mismatch (-want +got):\n"
	// success=true
}

func ExamplePathEqualBytes() {
	actual := []byte(`{"user":{"name":"Ada"}}`)
	expected := []byte(`"Ada"`)

	printExample(jsonassert.PathEqualBytes(actual, "user.name", expected))
	// Output:
	// message="JSON path <user.name> mismatch (-want +got):\n"
	// success=true
}

func ExampleArrayContains() {
	actual := `{"items":[{"id":1},{"id":2}]}`

	printExample(jsonassert.ArrayContains(actual, "items", `{"id":2.0}`))
	// Output:
	// message="found matching element at JSON path <items> index <1>"
	// success=true
}

func ExampleArrayContainsBytes() {
	actual := []byte(`{"items":[1,2]}`)
	expected := []byte(`2.0`)

	printExample(jsonassert.ArrayContainsBytes(actual, "items", expected))
	// Output:
	// message="found matching element at JSON path <items> index <1>"
	// success=true
}

func ExampleIgnorePaths() {
	actual := `{"status":"ok","meta":{"request_id":"abc"}}`
	expected := `{"status":"ok","meta":{"request_id":"xyz"}}`

	printExample(jsonassert.EqualWithOptions(actual, expected, jsonassert.IgnorePaths("meta.request_id")))
	// Output:
	// message="JSON mismatch (-want +got):\n"
	// success=true
}

func ExampleUnorderedArraysAt() {
	actual := `{"tags":["go","test","json"]}`
	expected := `{"tags":["json","go","test"]}`

	printExample(jsonassert.EqualWithOptions(actual, expected, jsonassert.UnorderedArraysAt("tags")))
	// Output:
	// message="JSON mismatch (-want +got):\n"
	// success=true
}

func printExample(result hammy.AssertionMessage) {
	fmt.Printf("message=%q\nsuccess=%t\n", result.Message, result.IsSuccessful)
}
