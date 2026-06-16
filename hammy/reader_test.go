package hammy_test

import (
	"errors"
	"io"
	"strings"
	"testing"

	"github.com/gogunit/gunit/eye"
	a "github.com/gogunit/gunit/hammy"
)

func Test_reader_EqualToString_success(t *testing.T) {
	a.New(t).Is(a.Reader(strings.NewReader("hello world")).EqualToString("hello world"))
}

func Test_reader_EqualToString_failure(t *testing.T) {
	aSpy := eye.Spy()
	a.New(aSpy).Is(a.Reader(strings.NewReader("hello world")).EqualToString("goodbye"))
	aSpy.HadErrorContaining(t, "got reader string <hello world>, wanted <goodbye>")
}

func Test_reader_ContainsString_success(t *testing.T) {
	a.New(t).Is(a.Reader(strings.NewReader("hello world")).ContainsString("world"))
}

func Test_reader_ContainsString_failure(t *testing.T) {
	aSpy := eye.Spy()
	a.New(aSpy).Is(a.Reader(strings.NewReader("hello world")).ContainsString("goodbye"))
	aSpy.HadErrorContaining(t, "got reader string <hello world>, wanted substring <goodbye>")
}

func Test_reader_EqualToBytes_success(t *testing.T) {
	a.New(t).Is(a.Reader(strings.NewReader("abc")).EqualToBytes([]byte("abc")))
}

func Test_reader_EqualToBytes_failure(t *testing.T) {
	aSpy := eye.Spy()
	a.New(aSpy).Is(a.Reader(strings.NewReader("abc")).EqualToBytes([]byte("cba")))
	aSpy.HadErrorContaining(t, "got reader bytes <[97 98 99]>, wanted <[99 98 97]>")
}

func Test_reader_IsEmpty_success(t *testing.T) {
	a.New(t).Is(a.Reader(strings.NewReader("")).IsEmpty())
}

func Test_reader_IsEmpty_failure(t *testing.T) {
	aSpy := eye.Spy()
	a.New(aSpy).Is(a.Reader(strings.NewReader("hello")).IsEmpty())
	aSpy.HadErrorContaining(t, "got reader len()=5, wanted empty reader")
}

func Test_reader_NotEmpty_success(t *testing.T) {
	a.New(t).Is(a.Reader(strings.NewReader("hello")).NotEmpty())
}

func Test_reader_NotEmpty_failure(t *testing.T) {
	aSpy := eye.Spy()
	a.New(aSpy).Is(a.Reader(strings.NewReader("")).NotEmpty())
	aSpy.HadErrorContaining(t, "got empty reader, wanted non-empty reader")
}

func Test_reader_nil_failure(t *testing.T) {
	aSpy := eye.Spy()
	a.New(aSpy).Is(a.Reader(nil).EqualToString("hello"))
	aSpy.HadErrorContaining(t, "got nil reader, wanted reader body")
}

func Test_reader_read_error_failure(t *testing.T) {
	aSpy := eye.Spy()
	a.New(aSpy).Is(a.Reader(errorReader{}).EqualToString("hello"))
	aSpy.HadErrorContaining(t, "got reader read error: boom")
}

func Test_reader_reuses_buffered_body_after_first_read(t *testing.T) {
	reader := strings.NewReader("hello")
	assert := a.Reader(reader)

	a.New(t).Is(assert.EqualToString("hello"))
	a.New(t).Is(assert.ContainsString("ell"))

	remaining, err := io.ReadAll(reader)
	if err != nil {
		t.Fatal(err)
	}
	if len(remaining) != 0 {
		t.Fatalf("got remaining reader bytes <%v>, wanted empty reader after assertion", remaining)
	}
}

type errorReader struct{}

func (errorReader) Read([]byte) (int, error) {
	return 0, errors.New("boom")
}
