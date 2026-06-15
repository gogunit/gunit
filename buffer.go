package gunit

import (
	"bytes"
	"strings"
)

func Buffer(t T, actual *bytes.Buffer) *Buf {
	return &Buf{T: t, actual: actual}
}

type Buf struct {
	T
	actual *bytes.Buffer
}

func (b *Buf) EqualToString(expected string) {
	b.Helper()
	if b.actual == nil {
		b.Errorf("got nil buffer, wanted string <%s>", expected)
		return
	}
	if b.actual.String() != expected {
		b.Errorf("got buffer string <%s>, wanted <%s>", b.actual.String(), expected)
	}
}

func (b *Buf) ContainsString(expected string) {
	b.Helper()
	if b.actual == nil {
		b.Errorf("got nil buffer, wanted substring <%s>", expected)
		return
	}
	if !strings.Contains(b.actual.String(), expected) {
		b.Errorf("got buffer string <%s>, wanted substring <%s>", b.actual.String(), expected)
	}
}

func (b *Buf) EqualToBytes(expected []byte) {
	b.Helper()
	if b.actual == nil {
		b.Errorf("got nil buffer, wanted bytes <%v>", expected)
		return
	}
	if !bytes.Equal(b.actual.Bytes(), expected) {
		b.Errorf("got buffer bytes <%v>, wanted <%v>", b.actual.Bytes(), expected)
	}
}

func (b *Buf) Len(expected int) {
	b.Helper()
	if b.actual == nil {
		b.Errorf("got nil buffer, wanted len()=%d", expected)
		return
	}
	if b.actual.Len() != expected {
		b.Errorf("got buffer len()=%d, wanted %d", b.actual.Len(), expected)
	}
}

func (b *Buf) IsEmpty() {
	b.Helper()
	if b.actual == nil {
		b.Errorf("got nil buffer, wanted empty buffer")
		return
	}
	if b.actual.Len() != 0 {
		b.Errorf("got buffer len()=%d, wanted empty buffer", b.actual.Len())
	}
}

func (b *Buf) IsNotEmpty() {
	b.Helper()
	if b.actual == nil {
		b.Errorf("got nil buffer, wanted non-empty buffer")
		return
	}
	if b.actual.Len() == 0 {
		b.Errorf("got empty buffer, wanted non-empty buffer")
	}
}
