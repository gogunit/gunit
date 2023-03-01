package gunit_test

import (
	. "github.com/nfisher/gunit"
	"testing"
)

func Test_string_EqualTo_success(t *testing.T) {
	String(t, "Hello world").EqualTo("Hello world")
}

func Test_string_EqualTo_failure(t *testing.T) {
	aSpy := spy()
	String(aSpy, "Hello World").EqualTo("Good-bye")
	aSpy.HadError(t)
}

func Test_string_Contains_success(t *testing.T) {
	String(t, "Baz Foo Bar").Contains("Foo")
}

func Test_string_Contains_failure(t *testing.T) {
	aSpy := spy()
	String(aSpy, "Foo").Contains("Baz")
	aSpy.HadError(t)
}

func Test_string_HasPrefix_success(t *testing.T) {
	String(t, "bluefin").HasPrefix("blue")
}

func Test_string_HasPrefix_failure(t *testing.T) {
	aSpy := spy()
	String(aSpy, "bluefin").HasPrefix("red")
	aSpy.HadError(t)
}

func Test_string_HasSuffix_success(t *testing.T) {
	String(t, "bluefin").HasSuffix("fin")
}

func Test_string_HasSuffix_failure(t *testing.T) {
	aSpy := spy()
	String(aSpy, "bluefin").HasSuffix("fish")
	aSpy.HadError(t)
}

func Test_string_IsEmpty_success(t *testing.T) {
	String(t, "").IsEmpty()
}

func Test_string_IsEmpty_failure(t *testing.T) {
	aSpy := spy()
	String(aSpy, "hi").IsEmpty()
	aSpy.HadError(t)
}

func Test_string_IsNotEmpty_success(t *testing.T) {
	String(t, "hi").IsNotEmpty()
}

func Test_string_IsNotEmpty_failure(t *testing.T) {
	aSpy := spy()
	String(aSpy, "").IsNotEmpty()
	aSpy.HadError(t)
}
