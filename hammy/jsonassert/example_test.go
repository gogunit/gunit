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

	printExample(jsonassert.String(actual).EqualTo(expected))
	// Output:
	// message="JSON mismatch (-want +got):\n"
	// success=true
}

func ExampleEqual_reorderedObjectKeys() {
	actual := `{"name":"Ada","age":37}`
	expected := `{"age":37,"name":"Ada"}`

	printExample(jsonassert.String(actual).EqualTo(expected))
	// Output:
	// message="JSON mismatch (-want +got):\n"
	// success=true
}

func ExampleEqualWithOptions() {
	actual := `{"status":"ok","meta":{"request_id":"abc"}}`
	expected := `{"status":"ok","meta":{"request_id":"xyz"}}`

	printExample(jsonassert.String(actual).EqualToWithOptions(expected, jsonassert.IgnorePaths("meta.request_id")))
	// Output:
	// message="JSON mismatch (-want +got):\n"
	// success=true
}

func ExampleEqualLines() {
	actual := `{"name":"Ada","age":37}` + "\n" + `{"name":"Grace","age":85}`
	expected := `{"age":37.0,"name":"Ada"}` + "\n" + `{"age":85.0,"name":"Grace"}`

	printExample(jsonassert.String(actual).LinesEqualTo(expected))
	// Output:
	// message="JSONL mismatch (-want +got):\n"
	// success=true
}

func ExampleEqualLinesWithOptions() {
	actual := `{"status":"ok","meta":{"request_id":"abc"}}` + "\n" + `{"status":"ok","meta":{"request_id":"def"}}`
	expected := `{"status":"ok","meta":{"request_id":"uvw"}}` + "\n" + `{"status":"ok","meta":{"request_id":"xyz"}}`

	printExample(jsonassert.String(actual).LinesEqualToWithOptions(expected, jsonassert.IgnorePaths("meta.request_id")))
	// Output:
	// message="JSONL mismatch (-want +got):\n"
	// success=true
}

func ExampleEqualLinesBytes() {
	actual := []byte(`{"id":1}` + "\n" + `{"id":2}`)
	expected := []byte(`{"id":1.0}` + "\n" + `{"id":2.0}`)

	printExample(jsonassert.Bytes(actual).LinesEqualTo(expected))
	// Output:
	// message="JSONL mismatch (-want +got):\n"
	// success=true
}

func ExampleEqualLinesBytesWithOptions() {
	actual := []byte(`{"tags":["go","test"]}` + "\n" + `{"tags":["json","assert"]}`)
	expected := []byte(`{"tags":["test","go"]}` + "\n" + `{"tags":["assert","json"]}`)

	printExample(jsonassert.Bytes(actual).LinesEqualToWithOptions(expected, jsonassert.UnorderedArraysAt("tags")))
	// Output:
	// message="JSONL mismatch (-want +got):\n"
	// success=true
}

func ExampleEqualBytes() {
	actual := []byte(`{"one":1}`)
	expected := []byte(`{"one":1.0}`)

	printExample(jsonassert.Bytes(actual).EqualTo(expected))
	// Output:
	// message="JSON mismatch (-want +got):\n"
	// success=true
}

func ExampleEqualBytesWithOptions() {
	actual := []byte(`{"tags":["go","test"]}`)
	expected := []byte(`{"tags":["test","go"]}`)

	printExample(jsonassert.Bytes(actual).EqualToWithOptions(expected, jsonassert.UnorderedArraysAt("tags")))
	// Output:
	// message="JSON mismatch (-want +got):\n"
	// success=true
}

func ExampleEqualReader() {
	actual := strings.NewReader(`{"one":1}`)
	expected := strings.NewReader(`{"one":1.0}`)

	printExample(jsonassert.Reader(actual).EqualTo(expected))
	// Output:
	// message="JSON mismatch (-want +got):\n"
	// success=true
}

func ExampleValid() {
	printExample(jsonassert.String(`{"name":"Ada"}`).IsValid())
	// Output:
	// message="got valid JSON"
	// success=true
}

