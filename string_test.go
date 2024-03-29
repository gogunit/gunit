package gunit_test

import (
	"github.com/gogunit/gunit"
	"github.com/gogunit/gunit/eye"
	"testing"
)

func Test_string_EqualTo_success(t *testing.T) {
	gunit.String(t, "Hello world").EqualTo("Hello world")
}

func Test_string_EqualTo_failure(t *testing.T) {
	aSpy := eye.Spy()
	gunit.String(aSpy, "Hello World").EqualTo("Good-bye")
	aSpy.HadError(t)
}

func Test_string_Contains_success(t *testing.T) {
	gunit.String(t, "Baz Foo Bar").Contains("Foo")
}

func Test_string_Contains_failure(t *testing.T) {
	aSpy := eye.Spy()
	gunit.String(aSpy, "Foo").Contains("Baz")
	aSpy.HadError(t)
}

func Test_string_HasPrefix_success(t *testing.T) {
	gunit.String(t, "bluefin").HasPrefix("blue")
}

func Test_string_HasPrefix_failure(t *testing.T) {
	aSpy := eye.Spy()
	gunit.String(aSpy, "bluefin").HasPrefix("red")
	aSpy.HadError(t)
}

func Test_string_HasSuffix_success(t *testing.T) {
	gunit.String(t, "bluefin").HasSuffix("fin")
}

func Test_string_HasSuffix_failure(t *testing.T) {
	aSpy := eye.Spy()
	gunit.String(aSpy, "bluefin").HasSuffix("fish")
	aSpy.HadError(t)
}

func Test_string_IsEmpty_success(t *testing.T) {
	gunit.String(t, "").IsEmpty()
}

func Test_string_IsEmpty_failure(t *testing.T) {
	aSpy := eye.Spy()
	gunit.String(aSpy, "hi").IsEmpty()
	aSpy.HadError(t)
}

func Test_string_IsNotEmpty_success(t *testing.T) {
	gunit.String(t, "hi").IsNotEmpty()
}

func Test_string_IsNotEmpty_failure(t *testing.T) {
	aSpy := eye.Spy()
	gunit.String(aSpy, "").IsNotEmpty()
	aSpy.HadError(t)
}

func Test_string_subtype_EqualTo(t *testing.T) {
	type S string
	var s S = "Hello world"
	gunit.String(t, s).EqualTo("Hello world")
}
