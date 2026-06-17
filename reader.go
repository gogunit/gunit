package gunit

import (
	"bytes"
	"io"
	"strings"
)

func Reader(t T, actual io.Reader) *ReaderAssert { return &ReaderAssert{T: t, actual: actual} }

type ReaderAssert struct {
	T
	actual io.Reader
	body   []byte
	read   bool
}

func (r *ReaderAssert) EqualToString(expected string) {
	r.Helper()
	actual, ok := r.readAll()
	if !ok {
		return
	}
	if string(actual) != expected {
		r.Errorf("got reader string <%s>, wanted <%s>", string(actual), expected)
	}
}

func (r *ReaderAssert) ContainsString(expected string) {
	r.Helper()
	actual, ok := r.readAll()
	if !ok {
		return
	}
	if !strings.Contains(string(actual), expected) {
		r.Errorf("got reader string <%s>, wanted substring <%s>", string(actual), expected)
	}
}

func (r *ReaderAssert) EqualToBytes(expected []byte) {
	r.Helper()
	actual, ok := r.readAll()
	if !ok {
		return
	}
	if !bytes.Equal(actual, expected) {
		r.Errorf("got reader bytes <%v>, wanted <%v>", actual, expected)
	}
}

func (r *ReaderAssert) IsEmpty() {
	r.Helper()
	actual, ok := r.readAll()
	if !ok {
		return
	}
	if len(actual) != 0 {
		r.Errorf("got reader len()=%d, wanted empty reader", len(actual))
	}
}

func (r *ReaderAssert) IsNotEmpty() {
	r.Helper()
	actual, ok := r.readAll()
	if !ok {
		return
	}
	if len(actual) == 0 {
		r.Errorf("got empty reader, wanted non-empty reader")
	}
}

func (r *ReaderAssert) NotEmpty() { r.IsNotEmpty() }

func (r *ReaderAssert) readAll() ([]byte, bool) {
	if r.actual == nil {
		r.Errorf("got nil reader, wanted reader body")
		return nil, false
	}
	if r.read {
		return r.body, true
	}
	body, err := io.ReadAll(r.actual)
	if err != nil {
		r.Errorf("got reader read error: %v", err)
		return nil, false
	}
	r.body = body
	r.read = true
	return body, true
}
