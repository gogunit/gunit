package jsonassert_test

import (
	"fmt"

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

func printExample(result hammy.AssertionMessage) {
	fmt.Printf("message=%q\nsuccess=%t\n", result.Message, result.IsSuccessful)
}
