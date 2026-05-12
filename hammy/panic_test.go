package hammy_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/gogunit/gunit/eye"
	a "github.com/gogunit/gunit/hammy"
)

func Test_Panics_success(t *testing.T) {
	assert := a.New(t)

	assert.Is(a.Panics(func() {
		panic("boom")
	}))
}

func Test_Panics_failure(t *testing.T) {
	aSpy := eye.Spy()
	assert := a.New(aSpy)

	assert.Is(a.Panics(func() {}))

	aSpy.HadErrorContaining(t, "wanted panic")
}

func Test_NotPanics_success(t *testing.T) {
	assert := a.New(t)

	assert.Is(a.NotPanics(func() {}))
}

func Test_NotPanics_failure(t *testing.T) {
	aSpy := eye.Spy()
	assert := a.New(aSpy)

	assert.Is(a.NotPanics(func() {
		panic("boom")
	}))

	aSpy.HadErrorContaining(t, "got panic <boom>, wanted no panic")
}

func Test_PanicsWithValue_success(t *testing.T) {
	assert := a.New(t)

	assert.Is(a.PanicsWithValue("boom", func() {
		panic("boom")
	}))
}

func Test_PanicsWithValue_failure_value(t *testing.T) {
	aSpy := eye.Spy()
	assert := a.New(aSpy)

	assert.Is(a.PanicsWithValue("expected", func() {
		panic("actual")
	}))

	aSpy.HadErrorContaining(t, "got panic value <actual>, wanted <expected>")
}

func Test_PanicsWithValue_failure_no_panic(t *testing.T) {
	aSpy := eye.Spy()
	assert := a.New(aSpy)

	assert.Is(a.PanicsWithValue("expected", func() {}))

	aSpy.HadErrorContaining(t, "got no panic, wanted panic value <expected>")
}

func Test_PanicsWithError_success(t *testing.T) {
	assert := a.New(t)

	assert.Is(a.PanicsWithError("boom", func() {
		panic(errors.New("boom"))
	}))
}

func Test_PanicsWithError_failure_non_error(t *testing.T) {
	aSpy := eye.Spy()
	assert := a.New(aSpy)

	assert.Is(a.PanicsWithError("boom", func() {
		panic("boom")
	}))

	aSpy.HadErrorContaining(t, "wanted error panic <boom>")
}

func Test_PanicsWithError_failure_message(t *testing.T) {
	aSpy := eye.Spy()
	assert := a.New(aSpy)

	assert.Is(a.PanicsWithError("expected", func() {
		panic(errors.New("actual"))
	}))

	aSpy.HadErrorContaining(t, "got panic error <actual>, wanted <expected>")
}

func Test_PanicErrorIs_success(t *testing.T) {
	assert := a.New(t)
	target := errors.New("target")

	assert.Is(a.PanicErrorIs(target, func() {
		panic(fmt.Errorf("wrapped: %w", target))
	}))
}

func Test_PanicErrorIs_failure(t *testing.T) {
	aSpy := eye.Spy()
	assert := a.New(aSpy)
	target := errors.New("target")

	assert.Is(a.PanicErrorIs(target, func() {
		panic(errors.New("other"))
	}))

	aSpy.HadErrorContaining(t, "wanted matching <target>")
}
