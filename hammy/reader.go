package hammy

import (
	"bytes"
	"io"
	"strings"
)

// Reader creates assertions for an io.Reader.
//
// Assertions read and buffer the reader body on first use. Reading consumes the
// supplied reader unless the reader's implementation handles buffering or
// resetting itself; repeated assertions on the same ReaderAssert reuse the
// buffered bytes.
func Reader(actual io.Reader) *ReaderAssert {
	return &ReaderAssert{actual: actual}
}

type ReaderAssert struct {
	actual io.Reader
	body   []byte
	read   bool
}

func (r *ReaderAssert) EqualToString(expected string) AssertionMessage {
	actual, result := r.readAll()
	if !result.IsSuccessful {
		return result
	}
	actualString := string(actual)
	return Assert(actualString == expected, "got reader string <%s>, wanted <%s>", actualString, expected)
}

func (r *ReaderAssert) ContainsString(expected string) AssertionMessage {
	actual, result := r.readAll()
	if !result.IsSuccessful {
		return result
	}
	actualString := string(actual)
	return Assert(strings.Contains(actualString, expected), "got reader string <%s>, wanted substring <%s>", actualString, expected)
}

func (r *ReaderAssert) EqualToBytes(expected []byte) AssertionMessage {
	actual, result := r.readAll()
	if !result.IsSuccessful {
		return result
	}
	return Assert(bytes.Equal(actual, expected), "got reader bytes <%v>, wanted <%v>", actual, expected)
}

func (r *ReaderAssert) IsEmpty() AssertionMessage {
	actual, result := r.readAll()
	if !result.IsSuccessful {
		return result
	}
	return Assert(len(actual) == 0, "got reader len()=%d, wanted empty reader", len(actual))
}

func (r *ReaderAssert) NotEmpty() AssertionMessage {
	actual, result := r.readAll()
	if !result.IsSuccessful {
		return result
	}
	return Assert(len(actual) != 0, "got empty reader, wanted non-empty reader")
}

func (r *ReaderAssert) readAll() ([]byte, AssertionMessage) {
	if r.actual == nil {
		return nil, Assert(false, "got nil reader, wanted reader body")
	}
	if r.read {
		return r.body, Assert(true, "read reader")
	}

	body, err := io.ReadAll(r.actual)
	if err != nil {
		return nil, Assert(false, "got reader read error: %v", err)
	}
	r.body = body
	r.read = true
	return body, Assert(true, "read reader")
}
