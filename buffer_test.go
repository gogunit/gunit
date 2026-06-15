package gunit_test

import (
	"bytes"
	"testing"

	"github.com/gogunit/gunit"
	"github.com/gogunit/gunit/eye"
)

func Test_buffer_EqualToString_success(t *testing.T) {
	gunit.Buffer(t, bytes.NewBufferString("hello world")).EqualToString("hello world")
}

func Test_buffer_EqualToString_failure(t *testing.T) {
	aSpy := eye.Spy()
	gunit.Buffer(aSpy, bytes.NewBufferString("hello world")).EqualToString("goodbye")
	aSpy.HadErrorContaining(t, "got buffer string <hello world>, wanted <goodbye>")
}

func Test_buffer_ContainsString_success(t *testing.T) {
	gunit.Buffer(t, bytes.NewBufferString("hello world")).ContainsString("world")
}

func Test_buffer_ContainsString_failure(t *testing.T) {
	aSpy := eye.Spy()
	gunit.Buffer(aSpy, bytes.NewBufferString("hello world")).ContainsString("goodbye")
	aSpy.HadErrorContaining(t, "got buffer string <hello world>, wanted substring <goodbye>")
}

func Test_buffer_EqualToBytes_success(t *testing.T) {
	gunit.Buffer(t, bytes.NewBuffer([]byte{1, 2, 3})).EqualToBytes([]byte{1, 2, 3})
}

func Test_buffer_EqualToBytes_failure(t *testing.T) {
	aSpy := eye.Spy()
	gunit.Buffer(aSpy, bytes.NewBuffer([]byte{1, 2, 3})).EqualToBytes([]byte{3, 2, 1})
	aSpy.HadErrorContaining(t, "got buffer bytes <[1 2 3]>, wanted <[3 2 1]>")
}

func Test_buffer_Len_success(t *testing.T) {
	gunit.Buffer(t, bytes.NewBufferString("hello")).Len(5)
}

func Test_buffer_Len_failure(t *testing.T) {
	aSpy := eye.Spy()
	gunit.Buffer(aSpy, bytes.NewBufferString("hello")).Len(4)
	aSpy.HadErrorContaining(t, "got buffer len()=5, wanted 4")
}

func Test_buffer_IsEmpty_success(t *testing.T) {
	gunit.Buffer(t, bytes.NewBuffer(nil)).IsEmpty()
}

func Test_buffer_IsEmpty_failure(t *testing.T) {
	aSpy := eye.Spy()
	gunit.Buffer(aSpy, bytes.NewBufferString("hello")).IsEmpty()
	aSpy.HadErrorContaining(t, "got buffer len()=5, wanted empty buffer")
}

func Test_buffer_IsNotEmpty_success(t *testing.T) {
	gunit.Buffer(t, bytes.NewBufferString("hello")).IsNotEmpty()
}

func Test_buffer_IsNotEmpty_failure(t *testing.T) {
	aSpy := eye.Spy()
	gunit.Buffer(aSpy, bytes.NewBuffer(nil)).IsNotEmpty()
	aSpy.HadErrorContaining(t, "got empty buffer, wanted non-empty buffer")
}

func Test_buffer_nil_failure(t *testing.T) {
	aSpy := eye.Spy()
	gunit.Buffer(aSpy, nil).EqualToString("hello")
	aSpy.HadErrorContaining(t, "got nil buffer, wanted string <hello>")
}
