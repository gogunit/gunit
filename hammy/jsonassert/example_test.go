package jsonassert_test

import (
	"fmt"
	"strings"

	"github.com/gogunit/gunit/hammy"
	"github.com/gogunit/gunit/hammy/jsonassert"
)

func ExampleStringAssert_EqualTo() {
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

func ExampleStringAssert_EqualTo_reorderedObjectKeys() {
	actual := `{"name":"Ada","age":37}`
	expected := `{"age":37,"name":"Ada"}`

	printExample(jsonassert.String(actual).EqualTo(expected))
	// Output:
	// message="JSON mismatch (-want +got):\n"
	// success=true
}

func ExampleStringAssert_EqualToWithOptions() {
	actual := `{"status":"ok","meta":{"request_id":"abc"}}`
	expected := `{"status":"ok","meta":{"request_id":"xyz"}}`

	printExample(jsonassert.String(actual).EqualToWithOptions(expected, jsonassert.IgnorePaths("meta.request_id")))
	// Output:
	// message="JSON mismatch (-want +got):\n"
	// success=true
}

func ExampleStringAssert_LinesEqualTo() {
	actual := `{"name":"Ada","age":37}` + "\n" + `{"name":"Grace","age":85}`
	expected := `{"age":37.0,"name":"Ada"}` + "\n" + `{"age":85.0,"name":"Grace"}`

	printExample(jsonassert.String(actual).LinesEqualTo(expected))
	// Output:
	// message="JSONL mismatch (-want +got):\n"
	// success=true
}

func ExampleStringAssert_LinesEqualToWithOptions() {
	actual := `{"status":"ok","meta":{"request_id":"abc"}}` + "\n" + `{"status":"ok","meta":{"request_id":"def"}}`
	expected := `{"status":"ok","meta":{"request_id":"uvw"}}` + "\n" + `{"status":"ok","meta":{"request_id":"xyz"}}`

	printExample(jsonassert.String(actual).LinesEqualToWithOptions(expected, jsonassert.IgnorePaths("meta.request_id")))
	// Output:
	// message="JSONL mismatch (-want +got):\n"
	// success=true
}

func ExampleBytesAssert_LinesEqualTo() {
	actual := []byte(`{"id":1}` + "\n" + `{"id":2}`)
	expected := []byte(`{"id":1.0}` + "\n" + `{"id":2.0}`)

	printExample(jsonassert.Bytes(actual).LinesEqualTo(expected))
	// Output:
	// message="JSONL mismatch (-want +got):\n"
	// success=true
}

func ExampleBytesAssert_LinesEqualToWithOptions() {
	actual := []byte(`{"tags":["go","test"]}` + "\n" + `{"tags":["json","assert"]}`)
	expected := []byte(`{"tags":["test","go"]}` + "\n" + `{"tags":["assert","json"]}`)

	printExample(jsonassert.Bytes(actual).LinesEqualToWithOptions(expected, jsonassert.UnorderedArraysAt("tags")))
	// Output:
	// message="JSONL mismatch (-want +got):\n"
	// success=true
}

func ExampleBytesAssert_EqualTo() {
	actual := []byte(`{"one":1}`)
	expected := []byte(`{"one":1.0}`)

	printExample(jsonassert.Bytes(actual).EqualTo(expected))
	// Output:
	// message="JSON mismatch (-want +got):\n"
	// success=true
}

func ExampleBytesAssert_EqualToWithOptions() {
	actual := []byte(`{"tags":["go","test"]}`)
	expected := []byte(`{"tags":["test","go"]}`)

	printExample(jsonassert.Bytes(actual).EqualToWithOptions(expected, jsonassert.UnorderedArraysAt("tags")))
	// Output:
	// message="JSON mismatch (-want +got):\n"
	// success=true
}

func ExampleReaderAssert_EqualTo() {
	actual := strings.NewReader(`{"one":1}`)
	expected := strings.NewReader(`{"one":1.0}`)

	printExample(jsonassert.Reader(actual).EqualTo(expected))
	// Output:
	// message="JSON mismatch (-want +got):\n"
	// success=true
}

func ExampleStringAssert_IsValid() {
	printExample(jsonassert.String(`{"name":"Ada"}`).IsValid())
	// Output:
	// message="got valid JSON"
	// success=true
}

func ExampleBytesAssert_IsValid() {
	printExample(jsonassert.Bytes([]byte(`{"name":"Ada"}`)).IsValid())
	// Output:
	// message="got valid JSON"
	// success=true
}

func ExampleReaderAssert_IsValid() {
	printExample(jsonassert.Reader(strings.NewReader(`{"name":"Ada"}`)).IsValid())
	// Output:
	// message="got valid JSON"
	// success=true
}

func ExampleStringAssert_Contains() {
	actual := `{"status":"ok","meta":{"page":1,"request_id":"abc"}}`
	expected := `{"meta":{"page":1.0}}`

	printExample(jsonassert.String(actual).Contains(expected))
	// Output:
	// message="JSON contained expected subset"
	// success=true
}

func ExampleBytesAssert_Contains() {
	actual := []byte(`{"status":"ok","extra":true}`)
	expected := []byte(`{"status":"ok"}`)

	printExample(jsonassert.Bytes(actual).Contains(expected))
	// Output:
	// message="JSON contained expected subset"
	// success=true
}

func ExampleStringAssert_LinesContain() {
	actual := `{"id":1,"name":"Ada"}` + "\n" + `{"id":2,"name":"Grace"}`

	printExample(jsonassert.String(actual).LinesContain(`{"name":"Grace","id":2.0}`))
	// Output:
	// message="found matching JSONL line <1>"
	// success=true
}

func ExampleStringAssert_LinesContainSubset() {
	actual := `{"status":"ok","meta":{"page":1,"request_id":"abc"}}` + "\n" + `{"status":"done"}`

	printExample(jsonassert.String(actual).LinesContainSubset(`{"meta":{"page":1.0}}`))
	// Output:
	// message="found JSONL line <0> containing expected subset"
	// success=true
}

func ExampleStringAssert_PathExists() {
	printExample(jsonassert.String(`{"user":{"name":"Ada"}}`).PathExists("user.name"))
	// Output:
	// message="JSON path <user.name> exists"
	// success=true
}

func ExampleStringAssert_PathMissing() {
	printExample(jsonassert.String(`{"user":{"name":"Ada"}}`).PathMissing("user.email"))
	// Output:
	// message="JSON path <user.email> missing"
	// success=true
}

func ExampleStringAssert_PathEqual() {
	printExample(jsonassert.String(`{"user":{"age":37}}`).PathEqual("user.age", `37.0`))
	// Output:
	// message="JSON path <user.age> mismatch (-want +got):\n"
	// success=true
}

func ExampleBytesAssert_PathEqual() {
	actual := []byte(`{"user":{"name":"Ada"}}`)
	expected := []byte(`"Ada"`)

	printExample(jsonassert.Bytes(actual).PathEqual("user.name", expected))
	// Output:
	// message="JSON path <user.name> mismatch (-want +got):\n"
	// success=true
}

func ExampleStringAssert_ArrayContains() {
	actual := `{"items":[{"id":1},{"id":2}]}`

	printExample(jsonassert.String(actual).ArrayContains("items", `{"id":2.0}`))
	// Output:
	// message="found matching element at JSON path <items> index <1>"
	// success=true
}

func ExampleBytesAssert_ArrayContains() {
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
