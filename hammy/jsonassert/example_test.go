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

func ExampleEqualBytes() {
	actual := []byte(`{"one":1}`)
	expected := []byte(`{"one":1.0}`)

	printExample(jsonassert.EqualBytes(actual, expected))
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

func printExample(result hammy.AssertionMessage) {
	fmt.Printf("message=%q\nsuccess=%t\n", result.Message, result.IsSuccessful)
}
