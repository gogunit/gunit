package hammy_test

import (
	"fmt"
	"github.com/nfisher/gunit/eye"
	"github.com/nfisher/gunit/hammy"
	"testing"
)

func Test_string_EqualTo_failure(t *testing.T) {
	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(String("hi").EqualTo("by"))
	aSpy.HadError(t)
}

func String[S hammy.Stringy](actual S) *Str[S] {
	return &Str[S]{actual: actual}
}

type Str[S hammy.Stringy] struct {
	actual S
}

func (s Str[S]) EqualTo(expected S) hammy.AssertionMessage {
	return hammy.AssertionMessage{
		IsSuccessful: s.actual == expected,
		Message:      fmt.Sprintf("want <%v> equal to <%v>", s.actual, expected),
	}
}
