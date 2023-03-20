package hammy_test

import (
	"testing"

	"github.com/gogunit/gunit/eye"
	"github.com/gogunit/gunit/hammy"
)

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
