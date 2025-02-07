package hammy_test

import (
	"testing"

	"github.com/gogunit/gunit/eye"
	"github.com/gogunit/gunit/hammy"
)

func Test_Not_False_failure(t *testing.T) {
	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.IsNot(hammy.False(false))
	aSpy.HadErrorContaining(t, "not(got true, wanted false)")
}

func Test_Not_False_success(t *testing.T) {
	assert := hammy.New(t)
	assert.IsNot(hammy.False(true))
}

func Test_Nil_failure(t *testing.T) {
	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(hammy.Nil(t))
	aSpy.HadErrorContaining(t, "got <*testing.T>, wanted nil")
}

func Test_Nil_success(t *testing.T) {
	assert := hammy.New(t)
	var i *int = nil
	assert.Is(hammy.Nil(i))
}

func Test_NotNil_success(t *testing.T) {
	assert := hammy.New(t)
	var i = 1
	assert.Is(hammy.NotNil(&i))
}

func Test_NotNil_failure(t *testing.T) {
	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	var i *int = nil
	assert.Is(hammy.NotNil(i))
	aSpy.HadErrorContaining(t, "got nil, wanted <*int>")
}

func Test_True_success(t *testing.T) {
	assert := hammy.New(t)
	assert.Is(hammy.True(true))
}

func Test_True_failure(t *testing.T) {
	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(hammy.True(false))
	aSpy.HadErrorContaining(t, "got false, wanted true")
}

func Test_False_success(t *testing.T) {
	assert := hammy.New(t)
	assert.Is(hammy.False(false))
}

func Test_False_failure(t *testing.T) {
	aSpy := eye.Spy()
	assert := hammy.New(aSpy)
	assert.Is(hammy.False(true))
	aSpy.HadErrorContaining(t, "got true, wanted false")
}
