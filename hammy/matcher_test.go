package hammy_test

import (
	"strings"
	"testing"

	"github.com/gogunit/gunit/eye"
	a "github.com/gogunit/gunit/hammy"
)

func Test_MatchFunc_success(t *testing.T) {
	assert := a.New(t)
	assert.Is(a.Match(42, a.MatchFunc(func(actual int) a.AssertionMessage {
		return a.Assert(actual == 42, "wrong answer")
	})))
}

func Test_MatchFunc_failure(t *testing.T) {
	aSpy := eye.Spy()
	assert := a.New(aSpy)
	assert.Is(a.Match(42, a.MatchFunc(func(actual int) a.AssertionMessage {
		return a.Assert(actual == 7, "got <%d>, wanted 7", actual)
	})))
	aSpy.HadErrorContaining(t, "got <42>, wanted 7")
}

func Test_Not_rewrites_failure_message(t *testing.T) {
	aSpy := eye.Spy()
	assert := a.New(aSpy)
	assert.Is(a.Match(42, a.Not(a.EqualTo(42))))
	aSpy.HadErrorContaining(t, "not(got <42>, wanted equal to <42>)")
}

func Test_AllOf_short_circuits_on_first_failure(t *testing.T) {
	secondCalled := false
	first := a.MatchFunc(func(actual int) a.AssertionMessage {
		return a.Assert(false, "first failed")
	})
	second := a.MatchFunc(func(actual int) a.AssertionMessage {
		secondCalled = true
		return a.Assert(true, "second matched")
	})

	result := a.Match(42, a.AllOf(first, second))
	if result.IsSuccessful {
		t.Fatalf("expected failure")
	}
	if secondCalled {
		t.Fatalf("expected second matcher not to run")
	}
	if !strings.Contains(result.Message, "first failed") {
		t.Fatalf("got %q", result.Message)
	}
}

func Test_AnyOf_short_circuits_on_first_success(t *testing.T) {
	secondCalled := false
	first := a.MatchFunc(func(actual int) a.AssertionMessage {
		return a.Assert(true, "first matched")
	})
	second := a.MatchFunc(func(actual int) a.AssertionMessage {
		secondCalled = true
		return a.Assert(false, "second failed")
	})

	result := a.Match(42, a.AnyOf(first, second))
	if !result.IsSuccessful {
		t.Fatalf("expected success, got %q", result.Message)
	}
	if secondCalled {
		t.Fatalf("expected second matcher not to run")
	}
}

func Test_Describe_prefixes_failure_message(t *testing.T) {
	aSpy := eye.Spy()
	assert := a.New(aSpy)
	assert.Is(a.Match(42, a.Describe("answer check", a.EqualTo(7))))
	aSpy.HadErrorContaining(t, "answer check: got <42>, wanted equal to <7>")
}