func ExampleValidBytes() {
	printExample(jsonassert.Bytes([]byte(`{"name":"Ada"}`)).IsValid())
	// Output:
	// message="got valid JSON"
	// success=true
}

func ExampleValidReader() {
	printExample(jsonassert.Reader(strings.NewReader(`{"name":"Ada"}`)).IsValid())
	// Output:
	// message="got valid JSON"
	// success=true
}

func ExampleContains() {
	actual := `{"status":"ok","meta":{"page":1,"request_id":"abc"}}`
	expected := `{"meta":{"page":1.0}}`

	printExample(jsonassert.String(actual).Contains(expected))
	// Output:
	// message="JSON contained expected subset"
	// success=true
}

func ExampleContainsBytes() {
	actual := []byte(`{"status":"ok","extra":true}`)
	expected := []byte(`{"status":"ok"}`)

	printExample(jsonassert.Bytes(actual).Contains(expected))
	// Output:
	// message="JSON contained expected subset"
	// success=true
}

func ExampleLinesContain() {
	actual := `{"id":1,"name":"Ada"}` + "\n" + `{"id":2,"name":"Grace"}`

	printExample(jsonassert.String(actual).LinesContain(`{"name":"Grace","id":2.0}`))
	// Output:
	// message="found matching JSONL line <1>"
	// success=true
}

func ExampleLinesContainSubset() {
	actual := `{"status":"ok","meta":{"page":1,"request_id":"abc"}}` + "\n" + `{"status":"done"}`

	printExample(jsonassert.String(actual).LinesContainSubset(`{"meta":{"page":1.0}}`))
	// Output:
	// message="found JSONL line <0> containing expected subset"
	// success=true
}

func ExamplePathExists() {
	printExample(jsonassert.String(`{"user":{"name":"Ada"}}`).PathExists("user.name"))
	// Output:
	// message="JSON path <user.name> exists"
	// success=true
}

func ExamplePathMissing() {
	printExample(jsonassert.String(`{"user":{"name":"Ada"}}`).PathMissing("user.email"))
	// Output:
	// message="JSON path <user.email> missing"
	// success=true
}

func ExamplePathEqual() {
	printExample(jsonassert.String(`{"user":{"age":37}}`).PathEqual("user.age", `37.0`))
	// Output:
	// message="JSON path <user.age> mismatch (-want +got):\n"
	// success=true
}

func ExamplePathEqualBytes() {
	actual := []byte(`{"user":{"name":"Ada"}}`)
	expected := []byte(`"Ada"`)

	printExample(jsonassert.Bytes(actual).PathEqual("user.name", expected))
	// Output:
	// message="JSON path <user.name> mismatch (-want +got):\n"
	// success=true
}

func ExampleArrayContains() {
	actual := `{"items":[{"id":1},{"id":2}]}`

	printExample(jsonassert.String(actual).ArrayContains("items", `{"id":2.0}`))
	// Output:
	// message="found matching element at JSON path <items> index <1>"
	// success=true
}

func ExampleArrayContainsBytes() {
	actual := []byte(`{"items":[1,2]}`)
	expected := []byte(`2.0`)

	printExample(jsonassert.Bytes(actual).ArrayContains("items", expected))
	// Output:
	// message="found matching element at JSON path <items> index <1>"
	// success=true
}

func ExampleIgnorePaths() {
	actual := `{"status":"ok","meta":{"request_id":"abc"}}`
	expected := `{"status":"ok","meta":{"request_id":"xyz"}}`

	printExample(jsonassert.String(actual).EqualToWithOptions(expected, jsonassert.IgnorePaths("meta.request_id")))
	// Output:
	// message="JSON mismatch (-want +got):\n"
	// success=true
}

func ExampleUnorderedArraysAt() {
	actual := `{"tags":["go","test","json"]}`
	expected := `{"tags":["json","go","test"]}`

	printExample(jsonassert.String(actual).EqualToWithOptions(expected, jsonassert.UnorderedArraysAt("tags")))
	// Output:
	// message="JSON mismatch (-want +got):\n"
	// success=true
}

func printExample(result hammy.AssertionMessage) {
	fmt.Printf("message=%q\nsuccess=%t\n", result.Message, result.IsSuccessful)
}
