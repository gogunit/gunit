package hammy

import (
	"bytes"
	"strings"
)

func Buffer(actual *bytes.Buffer) *Buf {
	return &Buf{actual: actual}
}

type Buf struct {
	actual *bytes.Buffer
}

func (b *Buf) Matches(matcher Matcher[*bytes.Buffer]) AssertionMessage {
	if b.actual == nil {
		return Assert(false, "got nil buffer, wanted to match buffer")
	}
	return matcher.Match(b.actual)
}

func (b *Buf) EqualToString(expected string) AssertionMessage {
	if b.actual == nil {
		return Assert(false, "got nil buffer, wanted string <%s>", expected)
	}
	return Assert(b.actual.String() == expected, "got buffer string <%s>, wanted <%s>", b.actual.String(), expected)
}

func (b *Buf) ContainsString(expected string) AssertionMessage {
	if b.actual == nil {
		return Assert(false, "got nil buffer, wanted substring <%s>", expected)
	}
	return Assert(strings.Contains(b.actual.String(), expected), "got buffer string <%s>, wanted substring <%s>", b.actual.String(), expected)
}

func (b *Buf) EqualToBytes(expected []byte) AssertionMessage {
	if b.actual == nil {
		return Assert(false, "got nil buffer, wanted bytes <%v>", expected)
	}
	return Assert(bytes.Equal(b.actual.Bytes(), expected), "got buffer bytes <%v>, wanted <%v>", b.actual.Bytes(), expected)
}

func (b *Buf) Len(expected int) AssertionMessage {
	if b.actual == nil {
		return Assert(false, "got nil buffer, wanted len()=%d", expected)
	}
	return Assert(b.actual.Len() == expected, "got buffer len()=%d, wanted %d", b.actual.Len(), expected)
}

func (b *Buf) IsEmpty() AssertionMessage {
	if b.actual == nil {
		return Assert(false, "got nil buffer, wanted empty buffer")
	}
	return Assert(b.actual.Len() == 0, "got buffer len()=%d, wanted empty buffer", b.actual.Len())
}

func (b *Buf) NotEmpty() AssertionMessage {
	if b.actual == nil {
		return Assert(false, "got nil buffer, wanted non-empty buffer")
	}
	return Assert(b.actual.Len() != 0, "got empty buffer, wanted non-empty buffer")
}
