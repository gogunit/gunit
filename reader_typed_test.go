package gunit_test

import (
	"bytes"
	"github.com/gogunit/gunit"
	"github.com/gogunit/gunit/eye"
	"strings"
	"testing"
)

func Test_reader_ContainsString_success(t *testing.T) {
	gunit.Reader(t, strings.NewReader("hello world")).ContainsString("world")
}
func Test_reader_ContainsString_failure(t *testing.T) {
	aSpy := eye.Spy()
	gunit.Reader(aSpy, strings.NewReader("hello")).ContainsString("bye")
	aSpy.HadErrorContaining(t, "wanted substring")
}

func Test_reader_EqualToString_success(t *testing.T) {
	gunit.Reader(t, strings.NewReader("hello")).EqualToString("hello")
}
func Test_reader_EqualToString_failure(t *testing.T) {
	aSpy := eye.Spy()
	gunit.Reader(aSpy, strings.NewReader("hello")).EqualToString("bye")
	aSpy.HadErrorContaining(t, "wanted <bye>")
}
func Test_reader_EqualToString_failure_for_nil(t *testing.T) {
	aSpy := eye.Spy()
	gunit.Reader(aSpy, nil).EqualToString("hello")
	aSpy.HadErrorContaining(t, "got nil reader")
}

func Test_reader_EqualToBytes_success(t *testing.T) {
	gunit.Reader(t, bytes.NewBufferString("bytes")).EqualToBytes([]byte("bytes"))
}
func Test_reader_EqualToBytes_failure(t *testing.T) {
	aSpy := eye.Spy()
	gunit.Reader(aSpy, bytes.NewBufferString("bytes")).EqualToBytes([]byte("other"))
	aSpy.HadErrorContaining(t, "wanted")
}

func Test_reader_IsEmpty_success(t *testing.T) { gunit.Reader(t, strings.NewReader("")).IsEmpty() }
func Test_reader_IsEmpty_failure(t *testing.T) {
	aSpy := eye.Spy()
	gunit.Reader(aSpy, strings.NewReader("x")).IsEmpty()
	aSpy.HadErrorContaining(t, "wanted empty reader")
}

func Test_reader_IsNotEmpty_success(t *testing.T) {
	gunit.Reader(t, strings.NewReader("x")).IsNotEmpty()
}
func Test_reader_IsNotEmpty_failure(t *testing.T) {
	aSpy := eye.Spy()
	gunit.Reader(aSpy, strings.NewReader("")).IsNotEmpty()
	aSpy.HadErrorContaining(t, "wanted non-empty reader")
}

func Test_reader_NotEmpty_success(t *testing.T) { gunit.Reader(t, strings.NewReader("x")).NotEmpty() }
func Test_reader_NotEmpty_failure(t *testing.T) {
	aSpy := eye.Spy()
	gunit.Reader(aSpy, strings.NewReader("")).NotEmpty()
	aSpy.HadErrorContaining(t, "wanted non-empty reader")
}
