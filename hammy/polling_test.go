package hammy_test

import (
	"testing"
	"time"

	"github.com/gogunit/gunit/eye"
	a "github.com/gogunit/gunit/hammy"
)

func Test_Eventually_success(t *testing.T) {
	assert := a.New(t)
	attempts := 0

	assert.Is(a.Eventually(func() a.AssertionMessage {
		attempts++
		return a.True(attempts >= 3)
	}, 50*time.Millisecond, time.Millisecond))
}

func Test_Eventually_failure(t *testing.T) {
	aSpy := eye.Spy()
	assert := a.New(aSpy)

	assert.Is(a.Eventually(func() a.AssertionMessage {
		return a.False(true)
	}, 5*time.Millisecond, time.Millisecond))

	aSpy.HadErrorContaining(t, "condition did not succeed within")
}

func Test_Never_success(t *testing.T) {
	assert := a.New(t)

	assert.Is(a.Never(func() a.AssertionMessage {
		return a.False(true)
	}, 5*time.Millisecond, time.Millisecond))
}

func Test_Never_failure_immediate(t *testing.T) {
	aSpy := eye.Spy()
	assert := a.New(aSpy)

	assert.Is(a.Never(func() a.AssertionMessage {
		return a.True(true)
	}, 50*time.Millisecond, time.Millisecond))

	aSpy.HadErrorContaining(t, "condition succeeded within")
}

func Test_Never_failure_late(t *testing.T) {
	aSpy := eye.Spy()
	assert := a.New(aSpy)
	attempts := 0

	assert.Is(a.Never(func() a.AssertionMessage {
		attempts++
		return a.True(attempts >= 3)
	}, 50*time.Millisecond, time.Millisecond))

	aSpy.HadErrorContaining(t, "condition succeeded within")
}

func Test_Consistently_success(t *testing.T) {
	assert := a.New(t)

	assert.Is(a.Consistently(func() a.AssertionMessage {
		return a.True(true)
	}, 5*time.Millisecond, time.Millisecond))
}

func Test_Consistently_failure(t *testing.T) {
	aSpy := eye.Spy()
	assert := a.New(aSpy)
	attempts := 0

	assert.Is(a.Consistently(func() a.AssertionMessage {
		attempts++
		return a.True(attempts < 3)
	}, 50*time.Millisecond, time.Millisecond))

	aSpy.HadErrorContaining(t, "condition failed within")
}

func Test_Eventually_uses_default_tick_for_nonpositive_interval(t *testing.T) {
	assert := a.New(t)

	assert.Is(a.Eventually(func() a.AssertionMessage {
		return a.True(true)
	}, 0, 0))
}
