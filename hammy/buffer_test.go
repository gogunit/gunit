package hammy_test

import (
	"bytes"
	"testing"

	"github.com/gogunit/gunit/eye"
	a "github.com/gogunit/gunit/hammy"
)

func Test_buffer_EqualToString_success(t *testing.T) {
	a.New(t).Is(a.Buffer(bytes.NewBufferString("hello world")).EqualToString("hello world"))
}

func Test_buffer_EqualToString_failure(t *testing.T) {
	aSpy := eye.Spy()
	a.New(aSpy).Is(a.Buffer(bytes.NewBufferString("hello world")).EqualToString("goodbye"))
	aSpy.HadErrorContaining(t, "got buffer string <hello world>, wanted <goodbye>")
}

func Test_buffer_ContainsString_success(t *testing.T) {
	a.New(t).Is(a.Buffer(bytes.NewBufferString("hello world")).ContainsString("world"))
}

func Test_buffer_ContainsString_failure(t *testing.T) {
	aSpy := eye.Spy()
	a.New(aSpy).Is(a.Buffer(bytes.NewBufferString("hello world")).ContainsString("goodbye"))
	aSpy.HadErrorContaining(t, "got buffer string <hello world>, wanted substring <goodbye>")
}

func Test_buffer_EqualToBytes_success(t *testing.T) {
	a.New(t).Is(a.Buffer(bytes.NewBuffer([]byte{1, 2, 3})).EqualToBytes([]byte{1, 2, 3}))
}

func Test_buffer_EqualToBytes_failure(t *testing.T) {
	aSpy := eye.Spy()
	a.New(aSpy).Is(a.Buffer(bytes.NewBuffer([]byte{1, 2, 3})).EqualToBytes([]byte{3, 2, 1}))
	aSpy.HadErrorContaining(t, "got buffer bytes <[1 2 3]>, wanted <[3 2 1]>")
}

func Test_buffer_Len_success(t *testing.T) {
	a.New(t).Is(a.Buffer(bytes.NewBufferString("hello")).Len(5))
}

func Test_buffer_Len_failure(t *testing.T) {
	aSpy := eye.Spy()
	a.New(aSpy).Is(a.Buffer(bytes.NewBufferString("hello")).Len(4))
	aSpy.HadErrorContaining(t, "got buffer len()=5, wanted 4")
}

func Test_buffer_IsEmpty_success(t *testing.T) {
	a.New(t).Is(a.Buffer(bytes.NewBuffer(nil)).IsEmpty())
}

func Test_buffer_IsEmpty_failure(t *testing.T) {
	aSpy := eye.Spy()
	a.New(aSpy).Is(a.Buffer(bytes.NewBufferString("hello")).IsEmpty())
	aSpy.HadErrorContaining(t, "got buffer len()=5, wanted empty buffer")
}

func Test_buffer_NotEmpty_success(t *testing.T) {
	a.New(t).Is(a.Buffer(bytes.NewBufferString("hello")).NotEmpty())
}

func Test_buffer_NotEmpty_failure(t *testing.T) {
	aSpy := eye.Spy()
	a.New(aSpy).Is(a.Buffer(bytes.NewBuffer(nil)).NotEmpty())
	aSpy.HadErrorContaining(t, "got empty buffer, wanted non-empty buffer")
}

func Test_buffer_Matches_success(t *testing.T) {
	matcher := a.MatchFunc(func(actual *bytes.Buffer) a.AssertionMessage {
		return a.Assert(actual != nil && actual.Len() == 5, "got buffer len()=%d, wanted 5", actual.Len())
	})

	a.New(t).Is(a.Buffer(bytes.NewBufferString("hello")).Matches(matcher))
}

func Test_buffer_Matches_nil_failure(t *testing.T) {
	aSpy := eye.Spy()
	matcher := a.MatchFunc(func(actual *bytes.Buffer) a.AssertionMessage {
		return a.Assert(true, "matched")
	})

	a.New(aSpy).Is(a.Buffer(nil).Matches(matcher))
	aSpy.HadErrorContaining(t, "got nil buffer, wanted to match buffer")
}

func Test_buffer_nil_failure(t *testing.T) {
	aSpy := eye.Spy()
	a.New(aSpy).Is(a.Buffer(nil).EqualToString("hello"))
	aSpy.HadErrorContaining(t, "got nil buffer, wanted string <hello>")
}
